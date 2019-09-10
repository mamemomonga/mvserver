package volumioapi

import (
	"net/http"
	"time"
	"log"
	"encoding/json"
	"io/ioutil"
)

type State struct {
	Status     string `json:"status"`
	Title      string `json:"title"`
	Artist     string `json:"artist"`
	Album      string `json:"album"`
	Albumart   string `json:"albumart"`
	URI        string `json:"uri"`
	TrackType  string `json:"trackType"`
	Seek       int    `json:"seek"`
	Duration   int    `json:"duration"`
	Samplerate string `json:"samplerate"`
	Bitdepth   string `json:"bitdepth"`
	Channels   int    `json:"channels"`
	Consume    string `json:"consume"`
	Volume     int    `json:"volume"`
	Mute       bool   `json:"mute"`
	Stream     bool   `json:"stream"`
	Updatedb   bool   `json:"updatedb"`
	Volatile   bool   `json:"volatile"`
	Service    string `json:"service"`
}

func GetState() (d State) {
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
	err = json.Unmarshal(b, &d)
	if err != nil {
		log.Printf("ERR: %s",err)
	}
	return
}

