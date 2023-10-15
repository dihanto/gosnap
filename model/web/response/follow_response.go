package response

type Follow struct {
	FollowerCount int `json:"followerCount"`
}

type GetFollower struct {
	Username string `json:"username"`
}
type GetFollowing struct {
	Username string `json:"username"`
}
