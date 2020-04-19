package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

func main() {
	db := database{"shoes": 50, "socks": 5}

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/item", db.item)
	http.HandleFunc("/update", db.update)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

type Item struct {
	Name string
	Price dollars
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]

	if !ok {
		msg := fmt.Sprintf("no such item: %s\n", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "%s\n", price)
}

func (db database) item(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		item := req.URL.Query().Get("item")
		price, ok := db[item]

		if !ok {
			msg := fmt.Sprintf("no such item: %s\n", item)
			http.Error(w, msg, http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "%s: %s\n", item, price)
	case "POST":
		var itemData Item
		err := json.NewDecoder(req.Body).Decode(&itemData)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, ok := db[itemData.Name]

		if !ok {
			db[itemData.Name] = itemData.Price
	
			fmt.Fprintf(w, "add success")
		} else {
			msg := fmt.Sprintf("already exists item: %s\n", itemData.Name)
			http.Error(w, msg, http.StatusBadRequest)
		}
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	var itemData Item
	err := json.NewDecoder(req.Body).Decode(&itemData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	price, ok := db[itemData.Name]

	if !ok {
		msg := fmt.Sprintf("can not find item: %s\n", itemData.Name)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	if price == itemData.Price {
		msg := fmt.Sprintf("same price: %.2f\n", price)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	db[itemData.Name] = itemData.Price

	fmt.Fprintf(w, "update success old price: %.2f new price: %.2f", price, itemData.Price)
}