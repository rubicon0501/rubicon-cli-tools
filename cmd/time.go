package cmd

import (
	"log"
	"rubicon-cli-tools/internal/timer"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	calculateTime string
	duration      string
	timeCmd       = &cobra.Command{
		Use:   "time",
		Short: "时间格式处理",
		Long:  "时间格式处理",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	nowTimeCmd = &cobra.Command{
		Use:   "now",
		Short: "获取当前时间",
		Long:  "获取当前时间",
		Run: func(cmd *cobra.Command, args []string) {
			nowTime := timer.GetNowTime()
			log.Printf("当前时间: %s,	%d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
		},
	}
	calculateTimeCmd = &cobra.Command{
		Use:   "calc",
		Short: "计算所需时间",
		Long:  "计算所需时间",
		Run: func(cmd *cobra.Command, args []string) {
			var curTimer time.Time
			layout := "2006-01-02 15:04:05"
			if calculateTime == "" {
				curTimer = timer.GetNowTime()
			} else {
				var err error
				space := strings.Count(calculateTime, " ")
				if space == 0 {
					layout = "2006-01-02"
				}
				curTimer, err = time.ParseInLocation(layout, calculateTime, time.Local)
				if err != nil {
					t, _ := strconv.Atoi(calculateTime)
					curTimer = time.Unix(int64(t), 0)
				}
			}
			t, err := timer.GetCalculateTime(curTimer, duration)
			if err != nil {
				log.Fatalf("GetCalculateTime err: %v", err)
			}
			log.Printf("所需时间: %s,	%d", t.Format(layout), t.Unix())
		},
	}
)

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "",
		"需要计算的时间，有效单位为时间戳或者已格式化的时间字符串")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "",
		`间隔时间，有效单位为"ns","us","ms","s","m","h"`)
}
