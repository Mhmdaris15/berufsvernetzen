package repositories

// var (
// 	ItemCollection = mongodb.GetCollection(mongodb.DB, "Items")
// )

// func CreateItem(item *berufsvernetzen.Item) error {
// 	_, err := ItemCollection.InsertOne(context.Background(), item)
// 	return err
// }

// func GetItem(id string) (*berufsvernetzen.Item, error) {
// 	filter := bson.M{"id": id}
// 	var item *berufsvernetzen.Item
// 	err := ItemCollection.FindOne(context.Background(), filter).Decode(&item)
// 	return item, err
// }

// func UpdateItem(item *berufsvernetzen.Item) error {
// 	filter := bson.M{"id": item.Id}
// 	update := bson.M{"$set": bson.M{"name": item.Name, "description": item.Description}}
// 	_, err := ItemCollection.UpdateOne(context.Background(), filter, update)
// 	return err
// }

// func DeleteItem(id string) error {
// 	filter := bson.M{"id": id}
// 	_, err := ItemCollection.DeleteOne(context.Background(), filter)
// 	return err
// }
