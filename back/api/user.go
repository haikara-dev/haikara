package api

import (
	"context"
	"github.com/cubdesign/dailyfj/ent"
	"github.com/cubdesign/dailyfj/ent/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Client *ent.Client
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.Client.User.
		Query().
		All(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, users)
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
