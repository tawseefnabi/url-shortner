package controller

import (
	"fmt"
	"net/http"
	service "url-shortner/Service"
)

var (
	defaultUrl string = "http://localhost:8080/home/"
)

type Controller struct {
	ser *service.Service
}

func NewController(ser *service.Service) Controller {
	return Controller{
		ser: ser,
	}
}

func (c *Controller) GenerateTinyUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GenerateTinyUrl")
	return
}
