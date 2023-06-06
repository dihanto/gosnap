package domain

type Comment struct {
	Id        int
	UserId    int
	PhotoId   int
	Message   string
	CreatedAt int8
	UpdatedAt int8
	DeletedAt int8
}
