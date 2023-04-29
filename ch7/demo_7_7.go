package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shows": 50, "sockes": 100}
	// db 实现了 handler接口，所以可以作为参数传递给 ListenAndServe方法
	log.Fatal(http.ListenAndServe("localhost:8080", db))
}

type dollars float64
type database map[string]dollars

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

/** func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s %s\n", item, price)
	}
} **/

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "not such item:%q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page %s", req.URL)
		//msg := fmt.Sprintf("no such page %s", req.URL)
		//http.Error(w,msg,http.StatusNotFound)
	}
}
