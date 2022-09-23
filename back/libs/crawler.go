package libs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	"github.com/gorilla/feeds"
	"github.com/haikara-dev/haikara/config"
	"github.com/haikara-dev/haikara/ent"
	"github.com/haikara-dev/haikara/ent/article"
	"github.com/haikara-dev/haikara/ent/site"
	"github.com/haikara-dev/haikara/utils"
	"github.com/mmcdole/gofeed"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//type SiteCrawlRuleByFile struct {
//	Url                 string
//	ArticleSelector     string
//	TitleSelector       string
//	LinkSelector        string
//	DescriptionSelector string
//	HasDataToList       bool
//	DateSelector        string
//	DateLayout          string
//	IsTimeHumanize      bool
//	IsSpa               bool
//}

func GetRSSUrl(baseUrl string) (string, error) {
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
			return "", err
		}
		ref, err := url.Parse(rssUrl)
		if err != nil {
			return "", err
		}
		rssUrl = base.ResolveReference(ref).String()
	}

	return rssUrl, nil
}

func GetRSS(feedUrl string) (string, error) {
	var contents string
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
	s.Visit(feedUrl)

	if contents == "" {
		return "", fmt.Errorf("no contents")
	}
	return contents, nil
}

//	func GetSiteCrawlRule(siteUrl string) (SiteCrawlRuleByFile, error) {
//		var siteCrawlRule SiteCrawlRuleByFile
//
//		switch siteUrl {
//
//		case "https://www.snowpeak.co.jp/news/":
//			siteCrawlRule = SiteCrawlRuleByFile{
//				Url:                 siteUrl,
//				ArticleSelector:     ".un_newsList .un_newsList_itemDetail",
//				TitleSelector:       ".un_newsList_title",
//				LinkSelector:        " > a:last-child",
//				DescriptionSelector: "",
//				HasDataToList:       true,
//				DateSelector:        ".un_newsList_date",
//				DateLayout:          "2006.1.2",
//				IsTimeHumanize:      false,
//				IsSpa:               false,
//			}
//
//		case "https://www.fashion-press.net/news/":
//			siteCrawlRule = SiteCrawlRuleByFile{
//				Url:                 siteUrl,
//				ArticleSelector:     ".pc_only .fp_media_tile.news_media",
//				TitleSelector:       " > a > div",
//				LinkSelector:        " > a",
//				DescriptionSelector: "",
//				HasDataToList:       true,
//				DateSelector:        " > div > span",
//				DateLayout:          "2006.1.2",
//				IsTimeHumanize:      true,
//				IsSpa:               false,
//			}
//
//		case "https://www.vogue.co.jp/fashion/news":
//			siteCrawlRule = SiteCrawlRuleByFile{
//				Url:                 siteUrl,
//				ArticleSelector:     ".summary-item",
//				TitleSelector:       ".summary-item__content h2",
//				LinkSelector:        ".summary-item__content > a",
//				DescriptionSelector: "",
//				HasDataToList:       true,
//				DateSelector:        ".summary-item__publish-date",
//				DateLayout:          "2006年1月2日",
//				IsTimeHumanize:      false,
//				IsSpa:               false,
//			}
//
//		case "https://www.elle.com/jp/fashion-news/":
//			siteCrawlRule = SiteCrawlRuleByFile{
//				Url:                 siteUrl,
//				ArticleSelector:     ".custom-item",
//				TitleSelector:       ".custom-item-title",
//				LinkSelector:        ".custom-item-title",
//				DescriptionSelector: "",
//				HasDataToList:       false,
//				DateSelector:        ".content-info-date",
//				DateLayout:          "2006/01/02",
//				IsTimeHumanize:      false,
//				IsSpa:               false,
//			}
//
//		case "https://lee.hpplus.jp/category/fashion/":
//			siteCrawlRule = SiteCrawlRuleByFile{
//				Url:                 siteUrl,
//				ArticleSelector:     "main > article",
//				TitleSelector:       ".entry-title",
//				LinkSelector:        " > a",
//				DescriptionSelector: "",
//				HasDataToList:       true,
//				DateSelector:        ".post-date time",
//				DateLayout:          "2006/01/02",
//				IsTimeHumanize:      false,
//				IsSpa:               false,
//			}
//
//			// Chromeが必要
//		case "https://corp.zozo.com/news-top/":
//			siteCrawlRule = SiteCrawlRuleByFile{
//				Url:                 siteUrl,
//				ArticleSelector:     ".news-list .list-item",
//				TitleSelector:       ".item-ttl .ttl-inner",
//				LinkSelector:        " > a",
//				DescriptionSelector: "",
//				HasDataToList:       true,
//				DateSelector:        ".item-date",
//				DateLayout:          "2006年01月02日",
//				IsTimeHumanize:      false,
//				IsSpa:               true,
//			}
//
//		default:
//
//			return siteCrawlRule, fmt.Errorf("no crawl rule")
//		}
//		return siteCrawlRule, nil
//	}

