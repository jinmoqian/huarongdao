package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"huarongdao/game"
	"huarongdao/generated"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"runtime"
	"strings"
	"sync/atomic"
	"time"
	"unsafe"
)

var serverMode *bool
var serverPort *string

func init() {
	serverMode = flag.Bool("s", false, "Server mode flag")
	serverPort = flag.String("a", "localhost:0", "addr:port. eg. :4321, localhost:4321")
	flag.Parse()
}

func getStaticContentString(key string) string {
	return generated.Contents[key]
}
func getStaticContent(key string) []byte {
	dataStr := getStaticContentString(key)
	dataStrHeader := (*reflect.StringHeader)(unsafe.Pointer(&dataStr))
	var data []byte
	dataHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	dataHeader.Data = dataStrHeader.Data
	dataHeader.Cap = dataStrHeader.Len
	dataHeader.Len = dataStrHeader.Len
	return data
}

type picker struct {
	fh *game.FileHead
	bh game.BoardHash
}

func work() error {
	var data []byte = getStaticContent("data/data")

	df := &game.DataFile{}
	err := df.OpenBytes(data)
	if err != nil {
		return err
	}

	const randomSize = 1000
	easies := make([]picker, 0)
	mediums := make([]picker, 0)
	hards := make([]picker, 0)
	pickerCh := make(chan struct {
		level int
		ch    chan picker
	})
	go func() {
		for _, level := range []struct {
			level  int
			depth  []byte
			picker *[]picker
		}{{level: 4, picker: &easies}, {level: 5, depth: []byte{0, 100}, picker: &mediums}, {level: 5, depth: []byte{101, 0}, picker: &hards}} {
			for i := 0; i <= level.level; i++ {
				fh := game.FileHead{Hor: i, Ver: level.level - i, Little: 20 - 4 - level.level*2 - 2}
				bhDepths := df.All(fh)
				for bh, d := range bhDepths {
					if level.depth != nil {
						if level.depth[0] != 0 && d < level.depth[0] {
							continue
						}
						if level.depth[1] != 0 && d > level.depth[1] {
							continue
						}
					}
					*level.picker = append(*level.picker, picker{fh: &fh, bh: bh})
				}
			}
			rand.Shuffle(len(*level.picker), func(i, j int) {
				(*level.picker)[i], (*level.picker)[j] = (*level.picker)[j], (*level.picker)[i]
			})
			if len(*level.picker) > randomSize {
				*level.picker = (*level.picker)[0:randomSize]
			}
		}
		var levelMap = make(map[int][]picker)
		levelMap[0] = easies
		levelMap[1] = mediums
		levelMap[2] = hards
		for {
			p, ok := <-pickerCh
			if !ok {
				return
			}
			p.ch <- levelMap[p.level][rand.Int31n(int32(len(levelMap[p.level])))]
			close(p.ch)
		}
	}()

	http.HandleFunc("/ok", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	})
	outputBoard := func(level int, w http.ResponseWriter) {
		ch := make(chan picker)
		pickerCh <- struct {
			level int
			ch    chan picker
		}{level: level, ch: ch}
		pk := <-ch
		g := game.New(pk.fh.Hor, pk.fh.Ver, pk.fh.Little)
		b := g.EmptyBoard()
		g.UnBoardHash(pk.bh, b)
		cs := b.Cells()
		sb := bytes.NewBuffer(make([]byte, 0))
		sb.WriteString("[[")
		sb.WriteString(fmt.Sprintf("%d", cs[0]))
		sb.WriteByte(']')
		var idx = 1
		for _, n := range []int{pk.fh.Hor, pk.fh.Ver, pk.fh.Little} {
			sb.WriteString(",[")
			for i := 0; i < n; i++ {
				if i != 0 {
					sb.WriteByte(',')
				}
				sb.WriteString(fmt.Sprintf("%d", cs[idx]))
				idx++
			}
			sb.WriteByte(']')
		}
		sb.WriteByte(']')
		w.WriteHeader(200)
		w.Write(sb.Bytes())
	}
	var easy = func(w http.ResponseWriter) { outputBoard(0, w) }
	var medium = func(w http.ResponseWriter) { outputBoard(1, w) }
	var hard = func(w http.ResponseWriter) { outputBoard(2, w) }
	var solve = func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			if err != nil {
				outputError(err, w)
			}
		}()

		var q string
		q, err = url.QueryUnescape(r.RequestURI)
		if err != nil {
			return
		}
		qm := strings.Index(q, "?")
		if qm == -1 {
			err = fmt.Errorf("query param error[%s]", q)
			return
		}
		var p [][]byte
		err = json.Unmarshal([]byte(q)[qm+1:], &p)
		if err != nil {
			return
		}
		if len(p) != 4 {
			err = fmt.Errorf("query param error type 2[%s]", q)
			return
		}
		var sizes = [][2]int{{2, 2}, {2, 1}, {1, 2}, {1, 1}}
		g := game.New(len(p[1]), len(p[2]), len(p[3]))
		b := g.EmptyBoard()
		for idx, cells := range p {
			size := sizes[idx]
			for _, cell := range cells {
				if idx == 0 {
					b.Put(size[0], size[1], cell)
				} else {
					if !b.Can(size[0], size[1], cell) {
						err = fmt.Errorf("cannot put %d", cell)
						return
					}
					b.Put(size[0], size[1], cell)
				}
			}
		}
		win := b.Win()
		ret := struct {
			Win  bool      `json:"win"`
			Path [][2]byte `json:"path"`
		}{
			Win: win,
		}
		if !win {
			var path []*game.Board
			okchan := make(chan struct{})
			go func() {
				path = df.Search(game.FileHead{Hor: len(p[1]), Ver: len(p[2]), Little: len(p[3])}, b)
				close(okchan)
			}()
			<-okchan
			paths := make([][2]byte, 0)
			for idx, step := range path {
				if idx != 0 {
					u := make(map[byte]struct{})
					for idx, p := range path[idx-1].Cells() {
						if idx >= g.Pieces() {
							break
						}
						u[p] = struct{}{}
					}
					var dst byte
					for idx, p := range step.Cells() {
						if idx >= g.Pieces() {
							break
						}
						if _, ok := u[p]; ok {
							delete(u, p)
						} else {
							dst = p
						}
					}
					if len(u) != 1 {
						panic(fmt.Sprintf("More than 1 piece moved from [%v] to [%v]", path[idx-1].Cells(), step.Cells()))
					}
					var src byte
					for src = range u {
						break
					}
					paths = append(paths, [2]byte{game.Value(src), game.Value(dst)})
				}
			}
			ret.Path = paths
		}
		var outStr []byte
		outStr, err = json.Marshal(ret)
		if err != nil {
			return
		}
		outputString(string(outStr), w)
	}
	var shutdownFlag int32
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var s = r.URL.Path
		switch s {
		case "/":
			outputFile("statics/index.html", "text/html", w)
		case "/easy":
			easy(w)
		case "/medium":
			medium(w)
		case "/hard":
			hard(w)
		case "/solve":
			solve(w, r)
			runtime.GC()
		case "/keepalive":
			atomic.StoreInt32(&shutdownFlag, 0)
			w.WriteHeader(200)
			w.Write([]byte("<html><head>huarongdao</head><body>keep alive ok</body></html>"))
		default:
			outputFile(string(([]byte(r.URL.String()))[1:]), "", w)
		}
	})
	svr := &http.Server{
		Addr: *serverPort,
	}
	l, err := net.Listen("tcp", *serverPort)
	if err != nil {
		return err
	}
	url := "http://" + l.Addr().String()
	clearner, err := openUI(!*serverMode, url)
	if !*serverMode && err != nil {
		return err
	}
	defer clearner()
	if !*serverMode {
		go func() {
			for {
				atomic.StoreInt32(&shutdownFlag, 1)
				time.Sleep(keepAliveTimeout)
				if atomic.LoadInt32(&shutdownFlag) == 1 {
					svr.Shutdown(context.Background())
					return
				}
			}
		}()
	} else {
		fmt.Println("Server addr:", url)
	}
	err = svr.Serve(l)
	if err != nil {
		if err.Error() == "http: Server closed" {
			return nil
		}
		return err
	}
	return nil
}
func outputString(str string, w http.ResponseWriter) {
	w.WriteHeader(200)
	d := []byte(str)
	for {
		n, err := w.Write(d)
		if err != nil {
			outputError(err, w)
		} else if n < len(d) {
			d = d[n:]
		} else {
			return
		}
	}
}
func outputError(err error, w http.ResponseWriter) {
	w.WriteHeader(500)
	w.Write([]byte(err.Error()))
}

var vars map[string]string

func init() {
	vars = map[string]string{
		"keepAlive": fmt.Sprintf("%t", !*serverMode),
	}
}

func outputFile(filename string, contentType string, w http.ResponseWriter) {
	if filename == "favicon.ico" {
		filename = "statics/" + filename
	}
	buf := getStaticContentString(filename)
	if contentType == "" {
		idx := strings.LastIndex(filename, ".")
		if idx == -1 {
			contentType = "text/plain"
		} else {
			switch string([]byte(filename)[idx:]) {
			case ".js":
				contentType = "text/javascript"
			case ".ico":
				contentType = "image/ico"
			case ".png":
				contentType = "image/png"
			default:
				contentType = "text/plain"
			}
		}
	}
	if contentType == "text/plain" || contentType == "text/html" || contentType == "text/javascript" {
		for v := range vars {
			buf = strings.Replace(buf, "{"+v+"}", vars[v], -1)
		}
	}
	w.Header().Add("Content-Type", contentType)
	outputString(buf, w)
}
