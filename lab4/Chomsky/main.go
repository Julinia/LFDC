package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

// EPS ...
const EPS = "ε"

// Productions is the type of our rules format
type Productions map[string][]string

func readFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}
	return string(data)
}

func convertRules(raw string) Productions {
	result := make(Productions)
	rows := strings.Split(raw, "\r\n")

	for _, v := range rows {
		parts := strings.Split(v, " ")

		result[parts[0]] = append(result[parts[0]], parts[1])
	}

	return result
}

func removeEpsilons(rules Productions) Productions {
	for rule := range rules {

		for i, prod := range rules[rule] {
			if prod == EPS {
				rules[rule] = removeElement(rules[rule], i)
				removeEpsilonSymbol(rules, rune(rule[0]))
			}
		}
	}
	return rules
}

func removeEpsilonSymbol(rules Productions, symbol rune) {
	for rule := range rules {
		for _, prod := range rules[rule] {
			if strings.Contains(prod, string(symbol)) {
				var occ []int

				for idx, val := range prod {
					if val == symbol {
						occ = append(occ, idx)
					}
				}

				masks := powerSet(occ)[1:]

				/* for _, mask := range masks {
					temp := []rune(prod)
					for i, val := range mask {
						temp = append(temp[0:val-i], temp[val+1-i:]...)
						fmt.Println(string(temp), string(symbol), val)
					}
					rules[rule] = append(rules[rule], string(temp))
				} */

				for _, mask := range masks {
					temp := []rune(prod)
					for i := len(mask) - 1; i >= 0; i-- {
						val := mask[i]

						temp = append(temp[0:val], temp[val+1:]...)
					}
					rules[rule] = append(rules[rule], string(temp))
				}
			}
		}
	}
}

func removeElement(a []string, i int) []string {

	// Remove the element at index i from a.
	copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
	a[len(a)-1] = ""     // Erase last element (write zero value).
	a = a[:len(a)-1]     // Truncate slice.

	return a
}

func powerSet(original []int) [][]int {
	powerSetSize := int(math.Pow(2, float64(len(original))))
	result := make([][]int, 0, powerSetSize)

	var index int
	for index < powerSetSize {
		var subSet []int

		for j, elem := range original {
			if index&(1<<uint(j)) > 0 {
				subSet = append(subSet, elem)
			}
		}
		result = append(result, subSet)
		index++
	}
	return result
}

func haveUnit(key string, rules Productions) (bool, string) {
	for _, prod := range rules[key] {
		if _, found := rules[prod]; found && len(prod) == 1 {
			return true, prod
		}
	}

	return false, ""
}

func removeUnit(key string, initProd string, rules Productions) Productions {
	if unit, prod := haveUnit(key, rules); unit {
		removeUnit(initProd, prod, rules)
	}

	for i, el := range rules[key] {
		if el == initProd {
			rules[key] = removeElement(rules[key], i)
			rules[key] = append(rules[key], rules[initProd]...)
			rules[key] = unique(rules[key])
		}
	}

	return rules
}

func removeUnits(rules Productions) Productions {
	for key := range rules {
		if unit, prod := haveUnit(key, rules); unit {
			rules = removeUnit(key, prod, rules)
		}
	}

	return rules
}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func removeInaccesibles(rules Productions) Productions {
	var accessed []string
	for key := range rules {
		for _, prod := range rules[key] {
			for _, letter := range prod {
				if _, found := rules[string(letter)]; found {
					accessed = append(accessed, string(letter))
				}
			}
		}
	}
	accessed = unique(accessed)

	for key := range rules {
		var isPresent = false
		for _, el := range accessed {
			if key == el {
				isPresent = true
			}
		}
		if !isPresent {
			delete(rules, key)
		}
	}

	return rules
}

func removeUnproductive(rules Productions) Productions {
	var productive []string
	//check if key has terminal element
	for key := range rules {
		for _, el := range rules[key] {
			if _, found := rules[el]; !found && len(el) == 1 {
				productive = unique(append(productive, key))
			}
		}
	}

	//check other productions
	counter := 1
	for counter != 0 {
		for key := range rules {
			if !stringInSlice(key, productive) {
				for _, el := range rules[key] {
					for _, letters := range el {
						if stringInSlice(string(letters), productive) {
							productive = unique(append(productive, key))
							counter++
						}
					}
				}
			}
		}
		counter--
	}

	for key := range rules { //[A]
		if !stringInSlice(key, productive) {
			delete(rules, key)
			for inner := range rules { //[S[...], B[...]]
				for i, inner_el := range rules[inner] {
					if strings.Contains(inner_el, key) {
						rules[inner] = removeElement(rules[inner], i)
					}
				}
			}
		}
	}

	fmt.Println(rules)
	return rules
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {
	raw := readFile("input.txt")

	rules := convertRules(raw)
	rules = removeEpsilons(rules)
	rules = removeUnits(rules)
	rules = removeInaccesibles(rules)
	rules = removeUnproductive(rules)

	fmt.Println(rules)

}
