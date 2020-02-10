package main

import (
	"fmt"
	"regexp"
)

func main() {
	var subject = `I used MA35 regarding all questions about residence and citizenship and it is the weirdest experience of all my life.`

	matches := findMatches("(?i)(?P<first_part_of_word>reg).*(?P<second_part_of_word>exp)", subject)

	if matches["first_part_of_word"] == "" {
		fmt.Println("There is no first part of word :(")
		return
	}

	if matches["second_part_of_word"] == "" {
		fmt.Println("There is no first part of word :(")
		return
	}

	fmt.Println(fmt.Sprintf("First part of word: %s\n", matches["first_part_of_word"]))
	fmt.Println(fmt.Sprintf("Second part of word: %s\n", matches["second_part_of_word"]))
	fmt.Println(fmt.Sprintf("Result: %s%s", matches["first_part_of_word"], matches["second_part_of_word"]))
}

func findMatches(regex string, subject string) map[string]string {
	//Prepare the regexp object for selected regex
	re := regexp.MustCompile(regex)

	//Find matches for selected subject
	matches := re.FindStringSubmatch(subject)

	if len(matches) == 0 {
		return map[string]string{}
	}

	//Go through all found matches and fill the result map by passing name of group as index or id
	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 {
			if name != "" {
				result[name] = matches[i]
			} else {
				result[fmt.Sprintf("%d", i)] = matches[i]
			}
		}
	}

	return result
}