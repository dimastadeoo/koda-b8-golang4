package search

import "strings"

func SearchPerson(user []string, keyword string) []string{
	var result []string
	for x:= range len(user){
		if strings.EqualFold(user[x], keyword){
			result = append(result, user[x])
		}
	}
	return result
}