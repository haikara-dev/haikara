package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/haikara-dev/haikara/config"
	"github.com/haikara-dev/haikara/ent"
	"github.com/haikara-dev/haikara/ent/article"
	"github.com/haikara-dev/haikara/ent/site"
	"github.com/haikara-dev/haikara/libs"
	"math"
	"net/http"
	"strconv"
	"time"
)

type ArticleHandler struct {
	Client *ent.Client
}

func (h *ArticleHandler) GetAllArticles(c *gin.Context) {
	pageStr := c.Query("page")

	if pageStr == "" {
		pageStr = "1"
	}

	page, err := strconv.Atoi(pageStr)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var siteId int
	siteIdStr := c.Query("site_id")

	if siteIdStr != "" {
		siteId, err = strconv.Atoi(siteIdStr)

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}
	pageSize := config.Config.PageSize

	offset := (page - 1) * pageSize

	createArticleQuery := func() *ent.ArticleQuery {
		articleQuery := h.Client.Article.
			Query()

		if siteId != 0 {
			articleQuery = articleQuery.Where(article.HasSiteWith(site.ID(siteId)))
		}
		return articleQuery
	}

	articles, err := createArticleQuery().
		WithOgpImage().
		WithSite().
		Order(ent.Desc(article.FieldPublishedAt)).
		Offset(offset).
		Limit(pageSize).
		All(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	totalCount, err := createArticleQuery().
		Count(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	type ResponseArticleSite struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	type ResponseArticle struct {
		ID          int                 `json:"id"`
		Title       string              `json:"title"`
		URL         string              `json:"url"`
		PublishedAt time.Time           `json:"published_at"`
		OGPImageURL string              `json:"ogp_image_url"`
		Site        ResponseArticleSite `json:"site"`
	}

	type ResponseJson struct {
		TotalCount int               `json:"totalCount"`
		TotalPage  int               `json:"totalPage"`
		PageSize   int               `json:"pageSize"`
		Data       []ResponseArticle `json:"data"`
	}

	var resFeeds = make([]ResponseArticle, 0)

	for _, article := range articles {

		ogpImageURL := ""
		if article.Edges.OgpImage != nil && article.Edges.OgpImage.FilePath != "" {
			ogpImageURL = config.Config.AssetsUrl + "/" + article.Edges.OgpImage.FilePath
		}

		site := article.Edges.Site

		resFeeds = append(resFeeds, ResponseArticle{
			ID:          article.ID,
			Title:       article.Title,
			URL:         article.URL,
			PublishedAt: article.PublishedAt,
			OGPImageURL: ogpImageURL,
			Site: ResponseArticleSite{
				ID:   site.ID,
				Name: site.Name,
				URL:  site.URL,
			},
		})
	}

	c.JSON(http.StatusOK, ResponseJson{
		TotalCount: totalCount,
		TotalPage:  int(math.Ceil(float64(totalCount) / float64(pageSize))),
		PageSize:   pageSize,
		Data:       resFeeds,
	})
}

func (h *ArticleHandler) GetArticle(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resArticle, err := h.Client.Article.
		Get(context.Background(), id)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if resArticle == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, resArticle)
}

func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	var reqArticle ent.Article
	err := c.ShouldBindJSON(&reqArticle)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	siteId := reqArticle.Edges.Site.ID

	resArticle, err := h.Client.Article.
		Create().
		SetTitle(reqArticle.Title).
		SetURL(reqArticle.URL).
		SetSiteID(siteId).
		Save(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &resArticle)
}

func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var reqArticle ent.Article
	err = c.ShouldBindJSON(&reqArticle)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	existArticle, err := h.Client.Article.
		Get(context.Background(), id)

	if err != nil && !ent.IsNotFound(err) {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if existArticle == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	siteId := reqArticle.Edges.Site.ID
	resArticle, err := existArticle.Update().
		SetTitle(reqArticle.Title).
		SetURL(reqArticle.URL).
		SetSiteID(siteId).
		Save(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &resArticle)
}
func (h *ArticleHandler) DeleteArticle(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	existArticle, err := h.Client.Article.
		Get(context.Background(), id)

	if err != nil && !ent.IsNotFound(err) {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if existArticle == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	err = libs.DeleteOGPImage(existArticle, h.Client)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.Client.Article.
		DeleteOne(existArticle).
		Exec(context.Background())

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (h *ArticleHandler) RunGetOGPImageOfArticle(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	existArticle, err := h.Client.Article.
		Query().
		Where(article.ID(id)).
		Only(context.Background())

	if err != nil && !ent.IsNotFound(err) {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if existArticle == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	ogpImage, err := libs.SaveOGPImage(existArticle, h.Client)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": ogpImage.OriginURL})
}
