package hardware_test

// TestHelloだけキャッシュなしで実行
// go test -v --count=1 -run TestMCP23017

import (
	"time"
	"log"
	"testing"
	"github.com/mamemomonga/rpi-volumio-status-led/hardware"
)

func TestInit(t *testing.T) {
	hdwr := hardware.New()
	err := hdwr.Init()
	if err != nil {
		t.Fatal(err)
	}
}

func TestMCP23017(t *testing.T) {

	hdwr := hardware.New()
	err := hdwr.Init()
	if err != nil {
		t.Fatal(err)
	}

	m := hdwr.MCP23017

	m.DirectionA(0,0,0,0,0,0,0,0)
	m.DirectionB(1,1,1,1,1,1,1,1)

	m.StateA = 0xFF
	m.ApplyA()
	time.Sleep( time.Second )
	m.StateA = 0x00
	m.ApplyA()
	time.Sleep( time.Second )

	m.FetchB()
	if m.GetB(0) {
		log.Println("StateB Pin0: High")
	} else {
		log.Println("StateB Pin0: Low")
	}

	for i := byte(0); i<8; i++ {
		m.SetA(i,true, true)
		time.Sleep( 250 * time.Millisecond )
		m.SetA(i,false, true)
	}

}

