package main

import (
	"sort"
	"fmt"
)

type User struct {
	Name string
	Age int 
}

type UserSlice []User

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User {
		{"nozami", 21},
		{"fardan", 31},
		{"ajitama", 25},
		{"Wildan", 4},
		{"Firdaus", 40},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}