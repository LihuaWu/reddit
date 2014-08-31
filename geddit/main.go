package main

import (
	"log"
	"fmt"
	"github.com/LihuaWu/reddit"
)

func main() {
	items, err := reddit.Get("golang")
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Println(item)
	}
//	resp, err := http.Get("http://reddit.com/r/golang.json")
//	if err != nil {
//		log.Fatal(err)
//	}
//	if resp.StatusCode != http.StatusOK {
//		log.Fatal(resp.Status)
//	}
////	_, err = io.Copy(os.Stdout, resp.Body)
//	r := new(Response)
//	err = json.NewDecoder(resp.Body).Decode(r)
//	for _, child := range r.Data.Children {
//		fmt.Println(child.Data.Title)
//	//	fmt.Println(child.Data.URL)
//	}	
//	if err != nil {
//		log.Fatal(err)
//	}
}
