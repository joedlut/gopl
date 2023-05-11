package memo4

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func TestMemoSingle(t *testing.T) {
	m := New(httpGetBody)
	incomingURLS := []string{"https://www.baidu.com", "https://sina.cn", "https://www.bilibili.com/", "https://sina.cn", "https://www.baidu.com", "https://news.qq.com/", "https://news.qq.com/"}

	for _, url := range incomingURLS {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d  bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

func TestMemoMulti(t *testing.T) {
	m := New(httpGetBody)
	var n sync.WaitGroup

	incomingURLS := []string{"https://www.baidu.com", "https://sina.cn", "https://www.bilibili.com/", "https://sina.cn", "https://www.baidu.com", "https://news.qq.com/", "https://news.qq.com/"}

	for _, url := range incomingURLS {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d  bytes\n", url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()

}
