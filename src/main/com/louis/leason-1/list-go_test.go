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

func TestInt(t *testing.T) {
	var num01 int = 0b1100
	var num02 int = 0o14
	var num03 int = 0xC

	fmt.Printf("2进制数 %b 表示的是: %d \n", num01, num01)
	fmt.Printf("8进制数 %o 表示的是: %d \n", num02, num02)
	fmt.Printf("16进制数 %X 表示的是: %d \n", num03, num03)

	var a byte = 65
	var b uint8 = 66
	fmt.Printf("a的值：%c \nb 的值：%c\n", a, b)
	var country string = "hello,中国"
	fmt.Println(len(country))

	var mystr01 string = `\r\n`
	fmt.Print(`\r\n`)
	fmt.Printf("的解释型字符串是： %q", mystr01)

	type arr3 [3]int

	myarry := arr3{1, 2, 3}
	fmt.Printf("%d 的类型是 %T", myarry, myarry)

}

type Profile struct {
	name   string
	age    int
	gender string
	mother *Profile
	father *Profile
}

type company struct {
	name        string
	companyAddr string
}

type staff struct {
	name string
	age  int
	company
}

type Say interface {
	/**
	输入一个参数，没有返回值
	*/
	sayHello(msg string)
	// 没有输入参数，但是有一个返回值
	doSth() string
}

func (staff staff) sayHello(msg string) {
	fmt.Printf("staff %s  say : %v\n", staff.name, msg)
}
func (staff staff) doSth() string {
	return "staff " + staff.name + " do sth"
}

func TestInterface(t *testing.T) {

	starfInfo := staff{
		name: "张三",
		age:  28,
	}
	starfInfo.sayHello("你好，世界")
	fmt.Printf("starfInfo.doSth(): %v\n", starfInfo.doSth())
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

	xm := Profile{
		name:   "小命",
		age:    18,
		gender: "male",
	}
	fmt.Printf("xm: %v\n", xm)

	xm.FmtProfile()

	myComplany := company{
		companyAddr: "湖北省",
		name:        "tencent",
	}
	starfInfo := staff{
		name:    "张三",
		age:     28,
		company: myComplany,
	}
	fmt.Printf("starfInfo: %s 工作 %s %v\n", starfInfo.name, starfInfo.name, starfInfo)

}

func (person Profile) FmtProfile() {

	fmt.Printf("person.name: %v\n", person.name)
}

func TestMap(t *testing.T) {
	var map1 map[string]int
	map1 = map[string]int{"one": 1}

	mapCreated := make(map[string]int)

	mapCreated["one"] = 1
	mapCreated["two"] = 2

	fmt.Printf("map1: %v\n", map1["two"])

	fmt.Printf("mapCreated: %v\n", mapCreated)

	keys := make([]string, 0, len(mapCreated))
	for k := range mapCreated {
		keys = append(keys, k)
	}
	fmt.Printf("keys: %v\n", keys)

	fm := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		3: func() int { return 30 },
	}
	fmt.Printf("fm: %v\n", fm)

}
