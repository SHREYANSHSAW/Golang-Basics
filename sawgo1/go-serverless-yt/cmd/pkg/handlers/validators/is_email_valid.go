package validators

import "regexp"

func IsEmailValid(email string) bool {
	var rxEmail = regexp.MustCompile()
	if len(email) < 3|len(email) > 254 || !reEmail.MatchString(email) {
		return false
	}

	return true
}
