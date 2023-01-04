package main

import(
	"grep/regx"
)
func main() {
	files := []string{"input.txt"}
    flags := []string{"-x"}
    for _, line := range Search("hello",flags, files) {
       println(line)
	}
}

func Search(pattern string, flags, files []string) []string {
	Store := []regx.File{}
	res := []string{}
	//var x = make(map[string]int)
	for i := 0; i < len(files); i++ {
		Store = append(Store, regx.ReadFile(files[i]))
	}

	/*for i := 0; i < len(Store); i++ {
		x[Store[i].name] = i
	}*/

	if len(flags) == 0 {//No Flag
		if len(files) == 1 {//One File
			res=regx.FlagZeroFileOne(Store,pattern)
		} else { //Many File
			res=regx.FlagZeroFileMany(Store,pattern) 
		}
	}else if len(flags) == 1 { //One Flag
		if len(files) == 1 { //One File
			res=regx.FlagOneFileOne(flags[0],Store,pattern)
		} else { //Many File
			res=regx.FlagOneFileMany(flags[0],Store,pattern)
		}
	}
	return res
}







