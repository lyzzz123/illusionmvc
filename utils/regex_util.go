package utils

import "regexp"

var handlerPattern = regexp.MustCompile("{([a-zA-Z\\d]+)}")

func HasPathValue(path string) bool {
	return handlerPattern.MatchString(path)
}

func CreatePathValueRegex(path string) *regexp.Regexp {
	pathValuePattern := "{[a-zA-Z\\d]+}"
	pathValuePatternRegex := regexp.MustCompile(pathValuePattern)
	replacedPathValuePattern := pathValuePatternRegex.ReplaceAllString(path, "([a-zA-Z\\d]+)")
	return regexp.MustCompile("^" + replacedPathValuePattern + "$")
}
