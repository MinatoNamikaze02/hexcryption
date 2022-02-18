package se

import (
	"fmt"
	"math/rand"
	"strings"
	"strconv"
	"time"
)

func Encoder(inp string) (string, string){
	var key string
	rand.Seed(time.Now().UnixNano())
	temp := []byte(inp)
	// fmt.Println(temp)
	for i:=0; i < len(temp); i++{
		innerkey := ""
		n := rand.Intn(10-1) + 1
		for j:=0; j<n; j++{
			x := rand.Intn(10-1) + 1
			temp[i] = temp[i] + byte(x)
			innerkey += fmt.Sprint(x)
		}
		key += fmt.Sprint(innerkey, "-")

	}
	var final string
	for _, s := range temp{
		final += fmt.Sprint(int(s), "-")
		
	}
	final = strings.TrimSuffix(final, "-")
	key = strings.TrimSuffix(key, "-")
	return final, key
}

func Decoder(inp string, key string) (string) {
	count := 1
	for i:=0; i<len(inp); i++{
		if string(inp[i]) == `-`{
			count++
		}
	}
	final := make([]byte, count)
	var tempKey string
	var tempDec string
	for j:=1; j<=count; j++{
		tempKey = findPartialString(key, j)
		tempDec = findPartialString(inp, j)
		final[j-1] = decode(tempKey, tempDec)
	}
	decrypted := string(final[:])
	return decrypted
}

func findPartialString(key string, j int)(string){
	count := 0
	var temp string
	for _, s := range key{
		if string(s) == `-`{
			count++
			if count == j{
				break
			}else{
				temp = ""
			}
		}else{
			temp += string(s)
		}	
	}
	return temp
}

func decode(key string, dec string)(byte){
	ksize := len(key)
	tempKey := make([]int, ksize)
	val, _ := strconv.Atoi(dec)
	for i, s := range key{
		tempKey[i], _ = strconv.Atoi(string(s))
	}
	for j:= 0; j<len(tempKey); j++{
		val = val - tempKey[j]
	}	
	bt := byte(val)
	return bt
}