package collection

// The below code is provided by Mark McGranaghan at https://gobyexample.com/collection-functions
// and is licensed under a Creative Commons Attribution 3.0 Unported License - https://creativecommons.org/licenses/by/3.0/
// Some modifications have been made to the names of the functions and parameters.

// IndexOf() - returns the first index of the target string element, or -1 if no match is found
func IndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found
}

// Contains() - returns true if the target string str is found in the slice
func Contains(str string, sliceToSearch []string) bool {
	for _, a := range sliceToSearch {
		if a == str {
			return true
		}
	}
	return false //not found
}

// Any() - returns true if one of the strings in the slice satisfies the predicate f
func Any(sliceToSearch []string, f func(string) bool) bool {
	for _, v := range sliceToSearch {
		if f(v) {
			return true
		}
	}
	return false
}

// All() - returns true if all of the strings in the slice satisfy the predicate f.
func All(sliceToSearch []string, f func(string) bool) bool {
	for _, v := range sliceToSearch {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter() - returns a new slice containing all strings in the slice that satisfy the predicate f
func Filter(sliceToSearch []string, f func(string) bool) []string {
	results := make([]string, 0)
	for _, v := range sliceToSearch {
		if f(v) {
			results = append(results, v)
		}
	}
	return results
}

// Map() - returns a new slice containing the results of applying the function f to each string in the original slice
func Map(sliceToModify []string, f func(string) string) []string {
	results := make([]string, len(sliceToModify))
	for i, v := range sliceToModify {
		results[i] = f(v)
	}
	return results
}
