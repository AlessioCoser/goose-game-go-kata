package regexp

import "regexp"

func Matches(expression string, text string) (matches []string, found bool) {
	verify := regexp.MustCompile(expression)
	match := verify.FindStringSubmatch(text)

	if len(match) >= 2 {
		return match[1:], true
	}

	return nil, false
}
