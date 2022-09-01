package api

import (
	"net/http"
	db "simple-bank/db/sqlc"
	"time"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username    string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	Fullname string `json:"fullname" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type createUserResponse struct {
	Username    string `json:"username"`
	Fullname string `json:"fullname"`
	Email string `json:"email"`
	CreatedAt time.Time
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return;
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		Password: req.Password,
		Email: req.Email,
		FullName: req.Fullname,
	}

	res, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return;
	}

	user := createUserResponse{
		Username: res.Username,
		Fullname: res.FullName,
		Email: res.Email,
		CreatedAt: res.CreatedAt,
	}

	ctx.JSON(http.StatusCreated, user)
}

type getUserReq struct {
	Username string `uri:"username" binding:"required"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserReq

	err := ctx.ShouldBindUri(&req)

	if err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	res, err := server.store.GetUser(ctx, req.Username)

	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return;
	}

	user := createUserResponse{
		Username: res.Username,
		Fullname: res.FullName,
		Email: res.FullName,
		CreatedAt: res.CreatedAt,
	}

	ctx.JSON(http.StatusOK, user)
}