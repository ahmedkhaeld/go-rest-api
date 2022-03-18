package main

import (
	"github.com/ahmedkhaeld/rest-api/cache"
	"github.com/ahmedkhaeld/rest-api/controller"
	router "github.com/ahmedkhaeld/rest-api/http"
	"github.com/ahmedkhaeld/rest-api/repository"
	"github.com/ahmedkhaeld/rest-api/service"
	"os"
)

var (
	postRepository = repository.NewSQLiteRepository()
	postService    = service.NewPostService(postRepository)
	postCache      = cache.NewRedisCache("localhost:6379", 1, 10)
	postController = controller.NewPostController(postService, postCache)
	httpRouter     = router.NewMuxRouter()
)

func main() {

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.GET("/posts/{id}", postController.GetPostByID)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(os.Getenv("PORT"))
}
