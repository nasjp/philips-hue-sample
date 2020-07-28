package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/nasjp/philips-hue-sample/internal/logger"
	"github.com/nasjp/philips-hue-sample/internal/model"
)

var _ http.Handler = (*danceHandler)(nil)

type danceHandler struct {
	l *logger.Logger
}

const configFileName = "hue-config.json"

func (h *danceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := h.getConfig(configFileName)
	if err != nil {
		h.l.Error(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	ls, err := h.getLights(c)
	if err != nil {
		h.l.Error(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := h.danceRandom(c, ls); err != nil {
		h.l.Error(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (h *danceHandler) getConfig(fileName string) (*model.HueConfig, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	hc := &model.HueConfig{}

	if err := json.NewDecoder(f).Decode(hc); err != nil {
		return nil, err
	}

	return hc, nil
}

func (h *danceHandler) getLights(hc *model.HueConfig) (model.Lights, error) {
	req, err := http.NewRequest("GET", hc.URL()+"/lights", nil)
	if err != nil {
		return nil, err
	}

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ls := model.Lights{}

	if err := json.NewDecoder(resp.Body).Decode(&ls); err != nil {
		return nil, err
	}

	return ls, nil
}

func (h *danceHandler) danceRandom(hc *model.HueConfig, ls model.Lights) error {
	if !ls.ExistReachable() {
		h.l.Println("skip")
		return nil
	}

	start := time.Now()

	h.l.Println("dancing")

	for i := 0; time.Since(start).Seconds() < hc.ContinuesSec; i++ {
		id := ls.GetRandomReachableID()

		respBody, err := h.dance(id, true, model.RandomBri(), model.RandomCt(), hc.URL())
		if err != nil {
			return err
		}

		h.l.Printf("  id: %d %s\n", id, string(respBody))

		time.Sleep(100 * time.Millisecond)
	}

	h.l.Println("clean up")

	for idStr, l := range ls {
		id, _ := strconv.Atoi(idStr)
		respBody, err := h.dance(id, l.State.On, l.State.Bri, l.State.Ct, hc.URL())

		if err != nil {
			return err
		}

		h.l.Printf("id: %d %s\n", id, string(respBody))
	}

	return nil
}

func (h *danceHandler) dance(id int, on bool, bri int, ct int, url string) ([]byte, error) {
	bs, err := json.Marshal(map[string]interface{}{
		"on":  on,
		"bri": bri,
		"ct":  ct,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/lights/%d/state", url, id), bytes.NewBuffer(bs))
	if err != nil {
		return nil, err
	}

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}
