package api

import (
	"context"
	"fmt"
	"github.com/cubdesign/dailyfj/ent"
	"github.com/cubdesign/dailyfj/ent/site"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
	"github.com/gorilla/feeds"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type SiteHandler struct {
	Client *ent.Client
}

func (h *SiteHandler) GetAllSites(c *gin.Context) {
	sites, err := h.Client.Site.
		Query().
		Order(ent.Desc(site.FieldUpdatedAt)).
		All(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, sites)
}

func (h *SiteHandler) GetSite(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resSite, err := h.Client.Site.
		Get(context.Background(), id)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if resSite == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, resSite)
}

func (h *SiteHandler) CreateSite(c *gin.Context) {
	var reqSite ent.Site
	err := c.ShouldBindJSON(&reqSite)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resSite, err := h.Client.Site.
		Create().
		SetName(reqSite.Name).
		SetURL(reqSite.URL).
		SetFeedURL(reqSite.FeedURL).
		SetActive(reqSite.Active).
		Save(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &resSite)
}

func (h *SiteHandler) UpdateSite(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var reqSite ent.Site
	err = c.ShouldBindJSON(&reqSite)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	existSite, err := h.Client.Site.
		Get(context.Background(), id)

	if err != nil && !ent.IsNotFound(err) {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if existSite == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	resSite, err := existSite.Update().
		SetName(reqSite.Name).
		SetURL(reqSite.URL).
		SetFeedURL(reqSite.FeedURL).
		SetActive(reqSite.Active).
		Save(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &resSite)
}

func (h *SiteHandler) DeleteSite(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.Client.Site.
		DeleteOneID(id).
		Exec(context.Background())

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (h *SiteHandler) ActiveSite(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	existSite, err := h.Client.Site.
		Get(context.Background(), id)

	if err != nil && !ent.IsNotFound(err) {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if existSite == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	resSite, err := h.Client.Site.
		UpdateOneID(id).
		SetActive(true).
		Save(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &resSite)
}

func (h *SiteHandler) DeActiveSite(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	existSite, err := h.Client.Site.
		Get(context.Background(), id)

	if err != nil && !ent.IsNotFound(err) {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if existSite == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	resSite, err := h.Client.Site.
		UpdateOneID(id).
		SetActive(false).
		Save(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &resSite)
}

func (h *SiteHandler) RunCrawling(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	existSite, err := h.Client.Site.
		Get(context.Background(), id)

	if err != nil && !ent.IsNotFound(err) {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if existSite == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var contents = ""

	if existSite.FeedURL != "" {
		s := colly.NewCollector()
		s.OnError(func(_ *colly.Response, err error) {
			log.Println("Something went wrong:", err)
		})
		s.OnRequest(func(r *colly.Request) {
			fmt.Println("visiting", r.URL)
		})
		s.OnResponse(func(r *colly.Response) {
			contents = string(r.Body)
		})
		s.Visit(existSite.FeedURL)

		if contents == "" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
	} else {
		feed := &feeds.Feed{}
		feed.Link = &feeds.Link{Href: existSite.URL}
		feed.Created = time.Now()

		s := colly.NewCollector()
		s.OnError(func(_ *colly.Response, err error) {
			log.Println("Something went wrong:", err)
		})
		s.OnRequest(func(r *colly.Request) {
			fmt.Println("visiting", r.URL)
		})
		s.OnResponse(func(r *colly.Response) {
			contents = string(r.Body)
		})
		s.OnHTML("title", func(e *colly.HTMLElement) {
			feed.Title = e.Text
		})
		s.OnHTML("meta[name=\"description\"]", func(e *colly.HTMLElement) {
			feed.Description = e.Attr("content")
		})
		s.OnHTML(".un_newsList .un_newsList_itemDetail a:last-child", func(e *colly.HTMLElement) {
			loc, _ := time.LoadLocation("Asia/Tokyo")
			layout := "2006.1.2"
			title := e.ChildText(".un_newsList_title")
			url := e.Attr("href")
			date, err := time.ParseInLocation(layout, fmt.Sprintf("%s", e.ChildText(".un_newsList_date")), loc)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			item := &feeds.Item{
				Title:   title,
				Link:    &feeds.Link{Href: url},
				Created: date,
			}
			feed.Add(item)
		})
		s.Visit(existSite.URL)

		if contents == "" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		rss, err := feed.ToRss()
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		contents = rss
	}

	resFeed, err := h.Client.Feed.
		Create().
		SetContents(contents).
		SetSite(existSite).
		Save(context.Background())

	c.JSON(http.StatusOK, gin.H{"id": resFeed.ID, "url": existSite.URL, "rss": existSite.FeedURL})
}

func (h *SiteHandler) GetRssUrlBySiteId(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	existSite, err := h.Client.Site.
		Get(context.Background(), id)

	if err != nil && !ent.IsNotFound(err) {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if existSite == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	fmt.Println("start crawling")
	rssUrl := getRSSUrl(existSite.URL)
	if rssUrl != "" {
		fmt.Println("link found:", rssUrl)
	}
	fmt.Println("end crawling")

	c.JSON(http.StatusOK, gin.H{"url": rssUrl})
}

func (h *SiteHandler) GetRssUrlByUrl(c *gin.Context) {
	url := c.Query("url")

	if url == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fmt.Println("start crawling")
	rssUrl := getRSSUrl(url)
	if rssUrl != "" {
		fmt.Println("link found:", rssUrl)
	}
	fmt.Println("end crawling")

	c.JSON(http.StatusOK, gin.H{"url": rssUrl})
}

func getRSSUrl(baseUrl string) string {
	var rssUrl = ""
	s := colly.NewCollector()
	s.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	s.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})
	s.OnHTML("link[type=\"application/rss+xml\"], link[type=\"application/atom+xml\"]", func(e *colly.HTMLElement) {
		if rssUrl == "" {
			rssUrl = e.Attr("href")
		}
	})
	s.Visit(baseUrl)

	if rssUrl != "" {
		base, err := url.Parse(baseUrl)
		if err != nil {
			log.Fatal(err)
		}
		ref, err := url.Parse(rssUrl)
		if err != nil {
			log.Fatal(err)
		}
		rssUrl = base.ResolveReference(ref).String()
	}

	return rssUrl
}
