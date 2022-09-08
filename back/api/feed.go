package api

import (
	"context"
	"github.com/cubdesign/dailyfj/ent"
	"github.com/cubdesign/dailyfj/ent/article"
	"github.com/cubdesign/dailyfj/ent/feed"
	"github.com/cubdesign/dailyfj/ent/site"
	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"
	"net/http"
	"strconv"
	"time"
)

type FeedHandler struct {
	Client *ent.Client
}

func (h *FeedHandler) GetAllFeedsNoneContentsField(c *gin.Context) {
	feeds, err := h.Client.Feed.
		Query().
		Select(feed.FieldID, feed.FieldCreatedAt).
		WithSite(func(query *ent.SiteQuery) {
			query.Select(site.FieldName)
		}).
		Order(ent.Desc(feed.FieldCreatedAt)).
		All(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	type ResponseFeed struct {
		ID        int       `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		SiteID    int       `json:"site_id"`
		SiteName  string    `json:"site_name"`
	}
	var resFeeds []ResponseFeed

	for _, feed := range feeds {
		resFeeds = append(resFeeds, ResponseFeed{
			ID:        feed.ID,
			CreatedAt: feed.CreatedAt,
			SiteID:    feed.Edges.Site.ID,
			SiteName:  feed.Edges.Site.Name,
		})
	}
	c.JSON(http.StatusOK, resFeeds)
}

func (h *FeedHandler) GetAllFeeds(c *gin.Context) {
	articles, err := h.Client.Feed.
		Query().
		Order(ent.Desc(feed.FieldCreatedAt)).
		All(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, articles)
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
		existArticle, err := h.Client.Article.
			Query().
			Where(article.URL(item.Link)).
			Only(context.Background())

		if err != nil && !ent.IsNotFound(err) {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if existArticle == nil {
			_, err := h.Client.Article.
				Create().
				SetTitle(item.Title).
				SetURL(item.Link).
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
				SetURL(item.Link).
				SetPublishedAt(*item.PublishedParsed).
				SetSiteID(existFeed.Edges.Site.ID).
				Save(context.Background())
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		}

	}

	c.JSON(http.StatusOK, existFeed)
}
