package models

type GetResponse struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"Get Success"`
}

type PostResponse struct {
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Post Success"`
}

type ErrorBadRequestResponse struct {
	Code  int    `json:"code" example:"400"`
	Error string `json:"error" example:"Bad Request"`
}

type ErrorUnauthorizedResponse struct {
	Code  int    `json:"code" example:"401"`
	Error string `json:"error" example:"Unauthorized"`
}

type ErrorForbiddenResponse struct {
	Code  int    `json:"code" example:"403"`
	Error string `json:"error" example:"Forbidden"`
}

type ErrorNotFoundResponse struct {
	Code  int    `json:"code" example:"404"`
	Error string `json:"error" example:"Not Found"`
}

type ErrorInternalServerErrorResponse struct {
	Code  int    `json:"code" example:"500"`
	Error string `json:"error" example:"Internal Server Error"`
}

type LoginSuccess struct {
	PostResponse
	Data struct {
		AccessToken string `json:"access_token" example:"access_token"`
		ExpiresIn   string `json:"expires_in" example:"1h"`
	}
}

type RegisterSuccess struct {
	PostResponse
	Data struct {
		ID uint `json:"id" example:"1"`
	}
}

type UpdateAccessTokenSuccess struct {
	LoginSuccess
}
