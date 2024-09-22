package controller

import (
	"bytes"
	"dohabits/data"
	"dohabits/db"
	"dohabits/logger"
	"dohabits/model"
	"dohabits/view"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	logger := &logger.Logger{}
	db := &db.MyMockDB{}
	m := &model.HabitsModel{}
	v := &view.HabitsView{}
	c := NewHabitsController(logger)

	endpoint := fmt.Sprintf("%s/%s", os.Getenv("API_NAME"), os.Getenv("API_VERSION"))

	marshalledNewHabit, err := json.Marshal(data.NewHabit{
		Name:       "Test Create Habit",
		Days:       30,
		DaysTarget: 50,
	})

	if err != nil {
		t.Errorf("Fail err: %s", err)
	}

	testCases := []struct {
		name string
		want []byte
	}{
		{
			name: "Test successful Insert",
			want: marshalledNewHabit,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/%s/CreateHabit", endpoint), io.NopCloser(bytes.NewBuffer(marshalledNewHabit)))
			w := httptest.NewRecorder()

			c.Create(w, req, m, v, db, logger)

			res := w.Result()

			defer res.Body.Close()

			got, err := io.ReadAll(res.Body)

			if err != nil {
				t.Errorf("Fail err: %s", err)
			}

			if !bytes.Equal(val.want, got) {
				t.Errorf("Fail want doesn't match got")
			}
		})
	}
}

func TestRetrieve(t *testing.T) {
	logger := &logger.Logger{}
	db := &db.MyMockDB{}
	m := &model.HabitsModel{}
	v := &view.HabitsView{}
	c := NewHabitsController(logger)

	endpoint := fmt.Sprintf("%s/%s", os.Getenv("API_NAME"), os.Getenv("API_VERSION"))

	marshalledHabit, err := json.Marshal(data.MockHabit[0])

	if err != nil {
		t.Errorf("Fail err: %s", err)
	}

	testCases := []struct {
		name string
		want []byte
	}{
		{
			name: "Test success Retrieve",
			want: marshalledHabit,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s/RetrieveAllHabits", endpoint), nil)
			w := httptest.NewRecorder()
			q := req.URL.Query()
			q.Add("id", "1")
			req.URL.RawQuery = q.Encode()

			c.Retrieve(w, req, m, v, db, logger)

			res := w.Result()

			defer res.Body.Close()

			got, err := io.ReadAll(res.Body)

			if err != nil {
				t.Errorf("Fail err: %s", err)
			}

			if !bytes.Equal(val.want, got) {
				t.Errorf("Fail want doesn't match got")
			}
		})
	}
}

func TestRetrieveAll(t *testing.T) {
	logger := &logger.Logger{}
	db := &db.MyMockDB{}
	m := &model.HabitsModel{}
	v := &view.HabitsView{}
	c := NewHabitsController(logger)

	endpoint := fmt.Sprintf("%s/%s", os.Getenv("API_NAME"), os.Getenv("API_VERSION"))

	marshalledAllHabits, err := json.Marshal(data.MockHabit)

	if err != nil {
		t.Errorf("Fail err: %s", err)
	}

	testCases := []struct {
		name string
		want []byte
	}{
		{
			name: "Test success Retrieve All",
			want: marshalledAllHabits,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s/RetrieveAllHabits", endpoint), nil)
			w := httptest.NewRecorder()

			c.RetrieveAll(w, req, m, v, db, logger)

			res := w.Result()

			defer res.Body.Close()

			got, err := io.ReadAll(res.Body)

			if err != nil {
				t.Errorf("Fail err: %s", err)
			}

			if !bytes.Equal(val.want, got) {
				t.Errorf("Fail want doesn't match got")
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	logger := &logger.Logger{}
	db := &db.MyMockDB{}
	m := &model.HabitsModel{}
	v := &view.HabitsView{}
	c := NewHabitsController(logger)

	endpoint := fmt.Sprintf("%s/%s", os.Getenv("API_NAME"), os.Getenv("API_VERSION"))

	marshalledHabit, err := json.Marshal(data.Habit{
		ID:         data.MockHabit[0].ID,
		CreatedAt:  data.MockHabit[0].CreatedAt,
		Name:       "Test Update Habit",
		Days:       30,
		DaysTarget: 50,
	})

	if err != nil {
		t.Errorf("Fail err: %s", err)
	}

	testCases := []struct {
		name string
		want []byte
	}{
		{
			name: "Test successful Update",
			want: marshalledHabit,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			newHabit := data.NewHabit{
				Name:       "Test Update Habit",
				Days:       30,
				DaysTarget: 50,
			}

			marshalledNewHabit, err := json.Marshal(newHabit)

			if err != nil {
				t.Errorf("Fail err: %s", err)
			}

			req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/%s/UpdateHabit", endpoint), io.NopCloser(bytes.NewBuffer(marshalledNewHabit)))
			w := httptest.NewRecorder()
			q := req.URL.Query()
			q.Add("id", "1")
			req.URL.RawQuery = q.Encode()

			c.Update(w, req, m, v, db, logger)

			res := w.Result()

			defer res.Body.Close()

			got, err := io.ReadAll(res.Body)

			if err != nil {
				t.Errorf("Fail err: %s", err)
			}

			if !bytes.Equal(val.want, got) {
				t.Errorf("Fail want doesn't match got")
			}
		})
	}
}

func TestDelete(t *testing.T) {
	logger := &logger.Logger{}
	db := &db.MyMockDB{}
	m := &model.HabitsModel{}
	v := &view.HabitsView{}
	c := NewHabitsController(logger)

	endpoint := fmt.Sprintf("%s/%s", os.Getenv("API_NAME"), os.Getenv("API_VERSION"))

	testCases := []struct {
		name string
		want map[string]bool
	}{
		{
			name: "Test Success Delete",
			want: map[string]bool{"success": true},
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/%s/DeleteHabit", endpoint), nil)
			w := httptest.NewRecorder()
			q := req.URL.Query()
			q.Add("id", "1")
			req.URL.RawQuery = q.Encode()

			expect, err := json.Marshal(map[string]bool{"success": true})

			if err != nil {
				t.Errorf("Fail err: %s", err)
			}

			c.Delete(w, req, m, v, db, logger)

			res := w.Result()

			defer res.Body.Close()

			got, err := io.ReadAll(res.Body)

			if err != nil {
				t.Errorf("Fail err: %s", "cake")
			}

			if !bytes.Equal(expect, got) {
				t.Errorf("Fail want doesn't match got")
			}
		})
	}
}
