package volumioapi_test

import (
	"testing"
	"github.com/rpi-volumio-status-led/volumio-status/volumioapi"
	"github.com/davecgh/go-spew/spew"
//	"log"
)

func Test01(t *testing.T) {
	v := volumioapi.GetState()
	spew.Dump(v)
}

