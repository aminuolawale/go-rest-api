package main
import ( 
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

func main(){
	router := mux.NewRouter();
	const port string = ":8000";
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request){
		fmt.Fprintln(res, "Up and running so fast")
	})
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", createPost).Methods("POST")
	log.Println("server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}