package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	postdto "waysGallery/dto/post"
	dto "waysGallery/dto/result"
	"waysGallery/models"
	"waysGallery/repositories"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerPost struct {
	PostRepository repositories.PostRepository
}

func HandlerPost(PostRepository repositories.PostRepository) *handlerPost {
	return &handlerPost{PostRepository}
}

func (h *handlerPost) FindPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	post, err := h.PostRepository.FindPost()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// for i, p := range post {
	// 	imagePath := os.Getenv("PATH_FILE") + p.Image
	// 	post[i].Image = imagePath
	// }

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: post}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerPost) GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var post models.Post
	post, err := h.PostRepository.GetPost(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// product.Image = os.Getenv("PATH_FILE") + product.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponsePost(post)}
	json.NewEncoder(w).Encode(response)
}

// func (h *handlerPost) CreatePost(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	// get data user token
// 	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
// 	userId := int(userInfo["id"].(float64))

// 	request := new(postdto.PostRequest)
// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}
// 	// {
// 	// 	Title:       r.FormValue("title"),
// 	// 	Description: r.FormValue("description"),
// 	// 	UserID:      userId,
// 	// 	Images:      []string{r.FormValue("image")},
// 	// 	// ImagesID:    imagesId,
// 	// }

// 	validation := validator.New()
// 	err := validation.Struct(request)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	post := models.Post{
// 		Title:       request.Title,
// 		Description: request.Description,
// 		UserID:      userId,
// 	}

// 	for _, img := range request.Images {
// 		imgData := models.ImageResponse{
// 			Image: img,
// 		}
// 		post.Images = append(post.Images, imgData)
// 	}

// 	post, err = h.PostRepository.CreatePost(post)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	post, _ = h.PostRepository.GetPost(post.ID)

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: post}
// 	json.NewEncoder(w).Encode(response)
// }

func (h *handlerPost) CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	request := new(postdto.PostRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// {
	// 	Title:       r.FormValue("title"),
	// 	Description: r.FormValue("description"),
	// 	UserID:      userId,
	// 	Images:      []string{r.FormValue("image")},
	// 	// ImagesID:    imagesId,
	// }

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	post := models.Post{
		Title:       request.Title,
		Description: request.Description,
		UserID:      userId,
	}

	for _, img := range request.Images {
		imgData := models.ImageResponse{
			Image: img,
		}
		post.Images = append(post.Images, imgData)
	}

	post, err = h.PostRepository.CreatePost(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	post, _ = h.PostRepository.GetPost(post.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: post}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerPost) UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get product id
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// get image filename
	// dataContex := r.Context().Value("dataFile")
	// filename := dataContex.(string)

	var imagesId []int
	imageArr, _ := strconv.Atoi(r.FormValue("imagesId"))

	imagesId = append(imagesId, imageArr)

	// for _, r := range r.FormValue("imagesId") {
	// 	if int(r-'0') >= 0 {
	// 	}
	// }

	request := postdto.PostRequest{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		UserID:      userId,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Get all category data by id []
	// var image []models.Image
	// if len(imagesId) != 0 {
	// 	image, _ = h.PostRepository.FindImagesById(imageID)
	// }

	post, _ := h.PostRepository.GetPost(id)

	// post.Title = request.Title
	// post.Description = request.Description
	// post.Images = image

	// if filename != "false" {
	// 	product.Image = filename
	// }

	post, err = h.PostRepository.UpdatePost(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: post}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerPost) DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	// userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	// userId := int(userInfo["id"].(float64))

	// Get product id
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	post, err := h.PostRepository.GetPost(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	deletePost, err := h.PostRepository.DeletePost(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: deletePost}
	json.NewEncoder(w).Encode(response)
}

func convertResponsePost(p models.Post) models.PostResponse {
	return models.PostResponse{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		User:        p.User,
		Images:      p.Images,
	}
}
