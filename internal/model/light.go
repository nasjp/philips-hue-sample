package model

import (
	"math/rand"
	"strconv"
	"time"
)

type Light struct {
	State            State        `json:"state"`
	Swupdate         Swupdate     `json:"swupdate"`
	Type             string       `json:"type"`
	Name             string       `json:"name"`
	Modelid          string       `json:"modelid"`
	Manufacturername string       `json:"manufacturername"`
	Productname      string       `json:"productname"`
	Capabilities     Capabilities `json:"capabilities"`
	Config           Config       `json:"config"`
	Uniqueid         string       `json:"uniqueid"`
	Swversion        string       `json:"swversion"`
}

type State struct {
	Bri       int    `json:"bri"`
	Ct        int    `json:"ct"`
	Alert     string `json:"alert"`
	Colormode string `json:"colormode"`
	Mode      string `json:"mode"`
	Reachable bool   `json:"reachable"`
	On        bool   `json:"on"`
}

type Swupdate struct {
	State       string `json:"state"`
	Lastinstall string `json:"lastinstall"`
}

type Ct struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type Control struct {
	Mindimlevel int `json:"mindimlevel"`
	Maxlumen    int `json:"maxlumen"`
	Ct          Ct  `json:"ct"`
}

type Streaming struct {
	Renderer bool `json:"renderer"`
	Proxy    bool `json:"proxy"`
}

type Capabilities struct {
	Control   Control   `json:"control"`
	Streaming Streaming `json:"streaming"`
	Certified bool      `json:"certified"`
}

type Startup struct {
	Mode       string `json:"mode"`
	Configured bool   `json:"configured"`
}

type Config struct {
	Archetype string  `json:"archetype"`
	Function  string  `json:"function"`
	Direction string  `json:"direction"`
	Startup   Startup `json:"startup"`
}

type Lights map[string]*Light

func (ls Lights) GetRandomReachableID() int {
	for idStr, l := range ls {
		if !l.State.Reachable {
			continue
		}

		id, _ := strconv.Atoi(idStr)

		return id
	}

	return 0
}

func (ls Lights) ExistReachable() bool {
	for _, l := range ls {
		if l.State.Reachable {
			return true
		}
	}

	return false
}

const (
	BriMin = 0
	BriMax = 255

	CtMin = 153
	CtMax = 454
)

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func RandomBri() int {
	return random(BriMin, BriMax)
}

func RandomCt() int {
	return random(CtMin, CtMax)
}
