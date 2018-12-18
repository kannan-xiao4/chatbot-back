package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Response struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Results []Result `json:"results"`
}

type Result struct {
	Perplexity float32 `json:"perplexity"`
	Reply      string  `json:"reply"`
}

//リクルートのやつを使う
//https://a3rt.recruit-tech.co.jp/product/talkAPI/
func Talk(message string) string {
	requesturl := os.Getenv("RECRUIT_SMARTTALK_ENDPOINT")
	apikey := os.Getenv("RECRUIT_SMARTTALK_API_KEY")

	client := &http.Client{}
	values := url.Values{}
	values.Add("apikey", apikey)
	values.Add("query", message)

	req, err := http.NewRequest("POST", requesturl, strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Sprintf("Fail Make Request: %v\n", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Sprintf("Fail Client Request: %v\n", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Sprintf("Fail Http Requst: %v\n", req)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Sprintf("Fail Read Response Body: %v\n", err)
	}

	str := string(body)
	fmt.Println(str)

	// JSONデコード
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Sprintf("Fail Json Decode: %v\n", err)
	}

	if response.Status != 0 {
		return fmt.Sprintf("Fail Request Result: %v\n", response.Message)
	}

	return response.Results[0].Reply
}
