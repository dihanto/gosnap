package domain

type Comment struct {
	Id        int
	UserId    int
	PhotoId   int
	Message   string
	CreatedAt int32
	UpdatedAt int32
	DeletedAt int32
}
