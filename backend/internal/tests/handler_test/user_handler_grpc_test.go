package handlertest

// func TestCreateUser(t *testing.T) {
// 	s := &handlers.UserServiceServerImpl{}
// 	req := &berufsvernetzen.CreateUserRequest{
// 		User: &berufsvernetzen.User{
// 			Name:           "John Doe",
// 			Username:       "johndoe",
// 			Email:          "johndoe@example.com",
// 			WhatsappNumber: "1234567890",
// 			Password:       "password",
// 			NIK:            "1234567890",
// 			Address:        "123 Main St",
// 			YearGraduation: "2020",
// 			Birthday:       "01/01/1990",
// 			Major:          "Computer Science",
// 			Languages:      []string{"English", "German"},
// 			Experiences: []*berufsvernetzen.Experience{
// 				{
// 					Position:    "Software Engineer",
// 					Description: "Internship at Google",
// 					Year:        "2020 - 2022",
// 				},
// 				{
// 					Position:    "Computer Science Student",
// 					Description: "Freelance web developer",
// 					Year:        "2020 - 2022",
// 				},
// 			},
// 			SocialMedia: &berufsvernetzen.SocialMedia{
// 				Id:        primitive.NewObjectID().String(),
// 				LinkedIn:  "https://www.linkedin.com/in/johndoe/",
// 				Facebook:  "https://www.facebook.com/johndoe/",
// 				Instagram: "https://www.instagram.com/johndoe/",
// 				Github:    "https://www.github.com/johndoe/",
// 				Twitter:   "https://www.twitter.com/johndoe/",
// 				Youtube:   "https://www.youtube.com/johndoe/",
// 			},
// 			Role: "Software Engineer",
// 			Certifications: []*berufsvernetzen.Certification{
// 				{
// 					Name:        "AWS Certified Developer - Associate",
// 					Institution: "Amazon Web Services",
// 					StartDate:   "2020-01-01",
// 					ExpiredDate: "2022-01-01",
// 				},
// 				{
// 					Name:        "AWS Certified Solutions Architect - Associate",
// 					Institution: "Amazon Web Services",
// 					StartDate:   "2020-01-01",
// 					ExpiredDate: "2022-01-01",
// 				},
// 			},
// 			Photo: []byte("photo"),
// 		},
// 	}
// 	expected := &berufsvernetzen.CreateUserResponse{
// 		User: &berufsvernetzen.User{
// 			Name:           "John Doe",
// 			Username:       "johndoe",
// 			Email:          "johndoe@example.com",
// 			WhatsappNumber: "1234567890",
// 			Password:       "password",
// 			NIK:            "1234567890",
// 			Address:        "123 Main St",
// 			YearGraduation: "2020",
// 			Birthday:       "01/01/1990",
// 			Major:          "Computer Science",
// 			Languages:      []string{"English", "German"},
// 			Experiences: []*berufsvernetzen.Experience{
// 				{
// 					Position:    "Software Engineer",
// 					Description: "Internship at Google",
// 					Year:        "2020 - 2022",
// 				},
// 				{
// 					Position:    "Computer Science Student",
// 					Description: "Freelance web developer",
// 					Year:        "2020 - 2022",
// 				},
// 			},
// 			SocialMedia: &berufsvernetzen.SocialMedia{
// 				Id:        primitive.NewObjectID().String(),
// 				LinkedIn:  "https://www.linkedin.com/in/johndoe/",
// 				Facebook:  "https://www.facebook.com/johndoe/",
// 				Instagram: "https://www.instagram.com/johndoe/",
// 				Github:    "https://www.github.com/johndoe/",
// 				Twitter:   "https://www.twitter.com/johndoe/",
// 				Youtube:   "https://www.youtube.com/johndoe/",
// 			},
// 			Role: "Software Engineer",
// 			Certifications: []*berufsvernetzen.Certification{
// 				{
// 					Name:        "AWS Certified Developer - Associate",
// 					Institution: "Amazon Web Services",
// 					StartDate:   "2020-01-01",
// 					ExpiredDate: "2022-01-01",
// 				},
// 				{
// 					Name:        "AWS Certified Solutions Architect - Associate",
// 					Institution: "Amazon Web Services",
// 					StartDate:   "2020-01-01",
// 					ExpiredDate: "2022-01-01",
// 				},
// 			},
// 			Photo: []byte("photo"),
// 		},
// 	}

// 	res, err := s.CreateUser(context.Background(), req)

// 	assert.NoError(t, err)
// 	assert.Equal(t, expected, res)
// }
