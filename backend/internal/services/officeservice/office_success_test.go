package officeservice

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/evgeny-tokarev/office_app/backend/internal/repositories/office_repository"
	"github.com/evgeny-tokarev/office_app/backend/mocks"
	"github.com/evgeny-tokarev/office_app/backend/util"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type OfficeServiceUnitSuite struct {
	suite.Suite
	router  *mux.Router
	querier *mocks.MockOffice
	server  *httptest.Server
	office  office_repository.Office
}

func TestOfficeServiceUnitSuite(t *testing.T) {
	suite.Run(t, &OfficeServiceUnitSuite{})
}

func (osu *OfficeServiceUnitSuite) SetupSuite() {
	setupTestServer(osu)
}

func (osu *OfficeServiceUnitSuite) SetupTest() {
	osu.office = RandomOffice()
}

func setupTestServer(osu *OfficeServiceUnitSuite) {
	querierMock := mocks.NewMockOffice(osu.T())
	officeService := New(querierMock)
	router := mux.NewRouter()
	officeService.SetHandlers(router, router)
	osu.router = router
	osu.querier = querierMock
	osu.server = httptest.NewServer(router)
	defer osu.server.Close()
}

func (osu *OfficeServiceUnitSuite) SetupServer() {
	setupTestServer(osu)
}

func (osu *OfficeServiceUnitSuite) TestGetOffice() {
	osu.querier.EXPECT().GetOffice(mock.Anything, osu.office.ID).Return(osu.office, nil).Times(1)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/offices/%d", osu.server.URL, osu.office.ID), nil)
	if err != nil {
		osu.FailNowf("Error sending request: ", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	osu.router.ServeHTTP(recorder, req)
	osu.Assert().NoError(err)
	osu.Assert().Equal(http.StatusOK, recorder.Code)
	var response GetResponse
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	createdAt, err1 := time.Parse(util.TimeLayout, response.CreatedAt)
	UpdatedAt, err2 := time.Parse(util.TimeLayout, response.UpdatedAt)
	osu.Assert().NoError(err)
	osu.Assert().NoError(err1)
	osu.Assert().NoError(err2)
	osu.Assert().Equal(osu.office.ID, response.ID)
	osu.Assert().Equal(osu.office.Name, response.Name)
	osu.Assert().Equal(osu.office.Address, response.Address)
	osu.Assert().WithinDuration(osu.office.CreatedAt, createdAt, time.Second)
	osu.Assert().WithinDuration(osu.office.UpdatedAt, UpdatedAt, time.Second)
}

func (osu *OfficeServiceUnitSuite) TestCreateOffice() {
	osu.querier.EXPECT().CreateOffice(mock.Anything, office_repository.CreateOfficeParams{
		Name:    osu.office.Name,
		Address: osu.office.Address,
	}).Return(osu.office, nil).Times(1)
	requestBody := fmt.Sprintf(`{"Name": "%s", "Address": "%s"}`, osu.office.Name, osu.office.Address)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/offices", osu.server.URL), strings.NewReader(requestBody))
	if err != nil {
		osu.FailNowf("Error sending request: ", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	osu.router.ServeHTTP(recorder, req)
	osu.Assert().NoError(err)
	osu.Assert().Equal(http.StatusCreated, recorder.Code)
	var response CreateResponse
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	createdAt, err1 := time.Parse(util.TimeLayout, response.CreatedAt)
	UpdatedAt, err2 := time.Parse(util.TimeLayout, response.UpdatedAt)
	osu.Assert().NoError(err)
	osu.Assert().NoError(err1)
	osu.Assert().NoError(err2)
	osu.Assert().NotNil(response.ID)
	osu.Assert().Equal(osu.office.Name, response.Name)
	osu.Assert().Equal(osu.office.Address, response.Address)
	osu.Assert().NotNil(createdAt)
	osu.Assert().NotNil(UpdatedAt)
}

func (osu *OfficeServiceUnitSuite) TestDeleteOffice() {
	osu.querier.EXPECT().GetImagePath(mock.Anything, osu.office.ID).Return(osu.office.ImgFile.String, nil).Times(1)
	osu.querier.EXPECT().DeleteOffice(mock.Anything, osu.office.ID).Return(nil).Times(1)
	requestBody := fmt.Sprintf(`{"Name": "%s", "Address": "%s"}`, osu.office.Name, osu.office.Address)
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/offices/%d", osu.server.URL, osu.office.ID), strings.NewReader(requestBody))
	if err != nil {
		osu.FailNowf("Error sending request: ", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	osu.router.ServeHTTP(recorder, req)
	responseBody := recorder.Body.Bytes()
	var errorResponse util.ErrorResponse
	err = json.Unmarshal(responseBody, &errorResponse)
	if err != nil {
		osu.FailNowf("Error unmarshalling response: ", err.Error())
	}
	osu.Assert().NoError(err)
	osu.Assert().Equal(recorder.Code, http.StatusOK)
}

func RandomOffice() office_repository.Office {
	rt := util.RandomTime(time.Now(), time.Now().Add(time.Second))
	return office_repository.Office{
		ID:        util.RandomInt(0, 100),
		Name:      util.RandomString(10),
		Address:   util.RandomString(10),
		CreatedAt: rt,
		UpdatedAt: rt,
		ImgFile:   sql.NullString{String: util.RandomString(10), Valid: true},
	}
}
