package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
	"strconv"
)

func main() {

	user_map := map[int][][]string {}
	player_ids_array := []int{134815, 133171}
	for _, player_id := range player_ids_array {
		resp, err := soup.Get("https://pesdb.net/pes2021/?id=" + strconv.Itoa(player_id))
		
		if err != nil {
			os.Exit(1)
		}
		doc := soup.HTMLParse(resp)
		links := doc.FindAllStrict("tr", "data-percent", "100")
		var one_user [][]string // create a slice to hold data for one player
		for _, link := range links {
			var one_row []string 
			for _, x := range(link.Children()){
				one_row = append(one_row, x.FullText()) // add a row of data
			}
			one_user = append(one_user, one_row) // add the row of data into that player
		}
		user_map[player_id] = one_user // assign to a map with key player id
	}
	fmt.Println(user_map)
}