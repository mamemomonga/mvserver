package runners

import (
	"context"
	"github.com/mamemomonga/rpi-volumio-status-led/server/hardwares"
	"log"
)

type Runners struct {
	hw          *hardwares.Hardwares
	ctx         context.Context
	cancel      context.CancelFunc
	RunnersDone chan bool
	runner      [2]*runner
}

func New(hw *hardwares.Hardwares) *Runners {
	t := new(Runners)
	t.hw = hw
	t.ctx, t.cancel = context.WithCancel(context.Background())
	t.RunnersDone = make(chan bool)
	return t
}

func (t *Runners) Run() {

	// 停止調停が必要なゴルーチン
	t.runner = [...]*runner{
		NewRunner("led", t),
		NewRunner("bme280", t),
	}
	for i := 0; i < len(t.runner); i++ {
		go t.runner[i].Run()
	}

	// 停止調停が不要なゴルーチン
	go t.runButtonPushed()
	go t.runWaitShutdown()
}

func (t *Runners) Stop() {
	t.cancel()
	log.Println("Wait")
	for i := 0; i < len(t.runner); i++ {
		select {
		case <-t.runner[i].Done:
		}
	}
	log.Println("All Done!")
	t.RunnersDone <- true
}
