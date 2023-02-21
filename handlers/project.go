package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	projectdto "waysGallery/dto/project"
	dto "waysGallery/dto/result"
	"waysGallery/models"
	"waysGallery/repositories"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type handlerProject struct {
	ProjectRepository repositories.ProjectRepository
}

func HandlerProject(ProjectRepository repositories.ProjectRepository) *handlerProject {
	return &handlerProject{ProjectRepository}
}

func (h *handlerProject) FindProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	project, err := h.ProjectRepository.FindProject()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: project}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProject) FindHiredById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	hired, err := h.ProjectRepository.FindHiredById(id)
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

func (h *handlerProject) GetProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var project models.Project
	project, err := h.ProjectRepository.GetProject(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: project}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProject) CreateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(projectdto.Project)
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

	project := models.Project{
		Description: request.Description,
		File:        request.File,
		HiredId:     request.HiredID,
	}

	project, err = h.ProjectRepository.CreateProject(project)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	project, _ = h.ProjectRepository.GetProject(project.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: project}
	json.NewEncoder(w).Encode(response)
}
