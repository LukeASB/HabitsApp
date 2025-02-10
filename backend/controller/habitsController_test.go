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
	db := db.NewMockDB(logger)
	habitsModel := model.NewHabitsModel(logger, db)
	habitsView := view.NewHabitsView(logger)
	c := NewHabitsController(habitsModel, habitsView, logger)

	endpoint := fmt.Sprintf("%s/%s", os.Getenv("API_NAME"), os.Getenv("API_VERSION"))

	// Make a deep copy of the original state
	originalMockHabitState := make([]data.Habit, len(data.MockHabit))
	copy(originalMockHabitState, data.MockHabit)

	marshalledNewHabitRequest, err := json.Marshal(data.NewHabit{
		Name:       "Test Create Habit",
		DaysTarget: 50,
	})

	if err != nil {
		t.Errorf("Fail err: %s", err)
	}

	marshalledNewHabitResponse, err := json.Marshal(data.NewHabitResponse{
		HabitID:         "0",
		Name:            "Test Create Habit",
		DaysTarget:      50,
		CompletionDates: []string{},
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
			want: marshalledNewHabitResponse,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/%s/CreateHabit", endpoint), io.NopCloser(bytes.NewBuffer(marshalledNewHabitRequest)))
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

			newHabitResponse := &data.NewHabitResponse{}

			err = json.Unmarshal(val.want, newHabitResponse)

			if err != nil {
				t.Errorf("TestCreateHabitsHandler err: %s", err)
				return
			}

			getLatestHabitID := data.MockHabit[len(data.MockHabit)-1].HabitID

			if getLatestHabitID == "" {
				t.Errorf("TestCreate Failed - failed to get latest MockHabit HabitID")
			}

			newHabitResponse.HabitID = getLatestHabitID

			marshalledNewHabitResponse, err = json.Marshal(data.NewHabitResponse{
				HabitID:         newHabitResponse.HabitID,
				Name:            "Test Create Habit",
				DaysTarget:      50,
				CompletionDates: []string{},
			})

			if err != nil {
				t.Errorf("TestCreateHabitsHandler err: %s", err)
				return
			}

			val.want = marshalledNewHabitResponse

			if !bytes.Equal(val.want, got) {
				t.Errorf("Fail want doesn't match got")
			}
		})
	}

	data.MockHabit = originalMockHabitState
}

func TestRetrieveHabitsHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewMockDB(logger)
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
	db := db.NewMockDB(logger)
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
	db := db.NewMockDB(logger)
	habitsModel := model.NewHabitsModel(logger, db)
	habitsView := view.NewHabitsView(logger)
	c := NewHabitsController(habitsModel, habitsView, logger)

	endpoint := fmt.Sprintf("%s/%s", os.Getenv("API_NAME"), os.Getenv("API_VERSION"))

	// Make a deep copy of the original state
	originalMockHabitState := make([]data.Habit, len(data.MockHabit))
	copy(originalMockHabitState, data.MockHabit)

	habit := data.Habit{
		Name:            "Test Update Habit",
		Days:            30,
		DaysTarget:      50,
		CompletionDates: []string{"2021-09-01", "2021-09-02", "2021-09-03"},
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
				HabitID:         "1",
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

	data.MockHabit = originalMockHabitState
}

func TestUpdateAllHabitsHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewMockDB(logger)
	habitsModel := model.NewHabitsModel(logger, db)
	habitsView := view.NewHabitsView(logger)
	c := NewHabitsController(habitsModel, habitsView, logger)

	endpoint := fmt.Sprintf("%s/%s", os.Getenv("API_NAME"), os.Getenv("API_VERSION"))

	// Make a deep copy of the original state
	originalMockHabitState := make([]data.Habit, len(data.MockHabit))
	copy(originalMockHabitState, data.MockHabit)

	habit := []data.Habit{
		{
			Name: "Test Update Habit 1", Days: 30, DaysTarget: 50, CompletionDates: []string{"2021-09-01", "2021-09-02", "2021-09-03"},
		},
		{
			Name: "Test Update Habit 2", Days: 30, DaysTarget: 50, CompletionDates: []string{"2021-09-01", "2021-09-02", "2021-09-03"},
		},
		{
			Name: "Test Update Habit 3", Days: 30, DaysTarget: 50, CompletionDates: []string{"2021-09-01", "2021-09-02", "2021-09-03"},
		},
	}

	marshalledUpdatedHabit, err := json.Marshal([]data.UpdateHabit{
		{
			HabitID:         "1",
			Name:            &habit[0].Name,
			Days:            &habit[0].Days,
			DaysTarget:      &habit[0].DaysTarget,
			CompletionDates: &habit[0].CompletionDates,
		},
		{
			HabitID:         "2",
			Name:            &habit[1].Name,
			Days:            &habit[1].Days,
			DaysTarget:      &habit[1].DaysTarget,
			CompletionDates: &habit[1].CompletionDates,
		},
		{
			HabitID:         "6",
			Name:            &habit[2].Name,
			Days:            &habit[2].Days,
			DaysTarget:      &habit[2].DaysTarget,
			CompletionDates: &habit[2].CompletionDates,
		},
	})

	if err != nil {
		t.Errorf("Fail err: %s", err)
	}

	testCases := []struct {
		name        string
		updateHabit []data.Habit
		want        []byte
	}{
		{
			name: "Test successful Update All",
			updateHabit: []data.Habit{
				{
					Name: "Test Update Habit 1", Days: 30, DaysTarget: 50, CompletionDates: []string{"2021-09-01", "2021-09-02", "2021-09-03"},
				},
				{
					Name: "Test Update Habit 2", Days: 30, DaysTarget: 50, CompletionDates: []string{"2021-09-01", "2021-09-02", "2021-09-03"},
				},
				{
					Name: "Test Update Habit 3", Days: 30, DaysTarget: 50, CompletionDates: []string{"2021-09-01", "2021-09-02", "2021-09-03"},
				},
			},
			want: marshalledUpdatedHabit,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			newHabit := []data.UpdateHabit{
				{
					HabitID:         "1",
					Name:            &val.updateHabit[0].Name,
					Days:            &val.updateHabit[0].Days,
					DaysTarget:      &val.updateHabit[0].DaysTarget,
					CompletionDates: &val.updateHabit[0].CompletionDates,
				},
				{
					HabitID:         "2",
					Name:            &val.updateHabit[1].Name,
					Days:            &val.updateHabit[1].Days,
					DaysTarget:      &val.updateHabit[1].DaysTarget,
					CompletionDates: &val.updateHabit[1].CompletionDates,
				},
				{
					HabitID:         "6",
					Name:            &val.updateHabit[2].Name,
					Days:            &val.updateHabit[2].Days,
					DaysTarget:      &val.updateHabit[2].DaysTarget,
					CompletionDates: &val.updateHabit[2].CompletionDates,
				},
			}

			marshalledNewHabit, err := json.Marshal(newHabit)

			if err != nil {
				t.Errorf("Fail err: %s", err)
			}

			req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/%s/UpdateHabits", endpoint), io.NopCloser(bytes.NewBuffer(marshalledNewHabit)))
			w := httptest.NewRecorder()

			claims := &session.Claims{Username: "johndoe1@example.com"}

			ctx := context.WithValue(req.Context(), session.ClaimsKey, claims)

			req = req.WithContext(ctx)

			c.UpdateAllHabitsHandler(w, req)

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
				t.Errorf("Fail want doesn't match got. got=%s, want=%s", got, val.want)
			}
		})
	}

	data.MockHabit = originalMockHabitState
}

func TestDeleteHabitsHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewMockDB(logger)
	habitsModel := model.NewHabitsModel(logger, db)
	habitsView := view.NewHabitsView(logger)
	c := NewHabitsController(habitsModel, habitsView, logger)

	endpoint := fmt.Sprintf("%s/%s", os.Getenv("API_NAME"), os.Getenv("API_VERSION"))

	// Make a deep copy of the original state
	originalMockHabitState := make([]data.Habit, len(data.MockHabit))
	copy(originalMockHabitState, data.MockHabit)

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
				t.Errorf("Fail want doesn't match got.")
			}
		})
	}
	data.MockHabit = originalMockHabitState
}
