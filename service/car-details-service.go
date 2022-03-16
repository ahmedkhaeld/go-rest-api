package service

import (
	"encoding/json"
	"fmt"
	"github.com/ahmedkhaeld/rest-api/entity"
	"net/http"
)

// invoke the required constructors, make two channels to pass the response to the request
var (
	ownerService     = NewOwnerService()
	carService       = NewCarService()
	carDatachannel   = make(chan *http.Response)
	ownerDatachannel = make(chan *http.Response)
)

type CarDetailsService interface {
	GetDetails() entity.CarDetails
}

type service struct{}

func NewCarDetailsService() CarDetailsService {
	return &service{}
}

func (*service) GetDetails() entity.CarDetails {
	// go routine call endpoint one  https://myfakeapi.com/api/cars/1
	go carService.FetchData()
	// go routine call endpoint two https://myfakeapi.com/api/users/1
	go ownerService.FetchData()

	// get the data from the channels
	// invoke getCarData func to get data from endpoint 1
	car, _ := getCarData()
	// invoke getUserData func to get data from endpoint 2
	owner, _ := getOwnerData()

	return entity.CarDetails{
		ID:             car.CarData.ID,
		Brand:          car.CarData.Brand,
		Model:          car.CarData.Model,
		Year:           car.CarData.Year,
		Vin:            car.CarData.Vin,
		OwnerFirstName: owner.OwnerData.FirstName,
		OwnerLastName:  owner.OwnerData.LastName,
		OwnerEmail:     owner.OwnerData.Email,
		OwnerJobTitle:  owner.OwnerData.JobTitle,
	}
}

func getCarData() (entity.Car, error) {
	//r1 read car data from the channel
	r1 := <-carDatachannel
	var car entity.Car
	err := json.NewDecoder(r1.Body).Decode(&car)
	if err != nil {
		fmt.Print(err.Error())
		return car, err
	}
	return car, nil
}

func getOwnerData() (entity.Owner, error) {
	// r2 read owner data from the channel
	r2 := <-ownerDatachannel
	var owner entity.Owner
	err := json.NewDecoder(r2.Body).Decode(&owner)
	if err != nil {
		fmt.Print(err.Error())
		return owner, err
	}
	return owner, nil
}
