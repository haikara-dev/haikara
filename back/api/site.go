package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haikara-dev/haikara/ent"
	"github.com/haikara-dev/haikara/ent/site"
	"github.com/haikara-dev/haikara/libs"
	"github.com/mmcdole/gofeed"
	"net/http"
	"strconv"
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
		Query().
		Where(site.ID(id)).
		WithSiteCrawlRule().
		Only(context.Background())

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

/*
	site_crawl_rule: {
	  article_selector: "",
	  title_selector: "",
	  link_selector: "",
	  description_selector: "",
	  has_data_to_list: true,
	  date_selector: "",
	  date_layout: "",
	  is_time_humanize: false,
	  is_spa: false,
	},
*/
func (h *SiteHandler) UpdateSite(c *gin.Context) {

	type RequestSite struct {
		ID            int               `json:"id"`
		Name          string            `json:"name"`
		URL           string            `json:"url"`
		FeedURL       string            `json:"feed_url"`
		Active        bool              `json:"active"`
		SiteCrawlRule ent.SiteCrawlRule `json:"site_crawl_rule"`
	}

	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var reqSite RequestSite
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

	existSiteRule, err := resSite.QuerySiteCrawlRule().Only(context.Background())
	if err != nil && !ent.IsNotFound(err) {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if existSiteRule == nil {
		fmt.Println("create site crawl rule")
		_, err = h.Client.SiteCrawlRule.
			Create().
			SetSite(resSite).
			SetArticleSelector(reqSite.SiteCrawlRule.ArticleSelector).
			SetTitleSelector(reqSite.SiteCrawlRule.TitleSelector).
			SetLinkSelector(reqSite.SiteCrawlRule.LinkSelector).
			SetDescriptionSelector(reqSite.SiteCrawlRule.DescriptionSelector).
			SetHasDataToList(reqSite.SiteCrawlRule.HasDataToList).
			SetDateSelector(reqSite.SiteCrawlRule.DateSelector).
			SetDateLayout(reqSite.SiteCrawlRule.DateLayout).
			SetIsTimeHumanize(reqSite.SiteCrawlRule.IsTimeHumanize).
			SetIsSpa(reqSite.SiteCrawlRule.IsSpa).
			Save(context.Background())

		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		fmt.Println("update site crawl rule")
		_, err = existSiteRule.Update().
			SetArticleSelector(reqSite.SiteCrawlRule.ArticleSelector).
			SetTitleSelector(reqSite.SiteCrawlRule.TitleSelector).
			SetLinkSelector(reqSite.SiteCrawlRule.LinkSelector).
			SetDescriptionSelector(reqSite.SiteCrawlRule.DescriptionSelector).
			SetHasDataToList(reqSite.SiteCrawlRule.HasDataToList).
			SetDateSelector(reqSite.SiteCrawlRule.DateSelector).
			SetDateLayout(reqSite.SiteCrawlRule.DateLayout).
			SetIsTimeHumanize(reqSite.SiteCrawlRule.IsTimeHumanize).
			SetIsSpa(reqSite.SiteCrawlRule.IsSpa).
			Save(context.Background())

		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
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
		contents, err = libs.GetRSS(existSite.FeedURL)
		if err != nil || contents == "" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
	} else {
		siteCrawlRule, err := existSite.QuerySiteCrawlRule().Only(context.Background())
		if err != nil && !ent.IsNotFound(err) {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if siteCrawlRule == nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		contents, err = libs.GetRSSByHTML(existSite.URL, siteCrawlRule, h.Client)
		if err != nil || contents == "" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
	}

	fp := gofeed.NewParser()
	feed, err := fp.ParseString(contents)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	resFeed, err := h.Client.Feed.
		Create().
		SetContents(contents).
		SetCount(len(feed.Items)).
		SetSite(existSite).
		Save(context.Background())

	c.JSON(http.StatusOK, gin.H{"id": resFeed.ID, "url": existSite.URL, "rss": existSite.FeedURL})
}

func (h *SiteHandler) DryRunCrawling(c *gin.Context) {
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
		contents, err = libs.GetRSS(existSite.FeedURL)
		if err != nil || contents == "" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
	} else {
		siteCrawlRule, err := existSite.QuerySiteCrawlRule().Only(context.Background())
		if err != nil && !ent.IsNotFound(err) {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if siteCrawlRule == nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		contents, err = libs.GetRSSByHTML(existSite.URL, siteCrawlRule, h.Client)
		if err != nil || contents == "" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
	}

	fp := gofeed.NewParser()
	feed, err := fp.ParseString(contents)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": len(feed.Items), "contents": contents})
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
	rssUrl, err := libs.GetRSSUrl(existSite.URL)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
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
	rssUrl, err := libs.GetRSSUrl(url)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if rssUrl != "" {
		fmt.Println("link found:", rssUrl)
	}
	fmt.Println("end crawling")

	c.JSON(http.StatusOK, gin.H{"url": rssUrl})
}
