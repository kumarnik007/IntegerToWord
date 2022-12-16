// ------------------
// Handler Functions
// ------------------

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ApiHandler func(http.ResponseWriter, *http.Request) error

// ApiHandler implements http.Handler interface.
func (handler ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Call handler function
	errorResponse := handler(w, r)
	if errorResponse == nil {
		return
	}

	// Error handling logic starts.
	// log.Printf("An error occured: %v", errorResponse)

	// Check if it is a ClientError.
	clientError, ok := errorResponse.(ClientError)
	if !ok {
		// If the error is not ClientError, assume that it is ServerError.
		// return 500 Internal Server Error.
		w.WriteHeader(500)
		return
	}

	// Get response body of ClientError.
	body, errorResponse := clientError.ResponseBody()
	if errorResponse != nil {
		log.Printf("An error accured: %v", errorResponse)
		w.WriteHeader(500)
		return
	}

	// Get http status code and headers.
	status, headers := clientError.ResponseHeaders()
	for k, v := range headers {
		w.Header().Set(k, v)
	}

	w.WriteHeader(status)
	w.Write(body)
}

func Identity(w http.ResponseWriter, r *http.Request) error {
	// check if http GET method is called
	if r.Method != http.MethodGet {
		// Return 405 Method Not Allowed.
		return NewHTTPError(nil, 405, "Method not allowed - Expecting GET /identity")
	}

	res1D := &responseGetIdentity{
		Server_name: string("Integer to Word Converter")}

	res1B, errorResponse := json.Marshal(res1D)
	if errorResponse != nil {
		// Return 500 Internal Server Error.
		return NewHTTPError(errorResponse, 500, "unable to prepare JSON response")
	}

	w.Header().Set("content-type", "application/json")
	w.Write(res1B)

	return nil
}

func Convert(w http.ResponseWriter, r *http.Request) error {
	// check if http POST method is called
	if r.Method != http.MethodPost {
		// Return 405 Method Not Allowed.
		return NewHTTPError(nil, 405, "Method not allowed. - Expecting POST /convert")
	}

	contentType := r.Header.Get("Content-type")
	if contentType != "application/json" {
		// Return 400 Bad Request.
		return NewHTTPError(nil, 400, "Bad Request : Content-Type header is not application/json")
	}

	if r.ContentLength == 0 {
		// Return 400 Bad Request.
		return NewHTTPError(nil, 400, "Bad Request : No body present")
	}

	// Read body
	b, errorResponse := io.ReadAll(r.Body)
	if errorResponse != nil {
		// Return 500 Internal Server Error.
		return fmt.Errorf("request body read error : %v", errorResponse)
	}

	// Unmarshal
	var msg requestPostConvert
	errorResponse = json.Unmarshal(b, &msg)
	if errorResponse != nil {
		// Return 400 Bad Request.
		return NewHTTPError(errorResponse, 400, "Bad Request : JSON input has invalid value (decimal, character, strings are not accepted), try giving a positive integer or remove the preceeding zero")
	}

	// log.Printf("-- Input Number = %d --", msg.Value)

	result, errorResponse := NumberToWords(msg.Value)
	if errorResponse != nil {
		return errorResponse
	}

	res2D := &responsePostConvert{
		Value:          msg.Value,
		Value_in_words: result}

	output, errorResponse := json.Marshal(res2D)
	if errorResponse != nil {
		// Return 500 Internal Server Error.
		return NewHTTPError(errorResponse, 500, "unable to prepare JSON response")
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)

	return nil
}
