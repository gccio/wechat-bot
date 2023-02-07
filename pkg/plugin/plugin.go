package plugin

import (
	"github.com/eatmoreapple/openwechat"
	"k8s.io/klog/v2"
)

type MessageHandler interface {
	Match(*openwechat.Message) bool
	Handle(*openwechat.Message)
}

var plugs = map[string]MessageHandler{}

func Register(name string, handler MessageHandler) {
	plugs[name] = handler
}

func Plugins() []MessageHandler {
	var handlers []MessageHandler
	for name, handler := range plugs {
		// TODO add filter
		handlers = append(handlers, handler)
		klog.Infof("Bot support [%s]", name)
	}
	return handlers
}
