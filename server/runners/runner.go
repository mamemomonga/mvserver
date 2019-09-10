// 停止調停が必要なゴルーチン
package runners

import ()

type runner struct {
	Done chan bool
	p    *Runners
	name string
}

func NewRunner(name string, p *Runners) *runner {
	t := new(runner)
	t.Done = make(chan bool)
	t.p = p
	t.name = name
	return t
}

func (t *runner) Run() {
	switch t.name {
	case "led":    t.led()
	case "bme280": t.bme280()
	case "volumio-state": t.volumioState()
	}
	t.Done <- true
}
