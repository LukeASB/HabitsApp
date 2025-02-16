package db

import (
	"context"
	"dohabits/data"
	"dohabits/helper"
	"dohabits/logger"
	"errors"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

// Enfore interface compliance
var _ IDB = (*MongoDB)(nil)

type MongoDB struct {
	logger                logger.ILogger
	client                *mongo.Client
	habitsAppDBName       string
	usersCollection       string
	userSessionCollection string
	habitsCollection      string
}

func NewMongoDB(logger logger.ILogger) *MongoDB {
	return &MongoDB{
		logger:                logger,
		habitsAppDBName:       os.Getenv("DB_NAME"),
		usersCollection:       os.Getenv("USERS_COLLECTION"),
		userSessionCollection: os.Getenv("USER_SESSION_COLLECTION"),
		habitsCollection:      os.Getenv("HABITS_COLLECTION"),
	}
}

/*
See README.md for users document example
*/
func (db *MongoDB) NewUsersCollection() *mongo.Collection {
	return db.client.Database(db.habitsAppDBName).Collection(db.usersCollection)
}

/*
See README.md for user_session document example
*/
func (db *MongoDB) NewUsersSessionCollection() *mongo.Collection {
	return db.client.Database(db.habitsAppDBName).Collection(db.userSessionCollection)
}

/*
See README.md for habits document example
*/
func (db *MongoDB) NewHabitsCollection() *mongo.Collection {
	return db.client.Database(db.habitsAppDBName).Collection(db.habitsCollection)
}

func (db *MongoDB) Connect() error {

	connectionString := os.Getenv("DB_URL")

	if connectionString == "" {
		db.logger.ErrorLog(helper.GetFunctionName(), "connectionString is empty")
		return fmt.Errorf("%s - connectionString is empty", helper.GetFunctionName())
	}

	db.logger.InfoLog(helper.GetFunctionName(), "")
	db.logger.DebugLog(helper.GetFunctionName(), fmt.Sprintf("%s\n", connectionString))

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(opts)
	if err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("%s", err))
		return fmt.Errorf(fmt.Sprintf("%s - %s", helper.GetFunctionName(), err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Send a ping to confirm a successful connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("%s", err))
		return fmt.Errorf(fmt.Sprintf("%s - %s", helper.GetFunctionName(), err))
	}

	db.client = client
	db.logger.InfoLog(helper.GetFunctionName(), "Pinged your deployment. You successfully connected to MongoDB!")

	// Ensure the TTL index is created
	err = db.EnsureTTLIndex()
	if err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("failed to ensure TTL index: %v", err))
		return fmt.Errorf("failed to ensure TTL index: %v", err)
	}

	return nil
}

func (db *MongoDB) EnsureTTLIndex() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newUsersSessionCollection := db.NewUsersSessionCollection()

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"CreatedAt": 1},                       // Index on CreatedAt field
		Options: options.Index().SetExpireAfterSeconds(86400), // 24 hours
	}

	_, err := newUsersSessionCollection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to create TTL index: %v", err))
		return fmt.Errorf("failed to create TTL index: %v", err)
	}

	db.logger.InfoLog(helper.GetFunctionName(), "TTL index created successfully on CreatedAt field")
	return nil
}

func (db *MongoDB) Disconnect() error {
	if err := db.client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}

	db.logger.InfoLog(helper.GetFunctionName(), "Successfully disconnected from MongoDB")
	return nil
}

