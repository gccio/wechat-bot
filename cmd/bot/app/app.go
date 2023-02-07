package app

import (
	"flag"
	"github.com/gccio/wechat-bot/pkg/bot"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

func NewWechatBotCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "Stupid WeChat Bot",
		Long: "Just for fun",
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error

			b := bot.NewWeChatBot()
			if err = b.InitAndLogin(); err != nil {
				return err
			}

			klog.Info("Start wechat robot now")
			return b.Run()
		},
	}
	klog.InitFlags(flag.CommandLine)
	cmd.Flags().AddGoFlagSet(flag.CommandLine)
	flag.Parse()

	return cmd
}
