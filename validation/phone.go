package validation

import "regexp"

func isValidPhone(phone string) bool {
	rule := regexp.MustCompile(`^[+]?[(]?[0-9]{7,15}$`)
	return rule.MatchString(phone)
}
