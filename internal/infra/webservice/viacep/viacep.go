package viacep

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/batistondeoliveira/fullcycle_desafio_go_multithreading/internal/infra/webservice/dto"
	"github.com/batistondeoliveira/fullcycle_desafio_go_multithreading/pkg/formatter"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (v *ViaCEP) GetZipCode(zipcode string) dto.ZipcodeOutputDto {
	zipcode = formatter.Zipcode(zipcode)
	req, err := http.Get("http://viacep.com.br/ws/" + zipcode + "/json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
	}
	var data ViaCEP
	err = json.Unmarshal(res, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
	}
	return dto.ZipcodeOutputDto{
		Cep:        data.Cep,
		Logradouro: data.Logradouro,
		Bairro:     data.Bairro,
		Localidade: data.Localidade,
		Uf:         data.Uf,
	}
}
