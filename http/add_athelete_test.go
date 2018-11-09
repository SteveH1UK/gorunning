package http

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SteveH1UK/gorunning"

	"github.com/SteveH1UK/gorunning/mocks"
	"github.com/golang/mock/gomock"
)

// required to reset production value when using it in tests
func addAthleleteTestInit() {
	callDecodeJSONFromBody = decodeJSONFromBody
}

// TestNewAtheleteHappyPath - tests working path, check that the DAO is populated with the correct struct.
func TestNewAtheleteHappyPath(t *testing.T) {
	fmt.Println("Test")

	addAthleleteTestInit()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAtheleteDAO := mocks.NewMockAthleteDAOInterface(mockCtrl)
	newAthelete := root.NewAthelete{FriendyURL: "johndoe", Name: "John Doe", DateOfBirth: "18-10-1990"}

	mockAtheleteDAO.EXPECT().CreateAthelete(&newAthelete)

	env := NewEnv(mockAtheleteDAO)

	payload := []byte(`{"friendly-url":"johndoe", "name":"John Doe", "date-of-birth":"18-10-1990"}`)

	req, err := http.NewRequest("POST", "/gorunning/atheletes", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(env.createAthelete)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	result := rr.Body.String()
	if result != `{"http-code":201,"id":"johndoe","href":"/gorunning/atheletes/johndoe"}` {
		t.Error("Wrong result got ", result)
	}

	expectedHTTPCode := 201
	if rr.Code != expectedHTTPCode {
		t.Error("Wrong HTTPCode returned ", rr.Code)
	}
}

func TestNewAtheleteErrorFromParsingJSON(t *testing.T) {
	fmt.Println("Test Parse Error")
	addAthleleteTestInit()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAtheleteDAO := mocks.NewMockAthleteDAOInterface(mockCtrl)

	env := NewEnv(mockAtheleteDAO)
	callDecodeJSONFromBody = mockErrorFromDecodeJSONFromBody

	payload := []byte(`{"rubbish"}`)

	req, err := http.NewRequest("POST", "/gorunning/atheletes", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(env.createAthelete)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	result := rr.Body.String()
	if result != `{"http-code":400,"message":"Can not decode request"}` {
		t.Error("Wrong result got ", result)
	}

	expectedHTTPCode := 400
	if rr.Code != expectedHTTPCode {
		t.Error("Wrong HTTPCode returned ", rr.Code)
	}
}

// I could have avoided this (external function would give error but wanted to use this method of mocking as a contrast)
func mockErrorFromDecodeJSONFromBody(body io.ReadCloser, athelete *root.NewAthelete) error {
	fmt.Println("in mock decode JSON")
	return errors.New("Test error from JSON decode")
}

func TestNewAtheleteValidationErrors(t *testing.T) {
	addAthleleteTestInit()
	fmt.Println("Test Validation Error")

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAtheleteDAO := mocks.NewMockAthleteDAOInterface(mockCtrl)

	env := NewEnv(mockAtheleteDAO)

	payload := []byte(`{"friendly-url":"jo", "name":"John Doe", "date-of-birth":"189-10-1990"}`)

	req, err := http.NewRequest("POST", "/gorunning/atheletes", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(env.createAthelete)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	result := rr.Body.String()
	if result != `{"http-code":422,"message":"Validation Errors","validation-errors":[{"code":701,"field":"friendly_url","message":"must be at least 4 characters long"},{"code":703,"field":"date-of-birth","message":"must be in format dd-MM-yyyy"}]}` {
		t.Error("Wrong result got ", result)
	}

	expectedHTTPCode := 422
	if rr.Code != expectedHTTPCode {
		t.Error("Wrong HTTPCode returned ", rr.Code)
	}
}

//TestNewAtheleteAlreadyExists - use Gomock to generate correct error for the handler
func TestNewAtheleteAlreadyExists(t *testing.T) {
	fmt.Println("Test")

	addAthleleteTestInit()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAtheleteDAO := mocks.NewMockAthleteDAOInterface(mockCtrl)
	newAthelete := root.NewAthelete{FriendyURL: "johndoe", Name: "John Doe", DateOfBirth: "18-10-1990"}

	mockAtheleteDAO.EXPECT().CreateAthelete(&newAthelete).Return(root.ErrDBRecordExists).Times(1)

	env := NewEnv(mockAtheleteDAO)

	payload := []byte(`{"friendly-url":"johndoe", "name":"John Doe", "date-of-birth":"18-10-1990"}`)

	req, err := http.NewRequest("POST", "/gorunning/atheletes", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(env.createAthelete)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	result := rr.Body.String()
	if result != `{"http-code":422,"message":"Athelete with this friendly names already exists"}` {
		t.Error("Wrong result got ", result)
	}

	expectedHTTPCode := 422
	if rr.Code != expectedHTTPCode {
		t.Error("Wrong HTTPCode returned ", rr.Code)
	}
}

func TestNewAtheleteDatabaseError(t *testing.T) {
	fmt.Println("Test")

	addAthleleteTestInit()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAtheleteDAO := mocks.NewMockAthleteDAOInterface(mockCtrl)
	newAthelete := root.NewAthelete{FriendyURL: "johndoe", Name: "John Doe", DateOfBirth: "18-10-1990"}

	mockAtheleteDAO.EXPECT().CreateAthelete(&newAthelete).Return(errors.New("random database error")).Times(1)

	env := NewEnv(mockAtheleteDAO)

	payload := []byte(`{"friendly-url":"johndoe", "name":"John Doe", "date-of-birth":"18-10-1990"}`)

	req, err := http.NewRequest("POST", "/gorunning/atheletes", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(env.createAthelete)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	result := rr.Body.String()
	if result != `{"http-code":500,"message":"Database error"}` {
		t.Error("Wrong result got ", result)
	}

	expectedHTTPCode := 500
	if rr.Code != expectedHTTPCode {
		t.Error("Wrong HTTPCode returned ", rr.Code)
	}
}
