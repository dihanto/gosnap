package request

type Follow struct {
	FollowerUsername string `json:"FollowerUsername" validate:"required"`
	TargetUsername   string `json:"TargetUsername" validate:"required"`
}
