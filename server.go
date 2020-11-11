package main
import ( 
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"os"
	"path/filepath"
)

func main(){
	setEnv()
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


func setEnv(){
	cwd, err := os.Getwd()
	if err == nil{
		fmt.Println("we are here")
		envPath := filepath.Join(cwd, "config", "golang-rest-api-project-firebase-adminsdk-zjeqf-35ecb16f4b.json")
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", envPath)
		fmt.Println("env set")
		return
	}
}