package runners

import (
	"log"
	"time"
)

func (t *Runners) runButtonPushed() {

	bts := [8]bool{false,false,false,false,false,false,false,false}
	fnc := [8]func(){ t.btnP1, t.btnP2, t.btnP3, t.btnP4, t.btnP5, t.btnP6, t.btnP7, t.btnP8 }

	for {
		t.hw.BtnInt.WaitForEdge(-1)
		select {
		case <-t.ctx.Done():
			return
		default:
		}

		port := t.hw.Ex1.B
		port.Fetch()

		for i := byte(0); i < 8; i++ {
			if !port.Get(i) {
				bts[i]=true
			}
		}
		for i := byte(0); i < 8; i++ {
			if bts[i] {
				t.hw.LedProc.Set(true).Apply()
				fnc[i]()
				time.Sleep(250 * time.Millisecond)
				t.hw.LedProc.Set(false).Apply()
				bts[i] = false
			}
		}
	}
}

func (t *Runners) btnP1() {
	log.Printf("BTN P1")
}
func (t *Runners) btnP2() {
	log.Printf("BTN P2")
}
func (t *Runners) btnP3() {
	log.Printf("BTN P3")
}
func (t *Runners) btnP4() {
	log.Printf("BTN P4")
}
func (t *Runners) btnP5() {
	log.Printf("BTN P5")
}
func (t *Runners) btnP6() {
	log.Printf("BTN P6")
}
func (t *Runners) btnP7() {
	log.Printf("BTN P7")
}
func (t *Runners) btnP8() {
	log.Printf("BTN P8")
}

