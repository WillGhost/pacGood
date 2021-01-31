package main

import (
	"bufio"
	_ "fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

var path_proxy string = "./data/p_site"

type doData struct {
	Direct map[string]int
	Proxy  map[string]int
}

//{{$direct := .Direct}}
var temp_pac string = `payload:
{{- range $k, $_ := .Proxy }}
  - '+.{{$k}}'
{{- end }}
`

func doList(path string) map[string]int {
	dl := make(map[string]int)

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		if len(str) > 1 && string(str[0]) != "#" {
			dl[str] = 0
		}
		if err == io.EOF {
			break
		}
	}
	return dl
}

func main() {
	dl_p := doList(path_proxy)

	dd := doData{nil, dl_p}
	tmpl, _ := template.New("_").Parse(temp_pac)
	tmpl.Execute(os.Stdout, dd)

}
