package services

type UserSearchRequest struct {
	Id     int64
	Status string
	Email  string
}

type UserListSearchRequest struct {
	Id         int64
	Status     string
	Email      string
	PageSize   int
	PageNumber int
}
