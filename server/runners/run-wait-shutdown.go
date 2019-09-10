package runners

import (
	"log"
	"os/exec"
)

func (t *Runners) runWaitShutdown() {
	t.hw.Shutdown.WaitForEdge(-1)
	t.Stop()
	log.Println("[SHUTDOWN]")
	err := exec.Command("sudo","/sbin/halt").Run()
	if err != nil {
		log.Printf("Shutdown Failed: %s", err)
		return
	}
}
