package middlware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/evgeny-tokarev/office_app/backend/internal/config"
	"github.com/evgeny-tokarev/office_app/backend/internal/repositories/user_repository"
	"github.com/evgeny-tokarev/office_app/backend/internal/services/userservice"
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

// Todo: remove user after test
type MiddlewareUnitSuite struct {
	suite.Suite
	user       user_repository.User
	password   string
	router     *mux.Router
	server     *httptest.Server
	querier    *mocks.MockUser
	token      string
	tokenMaker token.Maker
}

func TestMiddlewareUnitSuite(t *testing.T) {
	suite.Run(t, &MiddlewareUnitSuite{})
}

func setupTestServer(mus *MiddlewareUnitSuite) {
	querierMock := mocks.NewMockUser(mus.T())
	cfg := config.Config{
		JwtSecret: util.RandomString(32),
	}
	userService, err := userservice.New(querierMock, cfg)
	require.NoError(mus.T(), err)
	mus.router = mux.NewRouter()
	userService.SetHandlers(mus.router, mus.router)
	mus.router.Use(TokenMiddleware(userService))
	mus.tokenMaker = userService.TokenMaker
	mus.router.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)
	mus.querier = querierMock
	mus.server = httptest.NewServer(mus.router)
	defer mus.server.Close()
}

func (mus *MiddlewareUnitSuite) SetupSuite() {
	setupTestServer(mus)
}

func (mus *MiddlewareUnitSuite) SetupTest() {
	var err error
	mus.password, mus.user, err = CreateRandomUser()
	require.NoError(mus.T(), err)
	mus.querier.EXPECT().CreateUser(mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, params user_repository.CreateUserParams) (user_repository.User, error) {
		err := util.CheckPassword(mus.password, params.HashedPassword)
		if err != nil {
			return user_repository.User{}, err
		}
		return mus.user, nil
	}).Times(1)
	requestBody := fmt.Sprintf(`{"Name": "%s", "Email": "%s", "Role": "%s", "Password": "%s"}`, mus.user.Name, mus.user.Email, mus.user.Role, mus.password)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/user", mus.server.URL), strings.NewReader(requestBody))
	if err != nil {
		mus.FailNowf("Error sending request: ", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	mus.router.ServeHTTP(recorder, req)
	responseBody := recorder.Body.Bytes()
	var userResponse userservice.CreateResponse
	err = json.Unmarshal(responseBody, &userResponse)
	if err != nil {
		mus.FailNowf("Error unmarshalling response: ", err.Error())
	}
	mus.Assert().NoError(err)
	mus.Assert().Equal(http.StatusCreated, recorder.Code)
	mus.token = userResponse.Token
}

func addAuthorization(
	t *testing.T,
	request *http.Request,
	tokenMaker token.Maker,
	authorizationType string,
	username string,
	role string,
	duration time.Duration,
) {
	authToken, err := tokenMaker.CreateToken(username, role, duration)
	require.NoError(t, err)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, authToken)
	request.Header.Set(util.AuthorizationType, authorizationHeader)
}

func (mus *MiddlewareUnitSuite) TestAuthMiddlware() {
	testCases := []struct {
		name          string
		setUpAuth     func(t *testing.T, request *http.Request, maker token.Maker)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{{
		name: "OK",
		setUpAuth: func(t *testing.T, request *http.Request, maker token.Maker) {
			addAuthorization(
				mus.T(),
				request,
				maker,
				util.AuthorizationTypeBearer,
				mus.user.Name,
				mus.user.Role,
				time.Minute,
			)
		},
		checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
			status := recorder.Result().StatusCode
			require.Equal(t, status, http.StatusOK)
		},
	},
		{
			name:      "NoToken",
			setUpAuth: func(t *testing.T, request *http.Request, maker token.Maker) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				status := recorder.Result().StatusCode
				require.Equal(t, http.StatusUnauthorized, status)
			},
		},
		{
			name: "UnsupportedAuthorization",
			setUpAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(
					mus.T(),
					request,
					tokenMaker,
					"unsupported",
					mus.user.Name,
					mus.user.Role,
					time.Minute,
				)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InvalidAuthorizationFormat",
			setUpAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(
					mus.T(),
					request,
					tokenMaker,
					"",
					mus.user.Name,
					mus.user.Role,
					time.Minute,
				)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "ExpiredToken",
			setUpAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(
					mus.T(),
					request,
					tokenMaker,
					"",
					mus.user.Name,
					mus.user.Role,
					-time.Minute,
				)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		}}
	for i := range testCases {
		fmt.Printf("Starting test case %d \n", i)
		tc := testCases[i]
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/auth", mus.server.URL), nil)
		if err != nil {
			mus.FailNowf("Error sending request: ", err.Error())
		}
		tc.setUpAuth(mus.T(), req, mus.tokenMaker)
		recorder := httptest.NewRecorder()
		mus.router.ServeHTTP(recorder, req)
		tc.checkResponse(mus.T(), recorder)
	}
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
	}, nil
}
