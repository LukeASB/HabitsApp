package controller

import (
	"bytes"
	"context"
	"dohabits/data"
	"dohabits/db"
	"dohabits/helper"
	"dohabits/logger"
	"dohabits/middleware/session"
	"dohabits/model"
	"dohabits/view"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateHabitsHandler(t *testing.T) {
	logger := logger.NewLogger(0)
	db := db.NewMockDB(logger)
	habitsModel := model.NewHabitsModel(logger, db)
	habitsView := view.NewHabitsView(logger)
	c := NewHabitsController(habitsModel, habitsView, logger)

	// Make a deep copy of the original state
	originalMockHabitState := make([]data.Habit, len(data.MockHabit))
	copy(originalMockHabitState, data.MockHabit)

	testCases := []struct {
		name        string
		newHabitReq string
		newHabitRes string
		want        []byte
	}{
		{
			name:        "Test successful Create",
			newHabitReq: `{"name": "Test Create Habit", "daysTarget": 50}`,
			want:        []byte("{\"habitId\":\"7\",\"name\":\"Test Create Habit\",\"daysTarget\":50,\"completionDates\":[]}"),
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/CreateHabit", io.NopCloser(bytes.NewBuffer([]byte(val.newHabitReq))))
			w := httptest.NewRecorder()

			claims := &session.Claims{Username: "johndoe1@example.com"}

			ctx := context.WithValue(req.Context(), session.ClaimsKey, claims)

			req = req.WithContext(ctx)

			c.CreateHabitsHandler(w, req)

			if status := w.Code; status == http.StatusInternalServerError {
				t.Errorf("%s - Failed - HTTP Status Code = %d", helper.GetFunctionName(), status)
				return
			}

			res := w.Result()

			defer res.Body.Close()

			got, err := io.ReadAll(res.Body)

			if err != nil {
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
				return
			}

			if !bytes.Equal(val.want, got) {
				t.Errorf("%s - Failed - got=%s, want=%s", helper.GetFunctionName(), string(got), string(val.want))
				return
			}
		})
	}

	data.MockHabit = originalMockHabitState
}

func TestRetrieveHabitsHandler(t *testing.T) {
	logger := logger.NewLogger(0)
	db := db.NewMockDB(logger)
	habitsModel := model.NewHabitsModel(logger, db)
	habitsView := view.NewHabitsView(logger)
	c := NewHabitsController(habitsModel, habitsView, logger)

	marshalledHabit, err := json.Marshal(data.MockHabit[0])

	if err != nil {
		t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
		return
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
			req := httptest.NewRequest(http.MethodGet, "/RetrieveAllHabits", nil)
			w := httptest.NewRecorder()
			q := req.URL.Query()
			q.Add("habitId", "1")
			req.URL.RawQuery = q.Encode()

			claims := &session.Claims{Username: "johndoe1@example.com"}

			ctx := context.WithValue(req.Context(), session.ClaimsKey, claims)

			req = req.WithContext(ctx)

			c.RetrieveHabitsHandler(w, req)

			if status := w.Code; status == http.StatusInternalServerError {
				t.Errorf("%s - Failed - HTTP Status Code = %d", helper.GetFunctionName(), status)
				return
			}

			res := w.Result()

			defer res.Body.Close()

			got, err := io.ReadAll(res.Body)

			if err != nil {
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
				return
			}

			if !bytes.Equal(val.want, got) {
				t.Errorf("%s - Failed - got=%s, want=%s", helper.GetFunctionName(), string(got), string(val.want))
				return
			}
		})
	}
}

