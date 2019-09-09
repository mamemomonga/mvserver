package hardware

import (
	"github.com/mamemomonga/rpi-volumio-status-led/hardware/mcp23017"
	"golang.org/x/exp/io/i2c"
//	"github.com/quhar/bme280"
//	"log"
)

type Hardware struct {
	MCP23017 *mcp23017.MCP23017
}

func New() *Hardware {
	return new(Hardware)
}

const (
	i2cdev string = "/dev/i2c-1"
)

func (t *Hardware) Init() error {
	d, err := i2c.Open(&i2c.Devfs{Dev: i2cdev}, mcp23017.I2CAddr)
	if err != nil {
		return err
	}
	t.MCP23017 = mcp23017.New(d)
	return nil
}

