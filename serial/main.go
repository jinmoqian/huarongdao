package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		panic(os.Args)
	}
	serial(os.Args[1:])
	// serial([]string{"../statics/game.js"})
}

func serial(ins []string) {
	sort.Slice(ins, func(i, j int) bool {
		return ins[i] < ins[j]
	})
	varnames := make(map[string]string)
	contents := make(map[string]string)
	for _, in := range ins {
		var binary bool
		if strings.Index(in, "+") == 0 {
			binary = true
			in = string(([]byte(in))[1:])
		}
		ih, err := os.OpenFile(in, os.O_RDONLY, 0644)
		if err != nil {
			panic(err)
		}
		varname := strings.Replace(in, "/", "_slash_", -1)
		varname = strings.Replace(varname, ".", "_dot_", -1)
		const bufsize = 4096
		buf := make([]byte, bufsize)
		outBuf := strings.Builder{}
		var all int
		outBuf.WriteString("package generated\nconst " + varname + " = \"")
		for {
			n, err := ih.Read(buf)
			if err != nil && err != io.EOF {
				panic(err)
			}
			for i := 0; i < n; i++ {
				if 32 <= buf[i] && buf[i] <= 126 && buf[i] != 34 && buf[i] != 92 {
					err = outBuf.WriteByte(buf[i])
				} else {
					if !binary && buf[i] == '\r' {
						continue
					}
					v := fmt.Sprintf("%02x", buf[i])
					v = "\\x" + strings.ToUpper(v)
					_, err = outBuf.WriteString(v)
				}
				if err != nil {
					panic(err)
				}
			}
			all += n
			if n < bufsize {
				break
			}
		}
		outBuf.WriteString("\"\n")
		contents[varname] = outBuf.String()
		err = ih.Close()
		if err != nil {
			panic(err)
		}
		fmt.Printf("file[%s], all bytes[%d]\n", in, all)
		varnames[in] = varname
	}
	for varname, content := range contents {
		out := "generated/" + varname + "_content.go"
		oh, err := os.OpenFile(out, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			panic(err)
		}
		_, err = oh.WriteString(content)
		if err != nil {
			panic(err)
		}
		err = oh.Close()
		if err != nil {
			panic(err)
		}
	}
	out := "generated/contents.go"
	oh, err := os.OpenFile(out, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	_, err = oh.WriteString("package generated\nvar Contents = map[string]string{\n")
	if err != nil {
		panic(err)
	}
	for in, varname := range varnames {
		_, err = oh.WriteString("\"" + in + "\":" + varname + ",\n")
		if err != nil {
			panic(err)
		}
	}
	_, err = oh.WriteString("}\n")
	if err != nil {
		panic(err)
	}
	err = oh.Close()
	if err != nil {
		panic(err)
	}
}
