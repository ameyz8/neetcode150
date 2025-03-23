package main

import "fmt"

func isValid(s string) bool {
	stackProcess := []string{}
	pairs := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}
	for i := range s {
		if s[i:i+1] == "{" || s[i:i+1] == "(" || s[i:i+1] == "[" {
			stackProcess = append(stackProcess, s[i:i+1])
		} else if s[i:i+1] == "}" || s[i:i+1] == ")" || s[i:i+1] == "]" {
			if len(stackProcess) > 0 && pairs[s[i:i+1]] == stackProcess[len(stackProcess)-1] {
				stackProcess = stackProcess[:len(stackProcess)-1]
			} else {
				return false
			}
		}
	}
	if len(stackProcess) != 0 {
		return false
	}
	return true
}

func main() {
	s := "}[}]"
	fmt.Println(isValid(s))
}
