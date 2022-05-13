package regex

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Test_string(thing string) ([]string, error) {
	// to catch the desired regex
	expression, err := regexp.Compile(`/datas/[0-9]`)
	if err != nil {
		return nil, fmt.Errorf("the expression %v is not valid %v", expression, err)
	}
	final_text := expression.FindAllString(thing, -1)
	return final_text, nil
}

func Find_int(str string) int {
	exp, _ := regexp.Compile(`[0-9]`)
	result := exp.FindAllString(str, -1)
	num, _ := strconv.Atoi(strings.Join(result, ""))
	return num
}