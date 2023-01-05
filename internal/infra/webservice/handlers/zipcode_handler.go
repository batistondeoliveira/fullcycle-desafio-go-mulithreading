package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/batistondeoliveira/fullcycle_desafio_go_multithreading/internal/infra/webservice/apicep"
	"github.com/batistondeoliveira/fullcycle_desafio_go_multithreading/internal/infra/webservice/dto"
	"github.com/batistondeoliveira/fullcycle_desafio_go_multithreading/internal/infra/webservice/viacep"
	"github.com/go-chi/chi/v5"
)

func GetZipCode(w http.ResponseWriter, r *http.Request) {
	zipcode := chi.URLParam(r, "zipcode")
	if zipcode == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c1 := make(chan dto.ZipcodeOutputDto)
	c2 := make(chan dto.ZipcodeOutputDto)

	go getByApiCep(c2, zipcode)
	go getByViaCep(c1, zipcode)

	select {
	case msg1 := <-c1:
		json, _ := json.Marshal(msg1)
		fmt.Printf("API = VIACEP\nPayload = %s\n", json)
	case msg2 := <-c2:
		json, _ := json.Marshal(msg2)
		fmt.Printf("API = APICEP\nPayload = %s\n", json)
	case <-time.After(time.Second):
		fmt.Println("Timeout")
	}
}

func getByViaCep(ch chan<- dto.ZipcodeOutputDto, zipcode string) {
	viaCep := viacep.ViaCEP{}
	ch <- viaCep.GetZipCode(zipcode)
}

func getByApiCep(ch chan<- dto.ZipcodeOutputDto, zipcode string) {
	apiCep := apicep.ApiCEP{}
	ch <- apiCep.GetZipCode(zipcode)
}