func (db *MongoDB) RegisterUserHandler(value interface{}) (interface{}, error) {
	db.logger.InfoLog(helper.GetFunctionName(), "")

	newUser, ok := value.(*data.RegisterUserRequest)

	if !ok {
		db.logger.ErrorLog(helper.GetFunctionName(), "value type is not data.UserData")
		return nil, fmt.Errorf("%s - value type is not data.UserData", helper.GetFunctionName())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newUsersCollection := db.NewUsersCollection()

	registerTime := time.Now()

	type registerUserData struct {
		Password     string    `bson:"Password"`
		FirstName    string    `bson:"FirstName"`
		LastName     string    `bson:"LastName"`
		EmailAddress string    `bson:"EmailAddress"`
		CreatedAt    time.Time `bson:"CreatedAt"`
		LastLogin    time.Time `bson:"LastLogin"`
	}

	registerUser := registerUserData{
		EmailAddress: newUser.EmailAddress,
		Password:     newUser.Password,
		FirstName:    newUser.FirstName,
		LastName:     newUser.LastName,
		CreatedAt:    registerTime,
		LastLogin:    registerTime,
	}

	insertResult, err := newUsersCollection.InsertOne(ctx, registerUser)
	if err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to insert new user: %v", err))
		return nil, fmt.Errorf("%s - Failed to insert new user: %v", helper.GetFunctionName(), err)
	}

	db.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("User registered successfully with EmailAddress: %s, Acknowledged: %v, InsertedID: %v", registerUser.EmailAddress, insertResult.Acknowledged, insertResult.InsertedID))

	return &data.UserData{
		FirstName:    newUser.FirstName,
		LastName:     newUser.LastName,
		Password:     newUser.Password,
		EmailAddress: newUser.EmailAddress,
		CreatedAt:    registerTime,
	}, nil
}

func (db *MongoDB) LoginUser(value interface{}) error {
	db.logger.InfoLog(helper.GetFunctionName(), "")

	userSession, ok := value.(*data.UserSession)
	if !ok {
		db.logger.ErrorLog(helper.GetFunctionName(), "value type is not data.UserSession")
		return fmt.Errorf("%s - value type is not data.UserSession", helper.GetFunctionName())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newUsersSessionCollection := db.NewUsersSessionCollection()
	newUsersCollection := db.NewUsersCollection()

	objectID, err := primitive.ObjectIDFromHex(userSession.UserID)
	if err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to insert user session for userId=%s", userSession.UserID))
		return fmt.Errorf("%s - Failed to insert user session for userId=%s: %v", helper.GetFunctionName(), userSession.UserID, err)
	}

	if userSession.CreatedAt.IsZero() {
		userSession.CreatedAt = time.Now()
	}

	type insertUserSessionData struct {
		UserID       bson.ObjectID `bson:"_id"`
		RefreshToken string        `bson:"RefreshToken"`
		Device       string        `bson:"Device"`
		IPAddress    string        `bson:"IpAddress"`
		CreatedAt    time.Time     `bson:"CreatedAt"`
	}

	insertUserSession := insertUserSessionData{
		UserID:       bson.ObjectID(objectID),
		RefreshToken: userSession.RefreshToken,
		Device:       userSession.Device,
		IPAddress:    userSession.IPAddress,
		CreatedAt:    userSession.CreatedAt,
	}

	_, err = newUsersSessionCollection.InsertOne(ctx, insertUserSession)
	if err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to insert user session for userId=%s", userSession.UserID))
		return fmt.Errorf("%s - Failed to insert user session for userId=%s: %v", helper.GetFunctionName(), userSession.UserID, err)
	}

	filter := bson.M{"_id": bson.ObjectID(objectID)}

	update := bson.M{
		"$set": bson.M{
			"LastLogin": userSession.CreatedAt,
		},
	}

	result, err := newUsersCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to update users collection for userId=%s", userSession.UserID))
		return fmt.Errorf("%s - Failed to update users collection for userId=%s: %v", helper.GetFunctionName(), userSession.UserID, err)
	}

	db.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("The document has been updated. ModifiedCount: %v, UpsertedCount: %v", result.ModifiedCount, result.UpsertedCount))

	return nil
}

func (db *MongoDB) LogoutUser(value interface{}) error {
	db.logger.InfoLog(helper.GetFunctionName(), "")

	userLoggedOut, ok := value.(*data.UserData)

	if !ok {
		db.logger.ErrorLog(helper.GetFunctionName(), "value type is not data.UserData")
		return fmt.Errorf("%s - value type is not data.UserData", helper.GetFunctionName())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newUsersSessionCollection := db.NewUsersSessionCollection()

	objectID, err := primitive.ObjectIDFromHex(userLoggedOut.UserID)
	if err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to retrieve the objectId from userId=%s", userLoggedOut.UserID))
		return fmt.Errorf("%s - Failed to retrieve the objectId from userId=%s: %v", helper.GetFunctionName(), userLoggedOut.UserID, err)
	}

	sessionFilter := bson.D{{Key: "_id", Value: bson.ObjectID(objectID)}}

	deleteResult, err := newUsersSessionCollection.DeleteOne(ctx, sessionFilter)
	if err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to delete user session for userId=%s", userLoggedOut.UserID))
		return fmt.Errorf("%s - Failed to delete user session for userId=%s: %v", helper.GetFunctionName(), userLoggedOut.UserID, err)
	}

	db.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("The document has been updated. Acknowledged: %v, DeleteCount: %v", deleteResult.Acknowledged, deleteResult.DeletedCount))

	return nil
}

