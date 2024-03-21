package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func getIP() (string, error) {
	urlIP := "https://api.ipify.org/"
	rep, err := http.Get(urlIP)
	if err != nil {
		return "", err
	}
	defer rep.Body.Close()

	body, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func riskchecker(ip string) {
	urlScam := "https://scamalytics.com/ip/" + ip
	doc, err := goquery.NewDocument(urlScam)
	if err != nil {
		log.Fatal(err)
	}

	var data map[string]interface{}
	doc.Find("pre").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		err = json.Unmarshal([]byte(text), &data)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}
	})
	for key, value := range data {
		if key == "ip" {
			fmt.Printf("%s: %v\n", key, value)
		} else if key == "score" {
			fmt.Printf("%s: %v\n", key, value)
		} else {
			fmt.Printf("%s: %v\n", key, value)
		}
	}
}

func main() {
	ip_flag := flag.String("ip", "", "IP address to check")
	flag.Parse()

	if *ip_flag != "" {
		riskchecker(*ip_flag)
	} else {
		ip, err := getIP()
		if err != nil {
			log.Fatal(err)
		}
		riskchecker(ip)
	}
}
