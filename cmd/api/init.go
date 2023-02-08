package main

import (
	"fmt"
	"sync"
)

type Download struct {
	Url        string
	TargetPath string
	MaxConn    int
}

func NewDownload(url string, targetPath string, maxConnection int) *Download {
	return &Download{
		Url:        url,
		TargetPath: targetPath,
		MaxConn:    maxConnection,
	}
}

func (d Download) Start(size int) error {
	fmt.Println("Making Connection...")

	//Create Sections/Partitions
	var sections = make([][2]int, d.MaxConn)
	eachSize := size / d.MaxConn

	//Setting starting And Ending Bytes for Each Section
	for i := range sections {
		if i == 0 {
			//First Section Starting
			sections[i][0] = 0
		} else {
			//Other Sections Starting will be previous end +1
			sections[i][0] = sections[i-1][1] + 1
		}
		if i < d.MaxConn-1 {
			//Ending Bytes of other sections
			sections[i][1] = sections[i][0] + eachSize
		} else {
			//Ending Bytes of Last Section
			sections[i][1] = size - 1
		}

	}
	//	fmt.Println(sections)

	//Download Each Sections
	var wg sync.WaitGroup

	for i, s := range sections {
		wg.Add(1)
		go func(i int, s [2]int) {
			defer wg.Done()
			err := d.downloadSection(i, s)
			if err != nil {
				panic(err)
			}
		}(i, s)
	}
	wg.Wait()
	fmt.Println("Download compleated..")
	return d.combine(sections)
}
