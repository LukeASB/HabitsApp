package db

import (
	"context"
	"dohabits/data"
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

// type ICollection interface {
// 	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
// 	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
// 	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cur *mongo.Cursor, err error)
// 	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
// 	DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
// }

type UsersCollection struct {
}

type UserSessionCollection struct {
}

type HabitsCollection struct {
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

func (db *MongoDB) NewUsersCollection() *mongo.Collection {
	return db.client.Database(db.habitsAppDBName).Collection(db.usersCollection)
}

func (db *MongoDB) NewUsersSessionCollection() *mongo.Collection {
	return db.client.Database(db.habitsAppDBName).Collection(db.userSessionCollection)
}

func (db *MongoDB) NewHabitsCollection() *mongo.Collection {
	return db.client.Database(db.habitsAppDBName).Collection(db.habitsCollection)
}

func (db *MongoDB) Connect() error {

	connectionString := os.Getenv("DB_URL")

	if connectionString == "" {
		db.logger.ErrorLog("mongodb.Connect - connectionString is empty")
		return fmt.Errorf("mongodb.Connect - connectionString is empty")
	}

	db.logger.InfoLog("mongodb.Connect")
	db.logger.DebugLog(fmt.Sprintf("db - Connect() - %s\n", connectionString))

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(opts)
	if err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.Connect - %s", err))
		return fmt.Errorf(fmt.Sprintf("mongodb.Connect - %s", err))
	}

	db.logger.InfoLog("Pinged your deployment. You successfully connected to MongoDB!")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Send a ping to confirm a successful connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.Connect - %s", err))
		return fmt.Errorf(fmt.Sprintf("mongodb.Connect - %s", err))
	}

	db.client = client

	return nil
}

func (db *MongoDB) Disconnect() error {
	if err := db.client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}

	db.logger.InfoLog("mongodb.Disconnect - Successfully disconnected from MongoDB")
	return nil
}

func (db *MongoDB) RegisterUserHandler(value interface{}) (interface{}, error) {
	db.logger.InfoLog("mongodb.RegisterUserHandler")

	newUser, ok := value.(*data.RegisterUserRequest)

	if !ok {
		db.logger.ErrorLog("mongodb.RegisterUserHandler - value type is not data.UserData")
		return nil, fmt.Errorf("mongodb.RegisterUserHandler - value type is not data.UserData")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newUsersCollection := db.NewUsersCollection()

	filter := bson.M{"EmailAddress": newUser.EmailAddress}

	var existingUser data.UserData
	err := newUsersCollection.FindOne(ctx, filter).Decode(&existingUser)
	if err == nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.RegisterUserHandler - User already exists with EmailAddress: %s", newUser.EmailAddress))
		return nil, fmt.Errorf("mongodb.RegisterUserHandler - User already exists with EmailAddress: %s", newUser.EmailAddress)
	}

	if err != mongo.ErrNoDocuments {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.RegisterUserHandler - Error checking for existing user: %v", err))
		return nil, fmt.Errorf("mongodb.RegisterUserHandler - Error checking for existing user: %v", err)
	}

	registerUser := data.UserData{
		FirstName:    newUser.FirstName,
		LastName:     newUser.LastName,
		Password:     newUser.Password,
		EmailAddress: newUser.EmailAddress,
		CreatedAt:    time.Now(),
	}

	insertResult, err := newUsersCollection.InsertOne(ctx, registerUser)
	if err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.RegisterUserHandler - Failed to insert new user: %v", err))
		return nil, fmt.Errorf("mongodb.RegisterUserHandler - Failed to insert new user: %v", err)
	}

	db.logger.InfoLog(fmt.Sprintf("mongodb.RegisterUserHandler - User registered successfully with EmailAddress: %s, Acknowledged: %v, InsertedID: %v", registerUser.EmailAddress, insertResult.Acknowledged, insertResult.InsertedID))

	return nil, nil
}

