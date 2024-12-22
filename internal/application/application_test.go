package application_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LootNex/Http_Calc/internal/application"
)

func TestCaclulatorHandler_StatusOK(t *testing.T){

	reqbody := `{"expression":"2+2"}`

	req := httptest.NewRequest("POST", "/api/v1/calculate", bytes.NewBufferString(reqbody))
	req.Header.Set("Content-Type", "app;ication/json")

	rr := httptest.NewRecorder()

	application.CalculatorHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var result application.Results
	err := json.NewDecoder(rr.Body).Decode(&result)
	
	if err != nil{
		t.Fatal(err)
	}

	expected := 4.0

	if result.Result != expected{
		t.Errorf("expected result %f, got %f", expected, result.Result)
	}

}

func TestCaclulatorHandler_StatusUnprocessableEntity(t *testing.T){

	reqbody := `{"expression":"2+(2"}`

	req := httptest.NewRequest("POST", "/api/v1/calculate", bytes.NewBufferString(reqbody))
	req.Header.Set("Content-Type", "app;ication/json")

	rr := httptest.NewRecorder()

	application.CalculatorHandler(rr, req)

	if rr.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected status code %d, got %d", http.StatusUnprocessableEntity, rr.Code)
	}

	var result application.Results
	err := json.NewDecoder(rr.Body).Decode(&result)
	
	if err != nil{
		t.Fatal(err)
	}

	expected := 0.0

	if result.Result != expected{
		t.Errorf("expected result %f, got %f", expected, result.Result)
	}

}

func TestCaclulatorHandler_StatusInternalServerError(t *testing.T){

	reqbody := `{"expression":"2+2}`

	req := httptest.NewRequest("POST", "/api/v1/calculate", bytes.NewBufferString(reqbody))
	req.Header.Set("Content-Type", "app;ication/json")

	rr := httptest.NewRecorder()

	application.CalculatorHandler(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("expected status code %d, got %d", http.StatusInternalServerError, rr.Code)
	}

}
