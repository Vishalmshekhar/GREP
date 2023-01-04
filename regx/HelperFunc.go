package regx

import (
	"fmt"
	"strings"
)
func FlagZeroFileOne(Store []File, pattern string) []string {
	res := []string{}
	for j := 0; j < len(Store[0].lines); j++ {
		if strings.Contains(Store[0].lines[j], pattern) {
			res = append(res,Store[0].lines[j])
		}
	}
	return res
}
func FlagZeroFileMany(Store []File, pattern string) []string {
	res := []string{}
	for i := 0; i < len(Store); i++ {
		for j := 0; j < len(Store[i].lines); j++ {
			if strings.Contains(Store[0].lines[j], pattern) {
				res = append(res, Store[i].name+":"+Store[i].lines[j])
			}
		}
	}
	return res
}

func FlagOneFileOne(flag string, Store []File, pattern string) []string {
	res := []string{}
	switch flag {
	case "-n":
		for j := 0; j < len(Store[0].lines); j++ {
			if strings.Contains(Store[0].lines[j], pattern) {
				res = append(res,fmt.Sprint(j+1)+":"+Store[0].lines[j])
			}
		}
	case "-l":
		for j := 0; j < len(Store[0].lines); j++ {
			if strings.Contains(Store[0].lines[j], pattern) {
				res = append(res, Store[0].name)
				break
			}
		}
	case "-i":
		for j := 0; j < len(Store[0].lines); j++ {
			res2 := strings.Split(Store[0].lines[j], " ")
			for k := 0; k < len(res2); k++ {
				if strings.EqualFold(res2[k], pattern) {
					res = append(res, Store[0].name+":"+Store[0].lines[j])
					break
				}
			}
		}
	case "-v":
		for j := 0; j < len(Store[0].lines); j++ {
			if !strings.Contains(Store[0].lines[j], pattern) {
				res = append(res, Store[0].name+":"+Store[0].lines[j])
			}
		}
	case "-x":
		for j := 0; j < len(Store[0].lines); j++ {
			if strings.Compare(Store[0].lines[j], pattern) == 0 {
				res = append(res, Store[0].name+":"+Store[0].lines[j])
			}
		}
	}
	return res
}
func FlagOneFileMany(flag string, Store []File, pattern string) []string {
	res := []string{}
	switch flag {
	case "-n":
		for i := 0; i < len(Store); i++ {
			for j := 0; j < len(Store[i].lines); j++ {
				if strings.Contains(Store[i].lines[j], pattern) {
					res = append(res, Store[i].name+":"+fmt.Sprint(j+1)+":"+Store[i].lines[j])
				}
			}
		}
	case "-l":
		for i := 0; i < len(Store); i++ {
			for j := 0; j < len(Store[i].lines); j++ {
				if strings.Contains(Store[i].lines[j], pattern) {
					res = append(res, Store[i].name)
					break
				}
			}
		}
	case "-i":
		for i := 0; i < len(Store); i++ {
			for j := 0; j < len(Store[i].lines); j++ {
				res2 := strings.Split(Store[i].lines[j], " ")
				for k := 0; k < len(res2); k++ {
					if strings.EqualFold(res2[k], pattern) {
						res = append(res, Store[i].name+":"+Store[i].lines[j])
						break
					}
				}
			}
		}
	case "-v":
		for i := 0; i < len(Store); i++ {
			for j := 0; j < len(Store[i].lines); j++ {
				if !strings.Contains(Store[i].lines[j], pattern) {
					res = append(res, Store[i].name+":"+Store[i].lines[j])
				}
			}
		}
	case "-x":
		for i := 0; i < len(Store); i++ {
			for j := 0; j < len(Store[i].lines); j++ {
				if strings.Compare(Store[i].lines[j], pattern) == 0 {
					res = append(res, Store[i].name+":"+Store[i].lines[j])
				}
			}
		}
	}
	return res
}