func (db *MongoDB) LoginUser(value interface{}) error {
	db.logger.InfoLog("mongodb.LoginUser")

	userSession, ok := value.(*data.UserSession)
	if !ok {
		db.logger.ErrorLog("mongodb.LoginUser - value type is not data.UserSession")
		return fmt.Errorf("mongodb.LoginUser - value type is not data.UserSession")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newUsersSessionCollection := db.NewUsersSessionCollection()
	newUsersCollection := db.NewUsersCollection()

	_, err := newUsersSessionCollection.InsertOne(ctx, userSession)
	if err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.LoginUser - Failed to insert user session for userId=%s", userSession.UserID))
		return fmt.Errorf("mongodb.LoginUser - Failed to insert user session for userId=%s: %v", userSession.UserID, err)
	}

	filter := bson.M{"UserID": userSession.UserID}

	update := bson.M{
		"$set": bson.M{
			"IsLoggedIn": true,
			"LastLogin":  userSession.CreatedAt,
		},
	}

	result, err := newUsersCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.LoginUser - Failed to update users collection for userId=%s", userSession.UserID))
		return fmt.Errorf("mongodb.LoginUser - Failed to update users collection for userId=%s: %v", userSession.UserID, err)
	}

	db.logger.InfoLog(fmt.Sprintf("The document has been updated. ModifiedCount: %v, UpsertedCount: %v", result.ModifiedCount, result.UpsertedCount))

	return nil
}

func (db *MongoDB) LogoutUser(value interface{}) error {
	db.logger.InfoLog("mongodb.LogoutUser")

	userLoggedOut, ok := value.(*data.UserData)

	if !ok {
		db.logger.ErrorLog("mongodb.LogoutUser - value type is not data.UserLoggedOutRequest")
		return fmt.Errorf("mongodb.LogoutUser - value type is not data.UserLoggedOutRequest")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newUsersSessionCollection := db.NewUsersSessionCollection()
	newUsersCollection := db.NewUsersCollection()

	sessionFilter := bson.M{"UserID": userLoggedOut.UserID}

	deleteResult, err := newUsersSessionCollection.DeleteOne(ctx, sessionFilter)
	if err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.LogoutUser - Failed to delete user session for userId=%s", userLoggedOut.UserID))
		return fmt.Errorf("mongodb.LogoutUser - Failed to delete user session for userId=%s: %v", userLoggedOut.UserID, err)
	}

	db.logger.InfoLog(fmt.Sprintf("The document has been updated. Acknowledged: %v, DeleteCount: %v", deleteResult.Acknowledged, deleteResult.DeletedCount))

	updateFilter := bson.M{"UserID": userLoggedOut.UserID}
	update := bson.M{
		"$set": bson.M{
			"IsLoggedIn": false,
		},
	}

	updateResult, err := newUsersCollection.UpdateOne(ctx, updateFilter, update)
	if err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.LogoutUser - Failed to update users collection for userId=%s", userLoggedOut.UserID))
		return fmt.Errorf("mongodb.LogoutUser - Failed to update users collection for userId=%s: %v", userLoggedOut.UserID, err)
	}

	db.logger.InfoLog(fmt.Sprintf("The document has been updated. ModifiedCount: %v, UpsertedCount: %v", updateResult.ModifiedCount, updateResult.UpsertedCount))

	return nil
}

