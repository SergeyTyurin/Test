package main

/*
HTTP-сервер

Реализовать HTTP-сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP-библиотекой.

В рамках задания необходимо:
1.	Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
2.	Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
3.	Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
4.	Реализовать middleware для логирования запросов
Методы API:
	POST /create_event
	POST /update_event
	POST /delete_event
	GET /events_for_day
	GET /events_for_week
	GET /events_for_month

Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09). В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON-документ содержащий либо {"result": "..."} в случае успешного выполнения метода, либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
1.	Реализовать все методы.
2.	Бизнес логика НЕ должна зависеть от кода HTTP сервера.
3.	В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"time"
)

var calendar map[int][]Event

type Response struct {
	Result []Event `json:"result,omitempty"`
	Error  error   `json:"error,omitempty"`
	Test   int     `json:"test,omitempty"`
}

type Event struct {
	Id        int       `json:"event id"`
	DateStart time.Time `json:"event time"`
	Name      string    `json:"event name"`
	Descr     string    `json:"description"`
}

type CreateEventRequestFrom struct {
	UserId    int    `json:"user id"`
	DateStart string `json:"event time"`
	Name      string `json:"event name"`
	Descr     string `json:"description"`
}

func (u *CreateEventRequestFrom) ValidateAndCreateEvent() (*CreateEventRequest, error) {

	DateStart, err := time.Parse("2006-01-02", u.DateStart)
	if err != nil {
		return nil, err
	}
	if len(u.Name) == 0 {
		return nil, errors.New("name is Empty")
	}
	if u.UserId <= 0 {
		return nil, errors.New("user is Empty")
	}

	event := &CreateEventRequest{
		UserId:    u.UserId,
		DateStart: DateStart,
		Name:      u.Name,
		Descr:     u.Descr,
	}
	return event, nil
}

type CreateEventRequest struct {
	UserId    int       `json:"user id"`
	DateStart time.Time `json:"event time"`
	Name      string    `json:"event name"`
	Descr     string    `json:"description"`
}

type UpdateEventRequestFrom struct {
	Id        int    `json:"event id"`
	UserId    int    `json:"user id"`
	DateStart string `json:"event time"`
	Name      string `json:"event name"`
	Descr     string `json:"description"`
}

func (u *UpdateEventRequestFrom) ValidateAndCreateEvent() (*UpdateEventRequest, error) {

	DateStart, err := time.Parse("2006-01-02", u.DateStart)
	if err != nil {
		return nil, err
	}
	if len(u.Name) == 0 {
		return nil, errors.New("name is Empty")
	}
	if u.UserId <= 0 {
		return nil, errors.New("user is Empty")
	}
	if u.Id <= 0 {
		return nil, errors.New("id is Empty")
	}
	event := &UpdateEventRequest{
		Id:        u.Id,
		UserId:    u.UserId,
		DateStart: DateStart,
		Name:      u.Name,
		Descr:     u.Descr,
	}
	return event, nil
}

type UpdateEventRequest struct {
	Id        int       `json:"event id"`
	UserId    int       `json:"user id"`
	DateStart time.Time `json:"event time"`
	Name      string    `json:"event name"`
	Descr     string    `json:"description"`
}

type DeleteEventRequest struct {
	Id     int `json:"event id"`
	UserId int `json:"user id"`
}

func getEvents(period string, dateFromStr string, userIdStr string) Response {
	response := Response{}
	var events []Event

	dateFrom, err := time.Parse("2006-01-02", dateFromStr)
	if err != nil {
		response.Error = err
		return response
	}
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		response.Error = err
		return response
	}

	userCalender := calendar[userId]
	fmt.Println(period)
	if strings.EqualFold(period, "day") {
		for _, el := range userCalender {
			if (el.DateStart == dateFrom || el.DateStart.After(dateFrom)) && el.DateStart.Before(dateFrom.Add(24*time.Hour)) {
				events = append(events, el)
			}
		}
	} else if strings.EqualFold(period, "week") {
		year, week := dateFrom.ISOWeek()
		for _, el := range userCalender {
			yearCheck, weekCheck := el.DateStart.ISOWeek()
			if yearCheck == year && weekCheck == week {
				events = append(events, el)
			}
		}
	} else if strings.EqualFold(period, "month") {
		year, month, _ := dateFrom.Date()
		for _, el := range userCalender {
			yearCheck, monthCheck, _ := el.DateStart.Date()
			if yearCheck == year && monthCheck == month {
				events = append(events, el)
			}
		}
	}
	response.Result = events
	return response

}

func deleteEvent(event *DeleteEventRequest) error {
	var status bool

	eventList := calendar[event.UserId]
	for i, el := range eventList {
		if event.Id == el.Id {
			eventList = append(eventList[:i], eventList[i+1:]...)
			status = true
		}
	}
	if status {
		return nil
	} else {
		return errors.New("Event not found")
	}
}

func createEvent(event *CreateEventRequestFrom) (int, error) {

	crEvent, err := event.ValidateAndCreateEvent()
	if err != nil {
		return 0, err
	}
	eventList := calendar[crEvent.UserId]
	max := 0
	for _, el := range eventList {
		if max < el.Id {
			max = el.Id
		}
	}
	newEvent := Event{Id: max + 1, DateStart: crEvent.DateStart, Name: crEvent.Name, Descr: crEvent.Descr}

	eventList = append(eventList, newEvent)
	calendar[crEvent.UserId] = eventList
	return max + 1, nil
}

func updateEvent(event *UpdateEventRequestFrom) error {
	updEvent, err := event.ValidateAndCreateEvent()
	if err != nil {
		return err
	}
	eventList := calendar[updEvent.UserId]
	var status bool
	for _, el := range eventList {
		if el.Id == updEvent.Id {
			el.DateStart = updEvent.DateStart
			el.Name = updEvent.Name
			el.Descr = updEvent.Descr
			status = true
			break
		}
	}
	if status {
		return nil
	} else {
		return errors.New("Event not found")
	}
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Start working with request", r.URL.Path)
		next.ServeHTTP(w, r)
		log.Print("End working with request", r.URL.Path)
	})
}

func create(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var event CreateEventRequestFrom
		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		newId, err := createEvent(&event)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(201)
		res := fmt.Sprintf("{\"result\":%d}", newId)
		w.Write([]byte(res))
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var event UpdateEventRequestFrom
		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		err = updateEvent(&event)
		if err == nil {
			w.WriteHeader(200)
			res := "{\"result\":\"ok\"}"
			w.Write([]byte(res))
		} else {
			w.WriteHeader(503)
			w.Write([]byte(err.Error()))
		}
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var event DeleteEventRequest
		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		err = deleteEvent(&event)
		if err == nil {
			w.WriteHeader(200)
			res := "{\"result\":\"ok\"}"
			w.Write([]byte(res))
		} else {
			w.WriteHeader(503)
			w.Write([]byte(err.Error()))
		}

	}
}

func get(period string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		date := query["date"]
		if len(date) == 0 {
			w.WriteHeader(400)
			w.Write([]byte("date is not found"))
			return
		}
		userId := query["user_id"]
		if len(userId) == 0 {
			w.WriteHeader(400)
			w.Write([]byte("user is not found"))
			return
		}
		response := getEvents(period, date[0], userId[0])
		bytes, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(bytes)
	}
}

func getMux() *http.ServeMux {
	mux := http.NewServeMux()

	createHandler := http.HandlerFunc(create)
	mux.Handle("/create_event", middleware(createHandler))

	deleteHandler := http.HandlerFunc(delete)
	mux.Handle("/delete_event", middleware(deleteHandler))

	updateHandler := http.HandlerFunc(update)
	mux.Handle("/update_event", middleware(updateHandler))

	getDayHandler := http.HandlerFunc(get("day"))
	mux.Handle("/events_for_day", middleware(getDayHandler))

	getWeekHandler := http.HandlerFunc(get("week"))
	mux.Handle("/events_for_week", middleware(getWeekHandler))

	getMonthHandler := http.HandlerFunc(get("month"))
	mux.Handle("/events_for_month", middleware(getMonthHandler))

	return mux
}

func main() {
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

	server := http.Server{
		Addr:              ":8080",
		Handler:           getMux(),
		ReadHeaderTimeout: 30 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
