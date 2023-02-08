package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//Combine files
func (d Download) combine(sections [][2]int) error {
	f, err := os.OpenFile(d.TargetPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	for i := range sections {
		d, err := ioutil.ReadFile(fmt.Sprintf("section-%v.tmp", i))
		if err != nil {
			return err
		}
		_, err = f.Write(d)
		if err != nil {
			return err
		}
		err = os.Remove(fmt.Sprintf("section-%v.tmp", i))
		if err != nil {
			return err
		}
		fmt.Printf("section-%v.tmp merged\n", i)
	}
	return nil
}
