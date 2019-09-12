package runners

import (
	"net/http"
	"time"
	"log"
	"encoding/json"
	"io/ioutil"
)

// http://localhost:3000/api/v1/getState

type VolumioAPIGetState struct {
	Status               string      `json:"status"`
	Position             interface{} `json:"position"`
	Title                interface{} `json:"title"`
	Artist               interface{} `json:"artist"`
	Album                interface{} `json:"album"`
	Albumart             interface{} `json:"albumart"`
	URI                  interface{} `json:"uri"`
	TrackType            interface{} `json:"trackType"`
	Seek                 interface{} `json:"seek"`
	Duration             interface{} `json:"duration"`
	Samplerate           interface{} `json:"samplerate"`
	Bitdepth             interface{} `json:"bitdepth"`
	Channels             interface{} `json:"channels"`
	RepeatSingle         interface{} `json:"repeatSingle"`
	Consume              interface{} `json:"consume"`
	Volume               interface{} `json:"volume"`
	DisableVolumeControl interface{} `json:"disableVolumeControl"`
	Mute                 interface{} `json:"mute"`
	Stream               interface{} `json:"stream"`
	Updatedb             interface{} `json:"updatedb"`
	Volatile             interface{} `json:"volatile"`
	Service              string      `json:"service"`
}

func getVolumioState() (data VolumioAPIGetState, err error) {
	data = VolumioAPIGetState{}
	timeout := time.Duration(time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	res, err := client.Get("http://localhost:3000/api/v1/getState")
	if err != nil {
		log.Printf("Error: %s", err)
		return data,err
	}
	if res.StatusCode != 200 {
		log.Printf("StatusCode: %d", res.StatusCode)
		return data,err
	}
	b, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Println(err)
		return data,err
	}
	return data,nil
}

func (t *Runners) runVolumioState() {
	hw := t.hw
	ticker := time.NewTicker(1 * time.Second)

	enable  := false
	playing := false
	service := ""

	for {
		volumio,err := getVolumioState()
		if err != nil {
			if enable {
				log.Println("[Volumio] Disable")
				hw.LedEnable.Set(false).Apply()
				enable = false
			}
			continue
		}

		// log.Printf("%#v\n",volumio)

		if !enable {
			log.Println("[Volumio] Enable")
			hw.LedEnable.Set(true)
			enable = true
		}
		if volumio.Status == "play" {
			if !playing {
				log.Println("[Volumio] Play")
				hw.LedPlaying.Set(true)
				playing = true
			}
		} else {
			if playing {
				log.Println("[Volumio] Stop")
				hw.LedPlaying.Set(false)
				playing = false
			}
		}

		if service != volumio.Service {
			switch volumio.Service {
			case "mpd":
				log.Println("[Volumio] Service:mpd")

			case "volspotconnect2":
				log.Println("[Volumio] Service:Spotify")

			case "airplay_emulation":
				log.Println("[Volumio] Service:AirPlay")

			default:
				log.Println("[Volumio] Service:Unknown")
			}
			service = volumio.Service
		}

		hw.Ex1.A.Apply()

		<-ticker.C
		select {
		case <-t.ctx.Done():
			return
		default:
		}

	}
}

