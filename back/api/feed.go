package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/haikara-dev/haikara/config"
	"github.com/haikara-dev/haikara/ent"
	"github.com/haikara-dev/haikara/ent/article"
	"github.com/haikara-dev/haikara/ent/feed"
	"github.com/haikara-dev/haikara/ent/site"
	"github.com/haikara-dev/haikara/utils"
	"github.com/mmcdole/gofeed"
	"math"
	"net/http"
	"strconv"
	"time"
)

type FeedHandler struct {
	Client *ent.Client
}

func (h *FeedHandler) GetAllFeedsNoneContentsField(c *gin.Context) {
	pageStr := c.Query("page")

	if pageStr == "" {
		pageStr = "1"
	}

	page, err := strconv.Atoi(pageStr)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	pageSize := config.Config.PageSize

	offset := (page - 1) * pageSize

	feeds, err := h.Client.Feed.
		Query().
		Select(feed.FieldID, feed.FieldCount, feed.FieldCreatedAt, feed.FieldIndexedAt).
		WithSite(func(query *ent.SiteQuery) {
			query.Select(site.FieldName)
		}).
		Order(ent.Desc(feed.FieldCreatedAt)).
		Offset(offset).
		Limit(pageSize).
		All(context.Background())

	if err != nil && !ent.IsNotFound(err) {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	totalCount, err := h.Client.Feed.
		Query().
		Count(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	type ResponseFeed struct {
		ID        int        `json:"id"`
		Count     int        `json:"count"`
		CreatedAt time.Time  `json:"created_at"`
		SiteID    int        `json:"site_id"`
		SiteName  string     `json:"site_name"`
		IndexedAt *time.Time `json:"indexed_at"`
	}

	type ResponseJson struct {
		TotalCount int            `json:"totalCount"`
		TotalPage  int            `json:"totalPage"`
		PageSize   int            `json:"pageSize"`
		Data       []ResponseFeed `json:"data"`
	}

	var resFeeds = make([]ResponseFeed, 0)

	for _, feed := range feeds {
		resFeeds = append(resFeeds, ResponseFeed{
			ID:        feed.ID,
			Count:     feed.Count,
			CreatedAt: feed.CreatedAt,
			SiteID:    feed.Edges.Site.ID,
			SiteName:  feed.Edges.Site.Name,
			IndexedAt: feed.IndexedAt,
		})
	}
	c.JSON(http.StatusOK, ResponseJson{
		TotalCount: totalCount,
		TotalPage:  int(math.Ceil(float64(totalCount) / float64(pageSize))),
		PageSize:   pageSize,
		Data:       resFeeds,
	})
}

func (h *FeedHandler) GetAllFeeds(c *gin.Context) {
	articles, err := h.Client.Feed.
		Query().
		Order(ent.Desc(feed.FieldCreatedAt)).
		All(context.Background())

	if err != nil && !ent.IsNotFound(err) {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, articles)
}

func (h *FeedHandler) DeleteFeed(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.Client.Feed.
		DeleteOneID(id).
		Exec(context.Background())

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (h *FeedHandler) ParseFeed(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	existFeed, err := h.Client.Feed.
		Query().
		WithSite(func(query *ent.SiteQuery) {
			query.Select(site.FieldID)
		}).
		Where(feed.ID(id)).
		Only(context.Background())

	if err != nil && !ent.IsNotFound(err) {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if existFeed == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	fp := gofeed.NewParser()
	feed, err := fp.ParseString(existFeed.Contents)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	for _, item := range feed.Items {
		link := utils.AddSchemeIfNotExists(item.Link)
		existArticle, err := h.Client.Article.
			Query().
			Where(article.URL(link)).
			Only(context.Background())

		if err != nil && !ent.IsNotFound(err) {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if existArticle == nil {
			_, err := h.Client.Article.
				Create().
				SetTitle(item.Title).
				SetURL(link).
				SetPublishedAt(*item.PublishedParsed).
				SetSiteID(existFeed.Edges.Site.ID).
				Save(context.Background())
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		} else {
			_, err := h.Client.Article.
				UpdateOne(existArticle).
				SetTitle(item.Title).
				SetURL(link).
				SetPublishedAt(*item.PublishedParsed).
				SetSiteID(existFeed.Edges.Site.ID).
				Save(context.Background())
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		}

	}

	loc, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(loc)

	existFeed, err = existFeed.Update().
		SetIndexedAt(now).
		Save(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, existFeed)
}
