package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	. "github.com/siongui/godom"
)

type responseMsg struct {
	ShortUrl string `json:"short_url"`
	Err      string `json:"err"`
}

var input = Document.GetElementById("input")
var msg = Document.GetElementById("msg")
var btn = Document.GetElementById("create-btn")

func create(jsonValue []byte) {
	req, err := http.NewRequest("POST", "/create", bytes.NewBuffer(jsonValue))
	if err != nil {
		msg.SetTextContent(err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		msg.SetTextContent(err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		msg.SetTextContent(err.Error())
		return
	}

	respMsg := responseMsg{}
	json.Unmarshal(body, &respMsg)
	if respMsg.Err == "" {
		msg.SetTextContent(Window.Location().Href() + respMsg.ShortUrl)
	} else {
		msg.SetTextContent(respMsg.Err)
	}
}

func main() {

	btn.AddEventListener("click", func(e Event) {
		url := strings.TrimSpace(input.Value())
		if url != "" {
			value := map[string]string{"url": url}
			jsonValue, err := json.Marshal(value)
			if err != nil {
				msg.SetTextContent(err.Error())
				return
			}
			go create(jsonValue)
		}
	})
}
