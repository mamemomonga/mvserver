package runners

import (
	"fmt"
)

func (t *Runners) runButtonPushed() {
	for {
		t.hw.BtnInt.WaitForEdge(-1)
		port := t.hw.Ex1.B

		port.Fetch()
		buf := ""
		for i := uint8(0); i < 8; i++ {
			r := "FALSE"
			if port.Get(i) {
				r = "TRUE"
			}
			buf = buf + fmt.Sprintf("[%d:%s] ", i, r)
		}
		fmt.Printf("BTN: %s\n", buf)
	}
}
