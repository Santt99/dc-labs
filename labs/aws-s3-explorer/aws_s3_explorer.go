package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type ListBucketResult struct {
	XMLName  xml.Name   `xml:"ListBucketResult"`
	Contents []Contents `xml:"Contents"`
	Name     string     `xml:"Name"`
}

type Contents struct {
	XMLName      xml.Name `xml:"Contents"`
	Key          string   `xml:"Key"`
	LastModified string   `xml:"LastModified"`
	ETag         string   `xml:"ETag"`
	Size         string   `xml:"Size"`
	StorageClass string   `xml:"StorageClass"`
}

func main() {
	var bucket string

	bucket = os.Args[1]
	resp, err := http.Get("https://" + bucket + ".s3.amazonaws.com")
	if err != nil {
		return
	}
	var contents ListBucketResult
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(body, &contents)
	extensions := make(map[string]int)
	types := make(map[string]int)

	for current := range contents.Contents {
		key := contents.Contents[current].Key

		if strings.Contains(key, ".") {
			parts := strings.Split(key, "/")

			fileName := parts[len(parts)-1]

			fileNameSplited := strings.Split(fileName, ".")

			extensions[fileNameSplited[len(fileNameSplited)-1]]++
			types["object"]++
		} else {
			types["dir"]++
		}

	}
	fmt.Println("AWS S3 Explorer")
	fmt.Printf("Bucket Name            : %s\n", contents.Name)
	fmt.Printf("Number of objects      : %d\n", types["object"])
	fmt.Printf("Number of directories  : %d\n", types["dir"])
	fmt.Print("Extensions             : ")
	for key, value := range extensions {
		fmt.Printf("%s(%d), ", key, value)
	}
	fmt.Print("\n")
}
