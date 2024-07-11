package models

type User struct {
	Id             string `json:"-"`
	Name           string `json:"name"`
	LastName       string `json:"last_name"`
	UserName       string `json:"user_name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PhoneNumber    string `json:"phone_number"`
	BirthDate      string `json:"birth_date"`
	Biography      string `json:"biography"`
	Gender         string `json:"gender"`
	ProfilePicture string `json:"profile_picture"`
}

type UserCreate struct {
	Name           string `json:"name"`
	LastName       string `json:"last_name"`
	UserName       string `json:"user_name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PhoneNumber    string `json:"phone_number"`
	BirthDate      string `json:"birth_date"`
	Biography      string `json:"biography"`
	Gender         string `json:"gender"`
	ProfilePicture string `json:"profile_picture"`
}

type UserBYtokens struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	LastName       string `json:"last_name"`
	UserName       string `json:"user_name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PhoneNumber    string `json:"phone_number"`
	BirthDate      string `json:"birth_date"`
	Biography      string `json:"biography"`
	Gender         string `json:"gender"`
	ProfilePicture string `json:"profile_picture"`
	RefreshToken   string `json:"refresh_token"`
	AccessToken    string `json:"access_token"`
}

type UserByAccess struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	UserName       string `json:"user_name"`
	PhoneNumber    string `json:"phone_number"`
	BirthDate      string `json:"birth_date"`
	Biography      string `json:"biography"`
	Gender         string `json:"gender"`
	ProfilePicture string `json:"profile_picture"`
	AccessToken    string `json:"access_token"`
}

type RegisterUser struct {
	Name           string `json:"name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	UserName       string `json:"user_name"`
	PhoneNumber    string `json:"phone_number"`
	BirthDate      string `json:"birth_date"`
	Biography      string `json:"biography"`
	Gender         string `json:"gender"`
	ProfilePicture string `json:"profile_picture"`
}

type UserUpdate struct {
	Id             string `json:"-"`
	Name           string `json:"name"`
	LastName       string `json:"last_name"`
	PhoneNumber    string `json:"phone_number"`
	BirthDate      string `json:"birth_date"`
	Biography      string `json:"biography"`
	Gender         string `json:"gender"`
	ProfilePicture string `json:"profile_picture"`
}
