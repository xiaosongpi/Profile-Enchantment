package pkg

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type FailedResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
