package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
)

func main() {
	resp, err := soup.Get("https://pesdb.net/pes2021/?id=134815") 
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	links := doc.FindAllStrict("tr", "data-percent", "100")
	for _, link := range links {
		var one_user []string // create a slice
		for _, x := range(link.Children()){
			one_user = append(one_user, x.FullText())
		}
		fmt.Println(one_user)
	}
}