func TestRetrieveAllHabitsHandler(t *testing.T) {
	logger := logger.NewLogger(0)
	db := db.NewMockDB(logger)
	habitsModel := model.NewHabitsModel(logger, db)
	habitsView := view.NewHabitsView(logger)
	c := NewHabitsController(habitsModel, habitsView, logger)

	userMockHabits := []data.Habit{}

	for _, val := range data.MockHabit {
		if val.UserID == "1" {
			userMockHabits = append(userMockHabits, val)
		}
	}

	marshalledAllHabits, err := json.Marshal(userMockHabits)

	if err != nil {
		t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
		return
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
			req := httptest.NewRequest(http.MethodGet, "/RetrieveAllHabits", nil)
			w := httptest.NewRecorder()

			claims := &session.Claims{Username: "johndoe1@example.com"}

			ctx := context.WithValue(req.Context(), session.ClaimsKey, claims)

			req = req.WithContext(ctx)

			c.RetrieveAllHabitsHandler(w, req)

			if status := w.Code; status == http.StatusInternalServerError {
				t.Errorf("%s - Failed - HTTP Status Code = %d", helper.GetFunctionName(), status)
				return
			}

			res := w.Result()

			defer res.Body.Close()

			got, err := io.ReadAll(res.Body)

			if err != nil {
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
				return
			}

			if !bytes.Equal(val.want, got) {
				t.Errorf("%s - Failed - got=%s, want=%s", helper.GetFunctionName(), string(got), string(val.want))
				return
			}
		})
	}
}

func TestUpdateHabitsHandler(t *testing.T) {
	logger := logger.NewLogger(0)
	db := db.NewMockDB(logger)
	habitsModel := model.NewHabitsModel(logger, db)
	habitsView := view.NewHabitsView(logger)
	c := NewHabitsController(habitsModel, habitsView, logger)

	// Make a deep copy of the original state
	originalMockHabitState := make([]data.Habit, len(data.MockHabit))
	copy(originalMockHabitState, data.MockHabit)

	testCases := []struct {
		name        string
		updateHabit data.Habit
		want        []byte
	}{
		{
			name:        "Test successful Update",
			updateHabit: data.Habit{Name: "Test Update Habit", Days: 30, DaysTarget: 50, CompletionDates: []string{"2021-09-01", "2021-09-02", "2021-09-03"}},
			want:        []byte("{\"habitId\":\"\",\"name\":\"Test Update Habit\",\"days\":30,\"daysTarget\":50,\"completionDates\":[\"2021-09-01\",\"2021-09-02\",\"2021-09-03\"]}"),
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
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
			}

			req := httptest.NewRequest(http.MethodPut, "/UpdateHabit", io.NopCloser(bytes.NewBuffer(marshalledNewHabit)))
			w := httptest.NewRecorder()

			claims := &session.Claims{Username: "johndoe1@example.com"}

			ctx := context.WithValue(req.Context(), session.ClaimsKey, claims)

			req = req.WithContext(ctx)

			c.UpdateHabitsHandler(w, req)

			if status := w.Code; status == http.StatusInternalServerError {
				t.Errorf("%s - Failed - HTTP Status Code = %d", helper.GetFunctionName(), status)
				return
			}

			res := w.Result()

			defer res.Body.Close()

			got, err := io.ReadAll(res.Body)

			if err != nil {
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
			}

			if !bytes.Equal(val.want, got) {
				t.Errorf("%s - Failed - got=%s, want=%s", helper.GetFunctionName(), string(got), string(val.want))
			}
		})
	}

	data.MockHabit = originalMockHabitState
}

func TestUpdateAllHabitsHandler(t *testing.T) {
	logger := logger.NewLogger(0)
	db := db.NewMockDB(logger)
	habitsModel := model.NewHabitsModel(logger, db)
	habitsView := view.NewHabitsView(logger)
	c := NewHabitsController(habitsModel, habitsView, logger)

	// Make a deep copy of the original state
	originalMockHabitState := make([]data.Habit, len(data.MockHabit))
	copy(originalMockHabitState, data.MockHabit)

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
			want: []byte("[{\"habitId\":\"1\",\"name\":\"Test Update Habit 1\",\"days\":30,\"daysTarget\":50,\"completionDates\":[\"2021-09-01\",\"2021-09-02\",\"2021-09-03\"]},{\"habitId\":\"2\",\"name\":\"Test Update Habit 2\",\"days\":30,\"daysTarget\":50,\"completionDates\":[\"2021-09-01\",\"2021-09-02\",\"2021-09-03\"]},{\"habitId\":\"6\",\"name\":\"Test Update Habit 3\",\"days\":30,\"daysTarget\":50,\"completionDates\":[\"2021-09-01\",\"2021-09-02\",\"2021-09-03\"]}]"),
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
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
				return
			}

			req := httptest.NewRequest(http.MethodPut, "/UpdateHabits", io.NopCloser(bytes.NewBuffer(marshalledNewHabit)))
			w := httptest.NewRecorder()

			claims := &session.Claims{Username: "johndoe1@example.com"}

			ctx := context.WithValue(req.Context(), session.ClaimsKey, claims)

			req = req.WithContext(ctx)

			c.UpdateAllHabitsHandler(w, req)

			if status := w.Code; status == http.StatusInternalServerError {
				t.Errorf("%s - Failed - HTTP Status Code = %d", helper.GetFunctionName(), status)
				return
			}

			res := w.Result()

			defer res.Body.Close()

			got, err := io.ReadAll(res.Body)

			if err != nil {
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
				return
			}

			if !bytes.Equal(val.want, got) {
				t.Errorf("%s - Failed - got=%s, want=%s", helper.GetFunctionName(), string(got), string(val.want))
				return
			}
		})
	}

	data.MockHabit = originalMockHabitState
}

