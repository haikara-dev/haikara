// Code generated by ent, DO NOT EDIT.

package site

import (
	"time"
)

const (
	// Label holds the string label denoting the site type in the database.
	Label = "site"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldFeedURL holds the string denoting the feed_url field in the database.
	FieldFeedURL = "feed_url"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldCannotCrawlAt holds the string denoting the cannot_crawl_at field in the database.
	FieldCannotCrawlAt = "cannot_crawl_at"
	// EdgeArticles holds the string denoting the articles edge name in mutations.
	EdgeArticles = "articles"
	// EdgeFeeds holds the string denoting the feeds edge name in mutations.
	EdgeFeeds = "feeds"
	// EdgeSiteCrawlRule holds the string denoting the site_crawl_rule edge name in mutations.
	EdgeSiteCrawlRule = "site_crawl_rule"
	// EdgeSiteCategories holds the string denoting the site_categories edge name in mutations.
	EdgeSiteCategories = "site_categories"
	// Table holds the table name of the site in the database.
	Table = "sites"
	// ArticlesTable is the table that holds the articles relation/edge.
	ArticlesTable = "articles"
	// ArticlesInverseTable is the table name for the Article entity.
	// It exists in this package in order to avoid circular dependency with the "article" package.
	ArticlesInverseTable = "articles"
	// ArticlesColumn is the table column denoting the articles relation/edge.
	ArticlesColumn = "site_articles"
	// FeedsTable is the table that holds the feeds relation/edge.
	FeedsTable = "feeds"
	// FeedsInverseTable is the table name for the Feed entity.
	// It exists in this package in order to avoid circular dependency with the "feed" package.
	FeedsInverseTable = "feeds"
	// FeedsColumn is the table column denoting the feeds relation/edge.
	FeedsColumn = "site_feeds"
	// SiteCrawlRuleTable is the table that holds the site_crawl_rule relation/edge.
	SiteCrawlRuleTable = "site_crawl_rules"
	// SiteCrawlRuleInverseTable is the table name for the SiteCrawlRule entity.
	// It exists in this package in order to avoid circular dependency with the "sitecrawlrule" package.
	SiteCrawlRuleInverseTable = "site_crawl_rules"
	// SiteCrawlRuleColumn is the table column denoting the site_crawl_rule relation/edge.
	SiteCrawlRuleColumn = "site_site_crawl_rule"
	// SiteCategoriesTable is the table that holds the site_categories relation/edge. The primary key declared below.
	SiteCategoriesTable = "site_category_sites"
	// SiteCategoriesInverseTable is the table name for the SiteCategory entity.
	// It exists in this package in order to avoid circular dependency with the "sitecategory" package.
	SiteCategoriesInverseTable = "site_categories"
)

// Columns holds all SQL columns for site fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldURL,
	FieldFeedURL,
	FieldActive,
	FieldCannotCrawlAt,
}

var (
	// SiteCategoriesPrimaryKey and SiteCategoriesColumn2 are the table columns denoting the
	// primary key for the site_categories relation (M2M).
	SiteCategoriesPrimaryKey = []string{"site_category_id", "site_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// URLValidator is a validator for the "url" field. It is called by the builders before save.
	URLValidator func(string) error
	// DefaultActive holds the default value on creation for the "active" field.
	DefaultActive bool
)
