package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"testing"
)

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func TestValue(t *testing.T) {
	var b = [4]int{12, 23, 43}

	b = [4]int{1: 10, 2: 30, 0: 20}
	for i, i2 := range b {
		fmt.Printf("变量内容 %d %d  %x \n", i, i2, &i2)

	}

	var book = Books{"title", "author", "subject", 123}

	fmt.Println(book)
	book.author = "zhangsan"
	fmt.Println(book)

	var slice1 = make([]int, 1, 5)

	fmt.Println(slice1)
}
func TestString(t *testing.T) {
	// 如果直接截取，可能会乱码
	name := "尝试截取中文字符串"
	var splitName = strings.Split(name, "")

	fmt.Printf("len(splitName): %v\n", len(splitName))
	fmt.Printf("len(splitName): %v\n", len(name))

	str2 := name[0:5]

	fmt.Printf("str2: %v\n", str2)

	nameRune := []rune(name)

	fmt.Printf("len(nameRune): %v\n", len(nameRune))

	fmt.Printf("string(nameRune[0:5]): %v\n", string(nameRune[0:5]))

}

func TestConvert(t *testing.T) {
	str := "123456789ABC"
	strRev := Reversal(str)
	fmt.Printf("str: %v\n", str)
	fmt.Printf("strRev: %v\n", strRev)

	test := "i,love,go"

	str2 := test

	log.Print()

	keywordSlice := strings.Split(test, ",")

	for _, v := range keywordSlice {
		reg := regexp.MustCompile("(?i)" + v)
		str2 = reg.ReplaceAllString(str2, strings.ToUpper(v))

		fmt.Printf("str2: %v\n", str2)

	}

}

func Reversal(a string) (re string) {

	b := []rune(a)

	for i := 0; i < len(b)/2; i++ {
		// swap
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
	re = string(b)
	return
}

type User struct {
	name string
	auth string
}

type Admin struct {
	User
}

func (u User) introduceMe() {
	fmt.Println("hello, my name is " + u.name)
}

func TestObject(t *testing.T) {
	u := User{"momo", "auth"}
	u.introduceMe()

	u1 := User{"momo1", "auth2"}
	u2 := User{"momo2", "auth3"}
	fmt.Println(u1 == u2)
	fmt.Println(&u1 == &u2)

	var f func()

	// 这两种也是经常调用方法的方式
	f = u.introduceMe

	f()

	f2 := (User).introduceMe

	f2(u)

	admin := Admin{}
	admin.User = u1

	fmt.Println(admin.name)

}
