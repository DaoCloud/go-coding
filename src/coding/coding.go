package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {

	// 命令行app开发的工具包
	// https://github.com/codegangsta/cli

	app := cli.NewApp()
	app.Name = "Coding"
	app.Usage = "make coding a lot easier"
	//app.Action = func(c *cli.Context) {
	//	println("Hi, welcome!")
	//}

	app.Run(os.Args)
}
