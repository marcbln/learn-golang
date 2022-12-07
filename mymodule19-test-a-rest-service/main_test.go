package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_myHandler(t *testing.T) {

	type args struct {
		url            string
		expectedBody   string
		expectedStatus int
	}

	// ---- define tests
	tests := []struct {
		name string
		args args
	}{
		{"t1", args{"/bla.json", `{"name":"bla"}`, http.StatusOK}},
		{"t2", args{"/bla.xml", "", http.StatusNotImplemented}},
	}

	// ---- iterate tests
	for _, test := range tests {
		t.Logf("%#v\n", test)
		req, err := http.NewRequest(http.MethodGet, test.args.url, nil)
		if err != nil {
			t.Errorf("creating request failed: %#v", err)
		}

		recorder := httptest.NewRecorder()
		myHandleFunc(recorder, req)

		// t.Logf("recorded response: %#v", recorder)

		// testing response
		if recorder.Result().StatusCode != test.args.expectedStatus {
			t.Errorf("Status fail: %v (expected %v)", recorder.Result().Status, test.args.expectedStatus)
		}
		if recorder.Body.String() != test.args.expectedBody {
			t.Errorf("Body fail: %v (expected %v)", recorder.Body.String(), test.args.expectedBody)
		}
	}
}
