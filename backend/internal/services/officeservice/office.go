package officeservice

import (
	"encoding/json"
	"fmt"
	"github.com/evgeny-tokarev/office_app/backend/internal/repositories"
	"github.com/evgeny-tokarev/office_app/backend/internal/repositories/officerepository"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

type OfficeService struct {
	officeRepository officerepository.OfficeRepository
}

func New(officeRepository officerepository.OfficeRepository) *OfficeService {
	return &OfficeService{
		officeRepository: officeRepository,
	}
}

func (os *OfficeService) SetHandlers(router *mux.Router) {
	router.HandleFunc("/offices", os.Create).Methods(http.MethodPost)
	router.HandleFunc("/offices/{id}", os.Get).Methods(http.MethodGet)
	router.HandleFunc("/offices", os.List).Methods(http.MethodGet)
	router.HandleFunc("/offices", os.Update).Methods(http.MethodPut)
	router.HandleFunc("/offices/{id}", os.Delete).Methods(http.MethodDelete)
}

type CreateRequest struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func SendTranscribedError(w http.ResponseWriter, msg string, status int) {
	errorResponse := ErrorResponse{
		Status:  status,
		Message: msg,
	}
	responseBody, err := json.Marshal(errorResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	_, err = w.Write(responseBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (os *OfficeService) Create(w http.ResponseWriter, r *http.Request) {

	req := &CreateRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		repositories.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}
	var createdAt, updatedAt time.Time
	var err error
	createdAt, err = time.Parse(time.RFC3339, req.CreatedAt)
	updatedAt, err = time.Parse(time.RFC3339, req.UpdatedAt)
	if err != nil {
		fmt.Println("Time error: ", err)
		repositories.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}

	o := officerepository.Office{
		Name:      req.Name,
		Address:   req.Address,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	fmt.Println("Request: ", createdAt)

	if req.Name == "" || req.Address == "" || req.UpdatedAt == "" || req.CreatedAt == "" {
		SendTranscribedError(w, "All fields are required", http.StatusBadRequest)
		return
	}

	if err := os.officeRepository.Create(r.Context(), &o); err != nil {
		SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type GetResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (os *OfficeService) Get(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}

	e, err := os.officeRepository.Get(r.Context(), int64(id))
	if err != nil {
		fmt.Println(err)
		SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(&GetResponse{
		ID:        e.ID,
		Name:      e.Name,
		Address:   e.Address,
		CreatedAt: e.CreatedAt.Format("2006-01-02 15:04:05 -0700 MST"),
		UpdatedAt: e.UpdatedAt.Format("2006-01-02 15:04:05 -0700 MST"),
	})
}

func (os *OfficeService) List(w http.ResponseWriter, r *http.Request) {
	list, err := os.officeRepository.List(r.Context())
	if err != nil {
		SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := make([]GetResponse, 0, len(list))
	for _, e := range list {
		response = append(response, GetResponse{
			ID:        e.ID,
			Name:      e.Name,
			Address:   e.Address,
			CreatedAt: e.CreatedAt.Format("2006-01-02 15:04:05 -0700 MST"),
			UpdatedAt: e.UpdatedAt.Format("2006-01-02 15:04:05 -0700 MST"),
		})
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)

}

type UpdateRequest struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	UpdatedAt string `json:"updated_at"`
}

func (os *OfficeService) Update(w http.ResponseWriter, r *http.Request) {
	req := &UpdateRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var updatedAt time.Time
	var err error
	updatedAt, err = time.Parse(time.RFC3339, req.UpdatedAt)
	if err != nil {
		repositories.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}

	o := officerepository.Office{
		ID:        req.ID,
		Name:      req.Name,
		Address:   req.Address,
		UpdatedAt: updatedAt,
	}

	if err := os.officeRepository.Update(r.Context(), &o); err != nil {
		SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (os *OfficeService) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := os.officeRepository.Delete(r.Context(), int64(id)); err != nil {
		SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
