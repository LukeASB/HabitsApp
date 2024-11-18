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

func (db *MyMockDB) RegisterUser(value interface{}) error {
	db.logger.InfoLog("mock_db.RegisterUser")

	newUser, ok := value.(data.UserData)

	if !ok {
		db.logger.ErrorLog("mock_db.RegisterUser - value type is not data.UserData")
		return fmt.Errorf("mock_db.RegisterUser - value type is not data.UserData")
	}

	data.MockUsers = append(data.MockUsers, newUser)

	return nil
}

func (db *MyMockDB) LoginUser(value interface{}) error {
	db.logger.InfoLog("mock_db.LoginUser")

	userSession, ok := value.(*data.UserSession)

	if !ok {
		db.logger.ErrorLog("mock_db.LoginUser - value type is not data.UserSession")
		return fmt.Errorf("mock_db.LoginUser - value type is not data.UserSession")
	}
	var sessionID int

	if len(data.MockUserSession) > 0 {
		id, err := strconv.Atoi(data.MockUserSession[len(data.MockUserSession)-1].ID)

		if err != nil {
			db.logger.ErrorLog("mock_db.LoginUser - failed to get latest id")
			return fmt.Errorf("mock_db.LoginUser - failed to get latest id")
		}

		sessionID = id
	}

	userSession.ID = fmt.Sprintf("%v", sessionID+1)

	data.MockUserSession = append(data.MockUserSession, *userSession)

	for i, val := range data.MockUsers {
		if val.UserID == userSession.UserID {
			data.MockUsers[i].IsLoggedIn = true
			data.MockUsers[i].LastLogin = userSession.CreatedAt
		}
	}

	return nil
}

func (db *MyMockDB) LogoutUser(value interface{}) error {
	db.logger.InfoLog("mock_db.LogoutUser")
	userLoggedOut, ok := value.(*data.UserLoggedOutRequest)

	if !ok {
		db.logger.ErrorLog("mock_db.LogoutUser - value type is not data.UserLoggedOutRequest")
		return fmt.Errorf("mock_db.LogoutUser - value type is not data.UserLoggedOutRequest")
	}

	// Remove user session from struct
	for i, val := range data.MockUserSession {
		if val.UserID == userLoggedOut.UserID {
			data.MockUserSession = append(data.MockUserSession[:i], data.MockUserSession[i+1:]...)
		}
	}

	for i, val := range data.MockUsers {
		if val.UserID == userLoggedOut.UserID {
			data.MockUsers[i].IsLoggedIn = false
		}
	}

	return nil
}

func (db *MyMockDB) GetUserDetails(value interface{}) (interface{}, error) {
	db.logger.InfoLog("mock_db.GetUserDetails")

	if userAuth, ok := value.(*data.UserAuth); ok {
		for _, val := range data.MockUsers {
			if val.EmailAddress == userAuth.EmailAddress {
				return val, nil
			}
		}

		return nil, fmt.Errorf("mock_db.GetUserData - User doesn't exist")
	}

	if userLoggedOutRequest, ok := value.(*data.UserLoggedOutRequest); ok {
		for _, val := range data.MockUsers {
			if val.EmailAddress == userLoggedOutRequest.EmailAddress {
				return val, nil
			}
		}

		return nil, fmt.Errorf("mock_db.GetUserData - User doesn't exist")
	}

	db.logger.ErrorLog("mock_db.GetUserData - value type is unsupported")
	return nil, fmt.Errorf("mock_db.GetUserData - value type is unsupported")
}

func (db *MyMockDB) CreateHabitsHandler(value interface{}) error {
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

func (db *MyMockDB) RetrieveAllHabitsHandler() (interface{}, error) {
	db.logger.InfoLog("mock_db.RetrieveAll")
	return data.MockHabit, nil
}

func (db *MyMockDB) RetrieveHabitsHandler(id string) (interface{}, error) {
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

func (db *MyMockDB) UpdateHabitsHandler(id string, value interface{}) error {
	db.logger.InfoLog("mock_db.Update")
	newHabit, ok := value.(data.Habit)

	if !ok {
		err := "mock_db.Update - value type is not data.Habit"
		db.logger.ErrorLog(err)
		return fmt.Errorf(err)
	}

	for i, val := range data.MockHabit {
		if val.ID == id {
			db.logger.InfoLog(fmt.Sprintf("mock_db.UpdateHabitsHandler() match id=%s, val=%s\n", val.ID, val.Name))
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

func (db *MyMockDB) DeleteHabitsHandler(id string) error {
	db.logger.InfoLog("mock_db.Delete")

	for i, val := range data.MockHabit {
		if val.ID == id {
			db.logger.InfoLog(fmt.Sprintf("mock_db.DeleteHabitsHandler() match id=%s, val=%s\n", val.ID, val.Name))
			data.MockHabit = append(data.MockHabit[:i], data.MockHabit[i+1:]...)
			return nil
		}
	}

	err := "mock_db.Delete - Failed to delete"
	db.logger.ErrorLog(err)
	return fmt.Errorf(err)
}