func (db *MongoDB) RetrieveUserSession(value interface{}, userID string) (string, error) {
	if userID == "" {
		emailAddress, ok := value.(string)

		if !ok {
			db.logger.ErrorLog(helper.GetFunctionName(), "value type is not string")
			return "", fmt.Errorf("%s - value type is not string", helper.GetFunctionName())
		}

		userDetails, err := db.RetrieveUserDetails(&data.UserAuth{EmailAddress: emailAddress})

		if err != nil {
			return "", err
		}

		currentUserData, ok := userDetails.(*data.UserData)

		if !ok {
			return "", fmt.Errorf("%s - data.UserData is invalid", helper.GetFunctionName())
		}

		userID = currentUserData.UserID
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newUsersSessionCollection := db.NewUsersSessionCollection()

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to retrieve user session for userId=%s", userID))
		return "", fmt.Errorf("%s - Failed to retrieve user session for userId=%s: %v", helper.GetFunctionName(), userID, err)
	}

	filter := bson.M{"_id": bson.ObjectID(objectID)}

	var userSession data.UserSession
	var result bson.M

	err = newUsersSessionCollection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", fmt.Errorf("%s - User session doesn't exist", helper.GetFunctionName())
		}

		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to get user session err=%s", err))
		return "", fmt.Errorf("%s - Failed to get user session err=%s", helper.GetFunctionName(), err)
	}

	if refreshToken, ok := result["RefreshToken"].(string); ok {
		userSession.RefreshToken = refreshToken
	}

	return userSession.RefreshToken, nil
}

func (db *MongoDB) RetrieveUserDetails(value interface{}) (interface{}, error) {
	db.logger.InfoLog(helper.GetFunctionName(), "")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newUsersCollection := db.NewUsersCollection()

	if userRegisterRequest, ok := value.(*data.RegisterUserRequest); ok {
		filter := bson.M{"EmailAddress": userRegisterRequest.EmailAddress}
		var result bson.M

		err := newUsersCollection.FindOne(ctx, filter).Decode(&result)

		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, nil
			}

			db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to get user details err=%s", err))
			return nil, fmt.Errorf(helper.GetFunctionName(), "Failed to get user details err=%s", err)
		}

		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to get user details err=%s", err))
		return nil, fmt.Errorf("%s - Failed to get user details err=%s", helper.GetFunctionName(), err)
	}

	if userAuth, ok := value.(*data.UserAuth); ok {
		filter := bson.M{"EmailAddress": userAuth.EmailAddress}

		user, err := db.findUser(ctx, filter, newUsersCollection)

		if err != nil {
			db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to get user details err=%s", err))
			return nil, fmt.Errorf("%s - Failed to get user details err=%s", helper.GetFunctionName(), err)
		}

		return user, nil
	}

	if userLoggedOutRequest, ok := value.(*data.UserLoggedOutRequest); ok {
		filter := bson.M{"EmailAddress": userLoggedOutRequest.EmailAddress}

		user, err := db.findUser(ctx, filter, newUsersCollection)

		if err != nil {
			db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to get user details err=%s", err))
			return nil, fmt.Errorf("%s - Failed to get user details err=%s", helper.GetFunctionName(), err)
		}

		return user, nil
	}

	db.logger.ErrorLog(helper.GetFunctionName(), "value type is unsupported")
	return nil, fmt.Errorf("%s - value type is unsupported", helper.GetFunctionName())
}

