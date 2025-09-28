package db

import (
	"dohabits/data"
	"dohabits/helper"
	"dohabits/logger"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Enforce interface compliance
var _ IDB = (*MyMockDB)(nil)

var refreshTokenPath = "data/mock_refresh_tokens"
var refreshTokenFile = "mock_refresh_token.txt"

type MyMockDB struct {
	logger logger.ILogger
}

func NewMockDB(logger logger.ILogger) *MyMockDB {
	return &MyMockDB{
		logger: logger,
	}
}

func (db *MyMockDB) Connect() error {
	connectionString := os.Getenv("DB_URL")
	db.logger.InfoLog(helper.GetFunctionName(), "")
	db.logger.DebugLog(helper.GetFunctionName(), fmt.Sprintf("%s\n", connectionString))
	return nil
}

func (db *MyMockDB) Disconnect() error {
	db.logger.InfoLog(helper.GetFunctionName(), "")
	return nil
}

func (db *MyMockDB) RegisterUserHandler(value interface{}) (interface{}, error) {
	db.logger.InfoLog(helper.GetFunctionName(), "")

	newUser, ok := value.(*data.RegisterUserRequest)

	if !ok {
		db.logger.ErrorLog(helper.GetFunctionName(), "value type is not data.UserData")
		return nil, fmt.Errorf("%s - value type is not data.UserData", helper.GetFunctionName())
	}

	latestUserID, err := strconv.Atoi(data.MockUsers[len(data.MockUsers)-1].UserID)

	if err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), "get latestUserID and convert to int")
		return nil, fmt.Errorf("%s - couldn't get latestUserID and convert to int", helper.GetFunctionName())
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
	db.logger.InfoLog(helper.GetFunctionName(), "")

	userSession, ok := value.(*data.UserSession)

	if !ok {
		db.logger.ErrorLog(helper.GetFunctionName(), "value type is not data.UserSession")
		return fmt.Errorf("%s - value type is not data.UserSession", helper.GetFunctionName())
	}

	data.MockUserSession = append(data.MockUserSession, *userSession)

	for i, val := range data.MockUsers {
		if val.UserID == userSession.UserID {

			err := os.WriteFile(fmt.Sprintf("../%s/%s_%s", refreshTokenPath, val.EmailAddress, refreshTokenFile), []byte(userSession.RefreshToken), 0644)
			if err != nil {
				db.logger.ErrorLog(helper.GetFunctionName(), "Failed to store the Refresh Token")
				return fmt.Errorf("%s - Failed to store the Refresh Token", helper.GetFunctionName())
			}
			data.MockUsers[i].LastLogin = userSession.CreatedAt
		}
	}

	return nil
}

func (db *MyMockDB) LogoutUser(value interface{}) error {
	db.logger.InfoLog(helper.GetFunctionName(), "")
	userLoggedOut, ok := value.(*data.UserData)

	if !ok {
		db.logger.ErrorLog(helper.GetFunctionName(), "value type is not data.UserLoggedOutRequest")
		return fmt.Errorf("%s - value type is not data.UserLoggedOutRequest", helper.GetFunctionName())
	}

	// Remove user session from struct
	for i, val := range data.MockUserSession {
		if val.UserID == userLoggedOut.UserID {
			data.MockUserSession = append(data.MockUserSession[:i], data.MockUserSession[i+1:]...)
		}
	}

	for i, val := range data.MockUsers {
		if val.UserID == userLoggedOut.UserID {
			if err := os.Remove(fmt.Sprintf("../%s/%s_%s", refreshTokenPath, data.MockUsers[i].EmailAddress, refreshTokenFile)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (db *MyMockDB) RetrieveUserSession(value interface{}, userID string) (string, error) {
	username, ok := value.(string)

	if !ok {
		db.logger.ErrorLog(helper.GetFunctionName(), "value type is not data.MockRefreshJWT")
		return "", fmt.Errorf("%s - value type is not data.MockRefreshJWT", helper.GetFunctionName())
	}

	refreshToken, err := os.ReadFile(fmt.Sprintf("../%s/%s_%s", refreshTokenPath, username, refreshTokenFile))

	return string(refreshToken), err
}

func (db *MyMockDB) RetrieveUserDetails(value interface{}) (interface{}, error) {
	db.logger.InfoLog(helper.GetFunctionName(), "")

	if userAuth, ok := value.(*data.RegisterUserRequest); ok {
		for _, val := range data.MockUsers {
			if val.EmailAddress == userAuth.EmailAddress {
				return nil, fmt.Errorf("%s - User already exists", helper.GetFunctionName())
			}
		}

		return nil, nil
	}

	if userAuth, ok := value.(*data.UserAuth); ok {
		for _, val := range data.MockUsers {
			if val.EmailAddress == userAuth.EmailAddress {
				return &val, nil
			}
		}

		return nil, fmt.Errorf("%s - User doesn't exist", helper.GetFunctionName())
	}

	if userLoggedOutRequest, ok := value.(*data.UserLoggedOutRequest); ok {
		for _, val := range data.MockUsers {
			if val.EmailAddress == userLoggedOutRequest.EmailAddress {
				return &val, nil
			}
		}

		return nil, fmt.Errorf("%s - User doesn't exist", helper.GetFunctionName())
	}

	db.logger.ErrorLog(helper.GetFunctionName(), "value type is unsupported")
	return nil, fmt.Errorf("%s - value type is unsupported", helper.GetFunctionName())
}

func (db *MyMockDB) CreateHabitsHandler(userId string, value interface{}) (*data.NewHabitResponse, error) {
	db.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("userId=%s", userId))
	newHabit, ok := value.(data.NewHabit)

	if !ok {
		db.logger.ErrorLog(helper.GetFunctionName(), "value type is not data.Habit")
		return nil, fmt.Errorf("%s - value type is not data.Habit", helper.GetFunctionName())
	}

	var id int
	var err error
	if len(data.MockHabit) == 0 {
		id = 0 // Initial HabitID
	} else {
		id, err = strconv.Atoi(data.MockHabit[len(data.MockHabit)-1].HabitID)
	}

	if err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), "failed to get latest id")
		return nil, fmt.Errorf("%s - failed to get latest id", helper.GetFunctionName())
	}

	newHabitId := fmt.Sprintf("%v", id+1)

	habit := data.Habit{
		HabitID:         newHabitId,
		UserID:          userId,
		CreatedAt:       time.Now(),
		Name:            newHabit.Name,
		Days:            0,
		DaysTarget:      newHabit.DaysTarget,
		CompletionDates: []string{},
	}

	data.MockHabit = append(data.MockHabit, habit)

	return &data.NewHabitResponse{
		HabitID:    newHabitId,
		Name:       newHabit.Name,
		DaysTarget: newHabit.DaysTarget,
	}, nil
}

