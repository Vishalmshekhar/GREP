package main

import (
	"testing"
	"grep/regx"
	"github.com/stretchr/testify/assert"
) 

func Test_Search(t *testing.T)  {
	t.Run("Testing No flags,One file match present", Test_FlagZeroFileOne)
	t.Run("Testing No flags,More than one file match present", Test_FlagZeroFileMany)
    t.Run("Testing One or more Flags,But no match in file",Test_FlagWithNoMatch)
	t.Run("Testing One file,One flag {-n} match present", Test_FlagOneFileOne1)
    t.Run("Testing One file,One flag {-l} match present", Test_FlagOneFileOne2)
	t.Run("Testing One file,One flag {-i} match present", Test_FlagOneFileOne3)
	t.Run("Testing One file,One flag {-v} match present", Test_FlagOneFileOne4)
	t.Run("Testing One file,One flag {-x} match present", Test_FlagOneFileOne5)
	t.Run("Testing Many files,One flag{-n} match present", Test_FlagOneFileMany1)
	t.Run("Testing Many files,One flag{-l} match present", Test_FlagOneFileMany2)
	t.Run("Testing Many files,One flag{-i} match present", Test_FlagOneFileMany3)
	t.Run("Testing Many files,One flag{-v} match present", Test_FlagOneFileMany4)
	t.Run("Testing Many files,One flag{-x} match present", Test_FlagOneFileMany5)
	t.Run("Testing One file,Many flags{-v,-n} match present", Test_FlagManyFileOne1)
	t.Run("Testing One file,Many flags{-v,-n,-x} match present", Test_FlagManyFileOne2)
	t.Run("Testing One file,Many flags{-v,-n,-i} match present", Test_FlagManyFileOne3)
	t.Run("Testing One file,Many flags{-v,-n,-i,-x} match present", Test_FlagManyFileOne4)
	t.Run("Testing One file,Many flags{-v,-n,-i,-x,-l} match present", Test_FlagManyFileOne5)
	t.Run("Testing Many file,Many flags{-i,-n,-v} match present", Test_FlagManyFileMany1)
	t.Run("Testing Many file,Many flags{-i,-n,-x} match present", Test_FlagManyFileMany2)
	t.Run("Testing Many file,Many flags{-i,-v,-x} match present", Test_FlagManyFileMany3)
	t.Run("Testing Many file,Many flags{-i,-n,-v,-x} match present", Test_FlagManyFileMany4)
	t.Run("Testing Many file,Many flags{-i,-n,-v,-x,-l} match present", Test_FlagManyFileMany5)
}

func Test_FlagZeroFileOne(t *testing.T){
	result := Search("hello", []string{}, []string{"input.txt"})
	result2 := []string{"hello","hello again"}	
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}

func Test_FlagZeroFileMany(t *testing.T){
	result := Search("hello", []string{}, []string{"input.txt", "greeting.txt"})
	result2 := []string{"input.txt:hello","input.txt:hello again","greeting.txt:hello","greeting.txt:hello again"}	
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
} 
func Test_FlagWithNoMatch(t *testing.T){
	result := Search("hello", []string{}, []string{"another.txt"})
	result2 := []string{}	
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}

