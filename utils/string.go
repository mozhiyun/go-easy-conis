package utils

import (
	"math/rand"
	"strconv"
	"time"
)

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyz")

func init() {
	rand.Seed(time.Now().UnixNano())
}

//GetRandomString ...
func GetRandomString(l int) string {
	b := make([]rune, l)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Bool2String boolean to string
func Bool2String(b bool) string {
	return strconv.FormatBool(b)
}

// Int2String Integer to string
func Int2String(i int) string {
	return strconv.Itoa(i)
}
