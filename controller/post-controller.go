package controller

import (
	"encoding/json"
	"github.com/ahmedkhaeld/rest-api/cache"
	"github.com/ahmedkhaeld/rest-api/entity"
	"github.com/ahmedkhaeld/rest-api/errors"
	"github.com/ahmedkhaeld/rest-api/service"
	"net/http"
	"strings"
)

type controller struct{}

var (
	postService service.PostService
	postCache   cache.PostCache
)

type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
	GetPostByID(response http.ResponseWriter, request *http.Request)
}

func NewPostController(service service.PostService, cache cache.PostCache) PostController {
	postService = service
	postCache = cache // include the caching layer within the controller
	return &controller{}
}

func (*controller) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the posts"})
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

// GetPostByID get post by id, extract the id from the request url, make
// Get call to the cache by that id as key, if the post in the cache then
// retrieve, if not then use the service to get it, and then store the post
// in the cache using Set function
func (*controller) GetPostByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	postID := strings.Split(request.URL.Path, "/")[2]
	var post = postCache.Get(postID)
	if post == nil {
		post, err := postService.FindByID(postID)
		if err != nil {
			response.WriteHeader(http.StatusNotFound)
			json.NewEncoder(response).Encode(errors.ServiceError{Message: "No posts found!"})
			return
		}
		postCache.Set(postID, post)
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(post)
	} else {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(post)
	}

}

func (*controller) AddPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}
	err1 := postService.Validate(&post)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := postService.Create(&post)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
