package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/haikara-dev/haikara/config"
	"github.com/haikara-dev/haikara/ent"
	"github.com/haikara-dev/haikara/ent/user"
	"math"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Client *ent.Client
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
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

	users, err := h.Client.User.
		Query().
		Offset(offset).
		Limit(pageSize).
		All(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	totalCount, err := h.Client.User.
		Query().
		Count(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	type ResponseUser struct {
		ID    int       `json:"id"`
		Email string    `json:"email"`
		Role  user.Role `json:"role"`
	}
	type ResponseJson struct {
		TotalCount int            `json:"totalCount"`
		TotalPage  int            `json:"totalPage"`
		PageSize   int            `json:"pageSize"`
		Data       []ResponseUser `json:"data"`
	}

	var resUsers = make([]ResponseUser, 0)

	for _, user := range users {
		resUsers = append(resUsers, ResponseUser{
			ID:    user.ID,
			Email: user.Email,
			Role:  user.Role,
		})
	}
	c.JSON(http.StatusOK, ResponseJson{
		TotalCount: totalCount,
		TotalPage:  int(math.Ceil(float64(totalCount) / float64(pageSize))),
		PageSize:   pageSize,
		Data:       resUsers,
	})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resUser, err := h.Client.User.
		Get(context.Background(), id)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if resUser == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, resUser)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	// firebase で作成して、　UUIDを取得して、　UUIDでDBに登録する
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var reqUser ent.User
	err = ctx.ShouldBindJSON(&reqUser)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resUser, err := h.Client.User.
		UpdateOneID(id).
		SetUUID(reqUser.UUID).
		Save(context.Background())

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, &resUser)
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.Client.User.
		DeleteOneID(id).
		Exec(context.Background())

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (h *UserHandler) CreateUserByUUIDAndEmail(c *gin.Context) {
	var reqUser ent.User
	err := c.ShouldBindJSON(&reqUser)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if reqUser.UUID == "" || reqUser.UUID != c.MustGet("UUID").(string) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if reqUser.Email == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	existUser, err := h.Client.User.
		Query().
		Where(user.UUID(reqUser.UUID)).
		Only(context.Background())

	if err != nil && !ent.IsNotFound(err) {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if existUser != nil {
		c.JSON(http.StatusOK, &existUser)
		return
	}

	newUser, err := h.Client.User.
		Create().
		SetUUID(reqUser.UUID).
		SetEmail(reqUser.Email).
		Save(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &newUser)
}

func (h *UserHandler) GetCurrentUser(c *gin.Context) {

	UUID := c.MustGet("UUID").(string)

	resUser, err := h.Client.User.
		Query().
		Where(user.UUID(UUID)).
		Only(context.Background())

	if err != nil && !ent.IsNotFound(err) {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if resUser == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, resUser)
}

func (h *UserHandler) UpdateUserRole(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var reqUser ent.User
	err = ctx.ShouldBindJSON(&reqUser)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resUser, err := h.Client.User.
		UpdateOneID(id).
		SetRole(reqUser.Role).
		Save(context.Background())

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, &resUser)
}
