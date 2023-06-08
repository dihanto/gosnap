package domain

type Photo struct {
	Id        int
	Title     string
	Caption   string
	PhotoUrl  string
	UserId    int
	CreatedAt int32
	UpdatedAt int32
	DeletedAt int32
}
