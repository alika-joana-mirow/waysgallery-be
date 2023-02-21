package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	imagesdto "waysGallery/dto/images"
	dto "waysGallery/dto/result"
	"waysGallery/models"
	"waysGallery/repositories"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type handlerImages struct {
	ImagesRepository repositories.ImagesRepository
}

func HandlerImages(ImagesRepository repositories.ImagesRepository) *handlerImages {
	return &handlerImages{ImagesRepository}
}

func (h *handlerImages) FindImages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	images, err := h.ImagesRepository.FindImages()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: images}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerImages) GetImages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var image models.Image
	image, err := h.ImagesRepository.GetImages(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// product.Image = os.Getenv("PATH_FILE") + product.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: image}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerImages) CreateImages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(imagesdto.ImagesRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// image := models.Image{
	// 	Image: request.Image,
	// }

	// data, err := h.ImagesRepository.CreateImages(image)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: request}
	json.NewEncoder(w).Encode(response)
}
