package controller

import (
	"bytes"
	"context"
	"dohabits/data"
	"dohabits/db"
	"dohabits/logger"
	"dohabits/middleware/session"
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

func TestCreateHabitsHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewDB(logger)
	habitsModel := model.NewHabitsModel(logger, db)
	habitsView := view.NewHabitsView(logger)
	c := NewHabitsController(habitsModel, habitsView, logger)

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
			name: "Test successful Create",
			want: marshalledNewHabit,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/%s/CreateHabit", endpoint), io.NopCloser(bytes.NewBuffer(marshalledNewHabit)))
			w := httptest.NewRecorder()

			claims := &session.Claims{Username: "johndoe1@example.com"}

			ctx := context.WithValue(req.Context(), session.ClaimsKey, claims)

			req = req.WithContext(ctx)

			c.CreateHabitsHandler(w, req)

			if status := w.Code; status == http.StatusInternalServerError {
				t.Errorf("TestCreateHabitsHandler - HTTP Status Code = %d", status)
				return
			}

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

func TestRetrieveHabitsHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewDB(logger)
	habitsModel := model.NewHabitsModel(logger, db)
	habitsView := view.NewHabitsView(logger)
	c := NewHabitsController(habitsModel, habitsView, logger)

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
			q.Add("habitId", "1")
			req.URL.RawQuery = q.Encode()

			claims := &session.Claims{Username: "johndoe1@example.com"}

			ctx := context.WithValue(req.Context(), session.ClaimsKey, claims)

			req = req.WithContext(ctx)

			c.RetrieveHabitsHandler(w, req)

			if status := w.Code; status == http.StatusInternalServerError {
				t.Errorf("TestRetrieveHabitsHandler - HTTP Status Code = %d", status)
				return
			}

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

func TestRetrieveAllHabitsHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewDB(logger)
	habitsModel := model.NewHabitsModel(logger, db)
	habitsView := view.NewHabitsView(logger)
	c := NewHabitsController(habitsModel, habitsView, logger)

	endpoint := fmt.Sprintf("%s/%s", os.Getenv("API_NAME"), os.Getenv("API_VERSION"))

	userMockHabits := []data.Habit{}

	for _, val := range data.MockHabit {
		if val.UserID == "1" {
			userMockHabits = append(userMockHabits, val)
		}
	}

	marshalledAllHabits, err := json.Marshal(userMockHabits)

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

			claims := &session.Claims{Username: "johndoe1@example.com"}

			ctx := context.WithValue(req.Context(), session.ClaimsKey, claims)

			req = req.WithContext(ctx)

			c.RetrieveAllHabitsHandler(w, req)

			if status := w.Code; status == http.StatusInternalServerError {
				t.Errorf("TestRetrieveAllHabitsHandler - HTTP Status Code = %d", status)
				return
			}

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

func TestUpdateHabitsHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewDB(logger)
	habitsModel := model.NewHabitsModel(logger, db)
	habitsView := view.NewHabitsView(logger)
	c := NewHabitsController(habitsModel, habitsView, logger)

	endpoint := fmt.Sprintf("%s/%s", os.Getenv("API_NAME"), os.Getenv("API_VERSION"))

	habit := data.Habit{
		Name:            "Test Update Habit",
		Days:            30,
		DaysTarget:      50,
		CompletionDates: append(data.MockHabit[0].CompletionDates, []string{"2021-09-01", "2021-09-02", "2021-09-03"}...),
	}

	marshalledUpdatedHabit, err := json.Marshal(data.UpdateHabit{
		Name:            &habit.Name,
		Days:            &habit.Days,
		DaysTarget:      &habit.DaysTarget,
		CompletionDates: &habit.CompletionDates,
	})

	if err != nil {
		t.Errorf("Fail err: %s", err)
	}

	testCases := []struct {
		name        string
		updateHabit data.Habit
		want        []byte
	}{
		{
			name:        "Test successful Update",
			updateHabit: data.Habit{Name: "Test Update Habit", Days: 30, DaysTarget: 50, CompletionDates: []string{"2021-09-01", "2021-09-02", "2021-09-03"}},
			want:        marshalledUpdatedHabit,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			newHabit := data.UpdateHabit{
				Name:            &val.updateHabit.Name,
				Days:            &val.updateHabit.Days,
				DaysTarget:      &val.updateHabit.DaysTarget,
				CompletionDates: &val.updateHabit.CompletionDates,
			}

			marshalledNewHabit, err := json.Marshal(newHabit)

			if err != nil {
				t.Errorf("Fail err: %s", err)
			}

			req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/%s/UpdateHabit", endpoint), io.NopCloser(bytes.NewBuffer(marshalledNewHabit)))
			w := httptest.NewRecorder()
			q := req.URL.Query()
			q.Add("habitId", "1")
			req.URL.RawQuery = q.Encode()

			claims := &session.Claims{Username: "johndoe1@example.com"}

			ctx := context.WithValue(req.Context(), session.ClaimsKey, claims)

			req = req.WithContext(ctx)

			c.UpdateHabitsHandler(w, req)

			if status := w.Code; status == http.StatusInternalServerError {
				t.Errorf("TestUpdateHabitsHandler - HTTP Status Code = %d", status)
				return
			}

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

func TestDeleteHabitsHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewDB(logger)
	habitsModel := model.NewHabitsModel(logger, db)
	habitsView := view.NewHabitsView(logger)
	c := NewHabitsController(habitsModel, habitsView, logger)

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
			q.Add("habitId", "1")
			req.URL.RawQuery = q.Encode()

			claims := &session.Claims{Username: "johndoe1@example.com"}

			ctx := context.WithValue(req.Context(), session.ClaimsKey, claims)

			req = req.WithContext(ctx)

			expect, err := json.Marshal(map[string]bool{"success": true})

			if err != nil {
				t.Errorf("Fail err: %s", err)
			}

			c.DeleteHabitsHandler(w, req)

			if status := w.Code; status == http.StatusInternalServerError {
				t.Errorf("TestDeleteHabitsHandler - HTTP Status Code = %d", status)
				return
			}

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
