package helperfunc

import (
	"fmt"
	"strings"
)

func addstring(str1 string, str2 string) string {
	return fmt.Sprintf("%s%s", str1, str2)
}
func ArrArr2String(array [][]float64) (str string) {
	var strarr [][]string
	var tstr []string
	str = ""
	for i, v := range array {
		for j, _ := range v {
			strarr[i][j] = fmt.Sprintf("%f", array[i][j])
		}
	}
	for i, v := range strarr {
		tstr[i] = fmt.Sprintf("[%s]", strings.Join(v, ","))
	}
	str = fmt.Sprintf("[%s]", strings.Join(tstr, ","))
	return str
}
