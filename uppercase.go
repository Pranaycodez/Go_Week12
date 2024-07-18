// Package stringuppercase is used for converting a string to uppercase.
package stringuppercase

// imported the strings pacakge to use its prebuil function ToUpper
import "strings"

/*
ToUpperCase function is used to convert a string to uppercase.
It takes a string as input and returns a string in uppercase.
*/
func ToUpperCase(s string) string {
	return strings.ToUpper(s) //used the ToUpper function from strings package.
}
