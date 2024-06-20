package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	dt, _ := time.Parse("2006-01-02", "2024-04-02")
	e1 := Event{DateStart: dt, Name: "Doctor", Descr: "Go to doctor in Moscow"}
	dt2, _ := time.Parse("2006-01-02", "2024-04-06")
	e2 := Event{DateStart: dt2, Name: "Concert", Descr: "Go to concert in Kazan"}
	dt3, _ := time.Parse("2006-01-02", "2024-04-21")
	e3 := Event{DateStart: dt3, Name: "Train", Descr: "Go to doctor to Tula"}
	dt4, _ := time.Parse("2006-01-02", "2024-05-21")
	e4 := Event{DateStart: dt4, Name: "Test", Descr: "tessa"}

	events := []Event{e1, e2, e3, e4}
	calendar = make(map[int][]Event)
	calendar[1] = events

	body := `{"user id":1, "event name":"Testing", "description":"gfgggggg", "event time":"2024-04-01"}`
	request, err := http.NewRequest(http.MethodPost, "/create_event", strings.NewReader(body))
	require.NoError(t, err) // check err is nil

	response := httptest.NewRecorder()
	mux := getMux()

	mux.ServeHTTP(response, request)
	require.Equal(t, 201, response.Code)
}

func TestUpdate(t *testing.T) {
	dt, _ := time.Parse("2006-01-02", "2024-04-02")
	e1 := Event{Id: 1, DateStart: dt, Name: "Doctor", Descr: "Go to doctor in Moscow"}
	dt2, _ := time.Parse("2006-01-02", "2024-04-06")
	e2 := Event{Id: 2, DateStart: dt2, Name: "Concert", Descr: "Go to concert in Kazan"}
	dt3, _ := time.Parse("2006-01-02", "2024-04-21")
	e3 := Event{Id: 3, DateStart: dt3, Name: "Train", Descr: "Go to doctor to Tula"}
	dt4, _ := time.Parse("2006-01-02", "2024-05-21")
	e4 := Event{Id: 4, DateStart: dt4, Name: "Test", Descr: "tessa"}
	events := []Event{e1, e2, e3, e4}
	calendar = make(map[int][]Event)
	calendar[1] = events

	body := `{"user id":1, "event id": 1, "event name":"Changing", "description":"dbgdchndhndhbndg", "event time":"2024-04-01"}`
	request, err := http.NewRequest(http.MethodPost, "/update_event", strings.NewReader(body))
	require.NoError(t, err) // check err is nil

	response := httptest.NewRecorder()
	mux := getMux()

	mux.ServeHTTP(response, request)
	require.Equal(t, 200, response.Code)
}

func TestDelete(t *testing.T) {
	dt, _ := time.Parse("2006-01-02", "2024-04-02")
	e1 := Event{Id: 1, DateStart: dt, Name: "Doctor", Descr: "Go to doctor in Moscow"}
	dt2, _ := time.Parse("2006-01-02", "2024-04-06")
	e2 := Event{Id: 2, DateStart: dt2, Name: "Concert", Descr: "Go to concert in Kazan"}
	dt3, _ := time.Parse("2006-01-02", "2024-04-21")
	e3 := Event{Id: 3, DateStart: dt3, Name: "Train", Descr: "Go to doctor to Tula"}
	dt4, _ := time.Parse("2006-01-02", "2024-05-21")
	e4 := Event{Id: 4, DateStart: dt4, Name: "Test", Descr: "tessa"}
	events := []Event{e1, e2, e3, e4}
	calendar = make(map[int][]Event)
	calendar[1] = events

	body := `{"user id":1, "event id": 1}`
	request, err := http.NewRequest(http.MethodPost, "/delete_event", strings.NewReader(body))
	require.NoError(t, err) // check err is nil

	response := httptest.NewRecorder()
	mux := getMux()

	mux.ServeHTTP(response, request)
	require.Equal(t, 200, response.Code)
}
