/*
 * @Author: ming
 * @LastEditors: ming
 * @Date: 2020-11-14 19:40:50
 * @LastEditTime: 2020-11-14 20:06:22
 * @FilePath: /zero-demo/rpc/add/internal/config/config.go
 * @Description: Description
 */
package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource string          // 手动代码
	Table      string          // 手动代码
	Cache      cache.CacheConf // 手动代码
}
