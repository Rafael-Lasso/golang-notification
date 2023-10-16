package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func testWebServer(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	WebServer(rr, req)

	if http.StatusOK != rr.Code {
		t.Errorf("statur code spected: %d got: %d", http.StatusOK, rr.Code)
	} 
	// else {
	// 	fmt.Fprintln("✅ | Statur code spected: %d got: %d", http.StatusOK, rr.Code)
	// }
}
