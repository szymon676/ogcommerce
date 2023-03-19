package types

type User struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type ReqUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
