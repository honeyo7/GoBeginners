package main

import ("fmt"
		"net/http"
	"html/template")

func main() {

	http.HandleFunc("/", index_handler)
	http.HandleFunc("/aggr/", aggr_handler)
	http.ListenAndServe(":8080",nil)

}

func index_handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w, "<h1>This is dummy data !</h1>")

	fmt.Fprintf(w, "<p>You %s even add %s</p>","can","<strong> Variables</strong>")
}

func aggr_handler(w http.ResponseWriter,r *http.Request){
	p := NewsAggPage {Title: "Amazing News Aggregator", News: "News1"}
	t, _ := template.ParseFiles("demoNews.html")
	t.Execute(w,p)
}

type NewsAggPage struct{
	Title string
	News string
}