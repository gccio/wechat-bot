package plugin

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"k8s.io/klog/v2"
)

var PingRegisterName = "ping"

func init() {
	Register(PingRegisterName, &PingHandler{})
}

type PingHandler struct {
}

func (h PingHandler) Match(message *openwechat.Message) bool {
	return message.IsText() && message.Content == "ping"
}

func (h PingHandler) Handle(msg *openwechat.Message) {
	replymsg := "pong"
	if msg.IsComeFromGroup() {
		user, _ := msg.SenderInGroup()
		replymsg = fmt.Sprintf("@%s pong", user.NickName)
	}
	_, err := msg.ReplyText(replymsg)
	if err != nil {
		klog.Errorf("[%s] reply text failed with error: %s", PingRegisterName, err.Error())
	}
}
