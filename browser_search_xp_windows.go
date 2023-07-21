//go:build windowsxp

package main

import (
	"os/exec"
	"strings"
)

func queryBrowserList(key string) ([]string, error) {
	cmd := exec.Command("reg", "query", key)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	values := strings.Split(string(output), "\r\n")
	ret := make([]string, 0)
	for _, value := range values {
		vI := strings.Index(value, "\\")
		if vI != -1 {
			vK := strings.Index(key, "\\")
			if vK != -1 {
				if strings.Contains(string(([]byte(value))[vI:]), string(([]byte(key))[vI:])+"\\") {
					ret = append(ret, value)
				}
			}
		}
	}
	return ret, nil
}
func queryBrowserCmd(key string) (string, error) {
	cmd := exec.Command("reg", "query", key, "/ve")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	values := strings.Split(string(output), "\r\n")
	for _, value := range values {
		if strings.Contains(value, "REG_SZ") {
			firstComma := strings.Index(value, "\"")
			lastComma := strings.LastIndex(value, "\"")
			if firstComma != -1 && lastComma != -1 && firstComma != lastComma {
				val := string([]byte(value)[firstComma+1 : lastComma])
				return val, nil
			}
		}
	}
	return "", nil
}
func searchBrowser(browserMatch func(regValues []string) (bool, string, *browserType)) (*browserType, string, error) {
	var err error
	for _, rootKey := range []string{"hklm", "hkcu"} {
		var values []string
		values, err = queryBrowserList(rootKey + key)
		if err != nil {
			continue
		}
		for {
			ok, value, option := browserMatch(values)
			if !ok {
				break
			}
			var val string
			val, err = queryBrowserCmd(rootKey + key + "\\" + value + "\\shell\\open\\command")
			if err != nil {
				continue
			}
			if val != "" {
				return option, val, nil
			}
		}
	}
	if err == nil {
		err = NoProperBrowserError{}
	}
	return nil, "", err
}
func extractBrowserFromReg(value string) string {
	lastComma := strings.LastIndex(value, "\\")
	if lastComma != -1 && lastComma < len(value)-1 {
		return string([]byte(value)[lastComma+1:])
	}
	return ""
}
