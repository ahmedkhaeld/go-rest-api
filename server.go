package main

import (
	"fmt"
	"github.com/ahmedkhaeld/rest-api/controller"
	router "github.com/ahmedkhaeld/rest-api/http"
	"github.com/ahmedkhaeld/rest-api/service"
	"net/http"
)

var (
	carDetailsService    service.CarDetailsService       = service.NewCarDetailsService()
	carDetailsController controller.CarDetailsController = controller.NewCarDetailsController(carDetailsService)
	httpRouter           router.Router                   = router.NewChiRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/carDetails", carDetailsController.GetCarDetails)

	httpRouter.SERVE(port)
}
