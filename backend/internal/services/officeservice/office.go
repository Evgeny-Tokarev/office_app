package officeservice

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"github.com/evgeny-tokarev/office_app/backend/internal/repositories/office_repository"
	"github.com/evgeny-tokarev/office_app/backend/util"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type OfficeService struct {
	officeRepository office_repository.Queries
}

func New(officeRepository office_repository.Queries) *OfficeService {
	return &OfficeService{
		officeRepository: officeRepository,
	}
}

func (ofs *OfficeService) SetHandlers(router *mux.Router) {
	router.HandleFunc("/offices", ofs.Create).Methods(http.MethodPost)
	router.HandleFunc("/offices/{id}", ofs.Get).Methods(http.MethodGet)
	router.HandleFunc("/offices", ofs.List).Methods(http.MethodGet)
	router.HandleFunc("/offices", ofs.Update).Methods(http.MethodPut)
	router.HandleFunc("/offices/{id}", ofs.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/offices/{id}/image", ofs.Upload).Methods(http.MethodPost)
}

type CreateRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (ofs *OfficeService) Upload(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}
	oldImagePath, err := ofs.officeRepository.GetImagePath(r.Context(), id)
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
	fmt.Println(filename, webpFilePath)
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
	params := office_repository.AttachePhotoParams{
		ImgFile: sql.NullString{String: filepath.Join("images", filepath.Base(webpFilePath)), Valid: true},
		ID:      id,
	}
	if err := ofs.officeRepository.AttachePhoto(r.Context(), params); err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	} else if oldImagePath != "" {
		err = os.Remove(oldImagePath)
		if err != nil {
			util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)
}

type CreateResponse struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	ImgFile   string `json:"img_file"`
}

func (ofs *OfficeService) Create(w http.ResponseWriter, r *http.Request) {

	req := &CreateRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}

	o := office_repository.CreateOfficeParams{
		Name:    req.Name,
		Address: req.Address,
	}

	if req.Name == "" || req.Address == "" {
		util.SendTranscribedError(w, "all fields are required", http.StatusBadRequest)
		return
	}
	office, err := ofs.officeRepository.CreateOffice(r.Context(), o)
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(&CreateResponse{
		Name:      office.Name,
		Address:   office.Address,
		CreatedAt: office.CreatedAt.Format("2006-01-02 15:04:05 -0700 MST"),
		UpdatedAt: office.UpdatedAt.Format("2006-01-02 15:04:05 -0700 MST"),
		ImgFile:   util.ConvertToRegularString(office.ImgFile),
	}); err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type GetResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Photo     string `json:"photo"`
}

func (ofs *OfficeService) Get(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}

	e, err := ofs.officeRepository.GetOffice(r.Context(), int64(id))
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
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
		Photo:     e.ImgFile.String,
	})
}

func (ofs *OfficeService) List(w http.ResponseWriter, r *http.Request) {
	list, err := ofs.officeRepository.ListOffices(r.Context())
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
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
			Photo:     e.ImgFile.String,
		})
	}
	sort.SliceStable(response, func(i, j int) bool {
		return response[i].ID < response[j].ID
	})
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

func (ofs *OfficeService) Update(w http.ResponseWriter, r *http.Request) {
	req := &UpdateRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		if err.Error() == errors.New("EOF").Error() {
			err = errors.New("empty office body")
		}
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if req.Name == "" || req.Address == "" {
		util.SendTranscribedError(w, "all fields are required", http.StatusBadRequest)
		return
	}

	o := office_repository.UpdateOfficeParams{
		ID:      req.ID,
		Name:    req.Name,
		Address: req.Address,
	}

	if err := ofs.officeRepository.UpdateOffice(r.Context(), o); err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ofs *OfficeService) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ofs.officeRepository.DeleteOffice(r.Context(), int64(id)); err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	imagePath, err := ofs.officeRepository.GetImagePath(r.Context(), int64(id))
	if err := os.Remove(imagePath); err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