func (db *MongoDB) findUser(ctx context.Context, filter bson.M, newUsersCollection *mongo.Collection) (*data.UserData, error) {
	var user data.UserData
	var result bson.M

	err := newUsersCollection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("%s - User doesn't exist", helper.GetFunctionName())
		}

		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to get user details err=%s", err))
		return nil, fmt.Errorf("%s - Failed to get user details err=%s", helper.GetFunctionName(), err)
	}

	if id, ok := result["_id"].(bson.ObjectID); ok {
		user.UserID = id.Hex()
	} else {
		return nil, fmt.Errorf("%s - _id field is missing or not an ObjectID", helper.GetFunctionName())
	}

	// Manually assign other fields from result to user struct
	if password, ok := result["Password"].(string); ok {
		user.Password = password
	}

	if firstName, ok := result["FirstName"].(string); ok {
		user.FirstName = firstName
	}

	if lastName, ok := result["LastName"].(string); ok {
		user.LastName = lastName
	}

	if emailAddress, ok := result["EmailAddress"].(string); ok {
		user.EmailAddress = emailAddress
	}

	if createdAt, ok := result["CreatedAt"].(bson.DateTime); ok {
		user.CreatedAt = createdAt.Time()
	}

	if lastLogin, ok := result["LastLogin"].(bson.DateTime); ok {
		user.LastLogin = lastLogin.Time()
	}

	return &user, nil
}

func (db *MongoDB) CreateHabitsHandler(userId string, value interface{}) (*data.NewHabitResponse, error) {
	db.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("userId=%s", userId))
	newHabit, ok := value.(data.NewHabit)

	if !ok {
		db.logger.ErrorLog(helper.GetFunctionName(), "value type is not data.Habit")
		return nil, fmt.Errorf("%s - value type is not data.Habit", helper.GetFunctionName())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newHabitsCollection := db.NewHabitsCollection()

	objectID, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to insert new habit: userId=%s, err=%v", userId, err))
		return nil, fmt.Errorf("%s - Failed to insert new habit: userId=%s, err=%v", helper.GetFunctionName(), userId, err)
	}

	type insertHabitData struct {
		UserID          bson.ObjectID `bson:"UserID"`
		CreatedAt       time.Time     `bson:"CreatedAt"`
		Name            string        `bson:"Name"`
		Days            int           `bson:"Days"`
		DaysTarget      int           `bson:"DaysTarget"`
		CompletionDates []string      `bson:"CompletionDates"`
	}

	insertHabit := insertHabitData{
		UserID:          bson.ObjectID(objectID),
		CreatedAt:       time.Now(),
		Name:            newHabit.Name,
		Days:            0,
		DaysTarget:      newHabit.DaysTarget,
		CompletionDates: []string{},
	}

	insertResult, err := newHabitsCollection.InsertOne(ctx, insertHabit)
	if err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to insert new habit: userId=%s, err=%v", userId, err))
		return nil, fmt.Errorf("%s - Failed to insert new habit: userId=%s, err=%v", helper.GetFunctionName(), userId, err)
	}

	db.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("habit inserted successfully with userId: %s, Acknowledged: %v, InsertedID: %v", userId, insertResult.Acknowledged, insertResult.InsertedID))

	newInsertHabitID, ok := insertResult.InsertedID.(bson.ObjectID)

	if !ok {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to insert new habit - insertResult.InsertedID is not bson.ObjectID: userId=%s", userId))
		return nil, fmt.Errorf("%s - Failed to insert new habit - insertResult.InsertedID is not bson.ObjectID: userId=%s", helper.GetFunctionName(), userId)
	}

	return &data.NewHabitResponse{
		HabitID:    newInsertHabitID.Hex(),
		Name:       newHabit.Name,
		DaysTarget: newHabit.DaysTarget,
	}, nil
}

