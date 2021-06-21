package network

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func BaseUrl() string {
	//return "http://hacknode.io/"
	return "https://bitclout.com/"
}

func DoGet(route string) string {
	agent := "agent"

	urlString := fmt.Sprintf("%s%s", BaseUrl(), route)
	request, _ := http.NewRequest("GET", urlString, nil)
	request.Header.Set("User-Agent", agent)
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 5}
	return DoHttpRead("GET", route, client, request)
}

func DoHttpRead(verb, route string, client *http.Client, request *http.Request) string {
	resp, err := client.Do(request)
	if err == nil {
		defer resp.Body.Close()
		//body, err := ioutil.ReadAll(resp.Body)
		var buff bytes.Buffer
		io.Copy(&buff, resp.Body)
		body := buff.Bytes()
		if err != nil {
			fmt.Printf("\n\nERROR: %d %s\n\n", resp.StatusCode, err.Error())
			return ""
		}
		if resp.StatusCode == 200 || resp.StatusCode == 201 || resp.StatusCode == 204 {
			return string(body)
		} else {
			text := string(body)
			if strings.Contains(text, "RuleErrorFollowEntryAlreadyExists") == false {
				fmt.Printf("\n\nERROR: %d %s\n\n", resp.StatusCode, string(body))
			}
			return ""
		}
	}
	fmt.Printf("\n\nERROR: %s\n\n", err.Error())
	return ""
}

func DoPost(route string, payload []byte) string {
	body := bytes.NewBuffer(payload)
	urlString := fmt.Sprintf("%s%s", BaseUrl(), route)
	request, _ := http.NewRequest("POST", urlString, body)
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 50}

	return DoHttpRead("POST", route, client, request)
}
func DoPostMultipart(route, ct string, payload []byte) string {
	body := bytes.NewBuffer(payload)
	urlString := fmt.Sprintf("%s%s", BaseUrl(), route)
	request, _ := http.NewRequest("POST", urlString, body)
	request.Header.Set("Content-Type", ct)
	client := &http.Client{Timeout: time.Second * 50}

	return DoHttpRead("POST", route, client, request)
}
