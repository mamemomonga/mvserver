package main

import (
	"os"
	"io/ioutil"
	"log"
	"strings"
	"time"
	leds "github.com/mamemomonga/rpi-go-74hc595led/simple"
	"net/http"
	"os/signal"
	"syscall"
	"encoding/json"
	// "github.com/davecgh/go-spew/spew"
)

var (
	format  = "DUMMY"
	rate    = ""
	service = ""
	update  = false
	playing = false
)

func main() {
	leds.Start(1)
	defer leds.Finalize()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT)

	leds.AllHigh()
	time.Sleep(time.Second)
	leds.AllLow()
	time.Sleep(time.Second)

	go func() {
		for {
			do_update()
			time.Sleep(time.Second)
		}
	}()

	<-quit
	leds.AllLow()
}

func do_update() {
	hwp:=hw_params()
	if _,ok := hwp["format"]; ok {
		if !playing {
			playing = true
		}
	} else {
		if playing {
			log.Printf("STOP\n")
			setLedRate(-1)
			setLedService(-1)
			format = ""
			service = ""
			playing = false
		}
		return
	}
	vm := getVolumioState()

	if hwp["format"] != format {
		format = hwp["format"]
		update = true
	}
	if hwp["rate"] != rate {
		rate = hwp["rate"]
		update = true
	}
	if playing {
		if vm.Service != service {
			service = vm.Service
			update = true
		}
	}

	if !update {
		return
	}

	rates := strings.Split(rate," ")
	log.Printf("PLAY Format: %s | Rate: %s | Service %s\n",format, rates[0], service)
	switch(rates[0]) {
	case "44100":
		setLedRate(0)
	case "48000":
		setLedRate(1)
	case "96000":
		setLedRate(2)
	default:
		setLedRate(3)
	}

	switch(service) {
	case "mpd":
		setLedService(4)
	case "volspotconnect2":
		setLedService(5)
	case "airplay_emulation":
		setLedService(6)
	default:
		setLedService(7)
	}
	update = false
}

func setLedRate(led int) {
	for i:=uint8(0); i<=3; i++ {
		if i==uint8(led) {
			leds.Set(0,i,1)
		} else {
			leds.Set(0,i,0)
		}
	}
}

func setLedService(led int) {
	for i:=uint8(4); i<=7; i++ {
		if i==uint8(led) {
			leds.Set(0,i,1)
		} else {
			leds.Set(0,i,0)
		}
	}
}

func getVolumioState() (data VolumioAPIGetState) {
	data = VolumioAPIGetState{}
	timeout := time.Duration(time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	res, err := client.Get("http://localhost:3000/api/v1/getState")
	if err != nil {
		log.Printf("Error: %s",err)
		return
	}
	if res.StatusCode != 200 {
		log.Printf("StatusCode: %d",res.StatusCode)
		return
	}
	b, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func hw_params() (m map[string]string) {
    m = make(map[string]string)
	f,err := os.Open("/proc/asound/sndrpihifiberry/pcm0p/sub0/hw_params")
	if err != nil {
		return
	}
	defer f.Close()
	b,err := ioutil.ReadAll(f)
	if string(b) == "closed\n" {
		return
	}
	items := strings.FieldsFunc(string(b),func(c rune) bool { return c == 10 })
	for _, item := range items {
		x := strings.Split(item, ": ")
		m[x[0]] = x[1]
	}
	return
}

