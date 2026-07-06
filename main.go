package main

import (
	"fmt"
	"searchPerson/search"
)

func main(){
	users := []string{"john", "dimas", "rafli", "firman", "zacky",}
	var keyword string
	fmt.Print("Masukkan nama yang ingin dicari: ")
	fmt.Scan(&keyword)
	fmt.Println(search.SearchPerson(users, keyword))
}