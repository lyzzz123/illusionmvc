package utils

import "regexp"

var handlerPattern = regexp.MustCompile("{[^/]+}")

func HasPathValue(path string) bool {
	return handlerPattern.MatchString(path)
}

func CreateReplacedPath(path string) string {
	pathValuePattern := "{[^/]+}"
	pathValuePatternRegex := regexp.MustCompile(pathValuePattern)
	replacedPath := pathValuePatternRegex.ReplaceAllString(path, "*")
	return replacedPath
}