func getChromeDevToolsWebSocketDebuggerUrl() (string, error) {

	type VersionResponse struct {
		Browser              string `json:"Browser"`
		ProtocolVersion      string `json:"Protocol-Version"`
		UserAgent            string `json:"User-Agent"`
		V8Version            string `json:"V8-Version"`
		WebKitVersion        string `json:"WebKit-Version"`
		WebSocketDebuggerUrl string `json:"webSocketDebuggerUrl"`
	}

	url := "http://" + config.Config.ChromeDevToolsHostAndPort + "/json/version"

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return "", err
	}
	// fix Host header is specified and is not an IP address or localhost
	req.Host = "localhost"

	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println(string(body))
	// JSONを構造体にエンコード
	var versionResponse VersionResponse
	err = json.Unmarshal(body, &versionResponse)

	if err != nil {
		return "", err
	}

	// fix Host header is specified and is not an IP address or localhost
	versionResponse.WebSocketDebuggerUrl = strings.Replace(versionResponse.WebSocketDebuggerUrl, "localhost", config.Config.ChromeDevToolsHostAndPort, 1)

	return versionResponse.WebSocketDebuggerUrl, nil
}

func GetRSSByHTMLUseChrome(siteUrl string, siteCrawlRule *ent.SiteCrawlRule, client *ent.Client) (string, error) {
	var err error
	var contents string

	loc, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(loc)

	feed := &feeds.Feed{}
	feed.Link = &feeds.Link{Href: siteUrl}
	feed.Created = now

	var chromedpContext context.Context
	if config.Config.ChromeDevToolsHostAndPort == "" {
		chromedpContext = context.Background()
	} else {
		webSocketDebuggerUrl, err := getChromeDevToolsWebSocketDebuggerUrl()
		if err != nil {
			log.Println(err)
			return "", err
		}
		//devtoolsWsURL := flag.String("devtools-ws-url", webSocketDebuggerUrl, "DevTools WebSocket URL")
		//flag.Parse()
		//fmt.Printf("webSocketDebuggerUrl: %v", webSocketDebuggerUrl)
		//fmt.Printf("devtoolsWsURL: %v", devtoolsWsURL)

		// create allocator context for use with creating a browser context later
		allocatorContext, cancel := chromedp.NewRemoteAllocator(context.Background(), webSocketDebuggerUrl)
		defer cancel()
		chromedpContext = allocatorContext
	}

	ctx, cancel := chromedp.NewContext(
		chromedpContext,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var hasDescription bool

	err = chromedp.Run(ctx,
		chromedp.Navigate(siteUrl),
		chromedp.Title(&feed.Title),
		chromedp.ScrollIntoView(siteCrawlRule.ArticleSelector, chromedp.ByQuery),
		chromedp.AttributeValue("meta[name=\"description\"]", "content", &feed.Description, &hasDescription, chromedp.ByQuery),
		chromedp.OuterHTML(`html`, &contents, chromedp.ByQuery),
	)

	if err != nil {
		log.Println(err)
		return "", err
	}

	stringReader := strings.NewReader(contents)
	doc, err := goquery.NewDocumentFromReader(stringReader)

	articles := doc.Find(siteCrawlRule.ArticleSelector)

	articles.Each(func(i int, dom *goquery.Selection) {

		layout := siteCrawlRule.DateLayout

		rootSelector := siteCrawlRule.ArticleSelector

		titleSelector := utils.CreateSelectorOnChildrenScopeFeatureSupport(
			siteCrawlRule.TitleSelector,
			rootSelector,
		)

		title := dom.Find(titleSelector).Text()
		title = strings.TrimSpace(title)

		linkSelector := utils.CreateSelectorOnChildrenScopeFeatureSupport(
			siteCrawlRule.LinkSelector,
			rootSelector,
		)

		link, _ := dom.Find(linkSelector).Attr("href")
		if link != "" {
			base, err := url.Parse(siteUrl)
			if err != nil {
				return
			}
			ref, err := url.Parse(link)
			if err != nil {
				return
			}
			link = base.ResolveReference(ref).String()
		}

		var description string
		if siteCrawlRule.DescriptionSelector != "" {

			descriptionSelector := utils.CreateSelectorOnChildrenScopeFeatureSupport(
				siteCrawlRule.DescriptionSelector,
				rootSelector,
			)

			description = dom.Find(descriptionSelector).Text()
			description = strings.TrimSpace(description)
		}

		var date time.Time
		if siteCrawlRule.HasDataToList {
			dateSelector := utils.CreateSelectorOnChildrenScopeFeatureSupport(
				siteCrawlRule.DateSelector,
				rootSelector,
			)

			dateStr := dom.Find(dateSelector).Text()
			dateStr = strings.TrimSpace(dateStr)
			date, err = time.ParseInLocation(layout, dateStr, loc)
			if err != nil {
				if siteCrawlRule.IsTimeHumanize {
					date, err = utils.HumanizeParseTime(dateStr, now)
				}

				if err != nil {
					log.Println(err)
					return
				}

			}
		} else {
			existArticle, err := client.Article.
				Query().
				Where(article.URL(link)).
				Only(context.Background())

			if err != nil && !ent.IsNotFound(err) {
				log.Println(err)
				return
			}

			if existArticle == nil {
				var pageContent string

				time.Sleep(1 * time.Second)

				err = chromedp.Run(ctx,
					chromedp.Navigate(link),
					chromedp.ScrollIntoView(siteCrawlRule.DateSelector, chromedp.ByQuery),
					chromedp.OuterHTML(`html`, &pageContent, chromedp.ByQuery),
				)

				stringReader := strings.NewReader(pageContent)
				page, err := goquery.NewDocumentFromReader(stringReader)

				dateStr := page.Find(siteCrawlRule.DateSelector).Text()
				dateStr = strings.TrimSpace(dateStr)
				date, err = time.ParseInLocation(layout, dateStr, loc)
				if err != nil {
					if siteCrawlRule.IsTimeHumanize {
						date, err = utils.HumanizeParseTime(dateStr, now)
					}

					if err != nil {
						log.Println(err)
						return
					}

				}

			} else {
				date = existArticle.PublishedAt
			}
		}

		item := &feeds.Item{
			Title:       title,
			Link:        &feeds.Link{Href: link},
			Description: description,
			Created:     date,
		}
		feed.Add(item)

	})
	if contents == "" {
		return "", errors.New("contents is empty")
	}

	rss, err := feed.ToRss()
	if err != nil {
		return "", err
	}
	contents = rss
	return contents, nil
}

func GetRSSByHTMLUseColly(siteUrl string, siteCrawlRule *ent.SiteCrawlRule, client *ent.Client) (string, error) {
	var err error
	var contents string
	loc, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(loc)

	feed := &feeds.Feed{}
	feed.Link = &feeds.Link{Href: siteUrl}
	feed.Created = now

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
		title = strings.TrimSpace(title)

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
			description = strings.TrimSpace(description)
		}

		var date time.Time
		if siteCrawlRule.HasDataToList {
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
					log.Println(err)
					return
				}

			}
		} else {
			existArticle, err := client.Article.
				Query().
				Where(article.URL(url)).
				Only(context.Background())

			if err != nil && !ent.IsNotFound(err) {
				log.Println(err)
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
							log.Println(err, dateStr)
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

	s.Visit(siteUrl)

	if contents == "" {
		return "", errors.New("contents is empty")
	}

	rss, err := feed.ToRss()
	if err != nil {
		return "", err
	}
	contents = rss
	return contents, nil
}

func GetRSSByHTML(siteUrl string, siteCrawlRule *ent.SiteCrawlRule, client *ent.Client) (string, error) {
	var err error
	var contents string

	if siteCrawlRule.IsSpa {
		contents, err = GetRSSByHTMLUseChrome(siteUrl, siteCrawlRule, client)
	} else {
		contents, err = GetRSSByHTMLUseColly(siteUrl, siteCrawlRule, client)
	}

	if err != nil {
		return "", err
	}

	return contents, nil
}

func CrawlSite(site *ent.Site, client *ent.Client) (*ent.Feed, error) {

	var contents = ""
	var err error

	if site.FeedURL != "" {
		contents, err = GetRSS(site.FeedURL)
		if err != nil || contents == "" {
			return nil, err
		}
	} else {
		siteCrawlRule, err := site.QuerySiteCrawlRule().Only(context.Background())
		if err != nil && !ent.IsNotFound(err) {
			return nil, err
		}

		if siteCrawlRule == nil {
			return nil, err
		}

		contents, err = GetRSSByHTML(site.URL, siteCrawlRule, client)
		if err != nil || contents == "" {
			return nil, err
		}
	}

	fp := gofeed.NewParser()
	feed, err := fp.ParseString(contents)

	if err != nil {
		return nil, err
	}

	respFeed, err := client.Feed.
		Create().
		SetContents(contents).
		SetCount(len(feed.Items)).
		SetSite(site).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return respFeed, nil
}

func CrawlAllSite(client *ent.Client) ([]*ent.Site, error) {
	sites, err := client.Site.
		Query().
		Where(site.Active(true)).
		Order(ent.Desc(site.FieldUpdatedAt)).
		All(context.Background())

	if err != nil {
		return nil, err
	}

	for _, site := range sites {
		feed, err := CrawlSite(site, client)
		if err != nil {
			log.Println(err)
			// エラーが発生しても続行する
		} else {
			fmt.Println(site.URL, feed.Count)
		}
	}

	return sites, nil
}
