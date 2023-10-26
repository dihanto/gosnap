package response

type Follow struct {
	FollowerCount int `json:"followerCount"`
}

type GetFollower struct {
	Username      []string `json:"username"`
	FollowerCount int      `json:"followerCount"`
}
type GetFollowing struct {
	Username string `json:"username"`
}
