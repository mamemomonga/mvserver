package mcp23017

import (
	"log"
)

const (
	I2CAddr int = 0x20
	Debug bool = true
)

type MCP23017 struct {
	dev bus
}

type bus interface {
	ReadReg(byte, []byte) error
	WriteReg(byte, []byte) error
}

func New(dev bus) *MCP23017 {
	t := new(MCP23017)
	t.dev=dev
	return t
}

func (t *MCP23017) Write(addr uint8, data uint8) {
	err := t.dev.WriteReg(addr, []byte{data})
	if err != nil {
		log.Fatal(err)
	}
}

func (t *MCP23017) Read(addr uint8) byte {
	v := make([]byte, 1,1)
	err := t.dev.ReadReg(addr,v)
	if err != nil {
		log.Fatal(err)
	}
	return v[0]
}

// Slice to Bit vector
func (t *MCP23017) S2BV(in []uint8) uint8 {
	var v uint8
	var i uint8
	for i = 0; i < 8; i++ {
		v = v | (in[i] << i)
	}
	if Debug {
		log.Printf("BV: 0x%02X\n", v)
	}
	return v
}

// Bit vector to Slice
func (t *MCP23017) BV2S(in uint8) []uint8 {
	v := t.AllLow()
	var i uint8
	for i = 0; i < 8; i++ {
		if (in & (1 << i)) > 0 {
			v[i] = 1
		} else {
			v[i] = 0
		}
	}
	return v
}

func (t *MCP23017) AllHigh() []uint8 {
	return []uint8{1, 1, 1, 1, 1, 1, 1, 1}
}

func (t *MCP23017) AllLow() []uint8 {
	return []uint8{0, 0, 0, 0, 0, 0, 0, 0}
}

func (t *MCP23017) Byte(b0, b1, b2, b3, b4, b5, b6, b7 uint8) []uint8 {
	return []uint8{b0, b1, b2, b3, b4, b5, b6, b7}
}

// ICON.BANK=0 専用

func (t *MCP23017) DirectionA(v ...uint8) {
	t.Write(0x00, t.S2BV(v)) // IODIRA
}
func (t *MCP23017) DirectionB(v ...uint8) {
	t.Write(0x01, t.S2BV(v)) // IODIRB
}
func (t *MCP23017) LatchA(v []uint8) {
	t.Write(0x14, t.S2BV(v)) // OLATA
}
func (t *MCP23017) LatchB(v []uint8) {
	t.Write(0x15, t.S2BV(v)) // OLATB
}
func (t *MCP23017) PullUpA(v []uint8) {
	t.Write(0x0c, t.S2BV(v)) // GPPUA
}
func (t *MCP23017) PullUpB(v []uint8) {
	t.Write(0x0d, t.S2BV(v)) // GPPUB
}
func (t *MCP23017) GpioA() (v []uint8) {
	return t.BV2S(t.Read(0x12)) // GPIOA
}
func (t *MCP23017) GpioB() (v []uint8) {
	return t.BV2S(t.Read(0x13)) // GPIOB
}