func TestDeleteHabitsHandler(t *testing.T) {
	logger := logger.NewLogger(0)
	db := db.NewMockDB(logger)
	habitsModel := model.NewHabitsModel(logger, db)
	habitsView := view.NewHabitsView(logger)
	c := NewHabitsController(habitsModel, habitsView, logger)

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
			req := httptest.NewRequest(http.MethodDelete, "/DeleteHabit", nil)
			w := httptest.NewRecorder()
			q := req.URL.Query()
			q.Add("habitId", "1")
			req.URL.RawQuery = q.Encode()

			claims := &session.Claims{Username: "johndoe1@example.com"}

			ctx := context.WithValue(req.Context(), session.ClaimsKey, claims)

			req = req.WithContext(ctx)

			expect, err := json.Marshal(map[string]bool{"success": true})

			if err != nil {
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
				return
			}

			c.DeleteHabitsHandler(w, req)

			if status := w.Code; status == http.StatusInternalServerError {
				t.Errorf("%s - Failed - HTTP Status Code = %d", helper.GetFunctionName(), status)
				return
			}

			res := w.Result()

			defer res.Body.Close()

			got, err := io.ReadAll(res.Body)

			if err != nil {
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
				return
			}

			if !bytes.Equal(expect, got) {
				t.Errorf("%s - Failed - got=%s, want=%s", helper.GetFunctionName(), string(got), string(expect))
				return
			}
		})
	}
	data.MockHabit = originalMockHabitState
}

func BenchmarkCreateHabitsHandler(b *testing.B) {
	logger := logger.NewLogger(0)
	db := db.NewMockDB(logger)
	habitsModel := model.NewHabitsModel(logger, db)
	habitsView := view.NewHabitsView(logger)
	c := NewHabitsController(habitsModel, habitsView, logger)

	// Make a deep copy of the original state
	originalMockHabitState := make([]data.Habit, len(data.MockHabit))
	copy(originalMockHabitState, data.MockHabit)

	marshalledNewHabitRequest, err := json.Marshal(data.NewHabit{
		Name:       "Test Create Habit",
		DaysTarget: 50,
	})

	if err != nil {
		b.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
		return
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			req := httptest.NewRequest(http.MethodPost, "/CreateHabit", io.NopCloser(bytes.NewBuffer(marshalledNewHabitRequest)))
			w := httptest.NewRecorder()

			claims := &session.Claims{Username: "johndoe1@example.com"}

			ctx := context.WithValue(req.Context(), session.ClaimsKey, claims)

			req = req.WithContext(ctx)

			c.CreateHabitsHandler(w, req)
		}
	})
}

func BenchmarkRetrieveHabitsHandler(b *testing.B) {
	logger := logger.NewLogger(0)
	db := db.NewMockDB(logger)
	habitsModel := model.NewHabitsModel(logger, db)
	habitsView := view.NewHabitsView(logger)
	c := NewHabitsController(habitsModel, habitsView, logger)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			req := httptest.NewRequest(http.MethodGet, "/RetrieveAllHabits", nil)
			w := httptest.NewRecorder()
			q := req.URL.Query()
			q.Add("habitId", "1")
			req.URL.RawQuery = q.Encode()

			claims := &session.Claims{Username: "johndoe1@example.com"}

			ctx := context.WithValue(req.Context(), session.ClaimsKey, claims)

			req = req.WithContext(ctx)

			c.RetrieveHabitsHandler(w, req)
		}
	})

}