func (db *MongoDB) RetrieveAllHabitsHandler(userId string) (interface{}, error) {
	db.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("userId=%s", userId))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newHabitsCollection := db.NewHabitsCollection()

	objectID, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to retrieve all habits: userId=%s, err=%v", userId, err))
		return nil, fmt.Errorf("%s - Failed to retrieve all habits: userId=%s, err=%v", helper.GetFunctionName(), userId, err)
	}

	filter := bson.M{"UserID": bson.ObjectID(objectID)}

	cur, err := newHabitsCollection.Find(ctx, filter)

	if err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to retrieve habit: err=%v", err))
		return nil, fmt.Errorf("%s - Failed to retrieve habit: err=%v", helper.GetFunctionName(), err)
	}
	var results []data.Habit

	for cur.Next(ctx) {
		var el bson.M
		var habit data.Habit
		err := cur.Decode(&el)

		if err != nil {
			db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to retrieve habit: err=%v", err))
			return nil, fmt.Errorf("%s - Failed to retrieve habit: err=%v", helper.GetFunctionName(), err)
		}

		if id, ok := el["_id"].(bson.ObjectID); ok {
			habit.HabitID = id.Hex()
		} else {
			return nil, fmt.Errorf("%s - _id field is missing or not an ObjectID", helper.GetFunctionName())
		}

		// Manually assign other fields from result to user struct
		if userId, ok := el["UserID"].(bson.ObjectID); ok {
			habit.UserID = userId.Hex()
		}

		if createdAt, ok := el["CreatedAt"].(bson.DateTime); ok {
			habit.CreatedAt = createdAt.Time()
		}

		if name, ok := el["Name"].(string); ok {
			habit.Name = name
		}

		if days, ok := el["Days"].(int32); ok {
			habit.Days = int(days)
		}

		if daysTarget, ok := el["DaysTarget"].(int32); ok {
			habit.DaysTarget = int(daysTarget)
		}

		if completionDates, ok := el["CompletionDates"].(bson.A); ok {
			var completionDate []string
			for _, date := range completionDates {
				if strDate, ok := date.(string); ok {
					completionDate = append(completionDate, strDate)
				} else {
					db.logger.ErrorLog(helper.GetFunctionName(), "non-string value in Completed Date")
				}
			}

			if len(completionDates) == 0 {
				habit.CompletionDates = []string{}
			} else {
				habit.CompletionDates = completionDate
			}
		} else {
			habit.CompletionDates = []string{}
		}

		results = append(results, habit)
	}

	if err := cur.Err(); err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to retrieve habit: err=%v", err))
		return nil, fmt.Errorf("%s - Failed to retrieve habit: err=%v", helper.GetFunctionName(), err)
	}

	return results, nil
}

func (db *MongoDB) RetrieveHabitsHandler(userId, habitId string) (interface{}, error) {
	db.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("userId=%s, habitId=%s\n", userId, habitId))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newHabitCollection := db.NewHabitsCollection()

	objectID, err := primitive.ObjectIDFromHex(habitId)
	if err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to retrieve habit: userId=%s, err=%v", userId, err))
		return nil, fmt.Errorf("%s - Failed to retrieve habit: userId=%s, err=%v", helper.GetFunctionName(), userId, err)
	}

	filter := bson.M{"_id": bson.ObjectID(objectID)}
	var result bson.M
	var habit data.Habit
	err = newHabitCollection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("%s - Habit doesn't exist", helper.GetFunctionName())
		}

		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to retrieve habit: userId=%s, err=%v", userId, err))
		return nil, fmt.Errorf("%s - Failed to retrieve habit: userId=%s, err=%v", helper.GetFunctionName(), userId, err)
	}

	// Manually assign other fields from result to user struct
	if habitId, ok := result["_id"].(bson.ObjectID); ok {
		habit.HabitID = habitId.Hex()
	}

	if userId, ok := result["UserID"].(bson.ObjectID); ok {
		habit.UserID = userId.Hex()
	}

	if createdAt, ok := result["CreatedAt"].(bson.DateTime); ok {
		habit.CreatedAt = createdAt.Time()
	}

	if name, ok := result["Name"].(string); ok {
		habit.Name = name
	}

	if days, ok := result["Days"].(int32); ok {
		habit.Days = int(days)
	}

	if daysTarget, ok := result["DaysTarget"].(int32); ok {
		habit.DaysTarget = int(daysTarget)
	}

	if completionDates, ok := result["CompletionDates"].(bson.A); ok {
		var completionDate []string
		for _, date := range completionDates {
			if strDate, ok := date.(string); ok {
				completionDate = append(completionDate, strDate)
			} else {
				db.logger.ErrorLog(helper.GetFunctionName(), "non-string value in Completed Date")
			}
		}

		habit.CompletionDates = completionDate
	}

	return habit, nil
}

