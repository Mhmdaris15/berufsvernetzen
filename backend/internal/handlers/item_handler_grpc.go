package handlers

// type ItemServiceServerImpl struct {
// 	berufsvernetzen.UnimplementedItemServiceServer
// }

// func (s *ItemServiceServerImpl) CreateItem(ctx context.Context, req *berufsvernetzen.CreateItemRequest) (*berufsvernetzen.Item, error) {
// 	item := &berufsvernetzen.Item{
// 		Id:          "unique_id", // Generate a unique ID here
// 		Name:        req.Name,
// 		Description: req.Description,
// 	}

// 	if err := repositories.CreateItem(item); err != nil {
// 		return nil, err
// 	}

// 	return item, nil
// }

// func (s *ItemServiceServerImpl) GetItem(ctx context.Context, req *berufsvernetzen.GetItemRequest) (*berufsvernetzen.Item, error) {
// 	item, err := repositories.GetItem(req.Id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return item, nil
// }

// func (s *ItemServiceServerImpl) UpdateItem(ctx context.Context, req *berufsvernetzen.UpdateItemRequest) (*berufsvernetzen.Item, error) {
// 	item := &berufsvernetzen.Item{
// 		Id:          req.Id,
// 		Name:        req.Name,
// 		Description: req.Description,
// 	}

// 	if err := repositories.UpdateItem(item); err != nil {
// 		return nil, err
// 	}

// 	return item, nil
// }

// func (s *ItemServiceServerImpl) DeleteItem(ctx context.Context, req *berufsvernetzen.DeleteItemRequest) (*berufsvernetzen.DeleteResponse, error) {
// 	if err := repositories.DeleteItem(req.Id); err != nil {
// 		return nil, err
// 	}

// 	return &berufsvernetzen.DeleteResponse{Success: true}, nil
// }
