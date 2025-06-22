package usecase

import (
	"github.com/AndersonOdilo/fullcycle-deploy-cloud-run/internal/entity"
)

type GetTempInputDTO struct {
	Cep    string  `json:"cep"`
}

type GetTempOutputDTO struct {
	Celsius         float64  	`json:"temp_C"`
	Fahrenheit      float64 	`json:"temp_F"`
	Kelvin        	float64 	`json:"temp_K"`
}

type GetTempUseCase struct {
	LocationRepository 	entity.LocationRepositoryInterface
	TempRepository 		entity.TempRepositoryInterface
}

func NewGetTempUseCase(locationRepository entity.LocationRepositoryInterface, tempRepository entity.TempRepositoryInterface) *GetTempUseCase {
	return &GetTempUseCase{
		LocationRepository: locationRepository,
		TempRepository: tempRepository,
	}
}

func (g *GetTempUseCase) Execute(input GetTempInputDTO) (GetTempOutputDTO, error) {

	cep, err := entity.NewCep(input.Cep);

	if (err != nil) {
		return GetTempOutputDTO{}, err;
	}

	location, err := g.LocationRepository.Get(cep);

	if (err != nil) {
		return GetTempOutputDTO{}, err;
	}

	temp, err := g.TempRepository.Get(&location);

	if (err != nil) {
		return GetTempOutputDTO{}, err;
	}

	outputDTO := GetTempOutputDTO{
		Celsius:        temp.Celsius,
		Fahrenheit:     temp.Celsius * 1.8 + 32,
		Kelvin:        	temp.Celsius + 273,
	}


	return outputDTO, nil
}


