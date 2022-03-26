package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Code    int     `json:"code"`
	Status  string  `json:"status"`
	Results Results `json:"results"`
}
type Times struct {
	Imsak    string `json:"Imsak"`
	Sunrise  string `json:"Sunrise"`
	Fajr     string `json:"Fajr"`
	Dhuhr    string `json:"Dhuhr"`
	Asr      string `json:"Asr"`
	Sunset   string `json:"Sunset"`
	Maghrib  string `json:"Maghrib"`
	Isha     string `json:"Isha"`
	Midnight string `json:"Midnight"`
}
type Date struct {
	Timestamp int    `json:"timestamp"`
	Gregorian string `json:"gregorian"`
	Hijri     string `json:"hijri"`
}

type Datetime struct {
	Times Times `json:"times"`
	Date  Date  `json:"date"`
}

type Results struct {
	Datetime []Datetime `json:"datetime"`
}

func main(){
	response, err := http.Get("https://api.pray.zone/v2/times/today.json?city=jakarta&school=10")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	prayer_time := responseObject.Results.Datetime[0].Times
	date_time := responseObject.Results.Datetime[0].Date

	fmt.Println("Shubuh\t: " + prayer_time.Fajr)
	fmt.Println("Dzuhr\t: " + prayer_time.Dhuhr)
	fmt.Println("Ashar\t: " + prayer_time.Asr)
	fmt.Println("Maghrib\t: " + prayer_time.Maghrib)
	fmt.Println("Isya\t: " + prayer_time.Isha)

	fmt.Println("Date\t: " + date_time.Gregorian)


}
