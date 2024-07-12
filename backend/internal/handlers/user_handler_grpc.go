package handlers

// type UserServiceServerImpl struct {
// 	berufsvernetzen.UnimplementedUserServiceServer
// }

// func (s *UserServiceServerImpl) GetUsers(ctx context.Context, req *berufsvernetzen.EmptyRequest) (*berufsvernetzen.GetAllUsersResponse, error) {
// 	users, err := repositories.GetUsers()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &berufsvernetzen.GetAllUsersResponse{Users: users}, nil
// }

// func (s *UserServiceServerImpl) GetUser(ctx context.Context, req *berufsvernetzen.GetUserRequest) (*berufsvernetzen.GetUserResponse, error) {
// 	user, err := repositories.GetUser(req.Id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &berufsvernetzen.GetUserResponse{User: user}, nil
// }

// func (s *UserServiceServerImpl) CreateUser(ctx context.Context, req *berufsvernetzen.CreateUserRequest) (*berufsvernetzen.CreateUserResponse, error) {
// 	// log.Printf("CreateUser request %v", req.String())
// 	// log.Printf("Username: %v", req.GetUser().GetUsername())
// 	// log.Printf("Email: %v", req.GetUser().GetEmail())
// 	// log.Printf("Password: %v", req.GetUser().GetPassword())
// 	log.Printf("User %v", req.GetUser())

// 	newUser := &berufsvernetzen.User{}

// 	newUser, err := repositories.CreateUser(req.GetUser())
// 	if err != nil {
// 		return nil, err
// 	}

// 	// log.Printf("Repositories %v", req.GetUser())
// 	log.Printf("New user %v", newUser)

// 	return &berufsvernetzen.CreateUserResponse{User: newUser}, nil
// }

// func (s *UserServiceServerImpl) UpdateUser(ctx context.Context, req *berufsvernetzen.UpdateUserRequest) (*berufsvernetzen.UpdateUserResponse, error) {
// 	user := &berufsvernetzen.User{
// 		Id:             req.User.Id,
// 		Name:           req.User.Name,
// 		Username:       req.User.Username,
// 		Email:          req.User.Email,
// 		WhatsappNumber: req.User.WhatsappNumber,
// 		Password:       req.User.Password,
// 		NIK:            req.User.NIK,
// 		Address:        req.User.Address,
// 		YearGraduation: req.User.YearGraduation,
// 		Birthday:       req.User.Birthday,
// 		Major:          req.User.Major,
// 		Languages:      req.User.Languages,
// 		Experiences:    req.User.Experiences,
// 		SocialMedia:    req.User.SocialMedia,
// 		Role:           req.User.Role,
// 		Certifications: req.User.Certifications,
// 		Photo:          req.User.Photo,
// 	}

// 	if err := repositories.UpdateUser(user); err != nil {
// 		return nil, err
// 	}

// 	return &berufsvernetzen.UpdateUserResponse{User: user}, nil
// }

// func (s *UserServiceServerImpl) DeleteUser(ctx context.Context, req *berufsvernetzen.DeleteUserRequest) (*berufsvernetzen.DeleteUserResponse, error) {
// 	if err := repositories.DeleteUser(req.Id); err != nil {
// 		return nil, err
// 	}

// 	return &berufsvernetzen.DeleteUserResponse{Success: true}, nil
// }
