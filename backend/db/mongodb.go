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

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Send a ping to confirm a successful connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.Connect - %s", err))
		return fmt.Errorf(fmt.Sprintf("mongodb.Connect - %s", err))
	}

	db.client = client
	db.logger.InfoLog("Pinged your deployment. You successfully connected to MongoDB!")

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

	registerTime := time.Now()

	type registerUserData struct {
		Password     string    `bson:"Password"`
		FirstName    string    `bson:"FirstName"`
		LastName     string    `bson:"LastName"`
		EmailAddress string    `bson:"EmailAddress"`
		CreatedAt    time.Time `bson:"CreatedAt"`
		LastLogin    time.Time `bson:"LastLogin"`
		IsLoggedIn   bool      `bson:"IsLoggedIn"`
	}

	registerUser := registerUserData{
		EmailAddress: newUser.EmailAddress,
		Password:     newUser.Password,
		FirstName:    newUser.FirstName,
		LastName:     newUser.LastName,
		CreatedAt:    registerTime,
		LastLogin:    registerTime,
		IsLoggedIn:   false,
	}

	insertResult, err := newUsersCollection.InsertOne(ctx, registerUser)
	if err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.RegisterUserHandler - Failed to insert new user: %v", err))
		return nil, fmt.Errorf("mongodb.RegisterUserHandler - Failed to insert new user: %v", err)
	}

	db.logger.InfoLog(fmt.Sprintf("mongodb.RegisterUserHandler - User registered successfully with EmailAddress: %s, Acknowledged: %v, InsertedID: %v", registerUser.EmailAddress, insertResult.Acknowledged, insertResult.InsertedID))

	return &data.UserData{
		FirstName:    newUser.FirstName,
		LastName:     newUser.LastName,
		Password:     newUser.Password,
		EmailAddress: newUser.EmailAddress,
		CreatedAt:    registerTime,
	}, nil
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

	objectID, err := primitive.ObjectIDFromHex(userSession.UserID)
	if err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.LoginUser - Failed to insert user session for userId=%s", userSession.UserID))
		return fmt.Errorf("mongodb.LoginUser - Failed to insert user session for userId=%s: %v", userSession.UserID, err)
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
		db.logger.ErrorLog(fmt.Sprintf("mongodb.LoginUser - Failed to insert user session for userId=%s", userSession.UserID))
		return fmt.Errorf("mongodb.LoginUser - Failed to insert user session for userId=%s: %v", userSession.UserID, err)
	}

	filter := bson.M{"_id": bson.ObjectID(objectID)}

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

	objectID, err := primitive.ObjectIDFromHex(userLoggedOut.UserID)
	if err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.Logout - Failed to retrieve the objectId from userId=%s", userLoggedOut.UserID))
		return fmt.Errorf("mongodb.Logout -  Failed to retrieve the objectId from userId=%s: %v", userLoggedOut.UserID, err)
	}

	sessionFilter := bson.D{{Key: "_id", Value: bson.ObjectID(objectID)}}

	deleteResult, err := newUsersSessionCollection.DeleteOne(ctx, sessionFilter)
	if err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.LogoutUser - Failed to delete user session for userId=%s", userLoggedOut.UserID))
		return fmt.Errorf("mongodb.LogoutUser - Failed to delete user session for userId=%s: %v", userLoggedOut.UserID, err)
	}

	db.logger.InfoLog(fmt.Sprintf("mongodb.LogoutUser - The document has been updated. Acknowledged: %v, DeleteCount: %v", deleteResult.Acknowledged, deleteResult.DeletedCount))

	updateFilter := bson.M{"_id": bson.ObjectID(objectID)}
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

	db.logger.InfoLog(fmt.Sprintf("mongodb.LogoutUser - The document has been updated. ModifiedCount: %v, UpsertedCount: %v", updateResult.ModifiedCount, updateResult.UpsertedCount))

	return nil
}

func (db *MongoDB) GetUserDetails(value interface{}) (interface{}, error) {
	db.logger.InfoLog("mongodb.GetUserDetails")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newUsersCollection := db.NewUsersCollection()

	if userAuth, ok := value.(*data.RegisterUserRequest); ok {
		filter := bson.M{"EmailAddress": userAuth.EmailAddress}
		var result bson.M

		err := newUsersCollection.FindOne(ctx, filter).Decode(&result)

		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, nil
			}

			db.logger.ErrorLog(fmt.Sprintf("mongodb.GetUserDetails - Failed to get user details err=%s", err))
			return nil, fmt.Errorf("mongodb.GetUserDetails - Failed to get user details err=%s", err)
		}

		db.logger.ErrorLog(fmt.Sprintf("mongodb.GetUserDetails - Failed to get user details err=%s", err))
		return nil, fmt.Errorf("mongodb.GetUserDetails - Failed to get user details err=%s", err)
	}

	if userAuth, ok := value.(*data.UserAuth); ok {
		filter := bson.M{"EmailAddress": userAuth.EmailAddress}

		user, err := db.findUser(ctx, filter, newUsersCollection)

		if err != nil {
			db.logger.ErrorLog(fmt.Sprintf("mongodb.GetUserDetails - Failed to get user details err=%s", err))
			return nil, fmt.Errorf("mongodb.GetUserDetails - Failed to get user details err=%s", err)
		}

		return user, nil
	}

	if userLoggedOutRequest, ok := value.(*data.UserLoggedOutRequest); ok {
		filter := bson.M{"EmailAddress": userLoggedOutRequest.EmailAddress}

		user, err := db.findUser(ctx, filter, newUsersCollection)

		if err != nil {
			db.logger.ErrorLog(fmt.Sprintf("mongodb.GetUserDetails - Failed to get user details err=%s", err))
			return nil, fmt.Errorf("mongodb.GetUserDetails - Failed to get user details err=%s", err)
		}

		return user, nil
	}

	db.logger.ErrorLog("mongodb.GetUserDetails - value type is unsupported")
	return nil, fmt.Errorf("mongodb.GetUserDetails - value type is unsupported")
}

