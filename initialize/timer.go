package initialize

import (
	"fmt"

	"github.com/robfig/cron/v3"

	"go-generate/config"
	"go-generate/global"
	"go-generate/utils"
)

func Timer() {
	if global.YAN_CONFIG.Timer.Start {
		for i := range global.YAN_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				var option []cron.Option
				if global.YAN_CONFIG.Timer.WithSeconds {
					option = append(option, cron.WithSeconds())
				}
				_, err := global.YAN_Timer.AddTaskByFunc("ClearDB", global.YAN_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.YAN_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				}, option...)
				if err != nil {
					fmt.Println("add timer error:", err)
				}
			}(global.YAN_CONFIG.Timer.Detail[i])
		}
	}
}
