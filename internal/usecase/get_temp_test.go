package usecase

import (
	"testing"

	"github.com/AndersonOdilo/fullcycle-deploy-cloud-run/configs"
	"github.com/AndersonOdilo/fullcycle-deploy-cloud-run/internal/api"
	"github.com/AndersonOdilo/fullcycle-deploy-cloud-run/internal/entity"
	"github.com/stretchr/testify/suite"
)

type GetTempUseCaseTestSuite struct {
	suite.Suite
	LocationRepository 	entity.LocationRepositoryInterface
	TempRepository 		entity.TempRepositoryInterface
}

func (suite *GetTempUseCaseTestSuite) SetupSuite() {

	suite.LocationRepository = api.NewLocationRepository();
	suite.TempRepository = api.NewTempRepository(configs.Conf{
		WeatherApiKey: "376a4a90b4074c07a26165101252206",
	});
}


func TestSuite(t *testing.T) {
	suite.Run(t, new(GetTempUseCaseTestSuite))
}

func (suite *GetTempUseCaseTestSuite) TestBuscarTempComCepMenor_RetornaErrorInvalidZipcode() {
	
	getTempUseCase := NewGetTempUseCase(suite.LocationRepository, suite.TempRepository);

	inputDTO := GetTempInputDTO{
		Cep: "1234567",
	}

	_, err := getTempUseCase.Execute(inputDTO);

	suite.Error(err, "invalid zipcode");
	
}


func (suite *GetTempUseCaseTestSuite) TestBuscarTempComCepMaior_RetornaErrorInvalidZipcode() {
	
	getTempUseCase := NewGetTempUseCase(suite.LocationRepository, suite.TempRepository);

	inputDTO := GetTempInputDTO{
		Cep: "123456789",
	}

	_, err := getTempUseCase.Execute(inputDTO);

	suite.Error(err, "invalid zipcode");
	
}


func (suite *GetTempUseCaseTestSuite) TestBuscarTempComCepInexistente_RetornaErrorInvalidZipcode() {
	
	getTempUseCase := NewGetTempUseCase(suite.LocationRepository, suite.TempRepository);

	inputDTO := GetTempInputDTO{
		Cep: "99999999",
	}

	_, err := getTempUseCase.Execute(inputDTO);

	suite.Error(err, "can not find zipcode");
	
}

func (suite *GetTempUseCaseTestSuite) TestBuscarTempComCepValido_RetornaOutputComTodasTemperaturas() {
	
	getTempUseCase := NewGetTempUseCase(suite.LocationRepository, suite.TempRepository);

	inputDTO := GetTempInputDTO{
		Cep: "90010170",
	}

	output, err := getTempUseCase.Execute(inputDTO);

	suite.NoError(err);
	suite.NotNil(output);
	suite.NotNil(output.Celsius);
	suite.NotNil(output.Fahrenheit);
	suite.NotNil(output.Kelvin);
	
}


