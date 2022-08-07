package main

import (
	_ "config-deliver/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"config-deliver/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
