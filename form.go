package main

import ("net/http"
	"html/template")

func main() {

	http.HandleFunc("/", form_handler)
	http.HandleFunc("/process", process_handler)
	http.ListenAndServe(":8080",nil)

}


func form_handler(w http.ResponseWriter,r *http.Request){
	t, _ := template.ParseFiles("form.html")
	t.Execute(w,nil)
}

func process_handler(w http.ResponseWriter,r *http.Request){
	
	if r.Method != "POST" {
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	fname := r.FormValue("firster")
	lname := r.FormValue("laster")

	d := struct{
		First string
		Last string
	}{
		First: fname,
		Last: lname,
	}
	t, _ := template.ParseFiles("formresult.html")
	t.Execute(w,d)
}

