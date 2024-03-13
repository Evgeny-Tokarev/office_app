package officeservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/evgeny-tokarev/office_app/backend/internal/repositories/office_repository"
	"github.com/evgeny-tokarev/office_app/backend/util"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//todo: write tests for all functions

type OfficeServiceUnitFailSuite struct {
	OfficeServiceUnitSuite
}

func TestOfficeServiceUnitFailSuite(t *testing.T) {
	suite.Run(t, &OfficeServiceUnitFailSuite{})
}

func (osu *OfficeServiceUnitFailSuite) SetupSuite() {
	osu.OfficeServiceUnitSuite.SetupServer()
}

func (osu *OfficeServiceUnitFailSuite) SetupTest() {
	osu.office = RandomOffice()
}

func (osu *OfficeServiceUnitFailSuite) TestCreateOffice() {
	testCases := []struct {
		name          string
		requestBody   string
		expectedCode  int
		expectedError string
	}{
		{
			name:          "Invalid JSON",
			requestBody:   "invalid_json",
			expectedCode:  http.StatusBadRequest,
			expectedError: "invalid character 'i' looking for beginning of value",
		},
		{
			name:          "Missing Fields",
			requestBody:   `{"Name": "", "Address": ""}`,
			expectedCode:  http.StatusBadRequest,
			expectedError: "all fields are required",
		},
		{
			name:          "Encoding Failure",
			requestBody:   `{"Name": "Name for encoding failure", "Address": "Address"}`,
			expectedCode:  http.StatusInternalServerError,
			expectedError: "encoding failure",
		},
		{
			name:          "DB Failure",
			requestBody:   `{"Name": "Name for db failure", "Address": "Address"}`,
			expectedCode:  http.StatusInternalServerError,
			expectedError: "db failure",
		},
	}
	osu.querier.EXPECT().CreateOffice(mock.Anything, office_repository.CreateOfficeParams{
		Name:    "Name for encoding failure",
		Address: "Address",
	}).Return(osu.office, nil).Times(1)
	osu.querier.EXPECT().CreateOffice(mock.Anything, office_repository.CreateOfficeParams{
		Name:    "Name for db failure",
		Address: "Address",
	}).Return(osu.office, errors.New("db failure")).Times(1)

	for _, testCase := range testCases {
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/offices", osu.server.URL), strings.NewReader(testCase.requestBody))
		if err != nil {
			osu.FailNowf("Error sending request: ", err.Error())
		}
		req.Header.Set("Content-Type", "application/json")
		var recorder http.ResponseWriter
		var responseBody []byte
		if testCase.name == "Encoding Failure" {
			recorder = util.NewCustomResponseRecorder()
		} else {
			recorder = httptest.NewRecorder()
		}
		osu.router.ServeHTTP(recorder, req)
		customRecorder, ok := recorder.(*util.CustomResponseRecorder)
		if !ok {
			responseBody = recorder.(*httptest.ResponseRecorder).Body.Bytes()
		} else {
			responseBody = customRecorder.Body.Bytes()
		}
		var errorResponse util.ErrorResponse
		err = json.Unmarshal(responseBody, &errorResponse)
		if err != nil {
			osu.FailNowf("Error unmarshalling response: ", err.Error())
		}
		osu.Require().Equal(errorResponse.Status, testCase.expectedCode)
		osu.Require().Equal(errorResponse.Message, testCase.expectedError)
	}
}

//	func (osu *OfficeServiceUnitFailSuite) TestGetOffice() {
//		testCases := []struct {
//			name         string
//			requestBody  string
//			expectedCode int
//			expectedError string
//		}{
//			{
//				name:         "Invalid JSON",
//				requestBody:  "invalid_json",
//				expectedCode: http.StatusBadRequest,
//				expectedError: "invalid character 'i' looking for beginning of value",
//			},
//			{
//				name:         "Missing Fields",
//				requestBody:  `{"Name": "", "Address": ""}`,
//				expectedCode: http.StatusBadRequest,
//				expectedError: "all fields are required",
//			},
//			// Add more test cases for different error scenarios
//		}
//	}
