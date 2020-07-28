package model

import "fmt"

type HueConfig struct {
	IPAddress    string  `json:"ipAddress"`
	UserName     string  `json:"userName"`
	ContinuesSec float64 `json:"continuesSec"`
}

func (hc *HueConfig) URL() string {
	return fmt.Sprintf("http://%s/api/%s", hc.IPAddress, hc.UserName)
}
