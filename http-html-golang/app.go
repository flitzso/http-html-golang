package main

import (
	"net/http"

	"fmt"

	"html/template"
)

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `<h1>Olá Mundo!</h1>`)
}

func getSecondRoute(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(`
       <h1>Segunda forma</h1>
       <a href="./3">Ir para pagina 3</a>
	`))
}

func getThirdRoute(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "page3.html")
}

func getFourRoute(w http.ResponseWriter, r *http.Request) {

	tpl, err := template.ParseFiles("page4.html")

	if err != nil {
		fmt.Println(`Ocorreu um erro na analise do template 
					  ou o arquivo não foi encontrado`)
		fmt.Fprintf(w, `
		   <h1>A pagina não foi encontrada</h1>
		   <a href="./">Ir para a pagina inicial</a>`)
		return
	}

	data := map[string]string{
		"Title": "Quarta forma",
		"Link":  "Voltar para a pagina inicial",
	}

	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, data)

}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/2", getSecondRoute)
	http.HandleFunc("/3", getThirdRoute)
	http.HandleFunc("/4", getFourRoute)

	http.ListenAndServe(":8080", nil)
}
