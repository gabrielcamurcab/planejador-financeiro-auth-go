package user

type User struct {
	UUID     string `bson:"uuid"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}
