// Package luhn provides a utility function for performing the
// "Luhn algorithm" or "mod 10" algorithm on a string of digits.
package luhn

import "regexp"

var doubles []int = []int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}

// Check returns true or false to indicate whether or not the supplied
// string of 2 or more digits passes or fails the Luhn check algorithm.
func Check(number string) bool {
	matched, err := regexp.MatchString("^[[:digit:]]{2,}$", number)
	if matched == false || err != nil {
		return false
	}

	checkSum := 0
	for i, double := len(number)-1, false; i >= 0; i-- {
		if double {
			checkSum += doubles[int(number[i]-'0')]
		} else {
			checkSum += int(number[i] - '0')
		}
		double = !double
	}
	return checkSum%10 == 0
}
