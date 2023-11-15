package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type Employee struct {
	ID     int    // 唯一标识符
	Name   string // 姓名
	Salary int
}

type Car struct {
	color string
	size  string
}

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func (e *Employee) GiveRaise(percent int) {
	// 值类型
	dilbert := Employee{ID: 1, Name: "Dilbert", Salary: 5000}

	// 指针类型
	dilbert2 := new(Employee)
	dilbert2.ID = 1
	// 也可以使用字面量 但是必须使用 & 取地址
	dilbert3 := &Employee{ID: 1, Name: "Dilbert", Salary: 5000}
	dilbert3.Salary = 5000

	dilbert.Salary -= 5000 // promoted, for writing excellent code

	dilbert.Salary = dilbert.Salary * 2 // promoted, for writing excellent code

	movies := []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	}
	movies = append(movies, Movie{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}})

	f := square
	fmt.Println(f(3)) // "9"

}
func square(n int) int {
	return n * n
}
func negative(n int) int {
	return -n
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)

		depth++

	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
func title(url string) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	ct := resp.Header.Get("Content-Type")

	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s has type %s,not text/html", url, ct)
	}
	return err

}
