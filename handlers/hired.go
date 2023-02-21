package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	hireddto "waysGallery/dto/hired"
	dto "waysGallery/dto/result"
	"waysGallery/models"
	"waysGallery/repositories"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handleHired struct {
	HiredRepository repositories.HiredRepository
}

func HandlerHired(HiredRepository repositories.HiredRepository) *handleHired {
	return &handleHired{HiredRepository}
}

func (h *handleHired) FindHired(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	hired, err := h.HiredRepository.FindHired()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: hired}
	json.NewEncoder(w).Encode(response)
}

func (h *handleHired) GetHired(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var hired models.Hired
	hired, err := h.HiredRepository.GetHired(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: hired}
	json.NewEncoder(w).Encode(response)
}

func (h *handleHired) CreateHired(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	request := new(hireddto.HiredRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	hired := models.Hired{
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
		OrderById:   userId,
		OrderToId:   request.OrderToId,
		Status:      "Waiting Accept",
	}

	hired.StartDate, err = time.Parse("2 January 2006", request.StartDate)
	hired.EndDate, err = time.Parse("2 January 2006", request.EndDate)

	hired, err = h.HiredRepository.CreateHired(hired)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	hired, _ = h.HiredRepository.GetHired(hired.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: hired}
	json.NewEncoder(w).Encode(response)
}

func (h *handleHired) FindHiredUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// var trips models.Trip
	hired, err := h.HiredRepository.FindHiredUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// transaction.Attachment = path_file + transaction.Attachment

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: hired}
	json.NewEncoder(w).Encode(response)
}

func (h *handleHired) FindHiredUserOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// var trips models.Trip
	hired, err := h.HiredRepository.FindHiredUserOrder(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// transaction.Attachment = path_file + transaction.Attachment

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: hired}
	json.NewEncoder(w).Encode(response)
}

func (h *handleHired) UpdateHired(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(hireddto.HiredRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	hired, err := h.HiredRepository.GetHired(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Status != "" {
		hired.Status = request.Status
	}

	data, err := h.HiredRepository.UpdateHired(hired)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
