package main

import (
	"fmt"
	"net/http"

	controller "web-service-gorilla/controller"
	router "web-service-gorilla/http"
	repositories "web-service-gorilla/repositories/post"
	services "web-service-gorilla/services"
)

var (
	postRepository repositories.PostRepository = repositories.NewFirestoreRepository()
	postService    services.PostService        = services.NewPostService(postRepository)
	postController controller.PostController   = controller.NewPostController(postService)
	httpRouter     router.Router               = /* router.NewMuxRouter() */ router.NewChiRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and running...")
	})

	//APIs
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)
}
