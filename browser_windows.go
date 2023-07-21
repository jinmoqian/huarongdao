package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

type browserType struct {
	browser string
	param   string
	params  []string
}

const msedge = "Microsoft Edge"
const chrome = "Google Chrome"
const firefx = "FIREFOX.EXE"
const operas = "OperaStable"

const key = "\\SOFTWARE\\Clients\\StartMenuInternet"

type NoProperBrowserError struct{}

func (e NoProperBrowserError) Error() string {
	return "No proper brower found. Chrome/Edge/Firefox/Opera are options."
}

func openUI(keepAlive bool, url string) (func(), error) {
	url = strings.Replace(url, "[::]", "localhost", -1)
	tmpDir := os.TempDir() + "\\huarongdao\\" + fmt.Sprintf("%d", time.Now().UnixNano())
	options := []browserType{
		{browser: msedge, param: "--app=", params: []string{"--user-data-dir=" + tmpDir, "--window-size=800,560", "--no-first-run"}},
		{browser: chrome, param: "--app=", params: []string{"--user-data-dir=" + tmpDir, "--window-size=800,560", "--no-first-run"}},
		{browser: firefx, param: "", params: []string{"-new-window", "-height 560", "-width 800", "-fullscreen"}},
		{browser: operas, param: "", params: []string{}},
	}
	var browser *browserType
	var shell string
	var err error
	var cleaner func()
	var oIndex, vIndex int
	browser, shell, err = searchBrowser(func(regValues []string) (bool, string, *browserType) {
		for oI := oIndex; oI < len(options); oI++ {
			for vI := vIndex; vI < len(regValues); vI++ {
				p1 := strings.LastIndex(regValues[vI], options[oI].browser)
				if p1 != -1 && p1 == len(regValues[vI])-len(options[oI].browser) {
					val := extractBrowserFromReg(regValues[vI])
					if val != "" {
						oIndex = oI
						vIndex = vI
						if vI == len(regValues)-1 {
							vIndex = 0
							oIndex++
						} else {
							vIndex++
						}
						return true, val, &(options[oI])
					}
				}
			}
			vIndex = 0
		}
		oIndex = 0
		return false, "", nil
	})
	if browser == nil {
		return nil, NoProperBrowserError{}
	}
	if err == nil {
		args := append([]string{browser.param + url}, browser.params...)
		err = exec.Command(shell, args...).Start()
		if err == nil {
			if browser.browser == msedge || browser.browser == chrome {
				cleaner = func() {
					os.Remove(tmpDir)
				}
			} else {
				cleaner = func() {}
			}
		}
	}
	return cleaner, err
}
func errorMessage(message string) {
	syscall.NewLazyDLL("user32.dll").NewProc("MessageBoxW").Call(
		uintptr(0),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(message))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Error"))),
		uintptr(0))
}
func main() {
	start()
}