func BenchmarkRetrieveAllHabitsHandler(b *testing.B) {
	logger := logger.NewLogger(0)
	db := db.NewMockDB(logger)
	habitsModel := model.NewHabitsModel(logger, db)
	habitsView := view.NewHabitsView(logger)
	c := NewHabitsController(habitsModel, habitsView, logger)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			req := httptest.NewRequest(http.MethodGet, "/RetrieveAllHabits", nil)
			w := httptest.NewRecorder()

			claims := &session.Claims{Username: "johndoe1@example.com"}

			ctx := context.WithValue(req.Context(), session.ClaimsKey, claims)

			req = req.WithContext(ctx)

			c.RetrieveAllHabitsHandler(w, req)
		}
	})
}

func BenchmarkUpdateHabitsHandler(b *testing.B) {
	logger := logger.NewLogger(0)
	db := db.NewMockDB(logger)
	habitsModel := model.NewHabitsModel(logger, db)
	habitsView := view.NewHabitsView(logger)
	c := NewHabitsController(habitsModel, habitsView, logger)

	// Make a deep copy of the original state
	originalMockHabitState := make([]data.Habit, len(data.MockHabit))
	copy(originalMockHabitState, data.MockHabit)

	habit := data.Habit{
		Name:            "Test Update Habit",
		Days:            30,
		DaysTarget:      50,
		CompletionDates: []string{"2021-09-01", "2021-09-02", "2021-09-03"},
	}

	newHabit := data.UpdateHabit{
		HabitID:         "1",
		Name:            &habit.Name,
		Days:            &habit.Days,
		DaysTarget:      &habit.DaysTarget,
		CompletionDates: &habit.CompletionDates,
	}

	marshalledNewHabit, err := json.Marshal(newHabit)

	if err != nil {
		b.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
		return
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			req := httptest.NewRequest(http.MethodPut, "/UpdateHabit", io.NopCloser(bytes.NewBuffer(marshalledNewHabit)))
			w := httptest.NewRecorder()

			claims := &session.Claims{Username: "johndoe1@example.com"}

			ctx := context.WithValue(req.Context(), session.ClaimsKey, claims)

			req = req.WithContext(ctx)

			c.UpdateHabitsHandler(w, req)
		}
	})
}

func BenchmarkUpdateAllHabitsHandler(b *testing.B) {
	logger := logger.NewLogger(0)
	db := db.NewMockDB(logger)
	habitsModel := model.NewHabitsModel(logger, db)
	habitsView := view.NewHabitsView(logger)
	c := NewHabitsController(habitsModel, habitsView, logger)

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
		b.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
		return
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			req := httptest.NewRequest(http.MethodPut, "/UpdateHabits", io.NopCloser(bytes.NewBuffer(marshalledUpdatedHabit)))
			w := httptest.NewRecorder()

			claims := &session.Claims{Username: "johndoe1@example.com"}

			ctx := context.WithValue(req.Context(), session.ClaimsKey, claims)

			req = req.WithContext(ctx)

			c.UpdateAllHabitsHandler(w, req)
		}
	})

}

func BenchmarkDeleteHabitsHandler(b *testing.B) {
	logger := logger.NewLogger(0)
	db := db.NewMockDB(logger)
	habitsModel := model.NewHabitsModel(logger, db)
	habitsView := view.NewHabitsView(logger)
	c := NewHabitsController(habitsModel, habitsView, logger)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			req := httptest.NewRequest(http.MethodDelete, "/DeleteHabit", nil)
			w := httptest.NewRecorder()
			q := req.URL.Query()
			q.Add("habitId", "1")
			req.URL.RawQuery = q.Encode()

			claims := &session.Claims{Username: "johndoe1@example.com"}

			ctx := context.WithValue(req.Context(), session.ClaimsKey, claims)

			req = req.WithContext(ctx)

			c.DeleteHabitsHandler(w, req)
		}
	})
}
