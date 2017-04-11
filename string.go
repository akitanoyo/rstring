package  rstring

import (
    "regexp"
)

//////////////////////////////////////////////////////////////
// String
type String string

func (str *String) Match(regex string) bool {
    r := regexp.MustCompile(regex)
	return r.MatchString(string(*str))
}	

func (str *String) ReplaceAll(regex, replace string) string {
    r := regexp.MustCompile(regex)
	*str = String(r.ReplaceAllString(string(*str), replace))
	return string(*str)
}

func (str *String) Split(regex string) []string {
    r := regexp.MustCompile(regex)
	result := r.Split(string(*str), -1)
	return result
}

func (str *String) Len() int {
	return len(*str)
}

func (str *String) String() string {
	return string(*str)
}

