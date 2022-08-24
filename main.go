package main

import (
	"fmt"
	"os"
)

func main() {
	var app = &cli.App{
		Name:  "钉钉 通知工具 ",
		Usage: "WebHook 通知工具",
		Commands: []*cli.Command{
			&Cmd,
			&server.Web,
		},
		Version: "0.1",
	}
	if err := app.Run(os.Args); nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
}
