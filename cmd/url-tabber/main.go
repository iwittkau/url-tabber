package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"net/http"

	"github.com/iwittkau/url-tabber/osascript"
	"github.com/sethvargo/go-password/password"
)

var (
	port      string
	secret    string
	templates string
)

type reqest struct {
	URL    string `json:"url"`
	Secret string `json:"secret"`
}

type response struct {
	Error string `json:"error"`
}

func main() {

	flag.StringVar(&port, "p", "8080", "set the application port")
	flag.StringVar(&secret, "s", "", "set the application secret")
	flag.StringVar(&templates, "t", "web/templates/", "set the templates path")
	flag.Parse()

	if secret == "" {
		fmt.Println("Generating secret ...")
		var err error
		secret, err = password.Generate(20, 10, 0, false, false)
		if err != nil {
			panic(err)
		}
		fmt.Println("Application secret:", secret)
	}

	tmpl := template.Must(template.ParseFiles(templates + "index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/url", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		decoder := json.NewDecoder(req.Body)
		var r reqest
		err := decoder.Decode(&r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if r.Secret != secret {
			data, _ := newJSONResponse("wrong secret")
			http.Error(w, string(data), http.StatusUnauthorized)
			return
		}
		err = osascript.OpenNewChromeTab(r.URL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	})

	fmt.Println("Starting to listen on :" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err.Error())
	}

}

func newJSONResponse(message string) ([]byte, error) {
	resp := response{
		Error: message,
	}

	return json.Marshal(resp)
}
