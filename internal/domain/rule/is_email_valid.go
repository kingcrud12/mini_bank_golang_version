package rule

import "regexp"

var regexEmailFormat = regexp.MustCompile(`^[a-z]+@[a-z]+\.[a-z]{2,}$`)

func IsEmailValid(email string) bool {
	return regexEmailFormat.MatchString(email)
}
