package main

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

type timeHandler struct {
	format string
}

func (th timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm))
}

func indexTmpl(w http.ResponseWriter, request *http.Request) {
	t1, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Fatalln("模板解析失败", err)
	}
	t1.Execute(w, "Hello World")
}

func main() {
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("./public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", indexTmpl)
	// Initialise the timeHandler in exactly the same way we would any normal
	// struct.
	/*th := timeHandler{format: time.RFC1123}
	// Like the previous example, we use the mux.Handle() fnction to register
	// this with our ServeMux.
	mux.Handle("/time", th)
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello World, $s!", request.URL.Path[1:])
	})*/

	http.ListenAndServe(":8080", mux)

}
