package search

import "strings"

func SearchPerson(user []string, keyword string) []string{
	var result []string
	x :=0
	for x < len(user){
		if strings.EqualFold(user[x], keyword){
			result = append(result, user[x])
		}
		x++
	}
	return result
}