package  rstring

import (
    "regexp"
    "fmt"
)

//////////////////////////////////////////////////////////////
// String
type String string

func (str *String) Match(regex string) bool {
    r := regexp.MustCompile(regex)
	return r.MatchString(string(*str))
}	

// regex string
func submatch(mlist [][]string) []string {
    res := []string{}
    // res[0] = regex
    for _,m := range mlist {
        res = append(res, m[1])
    }
    return res
}


func (str *String) MatchList(regex string) ([]string, bool) {
    res := [][]string{}
    r := regexp.MustCompile(regex)
    ok := r.MatchString(string(*str))
    if ok {
        res = r.FindAllStringSubmatch(string(*str), -1)
    }
	return submatch(res), ok
}	

func (str *String) ReplaceAll(regex, replace string) string {
    r := regexp.MustCompile(regex)
	*str = String(r.ReplaceAllString(string(*str), replace))
	return string(*str)
}

func (str *String) ReplaceAllList(regex, replace string) (string, []string) {
    res := [][]string{}
    r := regexp.MustCompile(regex)
	s := r.ReplaceAllString(string(*str), replace)
    res = r.FindAllStringSubmatch(string(*str), -1)
    *str = String(s)
    return s, submatch(res)
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

func Dump(v interface{}) {
    f := fmt.Sprintf("%#v\n", v)
    rf := String(f)
    rf.ReplaceAll("[{}]", " \n")
    rf.ReplaceAll("[,] ", "\n")
    rf.ReplaceAll("\n", "\n   ")

    fmt.Println(rf.String())
}
