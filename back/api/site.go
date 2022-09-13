package api

import (
	"context"
	"fmt"
	"github.com/cubdesign/dailyfj/ent"
	"github.com/cubdesign/dailyfj/ent/article"
	"github.com/cubdesign/dailyfj/ent/site"
	"github.com/cubdesign/dailyfj/utils"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
	"github.com/gorilla/feeds"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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
		type SiteCrawlRule struct {
			Url                 string
			ArticleSelector     string
			TitleSelector       string
			LinkSelector        string
			DescriptionSelector string
			hasDataToList       bool
			DateSelector        string
			DateLayout          string
			IsTimeHumanize      bool
		}

		var siteCrawlRule SiteCrawlRule

		loc, _ := time.LoadLocation("Asia/Tokyo")
		now := time.Now().In(loc)

		feed := &feeds.Feed{}
		feed.Link = &feeds.Link{Href: existSite.URL}
		feed.Created = now

		switch existSite.URL {

		case "https://www.snowpeak.co.jp/news/":
			siteCrawlRule = SiteCrawlRule{
				Url:                 existSite.URL,
				ArticleSelector:     ".un_newsList .un_newsList_itemDetail",
				TitleSelector:       ".un_newsList_title",
				LinkSelector:        " > a:last-child",
				DescriptionSelector: "",
				hasDataToList:       true,
				DateSelector:        ".un_newsList_date",
				DateLayout:          "2006.1.2",
				IsTimeHumanize:      false,
			}

		case "https://www.fashion-press.net/news/":
			siteCrawlRule = SiteCrawlRule{
				Url:                 existSite.URL,
				ArticleSelector:     ".pc_only .fp_media_tile.news_media",
				TitleSelector:       " > a > div",
				LinkSelector:        " > a",
				DescriptionSelector: "",
				hasDataToList:       true,
				DateSelector:        " > div > span",
				DateLayout:          "2006.1.2",
				IsTimeHumanize:      true,
			}

		case "https://www.vogue.co.jp/fashion/news":
			siteCrawlRule = SiteCrawlRule{
				Url:                 existSite.URL,
				ArticleSelector:     ".summary-item",
				TitleSelector:       ".summary-item__content h2",
				LinkSelector:        ".summary-item__content > a",
				DescriptionSelector: "",
				hasDataToList:       true,
				DateSelector:        ".summary-item__publish-date",
				DateLayout:          "2006年1月2日",
				IsTimeHumanize:      false,
			}

		case "https://www.elle.com/jp/fashion-news/":
			siteCrawlRule = SiteCrawlRule{
				Url:                 existSite.URL,
				ArticleSelector:     ".custom-item",
				TitleSelector:       ".custom-item-title",
				LinkSelector:        ".custom-item-title",
				DescriptionSelector: "",
				hasDataToList:       false,
				DateSelector:        ".content-info-date",
				DateLayout:          "2006/01/02",
				IsTimeHumanize:      false,
			}

		default:

			c.AbortWithStatus(http.StatusNotFound)
			return

		}

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

		s.OnHTML(siteCrawlRule.ArticleSelector, func(e *colly.HTMLElement) {
			layout := siteCrawlRule.DateLayout

			rootSelector := siteCrawlRule.ArticleSelector

			titleSelector := utils.CreateSelectorOnChildrenScopeFeatureSupport(
				siteCrawlRule.TitleSelector,
				rootSelector,
			)

			title := e.DOM.Find(titleSelector).Text()

			linkSelector := utils.CreateSelectorOnChildrenScopeFeatureSupport(
				siteCrawlRule.LinkSelector,
				rootSelector,
			)

			url, _ := e.DOM.Find(linkSelector).Attr("href")
			url = e.Request.AbsoluteURL(url)

			var description string
			if siteCrawlRule.DescriptionSelector != "" {

				descriptionSelector := utils.CreateSelectorOnChildrenScopeFeatureSupport(
					siteCrawlRule.DescriptionSelector,
					rootSelector,
				)

				description = e.DOM.Find(descriptionSelector).Text()
			}

			var date time.Time
			if siteCrawlRule.hasDataToList {
				dateSelector := utils.CreateSelectorOnChildrenScopeFeatureSupport(
					siteCrawlRule.DateSelector,
					rootSelector,
				)

				dateStr := e.DOM.Find(dateSelector).Text()
				dateStr = strings.TrimSpace(dateStr)
				date, err = time.ParseInLocation(layout, dateStr, loc)
				if err != nil {
					if siteCrawlRule.IsTimeHumanize {
						date, err = utils.HumanizeParseTime(dateStr, now)
					}

					if err != nil {
						c.AbortWithError(http.StatusInternalServerError, err)
						return
					}

				}
			} else {
				existArticle, err := h.Client.Article.
					Query().
					Where(article.URL(url)).
					Only(context.Background())

				if err != nil && !ent.IsNotFound(err) {
					c.AbortWithError(http.StatusBadRequest, err)
					return
				}

				if existArticle == nil {
					s2 := colly.NewCollector()
					s2.Limit(&colly.LimitRule{
						RandomDelay: 5 * time.Second,
					})
					s2.OnHTML("body", func(e *colly.HTMLElement) {
						dateStr := e.DOM.Find(siteCrawlRule.DateSelector).Text()
						dateStr = strings.TrimSpace(dateStr)
						date, err = time.ParseInLocation(layout, dateStr, loc)
						if err != nil {
							if siteCrawlRule.IsTimeHumanize {
								date, err = utils.HumanizeParseTime(dateStr, now)
							}

							if err != nil {
								c.AbortWithError(http.StatusInternalServerError, err)
								return
							}

						}

					})
					s2.Visit(url)
				} else {
					date = existArticle.PublishedAt
				}
			}

			item := &feeds.Item{
				Title:       title,
				Link:        &feeds.Link{Href: url},
				Description: description,
				Created:     date,
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
