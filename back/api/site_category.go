package api

import (
	"context"
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

	categories, err := h.Client.SiteCategory.
		Query().
		Order(ent.Desc(sitecategory.FieldUpdatedAt)).
		Offset(offset).
		Limit(pageSize).
		All(context.Background())

	if err != nil && !ent.IsNotFound(err) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if categories == nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	totalCount, err := h.Client.SiteCategory.Query().Count(context.Background())

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	type ResponseJson struct {
		TotalCount int                 `json:"totalCount"`
		TotalPage  int                 `json:"totalPage"`
		PageSize   int                 `json:"pageSize"`
		Data       []*ent.SiteCategory `json:"data"`
	}

	c.JSON(http.StatusOK, ResponseJson{
		TotalCount: totalCount,
		TotalPage:  int(math.Ceil(float64(totalCount) / float64(pageSize))),
		PageSize:   pageSize,
		Data:       categories,
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
