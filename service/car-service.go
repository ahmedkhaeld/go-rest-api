package service

import (
	"fmt"
	"net/http"
)

type CarService interface {
	FetchData()
}

const (
	carServiceUrl = "https://myfakeapi.com/api/cars/1"
)

// fetchCarDataService implement the CarService interface
type fetchCarDataService struct{}

func NewCarService() CarService {
	return &fetchCarDataService{}
}

func (*fetchCarDataService) FetchData() {
	client := http.Client{}
	fmt.Printf("Fetching %s \n", carServiceUrl)

	// Call external API
	resp, _ := client.Get(carServiceUrl)

	// Write response to the channel
	carDatachannel <- resp
}
