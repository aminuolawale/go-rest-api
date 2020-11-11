package main
import ( 
	"fmt"
	router "./http"
	"net/http"
	"os"
	"path/filepath"
	"./controllers"
	"./service"
	"./repos"
)

var httpRouter router.Router = router.NewChiRouter()
var postRepository repos.PostRepo = repos.NewFirestoreRepository()
var postService service.PostService = service.NewPostService(postRepository)
var postController controllers.PostController = controllers.NewPostController(postService)

func main(){
	setEnv()
	const port string = ":8000";
	httpRouter.GET("/", func(res http.ResponseWriter, req *http.Request){
		fmt.Fprintln(res, "Up and running so fast")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.CreatePost)
	httpRouter.SERVE(port)
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