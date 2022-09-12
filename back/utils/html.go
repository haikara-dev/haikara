package utils

import (
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
