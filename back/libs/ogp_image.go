package libs

import (
	"context"
	"errors"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/haikara-dev/haikara/config"
	"github.com/haikara-dev/haikara/ent"
	"github.com/haikara-dev/haikara/ent/article"
	"github.com/haikara-dev/haikara/ent/ogpimage"
	"github.com/haikara-dev/haikara/utils"
	"github.com/kennygrant/sanitize"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func getOGPImageUrl(articleUrl string) (string, error) {
	var ogpUrl = ""
	s := colly.NewCollector(
		colly.UserAgent(config.Config.UserAgent),
	)
	s.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	s.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})
	s.OnHTML("meta[property=\"og:image\"]", func(e *colly.HTMLElement) {
		ogpUrl = e.Attr("content")
	})
	s.Visit(articleUrl)

	if ogpUrl != "" {
		base, err := url.Parse(articleUrl)
		if err != nil {
			return "", err
		}
		ref, err := url.Parse(ogpUrl)
		if err != nil {
			return "", err
		}
		ogpUrl = base.ResolveReference(ref).String()
	}

	return ogpUrl, nil
}

type SaveOGPImageFileResponse struct {
	FileName string
	FilePath string
}

func saveOGPImageFile(ogpImageUrl string, saveDir string, articleID int) (*SaveOGPImageFileResponse, error) {
	var res *SaveOGPImageFileResponse
	var err error

	if ogpImageUrl == "" {
		return nil, errors.New("ogpImageUrl is empty")
	}

	if saveDir == "" {
		return nil, errors.New("savePath is empty")
	}

	if articleID == 0 {
		return nil, errors.New("articleID is empty")
	}

	s := colly.NewCollector(
		colly.UserAgent(config.Config.UserAgent),
	)
	s.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	s.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})
	s.OnResponse(func(r *colly.Response) {
		os.MkdirAll(saveDir, os.ModePerm)

		filename := strconv.Itoa(articleID)
		ext := filepath.Ext(r.Request.URL.String())
		cleanExt := sanitize.BaseName(ext)
		fileName := fmt.Sprintf("%s.%s", filename, cleanExt[1:])
		filePath := saveDir + fileName

		err = r.Save(filePath)
		if err == nil {
			res = &SaveOGPImageFileResponse{
				fileName,
				filePath,
			}
		}
	})

	s.Visit(ogpImageUrl)

	return res, err
}

func SaveOGPImage(targetArticle *ent.Article, client *ent.Client) (*ent.OGPImage, error) {
	var err error
	var ogpImageURL = ""
	var fileName = ""
	var filePath = ""

	ogpImageURL, err = getOGPImageUrl(targetArticle.URL)
	if err != nil {
		// 何もしない
	}

	if ogpImageURL != "" {
		saveDir := "uploads/ogp_images/" + utils.DirectoryNameFromTime(targetArticle.PublishedAt)
		saveOGPImageFileResponse, err := saveOGPImageFile(ogpImageURL, saveDir, targetArticle.ID)
		if err != nil {
			// 何もしない
		}

		if saveOGPImageFileResponse != nil {
			fileName = saveOGPImageFileResponse.FileName
			filePath = saveOGPImageFileResponse.FilePath
		}

	}

	ogpImage, err := client.OGPImage.
		Query().
		Where(ogpimage.HasArticleWith(article.ID(targetArticle.ID))).
		Only(context.Background())

	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if ogpImage == nil {
		ogpImage, err = client.OGPImage.
			Create().
			SetArticleID(targetArticle.ID).
			SetOriginURL(ogpImageURL).
			SetFileName(fileName).
			SetFilePath(filePath).
			Save(context.Background())

		if err != nil {
			return nil, err
		}
	} else {
		ogpImage, err = ogpImage.Update().
			SetArticleID(targetArticle.ID).
			SetOriginURL(ogpImageURL).
			SetFileName(fileName).
			SetFilePath(filePath).
			Save(context.Background())

		if err != nil {
			return nil, err
		}
	}
	return ogpImage, nil
}

func CrawlOGPImage(limit int, client *ent.Client) ([]*ent.Article, error) {
	articles, err := client.Article.
		Query().
		Where(article.Not(article.HasOgpImage())).
		Order(ent.Desc(article.FieldPublishedAt)).
		Limit(limit).
		All(context.Background())

	if err != nil {
		return nil, err
	}

	for _, article := range articles {

		ogpImage, err := SaveOGPImage(article, client)
		if err != nil {
			log.Println(err)
			// エラーが発生しても続行する
		} else {
			fmt.Println(ogpImage.OriginURL)
		}
		time.Sleep(10 * time.Second) // 時間	10秒休む
	}

	return articles, nil
}
