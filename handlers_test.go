// handlers_test.go
package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestIdentity_Success(t *testing.T) {
	expected := `{"server_name":"Integer to Word Converter"}`
	Helper_TestConvert(t, "GET", "/identity", nil, expected, http.StatusOK)
}

func TestIdentity_405MethodNotAllowed(t *testing.T) {
	expected := `{"error_detail":"Method not allowed - Expecting GET /identity"}`
	Helper_TestConvert(t, "POST", "/identity", nil, expected, http.StatusMethodNotAllowed)
}

func TestConvert_Failure_405MethodNotAllowed(t *testing.T) {
	requestString := []byte(`{"value":0}`)
	expected := `{"error_detail":"Method not allowed. - Expecting POST /convert"}`
	Helper_TestConvert(t, "GET", "/convert", requestString, expected, http.StatusMethodNotAllowed)
}

func TestConvert_Failure_NoBody(t *testing.T) {
	expected := `{"error_detail":"Bad Request : No body present"}`
	Helper_TestConvert(t, "POST", "/convert", nil, expected, http.StatusBadRequest)
}

func TestConvert_Failure_Negative_Integer(t *testing.T) {
	requestString := []byte(`{"value":-1}`)
	expected := `{"error_detail":"Bad Input - Negative Integer is not accepted"}`
	Helper_TestConvert(t, "POST", "/convert", requestString, expected, http.StatusBadRequest)
}

func TestConvert_Failure_Decimal_Number(t *testing.T) {
	requestString := []byte(`{"value":21.1}`)
	expected := `{"error_detail":"Bad Request : JSON input has invalid value (decimal, character, strings are not accepted), try giving a positive integer or remove the preceeding zero"}`
	Helper_TestConvert(t, "POST", "/convert", requestString, expected, http.StatusBadRequest)
}

func TestConvert_Failure_Character(t *testing.T) {
	requestString := []byte(`{"value":"a"}`)
	expected := `{"error_detail":"Bad Request : JSON input has invalid value (decimal, character, strings are not accepted), try giving a positive integer or remove the preceeding zero"}`
	Helper_TestConvert(t, "POST", "/convert", requestString, expected, http.StatusBadRequest)
}

func TestConvert_Success_Integer_0(t *testing.T) {
	requestString := []byte(`{"value":0}`)
	expected := `{"value":0,"value_in_words":"Zero "}`
	Helper_TestConvert(t, "POST", "/convert", requestString, expected, http.StatusOK)
}

func TestConvert_Success_Integer_1(t *testing.T) {
	requestString := []byte(`{"value":1}`)
	expected := `{"value":1,"value_in_words":"One "}`
	Helper_TestConvert(t, "POST", "/convert", requestString, expected, http.StatusOK)
}

func TestConvert_Success_Integer_163(t *testing.T) {
	requestString := []byte(`{"value":163}`)
	expected := `{"value":163,"value_in_words":"One Hundred Sixty Three "}`
	Helper_TestConvert(t, "POST", "/convert", requestString, expected, http.StatusOK)
}

func TestConvert_Success_Integer_2462(t *testing.T) {
	requestString := []byte(`{"value":2462}`)
	expected := `{"value":2462,"value_in_words":"Two Thousand Four Hundred Sixty Two "}`
	Helper_TestConvert(t, "POST", "/convert", requestString, expected, http.StatusOK)
}

func TestConvert_Success_Integer_749241(t *testing.T) {
	requestString := []byte(`{"value":749241}`)
	expected := `{"value":749241,"value_in_words":"Seven Hundred Forty Nine Thousand Two Hundred Forty One "}`
	Helper_TestConvert(t, "POST", "/convert", requestString, expected, http.StatusOK)
}

func TestConvert_Success_Integer_4234643(t *testing.T) {
	requestString := []byte(`{"value":4234643}`)
	expected := `{"value":4234643,"value_in_words":"Four Million Two Hundred Thirty Four Thousand Six Hundred Forty Three "}`
	Helper_TestConvert(t, "POST", "/convert", requestString, expected, http.StatusOK)
}

func TestConvert_Success_Integer_9999999999(t *testing.T) {
	requestString := []byte(`{"value":9999999999}`)
	expected := `{"error_detail":"Bad Input - Value greater than 999999999 cannot be processed"}`
	Helper_TestConvert(t, "POST", "/convert", requestString, expected, http.StatusBadRequest)
}

func Helper_TestConvert(t *testing.T, apiType string, apiName string, requestString []byte, expected string, statusCode int) {
	req, err := http.NewRequest(apiType, apiName, bytes.NewBuffer(requestString))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-type", "application/json")

	// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	if apiName == "/convert" {
		handler := http.Handler(ApiHandler(Convert))
		handler.ServeHTTP(rr, req)
	} else if apiName == "/identity" {
		handler := http.Handler(ApiHandler(Identity))
		handler.ServeHTTP(rr, req)
	} else {
		log.Printf("-- FROM UT STUB - Invalid API Name --")
		return
	}

	// Check the status code is what we expect.
	if status := rr.Code; status != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, statusCode)
	}

	// Check the response body is what we expect.
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestMain(m *testing.M) {
	log.Printf("-- FROM UT STUB : BEGIN TestMain --")
	statusCode := m.Run()
	if statusCode != 0 {
		log.Printf("-- FROM UT STUB : A Test Case Failed Ending TestMain --")
	} else {
		log.Printf("-- FROM UT STUB : All Test Cases Passed Ending TestMain --")
	}
	log.Printf("-- FROM UT STUB : END TestMain --")
	os.Exit(statusCode)
}
