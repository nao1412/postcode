package main

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
)

func main()  {
	values := url.Values{}
	values.Add("appid" , <appid>)
	values.Add("query" , <郵便番号>)
	values.Add("output", "json")

	url := "https://map.yahooapis.jp/search/zip/V1/zipCodeSearch"

	request, err := http.NewRequest("GET", url,, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	request.URL.RawQuery = values.Encode()

	client := new(http.Client)
	resp, _ := client.Do(request)
	defer resp.body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}