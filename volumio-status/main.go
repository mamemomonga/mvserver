package main

import (
	"encoding/json"
	leds "github.com/mamemomonga/rpi-go-74hc595led/simple"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"strconv"
	// "github.com/davecgh/go-spew/spew"
)

var (
	rate  = 0
	service = ""
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
	hwp()
	vlmo()
}

func hwp() {
	hwp := hw_params()
	if _, ok := hwp["rate"]; ok {
		if !playing {
			playing = true
		}
	} else {
		if playing {
			log.Printf("STOP\n")
			setLedRate(-1)
			rate=0
			playing = false
		}
		return
	}
	if !playing {
		return
	}
	var rate_new int
	{
		rt := strings.Split(hwp["rate"], " ")
		rate_new,_ = strconv.Atoi(rt[0])
	}
	if rate_new == rate {
		return
	}
	rate=rate_new
	log.Printf("Rate: %d\n", rate)
	switch rate {
	case 44100:
		setLedRate(0)
	case 48000:
		setLedRate(1)
	default:
		if rate >= 96000 {
			setLedRate(2)
		} else {
			setLedRate(3)
		}
	}
}

func vlmo() {
	vm := getVolumioState()
	if vm.Service == service {
		return
	}
	service = vm.Service
	log.Printf("Service: %s",service)

	switch service {
	case "":
		setLedService(-1)
	case "mpd":
		setLedService(4)
	case "volspotconnect2":
		setLedService(5)
	case "airplay_emulation":
		setLedService(6)
	default:
		setLedService(7)
	}
}

func setLedRate(led int) {
	for i := uint8(0); i <= 3; i++ {
		if i == uint8(led) {
			leds.Set(0, i, 1)
		} else {
			leds.Set(0, i, 0)
		}
	}
}

func setLedService(led int) {
	for i := uint8(4); i <= 7; i++ {
		if i == uint8(led) {
			leds.Set(0, i, 1)
		} else {
			leds.Set(0, i, 0)
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
		log.Printf("Error: %s", err)
		return
	}
	if res.StatusCode != 200 {
		log.Printf("StatusCode: %d", res.StatusCode)
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
	f, err := os.Open("/proc/asound/sndrpihifiberry/pcm0p/sub0/hw_params")
	if err != nil {
		return
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if string(b) == "closed\n" {
		return
	}
	items := strings.FieldsFunc(string(b), func(c rune) bool { return c == 10 })
	for _, item := range items {
		x := strings.Split(item, ": ")
		m[x[0]] = x[1]
	}
	return
}
