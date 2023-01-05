package main

import(
	"grep/regx"
)
func main() {
	files := []string{"another.txt"}
    flags := []string{"-n"}
    for _, line := range Search("hello",flags, files) {
       println(line)
	}
}

func Search(pattern string, flags, files []string) []string {
	Store := []regx.File{}
	res := []string{}
	
	for i := 0; i < len(files); i++ {
		Store = append(Store, regx.ReadFile(files[i]))
	}

	
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
	}else{
		if len(files)==1{//FlagManyFileOne
			res=regx.FlagManyFileOne(flags,Store[0],pattern)
		}else{
			 res=regx.FlagManyFileMany(flags,Store,pattern)
		}
	}
	return res
}








