package userservice

import (
	"encoding/json"
	"github.com/evgeny-tokarev/office_app/backend/internal/repositories/user_repository"
	"github.com/evgeny-tokarev/office_app/backend/util"
	"github.com/gorilla/mux"
	"net/http"
)

type UserService struct {
	userRepository user_repository.Querier
}

func New(userRepository user_repository.Querier) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (us *UserService) SetHandlers(router *mux.Router) {
	router.HandleFunc("/user", us.Create).Methods(http.MethodPost)
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
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	Email     string `db:"email"`
	Role      string `db:"role"`
	CreatedAt string `db:"created_at"`
	ImgFile   string `db:"img_file"`
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

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(&CreateResponse{
		ID:        user.ID,
		Name:      user.Name,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05 -0700 MST"),
		ImgFile:   util.ConvertToRegularString(user.ImgFile),
	}); err != nil {
		util.SendTranscribedError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
