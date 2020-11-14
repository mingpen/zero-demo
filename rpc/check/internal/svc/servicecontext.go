/*
 * @Author: ming
 * @LastEditors: ming
 * @Date: 2020-11-14 19:53:08
 * @LastEditTime: 2020-11-14 20:14:10
 * @FilePath: /zero-demo/rpc/check/internal/svc/servicecontext.go
 * @Description: Description
 */
package svc

import (
	"bookstore/rpc/check/internal/config"
	"bookstore/rpc/model"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	c     config.Config
	Model *model.BookModel // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c:     c,
		Model: model.NewBookModel(sqlx.NewMysql(c.DataSource), c.Cache, c.Table), // 手动代码
	}
}
