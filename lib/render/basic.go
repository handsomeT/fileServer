package render

import (
	"fmt"
	"strings"
)

func linkStr(baseUrl string, name string) string {
	format := "<a href=\"%s\">%s</a>"
	url := baseUrl + "/" + name
	return fmt.Sprintf(format, "http://" + strings.Replace(url, "//", "/", -1), name)
}