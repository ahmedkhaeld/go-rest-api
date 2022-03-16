package main

import (
	"fmt"
	"github.com/ahmedkhaeld/rest-api/controller"
	router "github.com/ahmedkhaeld/rest-api/http"
	"github.com/ahmedkhaeld/rest-api/repository"
	"github.com/ahmedkhaeld/rest-api/service"
	"net/http"
)

var (
	postRepository = repository.NewSQLiteRepository()
	postService    = service.NewPostService(postRepository)
	postController = controller.NewPostController(postService)
	httpRouter     = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)
}
