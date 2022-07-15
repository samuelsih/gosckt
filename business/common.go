package business

type CommonResponse struct {
	ErrorCode int
	ErrorMsg  string
}

type CommonRequest struct {
	AuthToken string
}
