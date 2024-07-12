package repositories

// func CreateUser(user *berufsvernetzen.User) (*berufsvernetzen.User, error) {

// 	user.Id = primitive.NewObjectID().Hex()

// 	// Convert to models.User

// 	newUser, err := mongodb.UserCollection.InsertOne(context.Background(), user)
// 	if err != nil {
// 		return nil, err
// 	}
// 	log.Printf("Inserted ID: %v", newUser.InsertedID)
// 	// user.Id = newUser.InsertedID.(primitive.ObjectID).Hex()

// 	return user, err
// }

// func GetUsers() ([]*berufsvernetzen.User, error) {
// 	var users []*berufsvernetzen.User
// 	// var users2 []*models.User
// 	cursor, err := mongodb.UserCollection.Find(context.Background(), bson.M{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err = cursor.All(context.Background(), &users); err != nil {
// 		return nil, err
// 	}

// 	log.Printf("User %v", users)

// 	return users, nil
// }

// func GetUser(id string) (*berufsvernetzen.User, error) {
// 	filter := bson.M{"id": id}

// 	var user *berufsvernetzen.User
// 	var userModel models.User

// 	userDoc := mongodb.UserCollection.FindOne(context.Background(), filter)

// 	err := userDoc.Decode(&user)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = userDoc.Decode(&userModel)
// 	if err != nil {
// 		return nil, err
// 	}

// 	user.Id = userModel.ID.Hex()

// 	return user, err
// }

// func UpdateUser(user *berufsvernetzen.User) error {
// 	filter := bson.M{"id": user.Id}
// 	update := bson.M{"$set": bson.M{
// 		"name":            user.Name,
// 		"username":        user.Username,
// 		"email":           user.Email,
// 		"whatsapp_number": user.WhatsappNumber,
// 		"password":        user.Password,
// 		"nik":             user.NIK,
// 		"address":         user.Address,
// 		"year_graduation": user.YearGraduation,
// 		"birthdate":       user.Birthday,
// 		"major":           user.Major,
// 		"languages":       user.Languages,
// 		"experiences":     user.Experiences,
// 		"social_media":    user.SocialMedia,
// 		"role":            user.Role,
// 		"certifications":  user.Certifications,
// 		"photo":           user.Photo,
// 	}}
// 	_, err := mongodb.UserCollection.UpdateOne(context.Background(), filter, update)
// 	return err
// }

// func DeleteUser(id string) error {
// 	filter := bson.M{"id": id}

// 	_, err := mongodb.UserCollection.DeleteOne(context.Background(), filter)
// 	return err
// }
