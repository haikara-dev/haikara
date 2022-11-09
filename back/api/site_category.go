package api

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
	"github.com/haikara-dev/haikara/config"
	"github.com/haikara-dev/haikara/ent"
	"github.com/haikara-dev/haikara/ent/sitecategory"
	"math"
	"net/http"
	"strconv"
)

type SiteCategoryHandler struct {
	Client *ent.Client
}

func (h *SiteCategoryHandler) GetSiteCategories(c *gin.Context) {

	pageStr := c.Query("page")

	if pageStr == "" {
		pageStr = "1"
	}

	page, err := strconv.Atoi(pageStr)

	pageSize := config.Config.PageSize

	offset := (page - 1) * pageSize

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	type ResponseSiteCategory struct {
		ID         int    `json:"id"`
		Label      string `json:"label"`
		SitesCount int    `json:"sites_count"`
	}

	var resSiteCategories []*ResponseSiteCategory

	err = h.Client.SiteCategory.
		Query().
		Order(ent.Desc(sitecategory.FieldUpdatedAt)).
		Offset(offset).
		Limit(pageSize).
		GroupBy(sitecategory.FieldID, sitecategory.FieldLabel).
		Aggregate(func(s *sql.Selector) string {
			t := sql.Table(sitecategory.SitesTable)
			s.LeftJoin(t).On(s.C(sitecategory.FieldID), t.C(sitecategory.SitesPrimaryKey[0]))
			return sql.As(sql.Count(t.C(sitecategory.SitesPrimaryKey[0])), "sites_count")
		}).
		Scan(context.Background(), &resSiteCategories)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if len(resSiteCategories) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	totalCount, err := h.Client.SiteCategory.Query().Count(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	type ResponseJson struct {
		TotalCount int                     `json:"totalCount"`
		TotalPage  int                     `json:"totalPage"`
		PageSize   int                     `json:"pageSize"`
		Data       []*ResponseSiteCategory `json:"data"`
	}

	c.JSON(http.StatusOK, ResponseJson{
		TotalCount: totalCount,
		TotalPage:  int(math.Ceil(float64(totalCount) / float64(pageSize))),
		PageSize:   pageSize,
		Data:       resSiteCategories,
	})
}

func (h *SiteCategoryHandler) CreateSiteCategory(c *gin.Context) {
	var reqCategory ent.SiteCategory
	err := c.ShouldBindJSON(&reqCategory)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	existCategory, err := h.Client.SiteCategory.
		Query().
		Where(sitecategory.LabelEQ(reqCategory.Label)).
		Only(context.Background())

	if err != nil && !ent.IsNotFound(err) {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if existCategory != nil {
		c.AbortWithStatus(http.StatusConflict)
		return
	}

	resCategory, err := h.Client.SiteCategory.
		Create().
		SetLabel(reqCategory.Label).
		Save(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, resCategory)
}

func (h *SiteCategoryHandler) UpdateSiteCategory(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var reqCategory ent.SiteCategory
	err = c.ShouldBindJSON(&reqCategory)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	existCategory, err := h.Client.SiteCategory.
		Get(context.Background(), id)

	if err != nil && !ent.IsNotFound(err) {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if existCategory == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	resCategory, err := existCategory.
		Update().
		SetLabel(reqCategory.Label).
		Save(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &resCategory)
}

func (h *SiteCategoryHandler) DeleteSiteCategory(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.Client.SiteCategory.
		DeleteOneID(id).
		Exec(context.Background())

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
