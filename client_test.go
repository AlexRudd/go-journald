package gojournald

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexrudd/go-journald/mock"
)

func TestMachine(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(mock.JournalGatewayd))
	defer mockServer.Close()

	j := NewJournal()
	j.Configure(func(cj *Journal) error {
		cj.c = mockServer.Client()
		cj.url = mockServer.URL
		return nil
	})

	_, err := j.Machine()
	if err != nil {
		t.Error(err)
		return
	}
}
