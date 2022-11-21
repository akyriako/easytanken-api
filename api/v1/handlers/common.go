package handlers

import "strings"

func replaceLicense(responseBody string) string {
	replaced := strings.Replace(string(responseBody), "tankerkoenig", "rhynosaur", -1)
	replaced = strings.Replace(replaced, "creativecommons", "easytanken", -1)

	return replaced
}
