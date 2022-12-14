package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haikara-dev/haikara/config"
	"github.com/haikara-dev/haikara/ent"
	"github.com/haikara-dev/haikara/ent/article"
	"github.com/haikara-dev/haikara/ent/feed"
	"github.com/haikara-dev/haikara/ent/site"
	"github.com/haikara-dev/haikara/ent/sitecategory"
	"github.com/haikara-dev/haikara/ent/sitecrawlrule"
	"github.com/haikara-dev/haikara/libs"
	"github.com/mmcdole/gofeed"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type SiteHandler struct {
	Client *ent.Client
}

func (h *SiteHandler) GetAllSites(c *gin.Context) {
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

	sites, err := h.Client.Site.
		Query().
		WithSiteCategories().
		Order(ent.Desc(site.FieldUpdatedAt)).
		Offset(offset).
		Limit(pageSize).
		All(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	totalCount, err := h.Client.Site.
		Query().
		Count(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	type ResponseSiteCategory struct {
		ID    int    `json:"id"`
		Label string `json:"label"`
	}

	type ResponseSite struct {
		ID             int                     `json:"id"`
		CreatedAt      time.Time               `json:"created_at"`
		UpdatedAt      time.Time               `json:"updated_at"`
		Name           string                  `json:"name"`
		URL            string                  `json:"url"`
		FeedURL        string                  `json:"feed_url"`
		Active         bool                    `json:"active"`
		CannotCrawlAt  *time.Time              `json:"cannot_crawl_at"`
		SiteCategories []*ResponseSiteCategory `json:"site_categories"`
	}

	type ResponseJson struct {
		TotalCount int            `json:"totalCount"`
		TotalPage  int            `json:"totalPage"`
		PageSize   int            `json:"pageSize"`
		Data       []ResponseSite `json:"data"`
	}

	var resSite = make([]ResponseSite, 0)

	createResponseSiteCategories := func(siteCategories []*ent.SiteCategory) []*ResponseSiteCategory {
		var resSiteCategories = make([]*ResponseSiteCategory, 0)

		for _, siteCategory := range siteCategories {
			resSiteCategories = append(resSiteCategories, &ResponseSiteCategory{
				ID:    siteCategory.ID,
				Label: siteCategory.Label,
			})
		}
		return resSiteCategories
	}

	for _, site := range sites {
		resSite = append(resSite, ResponseSite{
			ID:             site.ID,
			CreatedAt:      site.CreatedAt,
			UpdatedAt:      site.UpdatedAt,
			Name:           site.Name,
			URL:            site.URL,
			FeedURL:        site.FeedURL,
			Active:         site.Active,
			CannotCrawlAt:  site.CannotCrawlAt,
			SiteCategories: createResponseSiteCategories(site.Edges.SiteCategories),
		})
	}

	c.JSON(http.StatusOK, ResponseJson{
		TotalCount: totalCount,
		TotalPage:  int(math.Ceil(float64(totalCount) / float64(pageSize))),
		PageSize:   pageSize,
		Data:       resSite,
	})
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
		WithSiteCategories().
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

	type RequestSite struct {
		Name            string            `json:"name"`
		URL             string            `json:"url"`
		FeedURL         string            `json:"feed_url"`
		Active          bool              `json:"active"`
		SiteCrawlRule   ent.SiteCrawlRule `json:"site_crawl_rule"`
		SiteCategoryIds []int             `json:"site_category_ids"`
	}

	var reqSite RequestSite
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

	if len(reqSite.SiteCategoryIds) > 0 {
		_, err = resSite.Update().
			AddSiteCategoryIDs(reqSite.SiteCategoryIds...).
			Save(context.Background())
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	c.JSON(http.StatusOK, &resSite)
}

func (h *SiteHandler) UpdateSite(c *gin.Context) {

	type RequestSite struct {
		ID              int               `json:"id"`
		Name            string            `json:"name"`
		URL             string            `json:"url"`
		FeedURL         string            `json:"feed_url"`
		Active          bool              `json:"active"`
		CannotCrawl     bool              `json:"cannot_crawl"`
		SiteCrawlRule   ent.SiteCrawlRule `json:"site_crawl_rule"`
		SiteCategoryIds []int             `json:"site_category_ids"`
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

	var cannotCrawlAt *time.Time = nil

	if reqSite.CannotCrawl {
		if existSite.CannotCrawlAt == nil {
			loc, _ := time.LoadLocation("Asia/Tokyo")
			now := time.Now().In(loc)
			cannotCrawlAt = &now
		} else {
			cannotCrawlAt = existSite.CannotCrawlAt
		}
	}

	var resSite *ent.Site

	if cannotCrawlAt == nil {
		resSite, err = existSite.Update().
			SetName(reqSite.Name).
			SetURL(reqSite.URL).
			SetFeedURL(reqSite.FeedURL).
			SetActive(reqSite.Active).
			ClearCannotCrawlAt().
			Save(context.Background())
	} else {
		resSite, err = existSite.Update().
			SetName(reqSite.Name).
			SetURL(reqSite.URL).
			SetFeedURL(reqSite.FeedURL).
			SetActive(reqSite.Active).
			SetNillableCannotCrawlAt(cannotCrawlAt).
			Save(context.Background())
	}

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

	if len(reqSite.SiteCategoryIds) > 0 {
		_, err = resSite.Update().
			ClearSiteCategories().
			AddSiteCategoryIDs(reqSite.SiteCategoryIds...).
			Save(context.Background())
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		_, err = resSite.Update().
			ClearSiteCategories().
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

	existSite, err := h.Client.Site.
		Get(context.Background(), id)

	if err != nil && !ent.IsNotFound(err) {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if existSite == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	_, err = h.Client.SiteCrawlRule.
		Delete().
		Where(sitecrawlrule.HasSiteWith(site.ID(id))).
		Exec(context.Background())

	if err != nil && !ent.IsNotFound(err) {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err = h.Client.SiteCategory.
		Delete().
		Where(sitecategory.HasSitesWith(site.ID(id))).
		Exec(context.Background())

	if err != nil && !ent.IsNotFound(err) {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err = h.Client.Feed.
		Delete().
		Where(feed.HasSiteWith(site.ID(id))).
		Exec(context.Background())

	if err != nil && !ent.IsNotFound(err) {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	articles, err := h.Client.Article.
		Query().
		Where(article.HasSiteWith(site.ID(id))).
		All(context.Background())

	if err != nil && !ent.IsNotFound(err) {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for _, article := range articles {

		err = libs.DeleteOGPImage(article, h.Client)

		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		err = h.Client.Article.
			DeleteOne(article).
			Exec(context.Background())

		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	if err != nil && !ent.IsNotFound(err) {
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

type SiteCrawlRuleImportExport struct {
	ArticleSelector     string `json:"article_selector"`
	TitleSelector       string `json:"title_selector"`
	LinkSelector        string `json:"link_selector"`
	DescriptionSelector string `json:"description_selector"`
	HasDataToList       bool   `json:"has_data_to_list"`
	DateSelector        string `json:"date_selector"`
	DateLayout          string `json:"date_layout"`
	IsTimeHumanize      bool   `json:"is_time_humanize"`
	IsSpa               bool   `json:"is_spa"`
}

type SiteCategoryImportExport struct {
	Label string `json:"label"`
}

type SiteImportExport struct {
	Name          string                      `json:"name"`
	URL           string                      `json:"url"`
	FeedURL       string                      `json:"feed_url"`
	Active        bool                        `json:"active"`
	CannotCrawlAt *time.Time                  `json:"cannot_crawl_at"`
	SiteCrawlRule *SiteCrawlRuleImportExport  `json:"site_crawl_rule"`
	SiteCategory  []*SiteCategoryImportExport `json:"site_category"`
}

func (h *SiteHandler) ExportSites(c *gin.Context) {
	loadSites, err := h.Client.Site.
		Query().
		WithSiteCrawlRule().
		WithSiteCategories().
		All(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var sites []SiteImportExport
	for _, loadSite := range loadSites {
		site := SiteImportExport{
			Name:          loadSite.Name,
			URL:           loadSite.URL,
			FeedURL:       loadSite.FeedURL,
			Active:        loadSite.Active,
			CannotCrawlAt: loadSite.CannotCrawlAt,
		}

		if loadSite.Edges.SiteCrawlRule != nil {
			site.SiteCrawlRule = &SiteCrawlRuleImportExport{
				ArticleSelector:     loadSite.Edges.SiteCrawlRule.ArticleSelector,
				TitleSelector:       loadSite.Edges.SiteCrawlRule.TitleSelector,
				LinkSelector:        loadSite.Edges.SiteCrawlRule.LinkSelector,
				DescriptionSelector: loadSite.Edges.SiteCrawlRule.DescriptionSelector,
				HasDataToList:       loadSite.Edges.SiteCrawlRule.HasDataToList,
				DateSelector:        loadSite.Edges.SiteCrawlRule.DateSelector,
				DateLayout:          loadSite.Edges.SiteCrawlRule.DateLayout,
				IsTimeHumanize:      loadSite.Edges.SiteCrawlRule.IsTimeHumanize,
				IsSpa:               loadSite.Edges.SiteCrawlRule.IsSpa,
			}
		} else {
			site.SiteCrawlRule = nil
		}

		var sitesCategories = make([]*SiteCategoryImportExport, 0)
		if loadSite.Edges.SiteCategories != nil {
			for _, siteCategory := range loadSite.Edges.SiteCategories {
				sitesCategories = append(sitesCategories, &SiteCategoryImportExport{
					Label: siteCategory.Label,
				})
			}
		}

		site.SiteCategory = sitesCategories

		sites = append(sites, site)
	}

	c.JSON(http.StatusOK, sites)
}

func insertOrUpdateSite(siteImportExport *SiteImportExport, siteCategoryMap *map[string]int, client *ent.Client) (*ent.Site, error) {
	existSite, err := client.Site.
		Query().
		Where(site.URL(siteImportExport.URL)).
		Only(context.Background())

	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if existSite != nil {

		siteUpdateOne := client.Site.
			UpdateOne(existSite).
			SetName(siteImportExport.Name).
			SetURL(siteImportExport.URL).
			SetFeedURL(siteImportExport.FeedURL).
			SetActive(siteImportExport.Active)

		if siteImportExport.CannotCrawlAt != nil {
			siteUpdateOne.SetNillableCannotCrawlAt(siteImportExport.CannotCrawlAt)
		} else {
			siteUpdateOne.ClearCannotCrawlAt()
		}

		_, err = siteUpdateOne.Save(context.Background())

		if err != nil {
			return nil, err
		}
	} else {
		existSite, err = client.Site.
			Create().
			SetName(siteImportExport.Name).
			SetURL(siteImportExport.URL).
			SetFeedURL(siteImportExport.FeedURL).
			SetActive(siteImportExport.Active).
			SetNillableCannotCrawlAt(siteImportExport.CannotCrawlAt).
			Save(context.Background())
		if err != nil {
			return nil, err
		}
	}

	if siteImportExport.SiteCrawlRule != nil {
		existSiteCrawlRule, err := client.SiteCrawlRule.
			Query().
			Where(sitecrawlrule.HasSiteWith(site.ID(existSite.ID))).
			Only(context.Background())

		if err != nil && !ent.IsNotFound(err) {
			return nil, err
		}

		if existSiteCrawlRule != nil {
			_, err = client.SiteCrawlRule.
				UpdateOne(existSiteCrawlRule).
				SetArticleSelector(siteImportExport.SiteCrawlRule.ArticleSelector).
				SetTitleSelector(siteImportExport.SiteCrawlRule.TitleSelector).
				SetLinkSelector(siteImportExport.SiteCrawlRule.LinkSelector).
				SetDescriptionSelector(siteImportExport.SiteCrawlRule.DescriptionSelector).
				SetHasDataToList(siteImportExport.SiteCrawlRule.HasDataToList).
				SetDateSelector(siteImportExport.SiteCrawlRule.DateSelector).
				SetDateLayout(siteImportExport.SiteCrawlRule.DateLayout).
				SetIsTimeHumanize(siteImportExport.SiteCrawlRule.IsTimeHumanize).
				SetIsSpa(siteImportExport.SiteCrawlRule.IsSpa).
				Save(context.Background())
			if err != nil {
				return nil, err
			}
		} else {
			_, err = client.SiteCrawlRule.
				Create().
				SetArticleSelector(siteImportExport.SiteCrawlRule.ArticleSelector).
				SetTitleSelector(siteImportExport.SiteCrawlRule.TitleSelector).
				SetLinkSelector(siteImportExport.SiteCrawlRule.LinkSelector).
				SetDescriptionSelector(siteImportExport.SiteCrawlRule.DescriptionSelector).
				SetHasDataToList(siteImportExport.SiteCrawlRule.HasDataToList).
				SetDateSelector(siteImportExport.SiteCrawlRule.DateSelector).
				SetDateLayout(siteImportExport.SiteCrawlRule.DateLayout).
				SetIsTimeHumanize(siteImportExport.SiteCrawlRule.IsTimeHumanize).
				SetIsSpa(siteImportExport.SiteCrawlRule.IsSpa).
				SetSite(existSite).
				Save(context.Background())
			if err != nil {
				return nil, err
			}
		}
	}

	if siteImportExport.SiteCategory != nil && len(siteImportExport.SiteCategory) > 0 {
		var siteCategoryIDs = make([]int, 0)
		for _, siteCategory := range siteImportExport.SiteCategory {
			siteCategoryID, ok := (*siteCategoryMap)[siteCategory.Label]
			if ok {
				siteCategoryIDs = append(siteCategoryIDs, siteCategoryID)
			}
		}

		existSite.Update().
			ClearSiteCategories().
			AddSiteCategoryIDs(siteCategoryIDs...).
			Save(context.Background())

	}

	resSite, err := client.Site.
		Query().
		Where(site.ID(existSite.ID)).
		WithSiteCrawlRule().
		Only(context.Background())

	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	return resSite, nil
}

func (h *SiteHandler) ImportSites(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	dist := filepath.Join("tmp", "export-sites.json")
	err = c.SaveUploadedFile(fileHeader, dist)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data, err := os.ReadFile(dist)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var reqSites []*SiteImportExport
	err = json.Unmarshal(data, &reqSites)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Site Category
	_, err = h.Client.SiteCategory.
		Delete().
		Exec(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	siteCategoryMap := make(map[string]int)
	for _, reqSite := range reqSites {
		if reqSite.SiteCategory != nil && len(reqSite.SiteCategory) > 0 {
			for _, siteCategory := range reqSite.SiteCategory {
				_, ok := siteCategoryMap[siteCategory.Label]
				if !ok {
					newSiteCategory, err := h.Client.SiteCategory.
						Create().
						SetLabel(siteCategory.Label).
						Save(context.Background())

					if err != nil {
						c.AbortWithError(http.StatusBadRequest, err)
						return
					}

					siteCategoryMap[siteCategory.Label] = newSiteCategory.ID
				}
			}
		}
	}

	fmt.Println(siteCategoryMap)
	// Site

	var resSites []*ent.Site
	for _, reqSite := range reqSites {
		site, err := insertOrUpdateSite(reqSite, &siteCategoryMap, h.Client)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		resSites = append(resSites, site)
	}
	c.JSON(http.StatusOK, gin.H{
		"count":    len(resSites),
		"reqSites": resSites,
	})
}
