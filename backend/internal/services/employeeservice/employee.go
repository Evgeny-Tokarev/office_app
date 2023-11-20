package employeeservice

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"github.com/evgeny-tokarev/office_app/backend/internal/repositories/employee_repository"
	"github.com/evgeny-tokarev/office_app/backend/util"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type EmployeeService struct {
	employeeRepository employee_repository.Queries
}

func New(employeeRepository employee_repository.Queries) *EmployeeService {
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
	router.HandleFunc("/employees/{id}/image", es.Upload).Methods(http.MethodPost)
}

type CreateRequest struct {
	Name     string `json:"name"`
	Age      int32  `json:"age"`
	OfficeID int64  `json:"office_id"`
}

func (es *EmployeeService) Upload(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}
	oldImagePath, err := es.employeeRepository.GetImagePath(r.Context(), id)
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = r.ParseMultipartForm(10 << 20) // 10MB is the maximum size of the uploaded file
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	img, err := imaging.Decode(file)
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}
	var buffer bytes.Buffer

	err = webp.Encode(&buffer, img, nil)
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}
	parts := strings.Split(header.Filename, ".")
	var filename string
	if len(parts) > 1 {
		filename = strings.Join(parts[:len(parts)-1], ".")
	} else {
		filename = parts[0]
	}

	webpFilePath := util.GetUniqueFileName("./images/" + filename + ".webp")
	webpFile, err := os.Create(webpFilePath)
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer webpFile.Close()

	_, err = webpFile.Write(buffer.Bytes())
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}
	params := employee_repository.AttachePhotoParams{
		ImgFile: sql.NullString{String: filepath.Join("images", filepath.Base(webpFilePath)), Valid: true},
		ID:      id,
	}
	if err := es.employeeRepository.AttachePhoto(r.Context(), params); err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	} else if oldImagePath != "" {
		err = os.Remove(oldImagePath)
		if err != nil {
			util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}

func (es *EmployeeService) Create(w http.ResponseWriter, r *http.Request) {
	req := &CreateRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	e := employee_repository.CreateEmployeeParams{
		Name:     req.Name,
		Age:      req.Age,
		OfficeID: req.OfficeID,
	}
	if req.Name == "" {
		util.SendTranscribedError(w, "name field is required", http.StatusBadRequest)
		return
	}

	if req.Age <= 0 {
		util.SendTranscribedError(w, "age field is required and must be a positive value", http.StatusBadRequest)
		return
	}
	employee, err := es.employeeRepository.CreateEmployee(r.Context(), e)
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	responseBody := map[string]int64{"id": employee.ID}
	if err := json.NewEncoder(w).Encode(responseBody); err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type GetResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Age      int32  `json:"age"`
	OfficeID int64  `json:"office_id"`
	Photo    string `json:"photo"`
}

func (es *EmployeeService) Get(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}

	e, err := es.employeeRepository.GetEmployee(r.Context(), int64(id))
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(&GetResponse{
		ID:       e.ID,
		Name:     e.Name,
		Age:      e.Age,
		OfficeID: e.OfficeID,
		Photo:    e.ImgFile.String,
	})
}

type ListResponse struct {
	List []GetResponse
}

func (es *EmployeeService) List(w http.ResponseWriter, r *http.Request) {

	officeID, err := strconv.Atoi(r.URL.Query().Get("office_id"))
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}

	list, err := es.employeeRepository.ListEmployees(r.Context(), int64(officeID))
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := make([]GetResponse, 0, len(list))
	for _, e := range list {
		response = append(response, GetResponse{
			ID:       e.ID,
			Name:     e.Name,
			Age:      e.Age,
			OfficeID: e.OfficeID,
			Photo:    e.ImgFile.String,
		})
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)

}

type UpdateRequest struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Age      int32  `json:"age"`
	OfficeID int64  `json:"office_id"`
}

func (es *EmployeeService) Update(w http.ResponseWriter, r *http.Request) {
	req := &UpdateRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	e := employee_repository.UpdateEmployeeParams{
		ID:   req.ID,
		Name: req.Name,
		Age:  req.Age,
	}
	if err := es.employeeRepository.UpdateEmployee(r.Context(), e); err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (es *EmployeeService) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := es.employeeRepository.DeleteEmployee(r.Context(), int64(id)); err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	imagePath, err := es.employeeRepository.GetImagePath(r.Context(), int64(id))
	if err := os.Remove(imagePath); err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
