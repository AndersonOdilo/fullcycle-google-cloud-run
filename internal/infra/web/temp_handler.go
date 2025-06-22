package web

import (
	"encoding/json"
	"net/http"

	"github.com/AndersonOdilo/fullcycle-deploy-cloud-run/internal/entity"
	"github.com/AndersonOdilo/fullcycle-deploy-cloud-run/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type WebTempHandler struct {
	LocationRepository 	entity.LocationRepositoryInterface
	TempRepository 		entity.TempRepositoryInterface
}

func NewWebTempHandler(locationRepository entity.LocationRepositoryInterface, tempRepository entity.TempRepositoryInterface) *WebTempHandler {
	return &WebTempHandler{
		LocationRepository: locationRepository,
		TempRepository: tempRepository,
	};
}

func (h *WebTempHandler) Get(w http.ResponseWriter, r *http.Request) {
	
	inputDTO := usecase.GetTempInputDTO{
		Cep:chi.URLParam(r, "cep") ,
	};

	getTempUseCase := usecase.NewGetTempUseCase(h.LocationRepository, h.TempRepository);

	output, err := getTempUseCase.Execute(inputDTO);

	if err != nil {

		if (err.Error() == "can not find zipcode"){

			http.Error(w, err.Error(), http.StatusNotFound)

		} else if (err.Error() == "invalid zipcode"){

			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		
		}else {

			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
		return
	}
	
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
