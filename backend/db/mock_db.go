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
	logger logger.ILogger
}

func NewDB(logger logger.ILogger) *MyMockDB {
	return &MyMockDB{
		logger: logger,
	}
}

func (db *MyMockDB) Connect() error {
	connectionString := os.Getenv("DB_URL")
	db.logger.InfoLog("mock_db.Connect")
	db.logger.DebugLog(fmt.Sprintf("db - Connect() - %s\n", connectionString))
	return nil
}

func (db *MyMockDB) Disconnect() error {
	db.logger.InfoLog("mock_db.Disconnect")
	return nil
}

func (db *MyMockDB) Create(value interface{}) error {
	db.logger.InfoLog("mock_db.Create")
	newHabit, ok := value.(data.NewHabit)

	if !ok {
		db.logger.ErrorLog("mock_db.Create - value type is not data.Habit")
		return fmt.Errorf("mock_db.Create - value type is not data.Habit")
	}

	id, err := strconv.Atoi(data.MockHabit[len(data.MockHabit)-1].ID)

	if err != nil {
		db.logger.ErrorLog("mock_db.Create - failed to get latest id")
		return fmt.Errorf("mock_db.Create - failed to get latest id")
	}

	habit := data.Habit{
		ID:               fmt.Sprintf("%v", id+1),
		CreatedAt:        time.Now(),
		Name:             newHabit.Name,
		Days:             newHabit.Days,
		DaysTarget:       newHabit.DaysTarget,
		NumberOfAttempts: 0,
	}

	data.MockHabit = append(data.MockHabit, habit)

	return nil
}

func (db *MyMockDB) RetrieveAll() (interface{}, error) {
	db.logger.InfoLog("mock_db.RetrieveAll")
	return data.MockHabit, nil
}

func (db *MyMockDB) Retrieve(id string) (interface{}, error) {
	db.logger.InfoLog(fmt.Sprintf("mock_db.Retrieve id=%s\n", id))

	for _, val := range data.MockHabit {
		if val.ID == id {
			db.logger.InfoLog(fmt.Sprintf("mock_db.Retrieve match id=%s, val=%s\n", val.ID, val.Name))
			return val, nil
		}
	}

	err := "mock_db.Retrieve - habit not found"
	db.logger.ErrorLog(err)
	return nil, fmt.Errorf(err)
}

func (db *MyMockDB) Update(id string, value interface{}) error {
	db.logger.InfoLog("mock_db.Update")
	newHabit, ok := value.(data.Habit)

	if !ok {
		err := "mock_db.Update - value type is not data.Habit"
		db.logger.ErrorLog(err)
		return fmt.Errorf(err)
	}

	for i, val := range data.MockHabit {
		if val.ID == id {
			db.logger.InfoLog(fmt.Sprintf("mock_db.Update() match id=%s, val=%s\n", val.ID, val.Name))
			data.MockHabit[i].Name = newHabit.Name
			data.MockHabit[i].Days = newHabit.Days
			data.MockHabit[i].DaysTarget = newHabit.DaysTarget
			return nil
		}
	}

	err := "mock_db.Update - Failed to update"
	db.logger.ErrorLog(err)
	return fmt.Errorf(err)
}

func (db *MyMockDB) Delete(id string) error {
	db.logger.InfoLog("mock_db.Delete")

	for i, val := range data.MockHabit {
		if val.ID == id {
			db.logger.InfoLog(fmt.Sprintf("mock_db.Delete() match id=%s, val=%s\n", val.ID, val.Name))
			data.MockHabit = append(data.MockHabit[:i], data.MockHabit[i+1:]...)
			return nil
		}
	}

	err := "mock_db.Delete - Failed to delete"
	db.logger.ErrorLog(err)
	return fmt.Errorf(err)
}
