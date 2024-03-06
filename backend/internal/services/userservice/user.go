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
	"strconv"
	"time"
)

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

func (us *UserService) SetHandlers(router *mux.Router) {
	router.HandleFunc("/user", us.Create).Methods(http.MethodPost)
	router.HandleFunc("/user/login", us.Login).Methods(http.MethodPost)
	//router.HandleFunc("/offices/{id}", ofs.Get).Methods(http.MethodGet)
	//router.HandleFunc("/offices", ofs.List).Methods(http.MethodGet)
	//router.HandleFunc("/offices", ofs.Update).Methods(http.MethodPut)
	//router.HandleFunc("/offices/{id}", ofs.Delete).Methods(http.MethodDelete)
	//router.HandleFunc("/offices/{id}/image", ofs.Upload).Methods(http.MethodPost)
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

	token, err := us.TokenMaker.CreateToken(user.Name, time.Hour)
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
		Token: token,
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
		if err == sql.ErrNoRows {
			util.SendTranscribedError(w, err.Error(), http.StatusNotFound)
			return
		}
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user2, err := us.userRepository.GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
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

	token, err := us.TokenMaker.CreateToken(user1.Name, time.Hour)
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
		Token: token,
	}); err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

type GetResponse struct {
	ID                int64  `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Role              string `json:"role"`
	HashedPassword    string `json:"hashed_password"`
	PasswordChangedAt string `json:"password_changed_at"`
	CreatedAt         string `json:"created_at"`
	ImgFile           string `json:"img_file"`
}

func (us *UserService) Get(w http.ResponseWriter, r *http.Request) {
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
	_ = json.NewEncoder(w).Encode(&GetResponse{
		ID:                u.ID,
		Name:              u.Name,
		Email:             u.Email,
		Role:              u.Role,
		HashedPassword:    u.HashedPassword,
		PasswordChangedAt: u.PasswordChangedAt.Format(util.TimeLayout),
		CreatedAt:         u.CreatedAt.Format(util.TimeLayout),
		ImgFile:           u.ImgFile.String,
	})
}

//func LogoutHandler(w http.ResponseWriter, r *http.Request) {
//	http.SetCookie(w, &http.Cookie{
//		Name:     "token",
//		Value:    "",
//		Path:     "/",
//		HttpOnly: true,
//		Expires:  time.Unix(0, 0),
//	})
//
//	http.Redirect(w, r, "/login", http.StatusFound)
//}
