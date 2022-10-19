package libs

import (
	"context"
	"fmt"
	"github.com/haikara-dev/haikara/ent"
	"github.com/haikara-dev/haikara/ent/article"
	"github.com/haikara-dev/haikara/ent/feed"
	"github.com/haikara-dev/haikara/utils"
	"github.com/mmcdole/gofeed"
	"log"
	"time"
)

func IndexFeed(feed *ent.Feed, client *ent.Client) ([]*ent.Article, error) {

	fp := gofeed.NewParser()
	parsedFeed, err := fp.ParseString(feed.Contents)

	if err != nil {
		return nil, err
	}

	var articles []*ent.Article

	for _, item := range parsedFeed.Items {
		link := utils.AddSchemeIfNotExists(item.Link)
		existArticle, err := client.Article.
			Query().
			Where(article.URL(link)).
			Only(context.Background())

		if err != nil && !ent.IsNotFound(err) {
			return nil, err
		}
		if existArticle == nil {
			existArticle, err = client.Article.
				Create().
				SetTitle(item.Title).
				SetURL(link).
				SetPublishedAt(*item.PublishedParsed).
				SetSiteID(feed.Edges.Site.ID).
				Save(context.Background())
			if err != nil {
				return nil, err
			}
		} else {
			existArticle, err = client.Article.
				UpdateOne(existArticle).
				SetTitle(item.Title).
				SetURL(link).
				SetPublishedAt(*item.PublishedParsed).
				SetSiteID(feed.Edges.Site.ID).
				Save(context.Background())
			if err != nil {
				return nil, err
			}
		}
		articles = append(articles, existArticle)
	}

	loc, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(loc)

	_, err = feed.Update().
		SetIndexedAt(now).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return articles, nil
}

func IndexAllFeed(client *ent.Client) ([]*ent.Feed, error) {

	feeds, err := client.Feed.
		Query().
		WithSite().
		Where(feed.IndexedAtIsNil()).
		Where(feed.CountGT(0)).
		Order(ent.Asc(feed.FieldCreatedAt)).
		All(context.Background())

	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if feeds == nil {
		return nil, nil
	}
	for _, feed := range feeds {
		articles, err := IndexFeed(feed, client)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(feed.Edges.Site.Name, len(articles))
	}
	return feeds, nil
}
