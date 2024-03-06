package middlware

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MiddlewareUnitSuite struct {
	suite.Suite
	router *mux.Router
	server *httptest.Server
}

func TestMiddlewareUnitSuite(t *testing.T) {
	suite.Run(t, &MiddlewareUnitSuite{})
}

func (mus *MiddlewareUnitSuite) SetupSuite() {
	setupTestServer(mus)
}

//func (osu *MiddlewareUnitSuite) SetupTest() {
//	osu.office = RandomOffice()
//}

func setupTestServer(mus *MiddlewareUnitSuite) {
	router := mux.NewRouter()
	mus.router = router
	mus.server = httptest.NewServer(router)
	defer mus.server.Close()
}

func (mus *MiddlewareUnitSuite) TestAuthMiddlware(t *testing.T) {
	testCases := []struct {
		name          string
		handleFunc    func(http.ResponseWriter, *http.Request)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{{
		name:          "OK",
		handleFunc:    func(http.ResponseWriter, *http.Request) {},
		checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {},
	}}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			authPath := "/auth"
			mus.router.HandleFunc(authPath, tc.handleFunc)
		})
	}
}
