package db

import (
	"dohabits/data"
	"dohabits/logger"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Enforce interface compliance
var _ IDB = (*MyMockDB)(nil)

type MyMockDB struct {
}

func (db *MyMockDB) Connect(logger logger.ILogger) error {
	connectionString := os.Getenv("DB_URL")
	logger.InfoLog("mock_db.Connect")
	logger.DebugLog(fmt.Sprintf("db - Connect() - %s\n", connectionString))
	return nil
}

func (db *MyMockDB) Disconnect(logger logger.ILogger) error {
	logger.InfoLog("mock_db.Disconnect")
	return nil
}

func (db *MyMockDB) Create(logger logger.ILogger, value interface{}) error {
	logger.InfoLog("mock_db.Create")
	newHabit, ok := value.(data.NewHabit)

	if !ok {
		logger.ErrorLog("mock_db.Create - value type is not data.Habit")
		return fmt.Errorf("mock_db.Create - value type is not data.Habit")
	}

	id, err := strconv.Atoi(data.MockHabit[len(data.MockHabit)-1].ID)

	if err != nil {
		logger.ErrorLog("mock_db.Create - failed to get latest id")
		return fmt.Errorf("mock_db.Create - failed to get latest id")
	}

	habit := data.Habit{
		ID:         fmt.Sprintf("%v", id+1),
		CreatedAt:  time.Now(),
		Name:       newHabit.Name,
		Days:       newHabit.Days,
		DaysTarget: newHabit.DaysTarget,
	}

	data.MockHabit = append(data.MockHabit, habit)

	return nil
}

func (db *MyMockDB) RetrieveAll(logger logger.ILogger) (interface{}, error) {
	logger.InfoLog("mock_db.RetrieveAll")
	return data.MockHabit, nil
}

func (db *MyMockDB) Retrieve(logger logger.ILogger, id string) (interface{}, error) {
	logger.InfoLog(fmt.Sprintf("mock_db.Retrieve id=%s\n", id))

	for _, val := range data.MockHabit {
		if val.ID == id {
			logger.InfoLog(fmt.Sprintf("mock_db.Retrieve match id=%s, val=%s\n", val.ID, val.Name))
			return val, nil
		}
	}

	err := "mock_db.Retrieve - habit not found"
	logger.ErrorLog(err)
	return nil, fmt.Errorf(err)
}

func (db *MyMockDB) Update(logger logger.ILogger, id string, value interface{}) error {
	logger.InfoLog("mock_db.Update")
	newHabit, ok := value.(data.Habit)

	if !ok {
		err := "mock_db.Update - value type is not data.Habit"
		logger.ErrorLog(err)
		return fmt.Errorf(err)
	}

	for i, val := range data.MockHabit {
		if val.ID == id {
			logger.InfoLog(fmt.Sprintf("mock_db.Update() match id=%s, val=%s\n", val.ID, val.Name))
			data.MockHabit[i].Name = newHabit.Name
			data.MockHabit[i].Days = newHabit.Days
			data.MockHabit[i].DaysTarget = newHabit.DaysTarget
			return nil
		}
	}

	err := "mock_db.Update - Failed to update"
	logger.ErrorLog(err)
	return fmt.Errorf(err)
}

func (db *MyMockDB) Delete(logger logger.ILogger, id string) error {
	logger.InfoLog("mock_db.Delete")

	for i, val := range data.MockHabit {
		if val.ID == id {
			logger.InfoLog(fmt.Sprintf("mock_db.Delete() match id=%s, val=%s\n", val.ID, val.Name))
			data.MockHabit = append(data.MockHabit[:i], data.MockHabit[i+1:]...)
			return nil
		}
	}

	err := "mock_db.Delete - Failed to delete"
	logger.ErrorLog(err)
	return fmt.Errorf(err)
}
