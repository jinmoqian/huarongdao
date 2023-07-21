//go:build !windowsxp

package main

import (
	"strings"

	"golang.org/x/sys/windows/registry"
)

func queryBrowserList(rootKey registry.Key) ([]string, error) {
	newKey := strings.Replace(key, "\\", "", 1)
	k, err := registry.OpenKey(rootKey, newKey, registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS)
	if err != nil {
		return nil, err
	}
	defer k.Close()
	ki, err := k.Stat()
	if err != nil {
		return nil, err
	}
	values, err := k.ReadSubKeyNames(int(ki.SubKeyCount))
	if err != nil {
		return nil, err
	}
	return values, nil
}
func queryBrowserCmd(rootKey registry.Key, key, value string) (string, error) {
	newKey := strings.Replace(key, "\\", "", 1)
	k, err := registry.OpenKey(rootKey, newKey+"\\"+value, registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS)
	if err != nil {
		return "", err
	}
	defer k.Close()
	val, _, err := k.GetStringValue("")
	if err != nil {
		return "", err
	}
	val = strings.Trim(val, "\"")
	return val, nil
}
func searchBrowser(browserMatch func(regValues []string) (bool, string, *browserType)) (*browserType, string, error) {
	var err error
	for _, rootKey := range []registry.Key{registry.LOCAL_MACHINE, registry.CURRENT_USER} {
		var values []string
		values, err = queryBrowserList(rootKey)
		if err != nil {
			continue
		}
		for {
			ok, value, option := browserMatch(values)
			if !ok {
				break
			}
			var val string
			val, err = queryBrowserCmd(rootKey, key, value+"\\shell\\open\\command")
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
	return value
}
