package main

import (
	"github.com/ahmedkhaeld/rest-api/controller"
	router "github.com/ahmedkhaeld/rest-api/http"
	"github.com/ahmedkhaeld/rest-api/repository"
	"github.com/ahmedkhaeld/rest-api/service"
	"os"
)

var (
	postRepository = repository.NewSQLiteRepository()
	postService    = service.NewPostService(postRepository)
	postController = controller.NewPostController(postService)
	httpRouter     = router.NewMuxRouter()
)

func main() {

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(os.Getenv("PORT"))
}
