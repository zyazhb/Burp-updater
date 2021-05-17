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
	fmt.Printf("A Burpsuite Update Finder BY-ZYA\n ")
	fmt.Printf("   _____             _____     _ _          _____       _     _          _____ _       _         \n   | __  |_ _ ___ ___|   __|_ _|_| |_ ___   |  |  |___ _| |___| |_ ___   |   __|_|___ _| |___ ___ \n   | __ -| | |  _| . |__   | | | |  _| -_|  |  |  | . | . | .'|  _| -_|  |   __| |   | . | -_|  _|\n   |_____|___|_| |  _|_____|___|_|_| |___|  |_____|  _|___|__,|_| |___|  |__|  |_|_|_|___|___|_|  \n				 |_|                              |_|                                             \n")
	fmt.Printf("[-]Finding...\n")

	y := time.Now().Year()
	Maxmonth := 12
	if y == time.Now().Year() {
		Maxmonth = int(time.Now().Month())
	}
	for m := 1; m <= Maxmonth; m++ {
		for d := 1; d <= 3; d++ {
			ver := fmt.Sprint(y) + "." + fmt.Sprint(m) + "." + fmt.Sprint(d)
			wg.Add(1)
			go rescode(api, ver, &wg)

		}
	}

	wg.Wait()
	fmt.Printf("[+]Done!")
	time.Sleep(1 * time.Minute)
	os.Exit(0)
}

func rescode(api string, ver string, wg *sync.WaitGroup) {
	defer wg.Done()
	timeout := time.Duration(5 * time.Second)
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
		fmt.Println("[+]Found a version!", api+ver)
		return
	}
	resp.Body.Close()
}
