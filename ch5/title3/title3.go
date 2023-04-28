package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	//for c := n.FirstChild; c != nil; c = n.NextSibling {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		//递归
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		//recover函数在defer函数中调用
		switch p := recover(); p {
		case nil:

		case bailout{}:
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p)
		}
	}()

	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	fmt.Println(title)
	return title, nil
}

func main() {
	resp, err := http.Get("https://new.qq.com/ch/tech/")
	if err != nil {
		log.Fatalf("get url resp failed %v", err)
	}
	defer resp.Body.Close()
	doc, _ := html.Parse(resp.Body)
	title, err := soleTitle(doc)
	if err != nil {
		log.Fatalf("soletitle failed %v", err)
	} else {
		fmt.Println(title)
	}
}
