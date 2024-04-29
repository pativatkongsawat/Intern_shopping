package helper

type SuccessResponse struct {
	Code    int    `json:"code" example:"2000"`
	Message string `json:"message" example:"OK"`
}

type UnAuthorizeResponse struct {
	Code    int    `json:"code" example:"4010"`
	Message string `json:"message" example:"UNAUTHORIZED"`
}

type NotFoundResponse struct {
	Code    int    `json:"code" example:"4040"`
	Message string `json:"message" example:"NOT FOUND"`
}

type InternalServerErrorResponse struct {
	Code    int    `json:"code" example:"5000"`
	Message string `json:"message" example:"INTERNAL SERVER ERROR"`
}
