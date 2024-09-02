package userservice

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/evgeny-tokarev/office_app/backend/internal/config"
	"github.com/evgeny-tokarev/office_app/backend/internal/repositories/user_repository"
	"github.com/evgeny-tokarev/office_app/backend/internal/token"
	"github.com/evgeny-tokarev/office_app/backend/util"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"
)

//todo: create tests for all methods

type UserService struct {
	userRepository user_repository.Querier
	TokenMaker     token.Maker
}

func New(userRepository user_repository.Querier, cfg config.Config) (*UserService, error) {
	tokenMaker, err := token.NewJWTMaker(cfg.JwtSecret)
	if err != nil {
		return nil, err
	}
	return &UserService{
		userRepository: userRepository,
		TokenMaker:     tokenMaker,
	}, nil
}

func (us *UserService) SetHandlers(router, authRoutes *mux.Router) {
	authRoutes.HandleFunc("/user/current", us.GetCurrent).Methods(http.MethodGet)
	authRoutes.HandleFunc("/user/{id}", us.Get).Methods(http.MethodGet)
	router.HandleFunc("/user", us.Create).Methods(http.MethodPost)
	router.HandleFunc("/user/login", us.Login).Methods(http.MethodPost)
	router.HandleFunc("/user", us.List).Methods(http.MethodGet)
	router.HandleFunc("/user", us.Update).Methods(http.MethodPut)
	router.HandleFunc("/user/{id}", us.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/user/{id}/image", us.Upload).Methods(http.MethodPost)
}

type CreateRequest struct {
	Name     string `db:"name"`
	Email    string `db:"email"`
	Role     string `db:"role"`
	Password string `db:"password"`
}

type CreateResponse struct {
	user_repository.User
	Token string `db:"token"`
}

func (us *UserService) Create(w http.ResponseWriter, r *http.Request) {
	req := &CreateRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.Email == "" || req.Role == "" || req.Password == "" {
		util.SendTranscribedError(w, "all fields are required", http.StatusBadRequest)
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	o := user_repository.CreateUserParams{
		Name:           req.Name,
		Email:          req.Email,
		Role:           req.Role,
		HashedPassword: hashedPassword,
	}

	user, err := us.userRepository.CreateUser(r.Context(), o)
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t, err := us.TokenMaker.CreateToken(user.ID, user.Role, time.Hour*24)
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(&CreateResponse{
		User: user_repository.User{
			ID:                user.ID,
			Name:              user.Name,
			Email:             user.Email,
			Role:              user.Role,
			HashedPassword:    user.HashedPassword,
			PasswordChangedAt: user.PasswordChangedAt,
			CreatedAt:         user.CreatedAt,
			ImgFile:           user.ImgFile,
		},
		Token: t,
	}); err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type loginRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	User  UserForLoginResponse `json:"user"`
	Token string               `json:"token"`
}

type UserForLoginResponse struct {
	ID                int64          `db:"id"`
	Name              string         `db:"name"`
	Email             string         `db:"email"`
	Role              string         `db:"role"`
	PasswordChangedAt time.Time      `db:"password_changed_at"`
	CreatedAt         time.Time      `db:"created_at"`
	ImgFile           sql.NullString `db:"img_file"`
}

func (us *UserService) Login(w http.ResponseWriter, r *http.Request) {
	req := &loginRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		if err.Error() == errors.New("EOF").Error() {
			err = errors.New("empty office body")
		}
		util.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		util.SendTranscribedError(w, "all fields are required", http.StatusBadRequest)
		return
	}

	user1, err := us.userRepository.GetUserByName(r.Context(), req.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			util.SendTranscribedError(w, err.Error(), http.StatusNotFound)
			return
		}
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user2, err := us.userRepository.GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			util.SendTranscribedError(w, err.Error(), http.StatusNotFound)
			return
		}
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user1.ID != user2.ID {
		util.SendTranscribedError(w, "wrong username or email", http.StatusUnprocessableEntity)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user1.HashedPassword), []byte(req.Password))
	if err != nil {
		util.SendTranscribedError(w, "incorrect password", http.StatusUnauthorized)
		return
	}

	t, err := us.TokenMaker.CreateToken(user1.ID, user1.Role, time.Hour*24)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	user := UserForLoginResponse{
		ID:                user1.ID,
		Name:              user1.Name,
		Email:             user1.Email,
		Role:              user1.Role,
		PasswordChangedAt: user1.PasswordChangedAt,
		CreatedAt:         user1.CreatedAt,
		ImgFile:           user1.ImgFile,
	}
	if err := json.NewEncoder(w).Encode(&loginResponse{
		User:  user,
		Token: t,
	}); err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

