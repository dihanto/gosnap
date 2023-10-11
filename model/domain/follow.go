package domain

type Follow struct {
	Id               int
	FollowerCount    int
	TargetUsername   string
	FollowerUsername string
}
