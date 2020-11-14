/*
 * @Author: ming
 * @LastEditors: ming
 * @Date: 2020-11-14 19:40:50
 * @LastEditTime: 2020-11-14 20:07:19
 * @FilePath: /zero-demo/rpc/add/internal/svc/servicecontext.go
 * @Description: Description
 */
package svc

import (
	"bookstore/rpc/add/internal/config"
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
