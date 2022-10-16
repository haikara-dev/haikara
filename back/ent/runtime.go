// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/haikara-dev/haikara/ent/article"
	"github.com/haikara-dev/haikara/ent/feed"
	"github.com/haikara-dev/haikara/ent/image"
	"github.com/haikara-dev/haikara/ent/schema"
	"github.com/haikara-dev/haikara/ent/site"
	"github.com/haikara-dev/haikara/ent/sitecategory"
	"github.com/haikara-dev/haikara/ent/sitecrawlrule"
	"github.com/haikara-dev/haikara/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	articleMixin := schema.Article{}.Mixin()
	articleMixinFields0 := articleMixin[0].Fields()
	_ = articleMixinFields0
	articleFields := schema.Article{}.Fields()
	_ = articleFields
	// articleDescCreatedAt is the schema descriptor for created_at field.
	articleDescCreatedAt := articleMixinFields0[0].Descriptor()
	// article.DefaultCreatedAt holds the default value on creation for the created_at field.
	article.DefaultCreatedAt = articleDescCreatedAt.Default.(func() time.Time)
	// articleDescUpdatedAt is the schema descriptor for updated_at field.
	articleDescUpdatedAt := articleMixinFields0[1].Descriptor()
	// article.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	article.DefaultUpdatedAt = articleDescUpdatedAt.Default.(func() time.Time)
	// article.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	article.UpdateDefaultUpdatedAt = articleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// articleDescTitle is the schema descriptor for title field.
	articleDescTitle := articleFields[0].Descriptor()
	// article.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	article.TitleValidator = articleDescTitle.Validators[0].(func(string) error)
	// articleDescURL is the schema descriptor for url field.
	articleDescURL := articleFields[1].Descriptor()
	// article.URLValidator is a validator for the "url" field. It is called by the builders before save.
	article.URLValidator = articleDescURL.Validators[0].(func(string) error)
	feedMixin := schema.Feed{}.Mixin()
	feedMixinFields0 := feedMixin[0].Fields()
	_ = feedMixinFields0
	feedFields := schema.Feed{}.Fields()
	_ = feedFields
	// feedDescCreatedAt is the schema descriptor for created_at field.
	feedDescCreatedAt := feedMixinFields0[0].Descriptor()
	// feed.DefaultCreatedAt holds the default value on creation for the created_at field.
	feed.DefaultCreatedAt = feedDescCreatedAt.Default.(func() time.Time)
	// feedDescUpdatedAt is the schema descriptor for updated_at field.
	feedDescUpdatedAt := feedMixinFields0[1].Descriptor()
	// feed.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	feed.DefaultUpdatedAt = feedDescUpdatedAt.Default.(func() time.Time)
	// feed.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	feed.UpdateDefaultUpdatedAt = feedDescUpdatedAt.UpdateDefault.(func() time.Time)
	// feedDescContents is the schema descriptor for contents field.
	feedDescContents := feedFields[0].Descriptor()
	// feed.ContentsValidator is a validator for the "contents" field. It is called by the builders before save.
	feed.ContentsValidator = feedDescContents.Validators[0].(func(string) error)
	// feedDescCount is the schema descriptor for count field.
	feedDescCount := feedFields[1].Descriptor()
	// feed.DefaultCount holds the default value on creation for the count field.
	feed.DefaultCount = feedDescCount.Default.(int)
	imageMixin := schema.Image{}.Mixin()
	imageMixinFields0 := imageMixin[0].Fields()
	_ = imageMixinFields0
	imageFields := schema.Image{}.Fields()
	_ = imageFields
	// imageDescCreatedAt is the schema descriptor for created_at field.
	imageDescCreatedAt := imageMixinFields0[0].Descriptor()
	// image.DefaultCreatedAt holds the default value on creation for the created_at field.
	image.DefaultCreatedAt = imageDescCreatedAt.Default.(func() time.Time)
	// imageDescUpdatedAt is the schema descriptor for updated_at field.
	imageDescUpdatedAt := imageMixinFields0[1].Descriptor()
	// image.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	image.DefaultUpdatedAt = imageDescUpdatedAt.Default.(func() time.Time)
	// image.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	image.UpdateDefaultUpdatedAt = imageDescUpdatedAt.UpdateDefault.(func() time.Time)
	// imageDescName is the schema descriptor for name field.
	imageDescName := imageFields[0].Descriptor()
	// image.NameValidator is a validator for the "name" field. It is called by the builders before save.
	image.NameValidator = imageDescName.Validators[0].(func(string) error)
	// imageDescFileName is the schema descriptor for file_name field.
	imageDescFileName := imageFields[1].Descriptor()
	// image.FileNameValidator is a validator for the "file_name" field. It is called by the builders before save.
	image.FileNameValidator = imageDescFileName.Validators[0].(func(string) error)
	// imageDescFilePath is the schema descriptor for file_path field.
	imageDescFilePath := imageFields[2].Descriptor()
	// image.FilePathValidator is a validator for the "file_path" field. It is called by the builders before save.
	image.FilePathValidator = imageDescFilePath.Validators[0].(func(string) error)
	siteMixin := schema.Site{}.Mixin()
	siteMixinFields0 := siteMixin[0].Fields()
	_ = siteMixinFields0
	siteFields := schema.Site{}.Fields()
	_ = siteFields
	// siteDescCreatedAt is the schema descriptor for created_at field.
	siteDescCreatedAt := siteMixinFields0[0].Descriptor()
	// site.DefaultCreatedAt holds the default value on creation for the created_at field.
	site.DefaultCreatedAt = siteDescCreatedAt.Default.(func() time.Time)
	// siteDescUpdatedAt is the schema descriptor for updated_at field.
	siteDescUpdatedAt := siteMixinFields0[1].Descriptor()
	// site.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	site.DefaultUpdatedAt = siteDescUpdatedAt.Default.(func() time.Time)
	// site.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	site.UpdateDefaultUpdatedAt = siteDescUpdatedAt.UpdateDefault.(func() time.Time)
	// siteDescName is the schema descriptor for name field.
	siteDescName := siteFields[0].Descriptor()
	// site.NameValidator is a validator for the "name" field. It is called by the builders before save.
	site.NameValidator = siteDescName.Validators[0].(func(string) error)
	// siteDescURL is the schema descriptor for url field.
	siteDescURL := siteFields[1].Descriptor()
	// site.URLValidator is a validator for the "url" field. It is called by the builders before save.
	site.URLValidator = siteDescURL.Validators[0].(func(string) error)
	// siteDescActive is the schema descriptor for active field.
	siteDescActive := siteFields[3].Descriptor()
	// site.DefaultActive holds the default value on creation for the active field.
	site.DefaultActive = siteDescActive.Default.(bool)
	sitecategoryMixin := schema.SiteCategory{}.Mixin()
	sitecategoryMixinFields0 := sitecategoryMixin[0].Fields()
	_ = sitecategoryMixinFields0
	sitecategoryFields := schema.SiteCategory{}.Fields()
	_ = sitecategoryFields
	// sitecategoryDescCreatedAt is the schema descriptor for created_at field.
	sitecategoryDescCreatedAt := sitecategoryMixinFields0[0].Descriptor()
	// sitecategory.DefaultCreatedAt holds the default value on creation for the created_at field.
	sitecategory.DefaultCreatedAt = sitecategoryDescCreatedAt.Default.(func() time.Time)
	// sitecategoryDescUpdatedAt is the schema descriptor for updated_at field.
	sitecategoryDescUpdatedAt := sitecategoryMixinFields0[1].Descriptor()
	// sitecategory.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sitecategory.DefaultUpdatedAt = sitecategoryDescUpdatedAt.Default.(func() time.Time)
	// sitecategory.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sitecategory.UpdateDefaultUpdatedAt = sitecategoryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sitecategoryDescLabel is the schema descriptor for label field.
	sitecategoryDescLabel := sitecategoryFields[0].Descriptor()
	// sitecategory.LabelValidator is a validator for the "label" field. It is called by the builders before save.
	sitecategory.LabelValidator = sitecategoryDescLabel.Validators[0].(func(string) error)
	sitecrawlruleMixin := schema.SiteCrawlRule{}.Mixin()
	sitecrawlruleMixinFields0 := sitecrawlruleMixin[0].Fields()
	_ = sitecrawlruleMixinFields0
	sitecrawlruleFields := schema.SiteCrawlRule{}.Fields()
	_ = sitecrawlruleFields
	// sitecrawlruleDescCreatedAt is the schema descriptor for created_at field.
	sitecrawlruleDescCreatedAt := sitecrawlruleMixinFields0[0].Descriptor()
	// sitecrawlrule.DefaultCreatedAt holds the default value on creation for the created_at field.
	sitecrawlrule.DefaultCreatedAt = sitecrawlruleDescCreatedAt.Default.(func() time.Time)
	// sitecrawlruleDescUpdatedAt is the schema descriptor for updated_at field.
	sitecrawlruleDescUpdatedAt := sitecrawlruleMixinFields0[1].Descriptor()
	// sitecrawlrule.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sitecrawlrule.DefaultUpdatedAt = sitecrawlruleDescUpdatedAt.Default.(func() time.Time)
	// sitecrawlrule.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sitecrawlrule.UpdateDefaultUpdatedAt = sitecrawlruleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sitecrawlruleDescHasDataToList is the schema descriptor for has_data_to_list field.
	sitecrawlruleDescHasDataToList := sitecrawlruleFields[4].Descriptor()
	// sitecrawlrule.DefaultHasDataToList holds the default value on creation for the has_data_to_list field.
	sitecrawlrule.DefaultHasDataToList = sitecrawlruleDescHasDataToList.Default.(bool)
	// sitecrawlruleDescIsTimeHumanize is the schema descriptor for is_time_humanize field.
	sitecrawlruleDescIsTimeHumanize := sitecrawlruleFields[7].Descriptor()
	// sitecrawlrule.DefaultIsTimeHumanize holds the default value on creation for the is_time_humanize field.
	sitecrawlrule.DefaultIsTimeHumanize = sitecrawlruleDescIsTimeHumanize.Default.(bool)
	// sitecrawlruleDescIsSpa is the schema descriptor for is_spa field.
	sitecrawlruleDescIsSpa := sitecrawlruleFields[8].Descriptor()
	// sitecrawlrule.DefaultIsSpa holds the default value on creation for the is_spa field.
	sitecrawlrule.DefaultIsSpa = sitecrawlruleDescIsSpa.Default.(bool)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescUUID is the schema descriptor for UUID field.
	userDescUUID := userFields[0].Descriptor()
	// user.UUIDValidator is a validator for the "UUID" field. It is called by the builders before save.
	user.UUIDValidator = userDescUUID.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
}
