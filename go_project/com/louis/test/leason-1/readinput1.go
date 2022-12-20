package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/rpc"
	"os"
	"sync"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "57.12/5212/go"
	format                 = "%f / %d / %s"
)
var inputReader *bufio.Reader
var err error
var keyChar = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {

	// inputFunc()

	http.HandleFunc("/", Redirect)

	http.HandleFunc("/add", Add)

	http.ListenAndServe(":8080", nil)

}

/*
 */
func inputFunc() {
	/**
	os.Open 读取文件
		返回的是File 的引用类型
	*/
	inputFile, inputError := os.Open("/Users/louis/workspace/louis/louis-blockchain/src/main/com/input.dat")
	if inputError != nil {
		fmt.Printf("an error ecc %s", inputError)
		return

	}
	// 定义 执行左右语句，
	defer inputFile.Close()

	defer fmt.Print("this is the last")
	/*
	 bufio.NewReader
	 param   io.Reader 接口数据，File 实现了io.Reader 接口,go 中实现接口的方式 不是很方便找到哪个类实现了哪个接口的哪个方法
	 return  bufio.Reader struct
	*/
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readError := inputReader.ReadString('\n')
		fmt.Printf("the input was :%s \n", inputString)
		if readError == io.EOF {
			return
		}

	}
}

func input2() {

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	fmt.Printf("Your name is %s", input)
	// For Unix: test with delimiter "\n", for Windows: test with "\r\n"
	switch input {
	case "Philip\r\n":
		fmt.Println("Welcome Philip!")
	case "Chris\r\n":
		fmt.Println("Welcome Chris!")
	case "Ivo\r\n":
		fmt.Println("Welcome Ivo!")
	default:
		fmt.Printf("You are not welcome here! Goodbye!")
	}

	// version 2:
	switch input {
	case "Philip\r\n":
		fallthrough
	case "Ivo\r\n":
		fallthrough
	case "Chris\r\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}

	// version 3:
	switch input {
	case "Philip\r\n", "Ivo\r\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}
}

func input3() {

}

/*
sync.RWMutex
sync.
*/
type URLStore struct {
	urls map[string]string // map from short to long
	mu   sync.RWMutex
}

/*
 */
func (s *URLStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	url := s.urls[key]
	return url
}

func (s *URLStore) Set(key, url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, present := s.urls[key]

	if present {
		return false
	}
	s.urls[key] = url
	return true
}

func (s *URLStore) Count() int {
	s.mu.RLock()
	// 解锁RLock
	defer s.mu.RUnlock()

	return len(s.urls)

}
func NewUrlStore(fileName string) *URLStore {

	s := &URLStore{urls: make(map[string]string)}
	if fileName != "" {
		// s.save = make(chan record, saveQueueLength)
		// if err != s.load(fileName); err != nil {

		// }
		// go s.saveLoop(fileName)
	}
	return s

	// & 是取地址操作
	// return &URLStore{urls: make(map[string]string)}

}
func (s *URLStore) Put(url string) string {
	for {
		key := genKey(s.Count())
		if s.Set(key, url) {
			return key
		}
	}
	// shouldn’t get here
	return ""
}
func genKey(n int) string {
	if n == 0 {
		return string(keyChar[0])
	}
	l := len(keyChar)
	s := make([]byte, 20) // FIXME: will overflow. eventually.
	i := len(s)
	for n > 0 && i >= 0 {
		i--
		j := n % l
		n = (n - j) / l
		s[i] = keyChar[j]
	}
	return string(s[i:])
}

func init() {
	var store = NewUrlStore("")

	if store.Set("a", "http://google.com") {

	}

}

var store = NewUrlStore("")

func Add(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		fmt.Fprint(w, AddForm)
		return
	}
	key := store.Put(url)
	fmt.Fprintf(w, "http://localhost:8080/%s", key)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	url := store.Get(key)
	if url == "" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

const AddForm = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`

type ProxyStore struct {
	client *rpc.Client
	urls   *URLStore
}

func (s *ProxyStore) get(key, url *string) error {
	return s.client.Call("store.Get", key, url)
}
func (s *ProxyStore) Put(url, key *string) error {
	return s.client.Call("Store.Put", url, key)
}

func NewProxyStore(addr string) *ProxyStore {
	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Println("Error constructing ProxyStore:", err)
	}
	return &ProxyStore{client: client, urls: NewUrlStore("")}

}

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		fmt.Fprintf(os.Stdout, "%s", buf)
		if err == io.EOF {
			break
		}
		fmt.Print("input end")
	}
	return

}
