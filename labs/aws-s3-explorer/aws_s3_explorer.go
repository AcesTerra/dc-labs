package main

import(
	"fmt"
	"os"
	"net/http"
	"encoding/xml"
	"log"
	//"reflect"
	"io/ioutil"
	"bytes"
	"strconv"
	"strings"
)

type ListBucketResult struct {
	Content	[]Contents `xml:"Contents"`
}

type Contents struct {
	Key	string `xml:"Key"`
}

func getContent(content []Contents) (int, string){
	directoriesCntr := 0
	noExtension := 0
	extensionsMap := make(map[string]int)
	for _, item := range content{
		name := string(item.Key)
		//fmt.Println(reflect.TypeOf(name))
		if(name[len(name)-1] == '/'){
			//fmt.Println(name)
			directoriesCntr++
		} else{
			extension := strings.Split(name, ".")
			//fmt.Println(extension)
			if (len(extension) < 2){
				noExtension++
			} else{
				if value, ok := extensionsMap[extension[len(extension)-1]]; ok {
					value++
					extensionsMap[extension[len(extension)-1]] = value
				} else{
					extensionsMap[extension[len(extension)-1]] = 1
				}
			}
		}
	}
	extensionsMap["no-extension"] = noExtension
	//fmt.Println(extensionsMap)
	extensions := ""
	for key, value := range extensionsMap{
		extensions = extensions + key + "(" + strconv.Itoa(value) + "), "
	}
	//fmt.Println(extensions)
	extensions = extensions[0:len(extensions)-2]
	//fmt.Println(extensions)
	return directoriesCntr, extensions
}

func getXML(url string) (string){
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	rawXML, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(reflect.TypeOf(rawXML))
	return string(rawXML)
}

func decodeXML(data string) (int, int, string){
	//data := string(xmlData)
	//fmt.Println(data)
	buf := bytes.NewBufferString(data)
	list := new(ListBucketResult)
	decoded := xml.NewDecoder(buf)

	err := decoded.Decode(list)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	//fmt.Println(reflect.TypeOf(list.Content))
	directories, extensions := getContent(list.Content)
	//fmt.Printf("Key: %s\n", list.Content[0].Key)
	//fmt.Printf("Title: %s\n", rss.Channel.Items[0].Description)
	//fmt.Printf("Title: %s\n", rss.Channel.Items[1].Title)
	return len(list.Content), directories, extensions
}

func main(){
	if (len(os.Args) < 3){
		//fmt.Println("No flag")
		fmt.Println("Usage: go run aws_s3_explorer.go --bucket amazon_S3_file")
		return
	}
	bucket := os.Args[1]
	//fmt.Println(bucket)
	if (bucket != "--bucket"){
		fmt.Println("Incorrect flag")
		fmt.Println("Usage: go run aws_s3_explorer.go --bucket amazon_S3_file")
		return
	}

	url := "https://" + os.Args[2] + ".s3.amazonaws.com"
	//fmt.Println(url)
	xml := getXML(url)
	//getXML(url)
	//fmt.Println(reflect.TypeOf(xml))
	objects, directories, extensions := decodeXML(xml)
	//decodeXML(xml)
	fmt.Println("AWS S3 Explorer")
	fmt.Printf("Bucket Name\t\t: %s\n", os.Args[2])
	fmt.Printf("Number of objects\t: %d\n", objects)
	fmt.Printf("Number of directories\t: %d\n", directories)
	fmt.Printf("Extensions\t\t: %s\n", extensions)
}
