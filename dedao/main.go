package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
出师未捷身先死
网页版没有签到功能
手机上截取不到请求报文。。。
*/

func main() {
	url := "http://cloudconf.fengkongcloud.com/v2/device/conf"

	json := []byte(`{"organization":"oXGNW54W7cZcFS1goYH7","data":{"smid":"202003051201002d1c823c693fa4396cb021809f3e574a0111f7e0c272e753","sdkver":"2.5.0","os":"ios","enc":1,"md5":"1b1bccf3343e024e90e9b4c231571af1"}}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("resp status:", resp.Status)
	fmt.Println("resp Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("resp Body:", string(body))

	url2 := "https://piccdn.umiwi.com"

	resp2, err := http.Head(url2)
	if err != nil {
		panic(err)
	}
	defer resp2.Body.Close()
	fmt.Println("resp2 status:", resp2.Status)
	body2, _ := ioutil.ReadAll(resp2.Body)
	fmt.Println("resp2 Body:", string(body2))

	resp2, _ = http.Head("https://piccdn1.umiwi.com")
	fmt.Println("resp2 status:", resp2.Status)
	resp2, _ = http.Head("https://piccdn2.umiwi.com")
	fmt.Println("resp2 status:", resp2.Status)
	resp2, _ = http.Head("https://piccdn3.umiwi.com")
	fmt.Println("resp2 status:", resp2.Status)
	resp2, _ = http.Head("https://ulogs.umeng.com/")
	fmt.Println("resp2 status:", resp2.Status)
	resp2, _ = http.Head("https://logs.luojilab.com/")
	fmt.Println("resp2 status:", resp2.Status)
	resp2, _ = http.Head("https://entree.igetget.com")
	fmt.Println("resp2 status:", resp2.Status)
	resp2, _ = http.Head("https://adt.cpatrk.net")
	fmt.Println("resp2 status:", resp2.Status)
	resp2, _ = http.Head("https://ulogs.umeng.com")
	fmt.Println("resp2 status:", resp2.Status)
	resp2, _ = http.Head("https://logs.luojilab.com")
	fmt.Println("resp2 status:", resp2.Status)
	resp2, _ = http.Head("https://me.cpatrk.net")
	fmt.Println("resp2 status:", resp2.Status)
	//resp2, _ = http.Head("https://cmyd-10-132.getui.com")
	fmt.Println("resp2 status:", resp2.Status)
}
