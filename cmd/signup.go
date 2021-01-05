package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"
	"toolset/internal/faker"

	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
)

var desc = strings.Join([]string{
	"因为开源，无法提供自动发送邮件功能，有需求的自己写入授权码即可实现",
	"该子命令是每日自动提交智慧沈航打卡，主要功能如下",
	"√√√.每日定点全自动打卡，不需要任何人为操作，拒绝漏签，告别被骂",
	"√√√.自动将打卡结果发送至指定邮箱，万一哪天意外发生也能第一时间知道，进行补签，安逸轻松(公开版本已经被阉割)",
	"√√√.所有参数运行时设置，不会泄露个人隐私，不用担心使用其他工具打卡导致的隐私泄露问题",
	"√√√.等待添加......",
	"---- 所需信息：----",
	"姓名，电话号，学院，所在省份，所在城市，智慧深航账号（学号），密码",
}, "\n")

var punchCount uint = 1

func init() {
	//signupCmd.Flags().StringVarP(&stuID, "stuID", "s", "", "学生学号")
	//signupCmd.Flags().StringVarP(&stuName, "stuName", "n", "", "学生姓名")
	//signupCmd.Flags().StringVarP(&phoneNumber, "phoneNumber", "p", "", "学生手机号")
	//signupCmd.Flags().StringVarP(&stuProvince, "stuProvince", "r", "", "所在省份")
	//signupCmd.Flags().StringVarP(&stuCity, "stuCity", "c", "", "所在城市")
	//signupCmd.Flags().StringVarP(&stuCollege, "stuCollege", "o", "", "学院")
	//signupCmd.Flags().StringVarP(&email, "email", "e", "", "接收结果的email")
	//signupCmd.Flags().StringVarP(&account, "account", "a", "", "账号")
	//signupCmd.Flags().StringVarP(&passwd, "passwd", "w", "", "密码")
}

var signupCmd = &cobra.Command{
	Use:   "faker",
	Short: "智慧沈航健康打卡",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		c := cron.New()
		f := faker.NewFaker()

		// 设置打卡两次，防止系统故障
		_, err := c.AddFunc("* * * * *", func() {
			fmt.Printf("%s开始第%d次打卡:\n", time.Now(), punchCount)

			// TODO:目前先每次载入所有数据，后续改为文件监听
			f = faker.NewFaker()

			// 执行打卡逻辑
			f.Do()

			fmt.Printf("%s: 第%d次打卡完毕.\n", time.Now(), punchCount)
			punchCount++
		})

		if err != nil {
			log.Fatalf("Add job err: %v", err)
			return
		}

		//启动定时任务
		c.Start()

		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit
	},
}
