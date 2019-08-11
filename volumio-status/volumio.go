package main

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
