package main

import (
	"testing"
	"grep/regx"
	"github.com/stretchr/testify/assert"
) 

func Test_Search(t *testing.T)  {
	t.Run("Testing FlagZeroFileOne", Test_FlagZeroFileOne)
	t.Run("Testing FlagZeroFileMany", Test_FlagZeroFileMany)
	t.Run("Testing FlagOneFileOne", Test_FlagOneFileOne)
	t.Run("Testing FlagOneFileMany", Test_FlagOneFileMany)
	t.Run("Testing FlagManyFileOne", Test_FlagManyFileOne)
	t.Run("Testing FlagManyFileMany", Test_FlagManyFileMany)
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

func Test_FlagOneFileOne(t *testing.T){
	result := Search("hello", []string{"-n"}, []string{"input.txt"})
	result2 := []string{"1:hello","3:hello again"}
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}

func Test_FlagOneFileMany(t *testing.T){
	result := Search("hello", []string{"-v"}, []string{"input.txt", "greeting.txt"})
	result2 := []string{"input.txt:world","input.txt:HELLO again","greeting.txt:world"}
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}

func Test_FlagManyFileOne(t *testing.T){
	result := Search("hello", []string{"-v","-n"}, []string{"input.txt"})
	result2 := []string{"2:world","4:HELLO again"}
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}

func Test_FlagManyFileMany(t *testing.T){
	result := Search("hello", []string{"-i","-n","-v"}, []string{"input.txt", "greeting.txt"})
	result2 := []string{"input.txt:2:world","greeting.txt:3:world"}
	
	assert.True(t, regx.StringSlicesEqual(result, result2), "The string slices don't match")
}

