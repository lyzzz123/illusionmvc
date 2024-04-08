package utils

import "regexp"

var handlerPattern = regexp.MustCompile("{[^/]+}")

func HasPathValue(path string) bool {
	return handlerPattern.MatchString(path)
}

func CreatePathValueRegex(path string) *regexp.Regexp {
	//pathValuePattern := "{[a-zA-Z\\d]+}"
	pathValuePattern := "{[^/]+}"
	pathValuePatternRegex := regexp.MustCompile(pathValuePattern)
	//replacedPathValuePattern := pathValuePatternRegex.ReplaceAllString(path, "([a-zA-Z\\d]+)")
	replacedPathValuePattern := pathValuePatternRegex.ReplaceAllString(path, "([^/]+)")
	return regexp.MustCompile("^" + replacedPathValuePattern + "$")
}
