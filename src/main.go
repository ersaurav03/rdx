package main
import(
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
	"github.com/go-redis/redis"
)
var templates *template.Template
var client *redis.Client
func main(){
    client = redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
	})
	templates = template.Must(template.ParseGlob("template/*"))
	r := mux.NewRouter()
	r.HandleFunc("/",indexGetHandler).Methods("GET")
	r.HandleFunc("/",indexPostHandler).Methods("POST")
	// r.HandleFunc("/about",aboutHandler).Methods("GET")
	// r.HandleFunc("/contact",contactHandler).Methods("GET")
	fs:= http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static",fs))
	http.Handle("/",r)
	http.ListenAndServe(":8000",nil)
  }
func indexGetHandler(w http.ResponseWriter, r *http.Request){
	comments, err := client.LRange("comments",0,1).Result()
	if err != nil{
		return
	}
    templates.ExecuteTemplate(w,"index.html",comments)
}
func indexPostHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	comment:= r.PostForm.Get("comment")
	 client.LPush("comments",comment)
	 http.Redirect(w,r,"/",302)
}

// func contactHandler(w http.ResponseWriter, r *http.Request){
//     templates.ExecuteTemplate(w,"contact.html",nil)
// }
// func aboutHandler(w http.ResponseWriter, r *http.Request){
//     templates.ExecuteTemplate(w,"about.html",nil)
// }

