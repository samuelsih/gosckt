package business

import "log"

type CommonRequest struct {
	Cookie string `json:"-"`
}

type CommonResponse struct {
	StatusCode int    `json:"code"`
	StatusMsg  string `json:"message"`
	SetCookie  string `json:"-"`
}

func (c *CommonResponse) SetError(code int, msg string) {
	c.StatusCode = code
	c.StatusMsg = msg

	if code >= 500 {
		log.Println("ERROR INTERNAL:", msg)
	}
}
