/*
 * @Author: ming
 * @LastEditors: ming
 * @Date: 2020-11-14 17:47:27
 * @LastEditTime: 2020-11-14 19:56:25
 * @FilePath: /zero-demo/api/internal/config/config.go
 * @Description: Description
 */
package config

import (
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Add   zrpc.RpcClientConf // 手动代码
	Check zrpc.RpcClientConf // 手动代码
}
