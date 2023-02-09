package structure

type RetweetRequest struct {
	TweetID int `query:"id"`
}

type RetweetResponse struct{}
