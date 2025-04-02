package controllers

import (
	"admin-v1/app/models/dao"
	"admin-v1/app/models/requests"
	"admin-v1/app/models/responses"

	"github.com/gin-gonic/gin"
)

func Filter[T any](req *requests.Filter, res *responses.Filter[T], c *gin.Context, tableName string) error {
	//kiem tra gia tri dau vao
	if err := c.ShouldBindQuery(&req); err != nil {
		return err
	}

	//xu li filter tang database
	if err := dao.FilterExec(req, res, tableName); err != nil {
		return err
	}

	return nil
}