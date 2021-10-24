package users

type UserFilter struct {
	Id     int64
	Status string
	Email  string
	Password string
}


type UserListFilter struct {
	Filter UserFilter
	PageSize int
	PageNumber int
}