func (db *MyMockDB) RetrieveAllHabitsHandler(userId string) (interface{}, error) {
	db.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("userId=%s", userId))

	var userMockHabits []data.Habit

	for _, habit := range data.MockHabit {
		if habit.UserID == userId {
			userMockHabits = append(userMockHabits, habit)
		}
	}

	return userMockHabits, nil
}

func (db *MyMockDB) RetrieveHabitsHandler(userId, habitId string) (interface{}, error) {
	db.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("userId=%s, habitId=%s\n", userId, habitId))

	for _, val := range data.MockHabit {
		if val.UserID == userId && val.HabitID == habitId {
			db.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("match habitId=%s, val=%s\n", val.HabitID, val.Name))
			return val, nil
		}
	}

	err := "habit not found"
	db.logger.ErrorLog(helper.GetFunctionName(), err)
	return nil, fmt.Errorf("%s - %s", helper.GetFunctionName(), err)
}

func (db *MyMockDB) UpdateHabitsHandler(userId, habitId string, value interface{}) error {
	db.logger.InfoLog(helper.GetFunctionName(), "")
	newHabit, ok := value.(data.Habit)

	if !ok {
		err := "value type is not data.Habit"
		db.logger.ErrorLog(helper.GetFunctionName(), err)
		return fmt.Errorf("%s - %s", helper.GetFunctionName(), err)
	}

	for i, val := range data.MockHabit {
		if val.UserID == userId && val.HabitID == habitId {
			db.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("match userId=%s, habitId=%s, val=%s\n", val.UserID, val.HabitID, val.Name))
			data.MockHabit[i].Name = newHabit.Name
			data.MockHabit[i].Days = newHabit.Days
			data.MockHabit[i].DaysTarget = newHabit.DaysTarget
			data.MockHabit[i].CompletionDates = newHabit.CompletionDates

			return nil
		}
	}

	err := "Failed to update"
	db.logger.ErrorLog(helper.GetFunctionName(), err)
	return fmt.Errorf("%s - %s", helper.GetFunctionName(), err)
}

func (db *MyMockDB) UpdateAllHabitsHandler(userId string, value interface{}) error {
	db.logger.InfoLog(helper.GetFunctionName(), "")
	newHabits, ok := value.([]data.Habit)

	if !ok {
		err := "value type is not []data.Habit"
		db.logger.ErrorLog(helper.GetFunctionName(), err)
		return fmt.Errorf("%s - %s", helper.GetFunctionName(), err)
	}

	for i, val := range data.MockHabit {
		for _, newHabit := range newHabits {
			if val.UserID == userId && val.HabitID == newHabit.HabitID {
				db.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("match userId=%s, habitId=%s, val=%s\n", val.UserID, val.HabitID, val.Name))
				data.MockHabit[i].Name = newHabit.Name
				data.MockHabit[i].Days = newHabit.Days
				data.MockHabit[i].DaysTarget = newHabit.DaysTarget
				data.MockHabit[i].CompletionDates = newHabit.CompletionDates

				return nil
			}
		}
	}

	err := "Failed to update"
	db.logger.ErrorLog(helper.GetFunctionName(), err)
	return fmt.Errorf("%s - %s", helper.GetFunctionName(), err)
}

func (db *MyMockDB) DeleteHabitsHandler(userId, habitId string) error {
	db.logger.InfoLog(helper.GetFunctionName(), "")

	for i, val := range data.MockHabit {
		if val.UserID == userId && val.HabitID == habitId {
			db.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("match userId=%s, habitId=%s, val=%s\n", val.UserID, val.HabitID, val.Name))
			data.MockHabit = append(data.MockHabit[:i], data.MockHabit[i+1:]...)
			return nil
		}
	}

	err := "Failed to delete"
	db.logger.ErrorLog(helper.GetFunctionName(), err)
	return fmt.Errorf("%s - %s", helper.GetFunctionName(), err)
}