func (db *MongoDB) GetUserDetails(value interface{}) (interface{}, error) {
	db.logger.InfoLog("mongodb.GetUserDetails")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newUsersCollection := db.NewUsersCollection()

	if userAuth, ok := value.(*data.RegisterUserRequest); ok {
		filter := bson.M{"EmailAddress": userAuth.EmailAddress}

		var user data.UserData

		err := newUsersCollection.FindOne(ctx, filter).Decode(&user)

		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return data.UserData{}, nil
			}

			db.logger.ErrorLog(fmt.Sprintf("mongodb.GetUserDetails - Failed to get user details err=%s", err))
			return nil, fmt.Errorf("mongodb.GetUserDetails - Failed to get user details err=%s", err)
		}

		return &user, nil
	}

	if userAuth, ok := value.(*data.UserAuth); ok {
		filter := bson.M{"EmailAddress": userAuth.EmailAddress}

		var user data.UserData

		err := newUsersCollection.FindOne(ctx, filter).Decode(&user)

		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, fmt.Errorf("mock_db.GetUserData - User doesn't exist")
			}

			db.logger.ErrorLog(fmt.Sprintf("mongodb.GetUserDetails - Failed to get user details err=%s", err))
			return nil, fmt.Errorf("mongodb.GetUserDetails - Failed to get user details err=%s", err)
		}

		return &user, nil
	}

	if userLoggedOutRequest, ok := value.(*data.UserLoggedOutRequest); ok {
		filter := bson.M{"EmailAddress": userLoggedOutRequest.EmailAddress}

		var user data.UserData

		err := newUsersCollection.FindOne(ctx, filter).Decode(&user)

		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, fmt.Errorf("mock_db.GetUserData - User doesn't exist")
			}

			db.logger.ErrorLog(fmt.Sprintf("mongodb.GetUserDetails - Failed to get user details err=%s", err))
			return nil, fmt.Errorf("mongodb.GetUserDetails - Failed to get user details err=%s", err)
		}

		return &user, nil
	}

	db.logger.ErrorLog("mock_db.GetUserData - value type is unsupported")
	return nil, fmt.Errorf("mock_db.GetUserData - value type is unsupported")
}

func (db *MongoDB) CreateHabitsHandler(userId string, value interface{}) error {
	db.logger.InfoLog(fmt.Sprintf("mongodb.CreateHabitsHandler = userId=%s", userId))
	newHabit, ok := value.(data.NewHabit)

	if !ok {
		db.logger.ErrorLog("mongodb.CreateHabitsHandler - value type is not data.Habit")
		return fmt.Errorf("mongodb.CreateHabitsHandler - value type is not data.Habit")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newHabitsCollection := db.NewHabitsCollection()

	habit := data.Habit{
		UserID:          userId,
		CreatedAt:       time.Now(),
		Name:            newHabit.Name,
		Days:            newHabit.Days,
		DaysTarget:      newHabit.DaysTarget,
		CompletionDates: []string{},
	}

	insertResult, err := newHabitsCollection.InsertOne(ctx, habit)
	if err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.CreateHabitsHandler - Failed to insert new habit: userId=%s, err=%v", userId, err))
		return fmt.Errorf("mongodb.CreateHabitsHandler - Failed to insert new habit: userId=%s, err=%v", userId, err)
	}

	db.logger.InfoLog(fmt.Sprintf("mongodb.CreateHabitsHandler - User registered successfully with userId: %s, Acknowledged: %v, InsertedID: %v", userId, insertResult.Acknowledged, insertResult.InsertedID))

	return nil
}

func (db *MongoDB) RetrieveAllHabitsHandler(userId string) (interface{}, error) {
	db.logger.InfoLog(fmt.Sprintf("mongodb.RetrieveAllHabitsHandler - userId=%s", userId))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newHabitsCollection := db.NewHabitsCollection()

	cur, err := newHabitsCollection.Find(ctx, bson.D{})

	if err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.RetrieveAllHabitsHandler - Failed to retrieve habit: err=%v", err))
		return nil, fmt.Errorf("mongodb.RetrieveAllHabitsHandler - Failed to retrieve habit: err=%v", err)
	}

	var results []*data.Habit

	for cur.Next(ctx) {
		var el data.Habit
		err := cur.Decode(&el)

		if err != nil {
			db.logger.ErrorLog(fmt.Sprintf("mongodb.RetrieveAllHabitsHandler - Failed to retrieve habit: err=%v", err))
			return nil, fmt.Errorf("mongodb.RetrieveAllHabitsHandler - Failed to retrieve habit: err=%v", err)
		}

		objectIDString := el.HabitID

		objectID, err := primitive.ObjectIDFromHex(objectIDString)

		if err != nil {
			db.logger.ErrorLog(fmt.Sprintf("mongodb.RetrieveAllHabitsHandler - Failed to retrieve habit: err=%v", err))
			return nil, fmt.Errorf("mongodb.RetrieveAllHabitsHandler - Failed to retrieve habit: err=%v", err)
		}

		idStr := objectID.Hex()

		habit := data.Habit{
			HabitID:         idStr,
			UserID:          el.UserID,
			CreatedAt:       el.CreatedAt,
			Name:            el.Name,
			Days:            el.Days,
			DaysTarget:      el.DaysTarget,
			CompletionDates: el.CompletionDates,
		}

		results = append(results, &habit)
	}

	if err := cur.Err(); err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.RetrieveAllHabitsHandler - Failed to retrieve habit: err=%v", err))
		return nil, fmt.Errorf("mongodb.RetrieveAllHabitsHandler - Failed to retrieve habit: err=%v", err)
	}

	return results, nil
}

