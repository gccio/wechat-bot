package bot

import (
	"github.com/eatmoreapple/openwechat"
	"github.com/gccio/wechat-bot/pkg/plugin"
)

type WechatBot struct {
	bot *openwechat.Bot
}

func NewWeChatBot() *WechatBot {
	return &WechatBot{
		bot: openwechat.DefaultBot(openwechat.Desktop),
	}
}

func (wb *WechatBot) InitAndLogin() error {
	plugins := plugin.Plugins()
	wb.bot.MessageHandler = func(msg *openwechat.Message) {
		for _, handler := range plugins {
			if !handler.Match(msg) {
				return
			}
			handler.Handle(msg)
		}
	}

	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()
	wb.bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	if err := wb.bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
		return err
	}

	return nil
}

func (wb *WechatBot) Run() error {
	return wb.bot.Block()
}
