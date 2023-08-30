package employeeservice

import (
	"encoding/json"
	"github.com/evgeny-tokarev/office_app/backend/internal/repositories"
	"github.com/evgeny-tokarev/office_app/backend/internal/repositories/employeerepository"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type EmployeeService struct {
	employeeRepository employeerepository.EmployeeRepository
}

func New(employeeRepository employeerepository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		employeeRepository: employeeRepository,
	}
}

func (es *EmployeeService) SetHandlers(router *mux.Router) {

	router.HandleFunc("/employees", es.Create).Methods(http.MethodPost)
	router.HandleFunc("/employees/{id}", es.Get).Methods(http.MethodGet)
	router.HandleFunc("/employees", es.List).Methods(http.MethodGet)
	router.HandleFunc("/employees", es.Update).Methods(http.MethodPut)
	router.HandleFunc("/employees/{id}", es.Delete).Methods(http.MethodDelete)
}

type CreateRequest struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	OfficeID int64  `json:"office_id"`
}

func (es *EmployeeService) Create(w http.ResponseWriter, r *http.Request) {
	req := &CreateRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	e := employeerepository.Employee{
		Name:     req.Name,
		Age:      req.Age,
		OfficeID: req.OfficeID,
	}
	if req.Name == "" {
		repositories.SendTranscribedError(w, "name field is required", http.StatusBadRequest)
		return
	}

	if req.Age <= 0 {
		repositories.SendTranscribedError(w, "age field is required and must be a positive value", http.StatusBadRequest)
		return
	}

	if err := es.employeeRepository.Create(r.Context(), &e); err != nil {
		repositories.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type GetResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	OfficeID int64  `json:"office_id"`
}

func (es *EmployeeService) Get(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		repositories.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}

	e, err := es.employeeRepository.Get(r.Context(), int64(id))
	if err != nil {
		repositories.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(&GetResponse{
		ID:       e.ID,
		Name:     e.Name,
		Age:      e.Age,
		OfficeID: e.OfficeID,
	})
}

type ListResponse struct {
	List []GetResponse
}

func (es *EmployeeService) List(w http.ResponseWriter, r *http.Request) {

	officeID, err := strconv.Atoi(r.URL.Query().Get("office_id"))
	if err != nil {
		repositories.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}

	list, err := es.employeeRepository.List(r.Context(), int64(officeID))
	if err != nil {
		repositories.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := make([]GetResponse, 0, len(list))
	for _, e := range list {
		response = append(response, GetResponse{
			ID:       e.ID,
			Name:     e.Name,
			Age:      e.Age,
			OfficeID: e.OfficeID,
		})
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(ListResponse{List: response})

}

type UpdateRequest struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	OfficeID int64  `json:"office_id"`
}

func (es *EmployeeService) Update(w http.ResponseWriter, r *http.Request) {
	req := &UpdateRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		repositories.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	e := employeerepository.Employee{
		ID:       req.ID,
		Name:     req.Name,
		Age:      req.Age,
		OfficeID: req.OfficeID,
	}
	if err := es.employeeRepository.Update(r.Context(), &e); err != nil {
		repositories.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (es *EmployeeService) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		repositories.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := es.employeeRepository.Delete(r.Context(), int64(id)); err != nil {
		repositories.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