func (db *MongoDB) RetrieveHabitsHandler(userId, habitId string) (interface{}, error) {
	db.logger.InfoLog(fmt.Sprintf("mongodb.RetrieveHabitsHandler userId=%s, habitId=%s\n", userId, habitId))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newHabitCollection := db.NewHabitsCollection()

	filter := bson.M{"_id": habitId}
	var habit data.Habit
	err := newHabitCollection.FindOne(ctx, filter).Decode(&habit)

	if err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.RetrieveHabitsHandler - Failed to retrieve habit: userId=%s, err=%v", userId, err))
		return nil, fmt.Errorf("mongodb.RetrieveHabitsHandler - Failed to retrieve habit: userId=%s, err=%v", userId, err)
	}

	if err != mongo.ErrNoDocuments {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.RetrieveHabitsHandler - Error checking for existing user: err=%v", err))
		return nil, fmt.Errorf("mongodb.RetrieveHabitsHandler - Error checking for existing user: err=%v", err)
	}

	return habit, nil
}

func (db *MongoDB) UpdateHabitsHandler(userId, habitId string, value interface{}) error {
	db.logger.InfoLog("mongodb.UpdateHabitsHandler")

	newHabit, ok := value.(data.Habit)

	if !ok {
		err := "mongodb.UpdateHabitsHandler - value type is not data.Habit"
		db.logger.ErrorLog(err)
		return fmt.Errorf(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newHabitCollection := db.NewHabitsCollection()

	filter := bson.M{"_id": habitId}

	update := bson.M{
		"$set": bson.M{
			"Name":       newHabit.Name,
			"DaysTarget": newHabit.DaysTarget,
		},
	}

	result, err := newHabitCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.UpdateHabitsHandler - Failed to update habits collection for userId=%s, habitId=%s, err=%s", userId, habitId, err))
		return fmt.Errorf("mongodb.UpdateHabitsHandler - Failed to update habits collection for userId=%s, habitId=%s, err=%s", userId, habitId, err)
	}

	db.logger.InfoLog(fmt.Sprintf("The document has been updated. ModifiedCount: %v, UpsertedCount: %v", result.ModifiedCount, result.UpsertedCount))

	return nil
}

func (db *MongoDB) DeleteHabitsHandler(userId, habitId string) error {
	db.logger.InfoLog("mongodb.DeleteHabitsHandler")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newHabitCollection := db.NewHabitsCollection()

	objectID, err := primitive.ObjectIDFromHex(habitId)
	if err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.DeleteHabitsHandler - Failed to delete habit: err=%v", err))
		return fmt.Errorf("mongodb.DeleteHabitsHandler - Failed to delete habit: err=%v", err)
	}

	filter := bson.M{"_id": objectID}

	result, err := newHabitCollection.DeleteOne(ctx, filter)

	if err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.DeleteHabitsHandler - Failed to delete habit for userId=%s: %v", userId, err))
		return fmt.Errorf("mongodb.DeleteHabitsHandler - Failed to delete habit for userId=%s: %v", userId, err)
	}

	db.logger.InfoLog(fmt.Sprintf("The document has been updated. Acknowledged: %v, DeleteCount: %v", result.Acknowledged, result.DeletedCount))

	return nil
}
