package entities

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "ijasmoopan"
// 	dbname   = "users"
// )

type User struct {
	UserID   int    `json:"userid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status bool `json:"status"`
}
type Authentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Status bool `json:"status"`
}
type Token struct {
	Username    string
	TokenString string
}
type Admin struct {
	Username string
	Password string
}
type UserList struct {
	UserID int `json:"userid"`
	Username string `json:"username"`
	Status bool `json:"status"`
}
type UserEdit struct{
	UserID   int    `json:"userid"`
	Username string `json:"username"`
	Password string `json:"password"`
}
