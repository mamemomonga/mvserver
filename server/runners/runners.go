package runners

import (
	"context"
	"github.com/mamemomonga/rpi-volumio-status-led/server/hardwares"
)

type Runners struct {
	hw          *hardwares.Hardwares
	ctx         context.Context
	cancel      context.CancelFunc
	RunnersDone chan bool
}

func New(hw *hardwares.Hardwares) *Runners {
	t := new(Runners)
	t.hw = hw
	t.ctx, t.cancel = context.WithCancel(context.Background())
	t.RunnersDone = make(chan bool)
	return t
}

func (t *Runners) Run() {
	go t.runBME280()
	go t.runVolumioState()
	go t.runButtonPushed()
	go t.runWaitShutdown()
}

func (t *Runners) Stop() {
	t.cancel()
	t.RunnersDone <- true
}
