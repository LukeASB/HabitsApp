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

func (db *MyMockDB) RegisterUserHandler(value interface{}) (interface{}, error) {
	db.logger.InfoLog("mock_db.RegisterUserHandler")

	newUser, ok := value.(*data.RegisterUserRequest)

	if !ok {
		db.logger.ErrorLog("mock_db.RegisterUserHandler - value type is not data.UserData")
		return nil, fmt.Errorf("mock_db.RegisterUserHandler - value type is not data.UserData")
	}

	latestUserID, err := strconv.Atoi(data.MockUsers[len(data.MockUsers)-1].UserID)

	if err != nil {
		db.logger.ErrorLog("mock_db.RegisterUserHandler - get latestUserID and convert to int")
		return nil, fmt.Errorf("mock_db.RegisterUserHandler - couldn't get latestUserID and convert to int")
	}

	registerUser := data.UserData{
		UserID:       strconv.Itoa(latestUserID + 1),
		FirstName:    newUser.FirstName,
		LastName:     newUser.LastName,
		Password:     newUser.Password,
		EmailAddress: newUser.EmailAddress,
		CreatedAt:    time.Now(),
	}

	data.MockUsers = append(data.MockUsers, registerUser)

	return &registerUser, err
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
	userLoggedOut, ok := value.(*data.UserData)

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

	if userAuth, ok := value.(*data.RegisterUserRequest); ok {
		for _, val := range data.MockUsers {
			if val.EmailAddress == userAuth.EmailAddress {
				return val, nil
			}
		}

		return data.UserData{}, nil
	}

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

func (db *MyMockDB) CreateHabitsHandler(userId string, value interface{}) error {
	db.logger.InfoLog(fmt.Sprintf("mock_db.Create = userId=%s", userId))
	newHabit, ok := value.(data.NewHabit)

	if !ok {
		db.logger.ErrorLog("mock_db.Create - value type is not data.Habit")
		return fmt.Errorf("mock_db.Create - value type is not data.Habit")
	}

	id, err := strconv.Atoi(data.MockHabit[len(data.MockHabit)-1].HabitID)

	if err != nil {
		db.logger.ErrorLog("mock_db.Create - failed to get latest id")
		return fmt.Errorf("mock_db.Create - failed to get latest id")
	}

	habit := data.Habit{
		HabitID:         fmt.Sprintf("%v", id+1),
		UserID:          userId,
		CreatedAt:       time.Now(),
		Name:            newHabit.Name,
		Days:            newHabit.Days,
		DaysTarget:      newHabit.DaysTarget,
		CompletionDates: []string{},
	}

	data.MockHabit = append(data.MockHabit, habit)

	return nil
}

func (db *MyMockDB) RetrieveAllHabitsHandler(userId string) (interface{}, error) {
	db.logger.InfoLog(fmt.Sprintf("mock_db.RetrieveAll - userId=%s", userId))

	var userMockHabits []data.Habit

	for _, habit := range data.MockHabit {
		if habit.UserID == userId {
			userMockHabits = append(userMockHabits, habit)
		}
	}

	return userMockHabits, nil
}

func (db *MyMockDB) RetrieveHabitsHandler(userId, habitId string) (interface{}, error) {
	db.logger.InfoLog(fmt.Sprintf("mock_db.Retrieve userId=%s, habitId=%s\n", userId, habitId))

	for _, val := range data.MockHabit {
		if val.UserID == userId && val.HabitID == habitId {
			db.logger.InfoLog(fmt.Sprintf("mock_db.Retrieve match habitId=%s, val=%s\n", val.HabitID, val.Name))
			return val, nil
		}
	}

	err := "mock_db.Retrieve - habit not found"
	db.logger.ErrorLog(err)
	return nil, fmt.Errorf(err)
}

func (db *MyMockDB) UpdateHabitsHandler(userId, habitId string, value interface{}) error {
	db.logger.InfoLog("mock_db.Update")
	newHabit, ok := value.(data.Habit)

	if !ok {
		err := "mock_db.Update - value type is not data.Habit"
		db.logger.ErrorLog(err)
		return fmt.Errorf(err)
	}

	for i, val := range data.MockHabit {
		if val.UserID == userId && val.HabitID == habitId {
			db.logger.InfoLog(fmt.Sprintf("mock_db.UpdateHabitsHandler() match userId=%s, habitId=%s, val=%s\n", val.UserID, val.HabitID, val.Name))
			data.MockHabit[i].Name = newHabit.Name
			data.MockHabit[i].Days = newHabit.Days
			data.MockHabit[i].DaysTarget = newHabit.DaysTarget
			data.MockHabit[i].CompletionDates = newHabit.CompletionDates

			return nil
		}
	}

	err := "mock_db.Update - Failed to update"
	db.logger.ErrorLog(err)
	return fmt.Errorf(err)
}

func (db *MyMockDB) DeleteHabitsHandler(userId, habitId string) error {
	db.logger.InfoLog("mock_db.Delete")

	for i, val := range data.MockHabit {
		if val.UserID == userId && val.HabitID == habitId {
			db.logger.InfoLog(fmt.Sprintf("mock_db.DeleteHabitsHandler() match userId=%s, habitId=%s, val=%s\n", val.UserID, val.HabitID, val.Name))
			data.MockHabit = append(data.MockHabit[:i], data.MockHabit[i+1:]...)
			return nil
		}
	}

	err := "mock_db.Delete - Failed to delete"
	db.logger.ErrorLog(err)
	return fmt.Errorf(err)
}
