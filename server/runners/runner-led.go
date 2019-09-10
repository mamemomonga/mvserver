package runners

import (
	"time"
)

func (t *runner) led() {
	ticker := time.NewTicker(50 * time.Millisecond)
	port := t.p.hw.Ex1.A
	for {
		for i := uint8(0); i < 8; i++ {
			port.Set(i, true).Apply()

			<-ticker.C

			port.Set(i, false)

			select {
			case <-t.p.ctx.Done():
				port.Apply()
				return
			default:
			}
		}
	}
}
