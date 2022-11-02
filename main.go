package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Result struct {
	Size int
	Cars []Car
}

func rollHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("simple_list.html")
		if err != nil {
			log.Fatal(err)
		}

		r.ParseForm()
		mark := r.Form.Get("mark")

		cars := getCars(mark)
		amount := getCarAmount()
		var res Result
		res.Cars = cars
		res.Size = amount

		t.Execute(w, res)
	}
}
func addCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("simple_form.html")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		mark := r.Form.Get("mark")
		country := r.Form.Get("country")
		price := r.Form.Get("price")
		year := r.Form.Get("year")
		addCar(mark, country, price, year)
	}
}
func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
		fmt.Println(port)
	}
	return ":" + port
}
func main() {
	http.HandleFunc("/", rollHandler)
	http.HandleFunc("/add", addCarHandler)
	log.Fatal(http.ListenAndServe(GetPort(), nil))
}
