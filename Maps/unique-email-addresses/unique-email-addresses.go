package uniqueemailaddresses

import (
	"strings"
)

func numUniqueEmails(emails []string) int {
	dict := make(map[string]int)
	for index := range emails {
		if _, ok := dict[transform(emails[index])]; !ok {
			dict[transform(emails[index])] = 1
		}
	}
	return len(dict)
}

func transform(email string) string {
	local := email[:strings.Index(email, "@")]
	if strings.Contains(email, "+") {
		local = email[:strings.Index(local, "+")]
	}
	local = strings.Replace(local, ".", "", -1)
	result := local + email[strings.Index(email, "@"):len(email)]
	return result
}
