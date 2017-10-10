package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func makeKey(element string) string {
	shot_key := strings.ToLower(element)
	key := SortString(shot_key)
	return key
}

func main() {
	b, _ := ioutil.ReadFile("sample.txt")

	var input_strings string = string(b)
	extracted_data := make(map[string]string)
	input_strings_array := strings.Fields(input_strings)
	for _, element := range input_strings_array {
		key_string := makeKey(element)
		if _, ok := extracted_data[key_string]; ok {
			extracted_data[key_string] = (extracted_data[key_string] + "','" + element)
			fmt.Println("--Update--%x", element)
		} else {
			extracted_data[key_string] = element
			fmt.Println("--Add---%x", element)
		}
	}
	fmt.Println(extracted_data)
}
