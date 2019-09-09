package mcp23017

import (
	"log"
)

const (
	I2CAddr int = 0x20
)

type MCP23017 struct {
	dev    bus
	StateA  byte
	StateB  byte
}

type bus interface {
	ReadReg(byte, []byte) error
	WriteReg(byte, []byte) error
}

func New(dev bus) *MCP23017 {
	t := new(MCP23017)
	t.dev=dev
	t.StateA = 0
	t.StateB = 0
	return t
}

func (t *MCP23017) write(addr uint8, data uint8) {
	err := t.dev.WriteReg(addr, []byte{data})
	if err != nil {
		log.Fatal(err)
	}
}

func (t *MCP23017) read(addr uint8) byte {
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
	t.write(0x00, t.S2BV(v)) // IODIRA
}
func (t *MCP23017) DirectionB(v ...uint8) {
	t.write(0x01, t.S2BV(v)) // IODIRB
}
func (t *MCP23017) PullUpA(v ...uint8) {
	t.write(0x0c, t.S2BV(v)) // GPPUA
}
func (t *MCP23017) PullUpB(v ...uint8) {
	t.write(0x0d, t.S2BV(v)) // GPPUB
}
func (t *MCP23017) ApplyA() {
	t.write(0x14, t.StateA) // OLATA
}
func (t *MCP23017) ApplyB() {
	t.write(0x15, t.StateB) // OLATB
}
func (t *MCP23017) FetchA() {
	t.StateA = t.read(0x12) // GPIOA
}
func (t *MCP23017) FetchB() {
	t.StateB = t.read(0x13) // GPIOB
}

func (t *MCP23017) SetA(p byte, v,a bool) {
	if v {
		t.StateA = t.StateA |  ( 1 << p )
	} else {
		t.StateA = t.StateA &^ ( 1 << p )
	}
	if a {
		t.ApplyA()
	}
}

func (t *MCP23017) SetB(p byte, v,a bool) {
	if v {
		t.StateB = t.StateB |  ( 1 << p )
	} else {
		t.StateB = t.StateB &^ ( 1 << p )
	}
	if a {
		t.ApplyB()
	}
}

func (t *MCP23017) GetA(p byte) bool {
	if t.StateA & ( 1 << p ) > 0 {
		return true
	} else {
		return false
	}
}
func (t *MCP23017) GetB(p byte) bool {
	if t.StateB & ( 1 << p ) > 0 {
		return true
	} else {
		return false
	}
}
