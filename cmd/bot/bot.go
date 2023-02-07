package main

import "github.com/gccio/wechat-bot/cmd/bot/app"

func main() {
	cmd := app.NewWechatBotCommand()
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