func (db *MongoDB) UpdateHabitsHandler(userId, habitId string, value interface{}) error {
	db.logger.InfoLog(helper.GetFunctionName(), "")

	updateHabit, ok := value.(data.Habit)

	if !ok {
		err := "value type is not data.Habit"
		db.logger.ErrorLog(helper.GetFunctionName(), err)
		return fmt.Errorf("%s - %s", helper.GetFunctionName(), err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newHabitCollection := db.NewHabitsCollection()

	objectId, err := primitive.ObjectIDFromHex(habitId)

	if err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to update habits collection for userId=%s, habitId=%s, err=%s", userId, habitId, err))
		return fmt.Errorf("%s - Failed to update habits collection for userId=%s, habitId=%s, err=%s", helper.GetFunctionName(), userId, habitId, err)
	}

	filter := bson.M{"_id": bson.ObjectID(objectId)}

	update := bson.M{
		"$set": bson.M{
			"Name":            updateHabit.Name,
			"DaysTarget":      updateHabit.DaysTarget,
			"CompletionDates": updateHabit.CompletionDates,
		},
	}

	result, err := newHabitCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to update habits collection for userId=%s, habitId=%s, err=%s", userId, habitId, err))
		return fmt.Errorf("%s - Failed to update habits collection for userId=%s, habitId=%s, err=%s", helper.GetFunctionName(), userId, habitId, err)
	}

	db.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("The document has been updated. ModifiedCount: %v, UpsertedCount: %v", result.ModifiedCount, result.UpsertedCount))

	return nil
}

func (db *MongoDB) UpdateAllHabitsHandler(userId string, value interface{}) error {
	db.logger.InfoLog(helper.GetFunctionName(), "")

	updateHabits, ok := value.([]data.Habit)

	if !ok {
		err := "value type is not data.Habit"
		db.logger.ErrorLog(helper.GetFunctionName(), err)
		return fmt.Errorf("%s - %s", helper.GetFunctionName(), err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	newHabitCollection := db.NewHabitsCollection()

	objectUserId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to update habits collection for userId=%s, err=%s", userId, err))
		return fmt.Errorf("%s - Failed to update habits collection for userId=%s, err=%s", helper.GetFunctionName(), userId, err)
	}

	models := []mongo.WriteModel{}
	for _, habit := range updateHabits {
		objectId, err := primitive.ObjectIDFromHex(habit.HabitID)
		if err != nil {
			db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Invalid habitId=%s for userId=%s, err=%s", habit.HabitID, userId, err))
			return fmt.Errorf("%s - Invalid habitId=%s for userId=%s, err=%s", helper.GetFunctionName(), habit.HabitID, userId, err)
		}

		update := mongo.NewUpdateOneModel().
			SetFilter(bson.M{"_id": bson.ObjectID(objectId), "UserID": bson.ObjectID(objectUserId)}).
			SetUpdate(bson.M{
				"$set": bson.M{
					"Name":            habit.Name,
					"DaysTarget":      habit.DaysTarget,
					"CompletionDates": habit.CompletionDates,
				},
			})
		models = append(models, update)
	}

	result, err := newHabitCollection.BulkWrite(ctx, models)
	if err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to update habits collection for userId=%s, err=%s", userId, err))
		return fmt.Errorf("%s - Failed to update habits collection for userId=%s, err=%s", helper.GetFunctionName(), userId, err)
	}

	db.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("The document has been updated. ModifiedCount: %v, UpsertedCount: %v", result.ModifiedCount, result.UpsertedCount))

	return nil
}

func (db *MongoDB) DeleteHabitsHandler(userId, habitId string) error {
	db.logger.InfoLog(helper.GetFunctionName(), "")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newHabitCollection := db.NewHabitsCollection()

	objectID, err := primitive.ObjectIDFromHex(habitId)
	if err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to delete habit: err=%v", err))
		return fmt.Errorf("%s - Failed to delete habit: err=%v", helper.GetFunctionName(), err)
	}

	filter := bson.M{"_id": bson.ObjectID(objectID)}

	result, err := newHabitCollection.DeleteOne(ctx, filter)

	if err != nil {
		db.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Failed to delete habit for userId=%s: %v", userId, err))
		return fmt.Errorf("%s - Failed to delete habit for userId=%s: %v", helper.GetFunctionName(), userId, err)
	}

	db.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("The document has been updated. Acknowledged: %v, DeleteCount: %v", result.Acknowledged, result.DeletedCount))

	return nil
}
