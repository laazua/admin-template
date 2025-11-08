package main

import (
	"adtmp/cmd/server/app"
	"adtmp/pkg/xlog"
)

func main() {
	// 设置日志
	xlog.Set()

	app.Start()
}
