package main

import (
	"time"

	"golang.org/x/sys/windows/svc"
)

type myService struct{}

// Execute メソッドにサービスのメインロジックを記述します
func (m *myService) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	// サービスに「開始中」を通知
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown
	changes <- svc.Status{State: svc.StartPending}

	// 任意の初期化処理...

	// サービスに「実行中」を通知
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}

	// メインループ
	for {
		select {
		case c := <-r:
			switch c.Cmd {
			case svc.Interrogate:
				changes <- c.CurrentStatus
			case svc.Stop, svc.Shutdown:
				// 停止シグナルを受け取ったらループを抜ける
				changes <- svc.Status{State: svc.StopPending}
				return
			}
		case <-time.After(1 * time.Second):
			// ここに定期実行したいバックグラウンド処理を書く
		}
	}
}

func main() {
	// 通常のコマンドライン実行か、サービスとしての実行かを判別
	isInt, err := svc.IsAnInteractiveSession()
	if err != nil {
		return
	}

	if isInt {
		// 通常実行時の処理（デバッグ用など）
		return
	}

	// サービスとして実行
	svc.Run("MyGoService", &myService{})
}
