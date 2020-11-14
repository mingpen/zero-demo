/*
 * @Author: ming
 * @LastEditors: ming
 * @Date: 2020-11-14 19:40:50
 * @LastEditTime: 2020-11-14 20:10:29
 * @FilePath: /zero-demo/rpc/add/internal/logic/addlogic.go
 * @Description: Description
 */
package logic

import (
	"context"

	"bookstore/rpc/add/add"
	"bookstore/rpc/add/internal/svc"
	"bookstore/rpc/model"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *add.AddReq) (*add.AddResp, error) {
	// 手动代码开始
	_, err := l.svcCtx.Model.Insert(model.Book{
		Book:  in.Book,
		Price: in.Price,
	})
	if err != nil {
		return nil, err
	}

	return &add.AddResp{
		Ok: true,
	}, nil
	// 手动代码结束
}
