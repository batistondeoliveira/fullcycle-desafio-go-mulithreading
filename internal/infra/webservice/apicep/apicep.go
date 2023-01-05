package apicep

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/batistondeoliveira/fullcycle_desafio_go_multithreading/internal/infra/webservice/dto"
	"github.com/batistondeoliveira/fullcycle_desafio_go_multithreading/pkg/formatter"
)

type ApiCEP struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

func (a *ApiCEP) GetZipCode(zipcode string) dto.ZipcodeOutputDto {
	zipcode = formatter.Zipcode(zipcode)
	req, err := http.Get("https://cdn.apicep.com/file/apicep/" + zipcode + ".json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
	}
	var data ApiCEP
	err = json.Unmarshal(res, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
	}
	return dto.ZipcodeOutputDto{
		Cep:        data.Code,
		Logradouro: data.Address,
		Bairro:     data.District,
		Localidade: data.City,
		Uf:         data.State,
	}
}
