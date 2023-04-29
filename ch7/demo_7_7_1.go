package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shows": 50, "sockes": 100}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	//http.HandleFunc("/list", db.list)
	//http.HandleFunc("/price", db.price)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
	//log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float64
type database map[string]dollars

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

/** func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s %s\n", item, price)
	}
} **/

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "not such item:%q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)

}

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