func (db *MongoDB) findUser(ctx context.Context, filter bson.M, newUsersCollection *mongo.Collection) (*data.UserData, error) {
	var user data.UserData
	var result bson.M

	err := newUsersCollection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("mongodb.findUser - User doesn't exist")
		}

		db.logger.ErrorLog(fmt.Sprintf("mongodb.findUser - Failed to get user details err=%s", err))
		return nil, fmt.Errorf("mongodb.findUser - Failed to get user details err=%s", err)
	}

	if id, ok := result["_id"].(bson.ObjectID); ok {
		user.UserID = id.Hex()
	} else {
		return nil, fmt.Errorf("mongodb.findUser - _id field is missing or not an ObjectID")
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

	if isLoggedIn, ok := result["IsLoggedIn"].(bool); ok {
		user.IsLoggedIn = isLoggedIn
	}

	return &user, nil
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

	objectID, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.CreateHabitsHandler - Failed to insert new habit: userId=%s, err=%v", userId, err))
		return fmt.Errorf("mongodb.CreateHabitsHandler - Failed to insert new habit: userId=%s, err=%v", userId, err)
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
		Days:            newHabit.Days,
		DaysTarget:      newHabit.DaysTarget,
		CompletionDates: []string{},
	}

	insertResult, err := newHabitsCollection.InsertOne(ctx, insertHabit)
	if err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.CreateHabitsHandler - Failed to insert new habit: userId=%s, err=%v", userId, err))
		return fmt.Errorf("mongodb.CreateHabitsHandler - Failed to insert new habit: userId=%s, err=%v", userId, err)
	}

	db.logger.InfoLog(fmt.Sprintf("mongodb.CreateHabitsHandler - habit inserted successfully with userId: %s, Acknowledged: %v, InsertedID: %v", userId, insertResult.Acknowledged, insertResult.InsertedID))

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
	var results []data.Habit

	for cur.Next(ctx) {
		var el bson.M
		var habit data.Habit
		err := cur.Decode(&el)

		if err != nil {
			db.logger.ErrorLog(fmt.Sprintf("mongodb.RetrieveAllHabitsHandler - Failed to retrieve habit: err=%v", err))
			return nil, fmt.Errorf("mongodb.RetrieveAllHabitsHandler - Failed to retrieve habit: err=%v", err)
		}

		if id, ok := el["_id"].(bson.ObjectID); ok {
			habit.HabitID = id.Hex()
		} else {
			return nil, fmt.Errorf("mongodb.RetrieveAllHabitsHandler - _id field is missing or not an ObjectID")
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
					db.logger.ErrorLog("mongodb.RetrieveAllHabitsHandler - non-string value in Completed Date")
				}
			}

			habit.CompletionDates = completionDate
		}

		results = append(results, habit)
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

	objectID, err := primitive.ObjectIDFromHex(habitId)
	if err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.RetrieveHabitsHandler - Failed to retrieve habit: userId=%s, err=%v", userId, err))
		return nil, fmt.Errorf("mongodb.RetrieveHabitsHandler - Failed to retrieve habit: userId=%s, err=%v", userId, err)
	}

	filter := bson.M{"_id": bson.ObjectID(objectID)}
	var result bson.M
	var habit data.Habit
	err = newHabitCollection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("mongodb.RetrieveHabitsHandler - Habit doesn't exist")
		}

		db.logger.ErrorLog(fmt.Sprintf("mongodb.RetrieveHabitsHandler - Failed to retrieve habit: userId=%s, err=%v", userId, err))
		return nil, fmt.Errorf("mongodb.RetrieveHabitsHandler - Failed to retrieve habit: userId=%s, err=%v", userId, err)
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
				db.logger.ErrorLog("mongodb.RetrieveAllHabitsHandler - non-string value in Completed Date")
			}
		}

		habit.CompletionDates = completionDate
	}

	return habit, nil
}

func (db *MongoDB) UpdateHabitsHandler(userId, habitId string, value interface{}) error {
	db.logger.InfoLog("mongodb.UpdateHabitsHandler")

	updateHabit, ok := value.(data.Habit)

	if !ok {
		err := "mongodb.UpdateHabitsHandler - value type is not data.Habit"
		db.logger.ErrorLog(err)
		return fmt.Errorf(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newHabitCollection := db.NewHabitsCollection()

	objectId, err := primitive.ObjectIDFromHex(habitId)

	if err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.UpdateHabitsHandler - Failed to update habits collection for userId=%s, habitId=%s, err=%s", userId, habitId, err))
		return fmt.Errorf("mongodb.UpdateHabitsHandler - Failed to update habits collection for userId=%s, habitId=%s, err=%s", userId, habitId, err)
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

	filter := bson.M{"_id": bson.ObjectID(objectID)}

	result, err := newHabitCollection.DeleteOne(ctx, filter)

	if err != nil {
		db.logger.ErrorLog(fmt.Sprintf("mongodb.DeleteHabitsHandler - Failed to delete habit for userId=%s: %v", userId, err))
		return fmt.Errorf("mongodb.DeleteHabitsHandler - Failed to delete habit for userId=%s: %v", userId, err)
	}

	db.logger.InfoLog(fmt.Sprintf("The document has been updated. Acknowledged: %v, DeleteCount: %v", result.Acknowledged, result.DeletedCount))

	return nil
}
