package regex

import (
	"fmt"
	"regexp"
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
