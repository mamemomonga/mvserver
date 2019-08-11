package main

// http://localhost:3000/api/v1/getState

type VolumioAPIGetState struct {
	Status               string      `json:"status"`
	Position             int         `json:"position"`
	Title                string      `json:"title"`
	Artist               string      `json:"artist"`
	Album                string      `json:"album"`
	Albumart             string      `json:"albumart"`
	URI                  string      `json:"uri"`
	TrackType            string      `json:"trackType"`
	Seek                 int         `json:"seek"`
	Duration             int         `json:"duration"`
	Samplerate           string      `json:"samplerate"`
	Bitdepth             string      `json:"bitdepth"`
	Channels             int         `json:"channels"`
	RepeatSingle         bool        `json:"repeatSingle"`
	Consume              bool        `json:"consume"`
	Volume               interface{} `json:"volume"`
	DisableVolumeControl bool        `json:"disableVolumeControl"`
	Mute                 bool        `json:"mute"`
	Stream               interface{} `json:"stream"`
	Updatedb             bool        `json:"updatedb"`
	Volatile             bool        `json:"volatile"`
	Service              string      `json:"service"`
}
