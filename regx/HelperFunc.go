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
			if strings.Contains(Store[i].lines[j], pattern) {
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
func Inslice(n string, h []string) bool {
	
	for i:=0;i<len(h);i++{
		res2:=strings.Split(h[i],";")
		if(res2[0]==n){
			return true
		}
	}
	return false
}

func FlagManyFileOne(flags []string, Store File, pattern string)[]string {
	var Flagcheck = map[string]bool{"-n": false, "-l": false, "-i": false, "-v": false, "-x": false}
    
	for i:=0;i<len(flags);i++{
		Flagcheck[flags[i]]=true
	}
	res:=[]string{}
    temp1:=[]string{}
	temp2:=[]string{}
	if Flagcheck["-i"]{//case Insensi
		for j := 0; j < len(Store.lines); j++ {
			res2 := strings.Split(Store.lines[j], " ")
			for k := 0; k < len(res2); k++ {
				if strings.EqualFold(res2[k], pattern) {
					temp1 = append(temp1,Store.lines[j]+";"+fmt.Sprint(k+1))
					break
				}
			}
		}
	}
	if Flagcheck["-x"]&& !Flagcheck["-i"]{//case Insen + not strictly matching lines
		for j := 0; j < len(Store.lines); j++ {
			if strings.Compare(Store.lines[j], pattern) == 0 {
				temp1 = append(temp1, Store.lines[j]+";"+fmt.Sprint(j+1))
			}
		}
	}
	if Flagcheck["-x"] && Flagcheck["-i"]{//case Insen + strictly matching lines
		for i:=0;i<len(temp1);i++{
			res2:=strings.Split(temp1[i],";")
			if strings.Compare(res2[0],pattern)==0{
             temp2 = append(temp2,res2[0]+";"+res2[1])
			}
		}
		temp1=temp2
	}
	if !Flagcheck["-x"] && !Flagcheck["-i"]{//not both cases just find matching and store in temp1
		for j := 0; j < len(Store.lines); j++ {
			if strings.Contains(Store.lines[j], pattern) {
				temp1 = append(temp1, Store.lines[j]+";"+fmt.Sprint(j+1))
			}
		}
	}
	temp2=temp2[:0]
	if Flagcheck["-v"]{//inverse it and store in temp1
       
       for i:=0;i<len(Store.lines);i++{
		if !Inslice(Store.lines[i], temp1) {
			temp2 = append(temp2, Store.lines[i]+";"+fmt.Sprint(i+1))
		}
	}
	temp1=temp2
	}
    
	if Flagcheck["-n"] && !Flagcheck["-l"]{// line number included
       for i:=0;i<len(temp1);i++{
           res2:=strings.Split(temp1[i],";")
		   res=append(res,res2[1]+":"+res2[0])}
	   	}else if !Flagcheck["-n"] && Flagcheck["-l"]{
			if len(temp1)>0{
				res=append(res,Store.name)
				}
			}else if Flagcheck["-n"] && Flagcheck["-l"]{//only file name
				if len(temp1)>0{
					res=append(res,Store.name)
					}
			}else if !Flagcheck["-n"] && !Flagcheck["-l"]{//just empty temp
				for i:=0;i<len(temp1);i++{
					res2:=strings.Split(temp1[i],";")
					res=append(res,res2[0])
				}
			}
	return res
}

func FlagManyFileMany(flags []string, Store []File, pattern string)[]string{
    res:=[]string{}
	for i:=0;i<len(Store);i++{
		tempHere:=FlagManyFileOne(flags,Store[i],pattern)

		if len(tempHere)==1 && tempHere[0]==Store[i].name{
            res=append(res,Store[i].name)
		}else{
			for j:=0;j<len(tempHere);j++{
				res=append(res,Store[i].name+":"+tempHere[j])
			}
		}

	}
	return res
}
func StringSlicesEqual(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}
