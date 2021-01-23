package main

//获取http状态码

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func main() {
	//"https://portswigger.net/burp/releases/download?product=pro&version=2020.11.3&type=jar"
	api := "https://portswigger.net/burp/releases/download?product=pro&type=windowsx64&version="
	for y := 2019; y <= time.Now().Year(); y++ {
		for m := 0; m <= int(time.Now().Month()); m++ {
			for d := 0; d <= 31; d++ {
				apiurl := api + string(y) + "." + string(m) + "." + string(d)
				res, err := rescode(apiurl)
				if err != nil && res == 200 {
					fmt.Printf(apiurl)
				}
			}
		}
	}
}

func rescode(apiurl string) (int, error) {
	u, _ := url.Parse(apiurl)
	q := u.Query()
	u.RawQuery = q.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		fmt.Println("0")
		return 0, err
	}
	resCode := res.StatusCode
	res.Body.Close()
	if err != nil {
		fmt.Println("0")
		return 0, err
	}
	fmt.Printf("%d\r\n", resCode)
	return resCode, nil
}
