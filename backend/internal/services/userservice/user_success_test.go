package userservice

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/evgeny-tokarev/office_app/backend/internal/config"
	"github.com/evgeny-tokarev/office_app/backend/internal/repositories/user_repository"
	"github.com/evgeny-tokarev/office_app/backend/internal/token"
	"github.com/evgeny-tokarev/office_app/backend/mocks"
	"github.com/evgeny-tokarev/office_app/backend/util"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type UserServiceUnitSuite struct {
	suite.Suite
	user       user_repository.User
	password   string
	router     *mux.Router
	server     *httptest.Server
	querier    *mocks.MockUser
	token      string
	tokenMaker token.Maker
}

func TestUserServiceUnitSuite(t *testing.T) {
	suite.Run(t, &UserServiceUnitSuite{})
}

func (usu *UserServiceUnitSuite) SetupSuite() {
	setupTestServer(usu)
}

func (usu *UserServiceUnitSuite) SetupTest() {
	var err error
	usu.password, usu.user, err = CreateRandomUser()
	require.NoError(usu.T(), err)
}

func setupTestServer(usu *UserServiceUnitSuite) {
	querierMock := mocks.NewMockUser(usu.T())
	cfg := config.Config{
		JwtSecret: util.RandomString(32),
	}
	userService, err := New(querierMock, cfg)
	require.NoError(usu.T(), err)
	router := mux.NewRouter()
	userService.SetHandlers(router, router)
	usu.router = router
	usu.querier = querierMock
	usu.server = httptest.NewServer(router)
	defer usu.server.Close()
}

func (usu *UserServiceUnitSuite) SetupServer() {
	setupTestServer(usu)
}

func (usu *UserServiceUnitSuite) TestGetuser() {
	usu.querier.EXPECT().GetUserById(mock.Anything, usu.user.ID).Return(usu.user, nil).Times(1)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/user/%d", usu.server.URL, usu.user.ID), nil)
	if err != nil {
		usu.FailNowf("Error sending request: ", err.Error())
	}
	ctx := context.WithValue(req.Context(), "owner", "moderator")
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	usu.router.ServeHTTP(recorder, req)
	usu.Assert().NoError(err)
	usu.Assert().Equal(http.StatusOK, recorder.Code)
	var response GetResponse
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	createdAt, err1 := time.Parse(util.TimeLayout, response.CreatedAt)
	usu.Assert().NoError(err)
	usu.Assert().NoError(err1)
	usu.Assert().Equal(usu.user.ID, response.ID)
	usu.Assert().Equal(usu.user.Name, response.Name)
	usu.Assert().Equal(usu.user.Email, response.Email)
	usu.Assert().Equal(usu.user.Role, response.Role)
	usu.Assert().WithinDuration(usu.user.CreatedAt, createdAt, time.Second)
}

func (usu *UserServiceUnitSuite) TestCreateUser() {
	usu.querier.EXPECT().CreateUser(mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, params user_repository.CreateUserParams) (user_repository.User, error) {
		err := util.CheckPassword(usu.password, params.HashedPassword)
		if err != nil {
			return user_repository.User{}, err
		}
		return usu.user, nil
	}).Times(1)
	requestBody := fmt.Sprintf(`{"Name": "%s", "Email": "%s", "Role": "%s", "Password": "%s"}`, usu.user.Name, usu.user.Email, usu.user.Role, usu.password)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/user", usu.server.URL), strings.NewReader(requestBody))
	if err != nil {
		usu.FailNowf("Error sending request: ", err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	usu.router.ServeHTTP(recorder, req)
	usu.Assert().NoError(err)
	usu.Assert().Equal(http.StatusCreated, recorder.Code)
	var response CreateResponse
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	usu.Assert().NoError(err)
	usu.Assert().NotNil(response.ID)
	usu.Assert().Equal(usu.user.Name, response.Name)
	usu.Assert().Equal(usu.user.Email, response.Email)
}

func (usu *UserServiceUnitSuite) TestDeleteUser() {
	usu.querier.EXPECT().GetImagePath(mock.Anything, usu.user.ID).Return(usu.user.ImgFile.String, nil).Times(1)
	usu.querier.EXPECT().DeleteUser(mock.Anything, usu.user.ID).Return(nil).Times(1)
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/user/%d", usu.server.URL, usu.user.ID), nil)
	if err != nil {
		usu.FailNowf("Error sending request: ", err.Error())
	}
	ctx := context.WithValue(req.Context(), "owner", "admin")
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	usu.router.ServeHTTP(recorder, req)
	responseBody := recorder.Body.Bytes()
	var errorResponse util.ErrorResponse
	err = json.Unmarshal(responseBody, &errorResponse)
	if err != nil {
		usu.FailNowf("Error unmarshalling response: ", err.Error())
	}
	usu.Assert().NoError(err)
	usu.Assert().Equal(recorder.Code, http.StatusOK)
}

func (usu *UserServiceUnitSuite) TestUpdateUser() {
	usu.querier.EXPECT().UpdateUser(mock.Anything, user_repository.UpdateUserParams{
		ID:    usu.user.ID,
		Name:  usu.user.Name,
		Email: usu.user.Email,
		Role:  usu.user.Role,
	}).Return(nil).Times(1)
	requestBody := fmt.Sprintf(`{"ID": %d, "Name": "%s", "Email": "%s", "Role": "%s"}`, usu.user.ID, usu.user.Name, usu.user.Email, usu.user.Role)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/user", usu.server.URL), strings.NewReader(requestBody))
	if err != nil {
		usu.FailNowf("Error sending request: ", err.Error())
	}
	ctx := context.WithValue(req.Context(), "owner", "admin")
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	usu.router.ServeHTTP(recorder, req)
	usu.Assert().NoError(err)
	usu.Assert().Equal(http.StatusOK, recorder.Code)
}

func CreateRandomUser() (string, user_repository.User, error) {
	password := util.RandomString(32)
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return password, user_repository.User{}, err
	}
	return password, user_repository.User{
		ID:                util.RandomInt(10000, 100000),
		Name:              util.RandomString(10),
		Email:             util.RandomString(10),
		Role:              "admin",
		HashedPassword:    hashedPassword,
		PasswordChangedAt: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:         time.Now(),
		ImgFile:           sql.NullString{String: util.RandomString(10), Valid: true},
	}, nil
}
