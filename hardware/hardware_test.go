package hardware_test

// TestHelloだけキャッシュなしで実行
// go test -v --count=1 -run TestHello

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

func Test01(t *testing.T) {
	hdwr := hardware.New()
	err := hdwr.Init()
	if err != nil {
		t.Fatal(err)
	}

	hdwr.MCP23017.Write(0x00, 0x00) // IOIRA
	hdwr.MCP23017.Write(0x14, 0xFF) // OLATA
	time.Sleep( 1 * time.Second )
	hdwr.MCP23017.Write(0x14, 0x00) // OLATA

	hdwr.MCP23017.Write(0x00, 0x01) // IOIRB

	log.Printf("GPIOB %#v\n",hdwr.MCP23017.Read(0x13) ) // GPIOB

	if err != nil {
		t.Fatal(err)
	}
}

func Test02(t *testing.T) {
	hdwr := hardware.New()
	err := hdwr.Init()
	if err != nil {
		t.Fatal(err)
	}

	hdwr.MCP23017.DirectionA(0,0,0,0,0,0,0,0)
	hdwr.MCP23017.DirectionB(1,1,1,1,1,1,1,1)

	hdwr.MCP23017.LatchA(hdwr.MCP23017.AllHigh())
	time.Sleep( 1 * time.Second )
	hdwr.MCP23017.LatchA(hdwr.MCP23017.AllLow())




	if err != nil {
		t.Fatal(err)
	}
}

