package stubby4go

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type dataTest struct {
	request  Request
	response Response
}

var dataTable = []dataTest{
	dataTest{
		createRequest("GET", "/hello-world"),
		createResponse(200, "application/json", "Hello-world!"),
	},
	dataTest{
		createRequest("GET", "/hello-golang"),
		createResponse(200, "application/json", "Hello-Golang!"),
	},
	dataTest{
		createRequest("GET", "/hello-golang"),
		createResponse(200, "text/xml", "Hello-Golang!"),
	},
}

func TestCreateHandler(t *testing.T) {
	for _, data := range dataTable {

		var handler http.Handler = CreateHandler(data.request, data.response)

		var actualRequest, err = http.NewRequest(data.request.Method, data.request.Url, nil)
		if err != nil {
			t.Error(err)
		}

		var responseRecorder = httptest.NewRecorder()

		handler.ServeHTTP(responseRecorder, actualRequest)

		assertEquals(t, responseRecorder.Code, data.response.Status)

		assertStringEquals(t, responseRecorder.Body.String(), data.response.Body)

		assertStringEquals(t, responseRecorder.Header().Get("content-type"),
			data.response.Headers["content-type"])

	}
}

func assertEquals(t *testing.T, actual int, expected int) {
	if actual != expected {
		t.Errorf("expected %d but was %d", expected, actual)
	}
}

func assertStringEquals(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Errorf("expected %s but was %s", expected, actual)
	}
}

func createRequest(method string, url string) Request {
	return Request{
		Method: method,
		Url:    url,
	}

}

func createResponse(status int, contentType string, body string) Response {
	return Response{
		Status: status,
		Headers: map[string]string{
			"content-type": contentType,
		},
		Body: body,
	}

}
