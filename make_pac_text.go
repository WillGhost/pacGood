package main
import (
    _ "fmt"
    "os"
    "bufio"
    "io"
    "strings"
    "text/template"
)


var path_direct string = "./data/direct"
var path_proxy string = "./data/p_site"


type doData struct {
    Direct map[string]int
    Proxy map[string]int
}


//{{$direct := .Direct}}
var temp_pac string  = `var mainland = {
{{- range $k, $_ := .Direct }}
  "{{$k}}": 1,
{{- end }}
}


var domains = {
{{- range $k, $_ := .Proxy }}
  "{{$k}}": 1,
{{- end }}
}


var proxy = "SOCKS5 127.0.0.1:2080; DIRECT;";
var direct = "DIRECT;";

function FindProxyForURL(url, host) {
    var lastPos;
    do {
        if (domains.hasOwnProperty(host)) {
            return proxy;
        }
		if (mainland.hasOwnProperty(host)) {
            return direct;
        }
        lastPos = host.indexOf('.') + 1;
        host = host.slice(lastPos);
    }
    while (lastPos >= 1);
    if (
        (-1 != host.lastIndexOf('cn')) ||
        (-1 != host.lastIndexOf('com')) ||
        (-1 != host.lastIndexOf('net')) ||
        (-1 != host.lastIndexOf('fm')) ||
        (-1 != host.search("[0-9]")) ||
        (-1 != host.lastIndexOf('gs'))
    ){
        return direct;
    }
    return proxy;
}
`



func doList(path string) map[string]int {
    dl := make(map[string]int)

    file, err := os.Open(path)
    if err != nil { panic(err) }
    defer file.Close()

    reader := bufio.NewReader(file)
    for {
        str, err := reader.ReadString('\n')
        str = strings.TrimSpace(str)
        if len(str) > 1 && string(str[0]) != "#" {
            dl[str] = 0
        }
        if err == io.EOF { break }
    }
    return dl
}


func main() {
    dl_d := doList(path_direct)
    dl_p := doList(path_proxy)

    dd := doData{dl_d, dl_p}
    tmpl, _ := template.New("_").Parse(temp_pac)
    tmpl.Execute(os.Stdout, dd)

}









