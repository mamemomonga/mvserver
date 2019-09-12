package hardwares

import (
	"log"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"

	"github.com/mamemomonga/rpi-go-peripherals/bme280"
	"github.com/mamemomonga/rpi-go-peripherals/mcp23017"
)

type Hardwares struct {
	i2c        i2c.BusCloser
	Env        *bme280.BME280
	Ex1        *mcp23017.MCP23017
	BtnInt     gpio.PinIO
	Shutdown   gpio.PinIO

	LedEnable  *mcp23017.PinT
	LedPlaying *mcp23017.PinT
	LedHiRez   *mcp23017.PinT
	LedProc    *mcp23017.PinT
	Led5       *mcp23017.PinT
	Led6       *mcp23017.PinT
	Led7       *mcp23017.PinT
	Led8       *mcp23017.PinT
}

func New() *Hardwares {
	return new(Hardwares)
}

func (t *Hardwares) Init() (err error) {

	log.Println("Hardware Init")

	if _, err := host.Init(); err != nil {
		return err
	}

	t.i2c, err = i2creg.Open("I2C1")
	if err != nil {
		return err
	}

	t.Ex1 = mcp23017.New(t.i2c, mcp23017.AddrAllLow)
	t.Env, err = bme280.New(t.i2c, bme280.AddrLow)
	if err != nil {
		return err
	}

	t.Ex1.A.DirectionAllOutput()
	t.Ex1.A.SetAllHigh().Apply()

	t.Ex1.B.DirectionAllInput()
	t.Ex1.B.PullUpAll()
	t.Ex1.B.InitInterrupt()

	t.BtnInt = gpioreg.ByName("GPIO13")
	if err := t.BtnInt.In(gpio.Float, gpio.FallingEdge); err != nil {
		return err
	}

	t.Shutdown = gpioreg.ByName("GPIO5")
	if err := t.Shutdown.In(gpio.PullUp, gpio.FallingEdge); err != nil {
		return err
	}

	t.LedEnable  = mcp23017.NewPin(t.Ex1.A, 7)
	t.LedPlaying = mcp23017.NewPin(t.Ex1.A, 6)
	t.LedHiRez   = mcp23017.NewPin(t.Ex1.A, 5)
	t.LedProc    = mcp23017.NewPin(t.Ex1.A, 4)
	t.Led5       = mcp23017.NewPin(t.Ex1.A, 3)
	t.Led6       = mcp23017.NewPin(t.Ex1.A, 2)
	t.Led7       = mcp23017.NewPin(t.Ex1.A, 1)
	t.Led8       = mcp23017.NewPin(t.Ex1.A, 0)

	time.Sleep(time.Second)
	t.Ex1.A.SetAllLow().Apply()

	return nil
}

func (t *Hardwares) Close() {
	t.Ex1.A.SetAllLow().Apply()
	t.i2c.Close()
}
