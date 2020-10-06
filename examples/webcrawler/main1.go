package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func HttpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result := ""

	buf := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("read complete")
			break
		}
		if err != nil && err != io.EOF {
			return "", err
		}
		result += string(buf[:n])
	}
	return result, nil
}

// web crawl
func working(start, end int) {
	fmt.Printf("crawling %d to %d", start, end)
	for i := start; i <= end; i++ {
		url := fmt.Sprintf("https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=%d", (i-1)*50)
		result, err := HttpGet(url)
		if err != nil {
			fmt.Println("Http err: ", err)
			continue
		}
		f, err := os.Create(fmt.Sprintf("page%d.html", i))
		if err != nil {
			fmt.Println("create file error: ", err)
			continue
		}
		f.WriteString(result)
		f.Close()
	}
}

func main() {
	var start, end int
	fmt.Print("please enter the start page:")
	fmt.Scan(&start)
	fmt.Print("please enter the end page:")
	fmt.Scan(&end)

	working(start, end)
}
