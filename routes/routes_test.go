package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewRouter(t *testing.T) {
	handler := &MyHandler{}
	server := httptest.NewServer(handler)
	defer server.Close()

	for _, i := range []int{1, 2} {
		resp, err := http.Get(server.URL)
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != 200 {
			t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
		}
		expected := fmt.Sprintf("Visitor count: %d.", i)
		actual, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		if expected != string(actual) {
			t.Errorf("Expected the message '%s'\n", expected)
		}
	}
}