type GetUserResponse struct {
	ID                int64  `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Role              string `json:"role"`
	PasswordChangedAt string `json:"password_changed_at"`
	CreatedAt         string `json:"created_at"`
	ImgFile           string `json:"img_file"`
}

func (us *UserService) Get(w http.ResponseWriter, r *http.Request) {
	if !util.HasAccessRights(w, r, "moderator") {
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := us.userRepository.GetUserById(r.Context(), int64(id))
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(&GetUserResponse{
		ID:                u.ID,
		Name:              u.Name,
		Email:             u.Email,
		Role:              u.Role,
		PasswordChangedAt: u.PasswordChangedAt.Format(util.TimeLayout),
		CreatedAt:         u.CreatedAt.Format(util.TimeLayout),
		ImgFile:           u.ImgFile.String,
	})
}

func (us *UserService) GetCurrent(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(int64)

	u, err := us.userRepository.GetUserById(r.Context(), userId)
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	ticker200 := time.NewTicker(time.Second * 3)
	defer ticker200.Stop()
	<-ticker200.C
	if err = json.NewEncoder(w).Encode(&GetUserResponse{
		ID:                u.ID,
		Name:              u.Name,
		Email:             u.Email,
		Role:              u.Role,
		PasswordChangedAt: u.PasswordChangedAt.Format(util.TimeLayout),
		CreatedAt:         u.CreatedAt.Format(util.TimeLayout),
		ImgFile:           u.ImgFile.String,
	}); err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (us *UserService) List(w http.ResponseWriter, r *http.Request) {
	if !util.HasAccessRights(w, r, "moderator") {
		return
	}

	list, err := us.userRepository.ListUsers(r.Context())
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := make([]GetUserResponse, 0, len(list))
	for _, e := range list {
		response = append(response, GetUserResponse{
			ID:                e.ID,
			Name:              e.Name,
			Email:             e.Email,
			Role:              e.Role,
			PasswordChangedAt: e.PasswordChangedAt.Format(util.TimeLayout),
			CreatedAt:         e.CreatedAt.Format(util.TimeLayout),
			ImgFile:           e.ImgFile.String,
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
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func (us *UserService) Update(w http.ResponseWriter, r *http.Request) {
	if !util.HasAccessRights(w, r, "admin") {
		return
	}

	req := &UpdateRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		if err.Error() == errors.New("EOF").Error() {
			err = errors.New("empty user body")
		}
		util.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Name == "" || req.Email == "" || req.Role == "" {
		util.SendTranscribedError(w, "all fields are required", http.StatusBadRequest)
		return
	}

	o := user_repository.UpdateUserParams{
		ID:    req.ID,
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}
	if err := us.userRepository.UpdateUser(r.Context(), o); err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (us *UserService) Delete(w http.ResponseWriter, r *http.Request) {
	if !util.HasAccessRights(w, r, "admin") {
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}

	imagePath, err := us.userRepository.GetImagePath(r.Context(), int64(id))
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if imagePath != "" {
		if err := os.Remove(imagePath); err != nil {
			util.WriteResponse(w, http.StatusOK, map[string]interface{}{
				"Status":  http.StatusOK,
				"Message": "User deleted, but image file deletion failed",
			})
		}
	}

	if err := us.userRepository.DeleteUser(r.Context(), int64(id)); err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (us *UserService) Upload(w http.ResponseWriter, r *http.Request) {
	if !util.HasAccessRights(w, r, "admin") {
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}
	oldImagePath, err := us.userRepository.GetImagePath(r.Context(), id)
	if err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusBadRequest)
		return
	}

	webpFilePath, status, err := util.SaveImage(r)
	if err != nil {
		util.SendTranscribedError(w, err.Error(), status)
	}
	params := user_repository.AttachePhotoParams{
		ImgFile: sql.NullString{String: filepath.Join("images", filepath.Base(webpFilePath)), Valid: true},
		ID:      id,
	}
	if err := us.userRepository.AttachePhoto(r.Context(), params); err != nil {
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
