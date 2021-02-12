package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/fignocius/microsservices/view/model"
)

var status = map[int]map[int]string{
	0: {0: "Triagem", 1: "Pacote triado para transporte"},
	1: {0: "Em Transporte", 1: "Em transferência entre cidades"},
	2: {0: "Entregue", 1: "Entregue ao destinatário"},
}

// Parameters struct of request
type Parameters struct {
	Code        string `json:"code"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

// Buffer return json bytes buffer
func (m Parameters) Buffer() *bytes.Buffer {
	b, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	return bytes.NewBuffer(b)
}

func main() {
	var (
		trackings []model.Tracking
		err       error
	)
	if trackings, err = create(); err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	if err = update(trackings); err != nil {
		fmt.Printf("Error Update: %s", err)
		return
	}
	if err = view(trackings); err != nil {
		fmt.Printf("Error View: %s", err)
		return
	}
	fmt.Println("----------- Final do tracking ------------")
	return
}

func create() (trackings []model.Tracking, err error) {
	fmt.Println("-------- Tracking creates -----------")
	for i := 0; i < 6; i++ {
		var (
			t       model.Tracking
			res     *http.Response
			payload []byte
		)
		res, err = http.Get(os.Getenv("CREATE_URL"))
		if err != nil {
			return
		}
		payload, err = ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return
		}
		if err = json.Unmarshal(payload, &t); err != nil {
			return
		}
		trackings = append(trackings, t)
		fmt.Printf("tracking-%d: %s\n", i, t.String())
		time.Sleep(1 * time.Second)
	}
	return
}

func update(trackings []model.Tracking) (err error) {
	fmt.Println("-------- Trackings updates -----------")
	var (
		req    *http.Request
		client = http.Client{}
	)
	rand.Seed(time.Now().UnixNano())
	for _, tracking := range trackings {
		var (
			num  = rand.Intn(2)
			data = Parameters{
				Code:        tracking.Code,
				Status:      status[num][0],
				Description: status[num][1],
			}
		)

		if req, err = http.NewRequest("POST", os.Getenv("UPDATE_URL"), data.Buffer()); err != nil {
			return
		}
		req.Header.Set("Content-Type", "application/json")
		if _, err = client.Do(req); err != nil {
			return
		}
		fmt.Printf("tracking code: %s atualizado\n", data.Code)
		time.Sleep(1 * time.Second)
	}
	return
}

func view(trackings []model.Tracking) (err error) {
	fmt.Println("-------- Trackings status -----------")
	var (
		t       model.Tracking
		res     *http.Response
		payload []byte
	)
	for i, tracking := range trackings {
		res, err = http.Get(fmt.Sprintf("%s/%s", os.Getenv("VIEW_URL"), tracking.Code))
		if err != nil {
			return
		}
		payload, err = ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return
		}
		if err = json.Unmarshal(payload, &t); err != nil {
			return
		}
		fmt.Printf("tracking-%d: %s\n", i, t.String())
		time.Sleep(1 * time.Second)
	}
	return
}
