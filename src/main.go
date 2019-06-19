package main
import(
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
)
var templates *template.Template
func main(){
	templates = template.Must(template.ParseGlob("template/*"))
	r := mux.NewRouter()
	r.HandleFunc("/",indexHandler).Methods("GET")
	r.HandleFunc("/about",aboutHandler).Methods("GET")
	r.HandleFunc("/contact",contactHandler).Methods("GET")
	http.Handle("/",r)
	http.ListenAndServe(":8000",nil)
  }
func indexHandler(w http.ResponseWriter, r *http.Request){
    templates.ExecuteTemplate(w,"index.html",nil)
}
func contactHandler(w http.ResponseWriter, r *http.Request){
    templates.ExecuteTemplate(w,"contact.html",nil)
}
func aboutHandler(w http.ResponseWriter, r *http.Request){
    templates.ExecuteTemplate(w,"about.html",nil)
}

