package main

//获取http状态码

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	api := "https://portswigger.net/burp/releases/download?product=pro&type=windowsx64&version="
	for y := 2020; y <= time.Now().Year(); y++ {
		for m := 1; m <= 12; m++ {
			for d := 1; d <= 3; d++ {
				ver := fmt.Sprint(y) + "." + fmt.Sprint(m) + "." + fmt.Sprint(d)
				wg.Add(1)
				go rescode(api, ver, &wg)

			}
		}
	}
	wg.Wait()
	os.Exit(0)
}

func rescode(api string, ver string, wg *sync.WaitGroup) {
	defer wg.Done()
	timeout := time.Duration(3 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	u, _ := url.Parse(api + ver)
	q := u.Query()
	u.RawQuery = q.Encode()
	resp, err := client.Get(u.String())
	if err != nil {
		return
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Found a version!", api+ver)
		return
	}
	resp.Body.Close()
	return
}
