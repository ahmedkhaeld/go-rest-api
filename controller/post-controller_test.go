package controller

import (
	"bytes"
	"encoding/json"
	"github.com/ahmedkhaeld/rest-api/cache"
	"github.com/ahmedkhaeld/rest-api/entity"
	"github.com/ahmedkhaeld/rest-api/repository"
	"github.com/ahmedkhaeld/rest-api/service"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	ID    int64  = 123
	TITLE string = "title 1"
	TEXT  string = "text 1"
)

var (
	postRepo       = repository.NewSQLiteRepository()
	postSrv        = service.NewPostService(postRepo)
	postCh         = cache.NewRedisCache("localhost:6379", 0, 10)
	postController = NewPostController(postSrv, postCh)
)

func TestAddPost(t *testing.T) {
	// Create new HTTP request
	var jsonStr = []byte(`{"title":"` + TITLE + `","text":"` + TEXT + `"}`)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonStr))

	// Assign HTTP Request handler Function (controller function)
	handler := http.HandlerFunc(postController.AddPost)
	// Record the HTTP Response
	response := httptest.NewRecorder()
	// Dispatch(send off to a destination) the HTTP Request
	handler.ServeHTTP(response, req)

	// Assert HTTP status
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode HTTP response
	var post entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&post)

	// Assert HTTP response
	assert.NotNil(t, post.ID)
	assert.Equal(t, TITLE, post.Title)
	assert.Equal(t, TEXT, post.Text)

	// Cleanup database
	tearDown(post.ID)
}

func setup() {
	var post entity.Post = entity.Post{
		ID:    ID,
		Title: TITLE,
		Text:  TEXT,
	}
	postRepo.Save(&post)
}

func tearDown(postID int64) {
	var post entity.Post = entity.Post{
		ID: postID,
	}
	postRepo.Delete(&post)
}

func TestGetPosts(t *testing.T) {

	// Insert new post
	setup()

	// Create new HTTP request
	req, _ := http.NewRequest("GET", "/posts", nil)

	// Assign HTTP Request handler Function (controller function)
	handler := http.HandlerFunc(postController.GetPosts)
	// Record the HTTP Response
	response := httptest.NewRecorder()
	// Dispatch the HTTP Request
	handler.ServeHTTP(response, req)

	// Assert HTTP status
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode HTTP response
	var posts []entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&posts)

	// Assert HTTP response
	assert.Equal(t, ID, posts[0].ID)
	assert.Equal(t, TITLE, posts[0].Title)
	assert.Equal(t, TEXT, posts[0].Text)

	// Cleanup database
	tearDown(ID)
}
