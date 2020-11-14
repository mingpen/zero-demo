/*
 * @Author: ming
 * @LastEditors: ming
 * @Date: 2020-11-14 19:53:08
 * @LastEditTime: 2020-11-14 20:11:28
 * @FilePath: /zero-demo/rpc/check/internal/logic/checklogic.go
 * @Description: Description
 */
package logic

import (
	"context"

	"bookstore/rpc/check/check"
	"bookstore/rpc/check/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *check.CheckReq) (*check.CheckResp, error) {
	// 手动代码开始
	resp, err := l.svcCtx.Model.FindOne(in.Book)
	if err != nil {
		return nil, err
	}

	return &check.CheckResp{
		Found: true,
		Price: resp.Price,
	}, nil
	// 手动代码结束
}
