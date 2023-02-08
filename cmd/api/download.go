package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func (d Download) downloadSection(idx int, s [2]int) error {
	r, err := d.newRequest("GET")
	if err != nil {
		return err
	}
	r.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", s[0], s[1]))
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	if res.StatusCode > 299 {
		return fmt.Errorf("cannot process, response is %v", res.StatusCode)
	}
	fmt.Printf("Downloaded %v bytes for section %v\n", res.Header.Get("Content-Length"), idx)
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("section-%v.tmp", idx), b, os.ModePerm)
	if err != nil {
		return err
	}
	fmt.Printf("Downloaded %v bytes for part %v: %v\n", res.Header.Get("Content-Length"), idx, s)

	return nil
}
