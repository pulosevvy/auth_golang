package validation

import "regexp"

func IsGuid(guid string) bool {
	regex := regexp.MustCompile("[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}")
	return regex.MatchString(guid)
}
