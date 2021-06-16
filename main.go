package main

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
)

func main()  {
	values := url.Values{}
	values.Add("appid", <appid>)
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

	type YOLP struct {
    ResultInfo struct {
        Count       int     `json:"Count"`
        Total       int     `json:"Total"`
        Start       int     `json:"Start"`
        Status      int     `json:"Status"`
        Description string  `json:"Description"`
        Copyright   string  `json:"Copyright"`
        Latency     float64 `json:"Latency"`
    } `json:"ResultInfo"`
    Feature []struct {
        ID       string `json:"Id"`
        Gid      string `json:"Gid"`
        Name     string `json:"Name"`
        Geometry struct {
            Type        string `json:"Type"`
            Coordinates string `json:"Coordinates"`
        } `json:"Geometry"`
        Category    []string      `json:"Category"`
        Description string        `json:"Description"`
        Style       []interface{} `json:"Style"`
        Property    struct {
            UID        string `json:"Uid"`
            CassetteID string `json:"CassetteId"`
            Country    struct {
                Code string `json:"Code"`
                Name string `json:"Name"`
            } `json:"Country"`
            Address              string `json:"Address"`
            GovernmentCode       string `json:"GovernmentCode"`
            AddressMatchingLevel string `json:"AddressMatchingLevel"`
            PostalName           string `json:"PostalName"`
            Station              []struct {
                ID       string `json:"Id"`
                SubID    string `json:"SubId"`
                Name     string `json:"Name"`
                Railway  string `json:"Railway"`
                Exit     string `json:"Exit"`
                ExitID   string `json:"ExitId"`
                Distance string `json:"Distance"`
                Time     string `json:"Time"`
                Geometry struct {
                    Type        string `json:"Type"`
                    Coordinates string `json:"Coordinates"`
                } `json:"Geometry"`
            } `json:"Station"`
        } `json:"Property"`
    } `json:"Feature"`
	}
	
	var data1 YOLP
	json.Unmarshal(body, &data1)
	fmt.Println(data1)
}
