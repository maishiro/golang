package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/robfig/cron/v3"
)

// func main() {
// 	// 5秒ごとにイベントを発生
// 	ticker := time.NewTicker(time.Second * 5)
// 	for {
// 		select {
// 		case <-ticker.C:  // イベントを受け取る
// 			work()
// 		}
// 	}
// }

func main() {
	// シグナル用のチャネル定義
	quit := make(chan os.Signal)

	// 受け取るシグナルを設定
	signal.Notify(quit, os.Interrupt)

	c := cron.New()
	c.AddFunc("5,20,35,50 * * * *", work) // 15分刻み+5分
	// c.AddFunc("5 * * * *", work)		// 毎時5分
	// c.AddFunc("@every 10s", work)	// 10秒毎
	c.Start()

	<-quit // ここでシグナルを受け取るまで待つ
}

func work() {
	fmt.Println(time.Now().Format("2006-01-02T15:04:05.999Z07:00"))
}
