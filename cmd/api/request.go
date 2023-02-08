package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func request(d *Download) (int, error) {
	//Create Req
	r, err := d.newRequest("HEAD")
	if err != nil {
		return 0, err
	}

	//Send Req
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return 0, err
	}
	if res.StatusCode > 299 {
		return 0, fmt.Errorf("can't process response :%v", res.StatusCode)
	}

	//Get Size of File
	len := res.Header.Get("Content-Length")
	size, err := strconv.Atoi(len)
	if err != nil {
		return 0, err
	}
	fmt.Println("Size in bytes: ", size)
	return size, nil
}

func (d Download) newRequest(method string) (*http.Request, error) {
	r, err := http.NewRequest(method, d.Url, nil)
	if err != nil {
		return nil, err
	}
	r.Header.Set("User-Agent", "DownMan ")
	return r, nil
}
