package utils

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

// goquey selectorに、直下のセレクト機能を追加する
func CreateSelectorOnChildrenScopeFeatureSupport(selector string, rootSelector string) string {
	returnSelector := selector
	// 先頭が > で始まる場合は、:scopeを追加する
	reg := regexp.MustCompile(`^[\s]*>[\s]*(.*)`)
	returnSelector = reg.ReplaceAllString(returnSelector, ":scope > $1")
	// :scopeをrootSelectorに置き換える
	// rootSelectorと先頭のselectorが同じ場合は、selectorには直下のセレクタとみなす仕様があるため
	returnSelector = strings.Replace(returnSelector, ":scope", rootSelector, -1)
	return returnSelector
}

func GetArticleTitle(selector string, selection *goquery.Selection) string {
	if selector == "" {
		return ""
	}

	if selection == nil {
		return ""
	}

	titleDom := selection.Find(selector)
	if titleDom.Length() == 0 {
		return ""
	}

	// タイトル　<a>さらに詳しく</a>
	// のようなaタグを削除する
	titleDom.Find("a").Remove()

	title := titleDom.Text()
	title = strings.TrimSpace(title)
	return title
}

func GetArticleLink(selector string, selection *goquery.Selection) string {
	if selection == nil {
		return ""
	}

	var linkDom *goquery.Selection
	if selector == "" {
		// selectionがaタグの場合は、そのまま取得する
		linkDom = selection
	} else {
		linkDom = selection.Find(selector)
		if linkDom.Length() == 0 {
			return ""
		}
	}

	link, _ := linkDom.Attr("href")
	link = strings.TrimSpace(link)
	return link
}

func AddSchemeIfNotExists(url string) string {
	if url == "" {
		return ""
	}

	reg := regexp.MustCompile(`^https?://`)
	if reg.MatchString(url) {
		return url
	}

	reg = regexp.MustCompile(`^//`)
	if reg.MatchString(url) {
		return "https:" + url
	}

	return "https://" + url
}
