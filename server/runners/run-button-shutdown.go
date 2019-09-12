package runners

import (
	"log"
	"time"
	"os/exec"
)


func (t *Runners) goShutdown() {
	t.Stop()
	log.Println("[SHUTDOWN]")
	err := exec.Command("/usr/bin/sudo","/sbin/poweroff").Run()
	if err != nil {
		log.Printf("Shutdown Failed: %s", err)
		return
	}
	log.Println("Shutdown Start")
}


func (t *Runners) runWaitShutdown() {
	for {
		t.hw.Shutdown.WaitForEdge(-1)
		log.Println("Shutdown button down")

		flag := true
		for i:=0; i<30; i++ {
			if t.hw.Shutdown.Read() {
				flag = false
			}
			time.Sleep( 100 * time.Millisecond )
		}
		if flag {
			t.goShutdown()
			return
		}
	}
}
