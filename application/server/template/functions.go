package template

import (
	"fmt"
	tpl "html/template"
	"strings"
)

func AddFunction(a, b int) int {
	return a + b
}

func AsHTML(a string) tpl.HTML {
	return tpl.HTML(a)
}

var numbers = map[string]string{
	"۱": "1",
	"۲": "2",
	"۳": "3",
	"۴": "4",
	"۵": "5",
	"۶": "6",
	"۷": "7",
	"۸": "8",
	"۹": "9",
	"۰": "0",
}

func ToPersianNumber(number int) string {
	a := fmt.Sprintf("%d", number)
	for key, value := range numbers {
		a = strings.Replace(a, value, key, -1)
	}
	return a
}

func GetFirst(a string) string {
	parts := strings.Split(a, "<br/><br/>")
	return parts[0]
}

func GetType(a int) string {
	if a == 1 {
		return "ghazal"
	} else if a == 2 {
		return "robaei"
	}
	return ""
}
