package api

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "swiftiesoft.com/simplebank/db/sqlc"
	"swiftiesoft.com/simplebank/utils"
)

type createAccountReq struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

func (server *Server) CreateAccount(c *gin.Context) {
	var req createAccountReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(errors.New("bad request")))
		return
	}

	account, err := server.store.CreateAccounts(c, db.CreateAccountsParams{
		Owner:    req.Owner,
		Balance:  0,
		Currency: req.Currency,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(account))
}
func (server *Server) GetAccount(c *gin.Context) {
	idString, _ := c.Params.Get("id")
	id, _ := strconv.Atoi(idString)

	account, err := server.store.GetAccounts(c, int64(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, utils.SuccessResponse("not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(account))
}

type getAllReq struct {
	PageNo   int32 `form:"page_no" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

func (server *Server) GetAllAccounts(c *gin.Context) {
	var req getAllReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	accounts, err := server.store.GetAllAccounts(c, db.GetAllAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageNo - 1) * req.PageSize,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, utils.SuccessResponse(accounts))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	count, _ := server.store.GetCountAllAccounts(c)
	c.JSON(http.StatusOK, utils.SuccessResponseWithCount(accounts, count))
}