func Test_FlagOneFileOne1(t *testing.T){
	result := Search("hello", []string{"-n"}, []string{"input.txt"})
	result2 := []string{"1:hello","3:hello again"}
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}
func Test_FlagOneFileOne2(t *testing.T){
	result := Search("hello", []string{"-l"}, []string{"input.txt"})
	result2 := []string{"input.txt"}
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}
func Test_FlagOneFileOne3(t *testing.T){
	result := Search("hello", []string{"-i"}, []string{"input.txt"})
	result2 := []string{"hello","hello again","HELLO again"}
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}
func Test_FlagOneFileOne4(t *testing.T){
	result := Search("hello", []string{"-v"}, []string{"input.txt"})
	result2 := []string{"world","HELLO again"}
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}
func Test_FlagOneFileOne5(t *testing.T){
	result := Search("hello", []string{"-x"}, []string{"input.txt"})
	result2 := []string{"hello"}
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}
func Test_FlagOneFileMany1(t *testing.T){
	result := Search("hello", []string{"-n"}, []string{"input.txt", "greeting.txt"})
	result2 := []string{"input.txt:1:hello","input.txt:3:hello again","greeting.txt:1:hello","greeting.txt:2:hello again"}
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}
func Test_FlagOneFileMany2(t *testing.T){
	result := Search("hello", []string{"-l"}, []string{"input.txt", "greeting.txt"})
	result2 := []string{"input.txt","greeting.txt"}
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}
func Test_FlagOneFileMany3(t *testing.T){
	result := Search("hello", []string{"-i"}, []string{"input.txt", "greeting.txt"})
	result2 := []string{"input.txt:hello","input.txt:hello again","input.txt:HELLO again","greeting.txt:hello","greeting.txt:hello again"}
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}

func Test_FlagOneFileMany4(t *testing.T){
	result := Search("hello", []string{"-v"}, []string{"input.txt", "greeting.txt"})
	result2 := []string{"input.txt:world","input.txt:HELLO again","greeting.txt:world"}
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}
func Test_FlagOneFileMany5(t *testing.T){
	result := Search("hello", []string{"-x"}, []string{"input.txt", "greeting.txt"})
	result2 := []string{"input.txt:hello","greeting.txt:hello"}
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}

func Test_FlagManyFileOne1(t *testing.T){
	result := Search("hello", []string{"-v","-n"}, []string{"input.txt"})
	result2 := []string{"2:world","4:HELLO again"}
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}
func Test_FlagManyFileOne2(t *testing.T){
	result := Search("hello", []string{"-v","-n","-x"}, []string{"input.txt"})
	result2 := []string{"2:world","3:hello again","4:HELLO again"}
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}
func Test_FlagManyFileOne3(t *testing.T){
	result := Search("hello", []string{"-v","-n","-i"}, []string{"input.txt"})
	result2 := []string{"2:world"}
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}
func Test_FlagManyFileOne4(t *testing.T){
	result := Search("hello", []string{"-v","-n","-i","-x"}, []string{"input.txt"})
	result2 := []string{"2:world","3:hello again","4:HELLO again"}
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}
func Test_FlagManyFileOne5(t *testing.T){
	result := Search("hello", []string{"-v","-n","-i-","-x","-l"}, []string{"input.txt"})
	result2 := []string{"input.txt"}
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}

func Test_FlagManyFileMany1(t *testing.T){
	result := Search("hello", []string{"-i","-n","-v"}, []string{"input.txt", "greeting.txt"})
	result2 := []string{"input.txt:2:world","greeting.txt:3:world"}
	
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}
func Test_FlagManyFileMany2(t *testing.T){
	result := Search("hello", []string{"-i","-n","-x"}, []string{"input.txt", "greeting.txt"})
	result2 := []string{"input.txt:1:hello","greeting.txt:1:hello"}
	
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}
func Test_FlagManyFileMany3(t *testing.T){
	result := Search("hello", []string{"-v","-i-","-x"}, []string{"input.txt", "greeting.txt"})
	result2 := []string{"input.txt:world","input.txt:hello again","input.txt:HELLO again","greeting.txt:hello again","greeting.txt:world"}
	
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}
func Test_FlagManyFileMany4(t *testing.T){
	result := Search("hello", []string{"-i","-n","-v","-x"}, []string{"input.txt", "greeting.txt"})
	result2 := []string{"input.txt:2:world","input.txt:3:hello again","input.txt:4:HELLO again","greeting.txt:2:hello again","greeting.txt:3:world"}
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}
func Test_FlagManyFileMany5(t *testing.T){
	result := Search("hello", []string{"-i","-n","-v","-x","-l"}, []string{"input.txt", "greeting.txt"})
	result2 := []string{"input.txt", "greeting.txt"}
	
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}

