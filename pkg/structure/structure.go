package structure

type RetweetRequest struct {
	TweetID int64 `form:"id" binding:"required" validate:"required"`
}

type RetweetResponse struct {
	IsSuccess bool `json:"isSuccess"`
}

type DefaultErrorResponse struct {
	Code  int         `json:"code"`
	Error interface{} `json:"error"`
}
