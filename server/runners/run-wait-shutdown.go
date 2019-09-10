package runners

import (
	"log"
)

func (t *Runners) runWaitShutdown() {
	t.hw.Shutdown.WaitForEdge(-1)
	t.Stop()
	log.Println("[SHUTDOWN]")
}
