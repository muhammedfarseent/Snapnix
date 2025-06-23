package responce

type Response struct {
	Status_code int    `json:"status_code"`
	Messege     string `json:"messege"`
	Data        any    `json:"data"`
	Error       any    `json:"error"`
}

func ClientResponse(status_code int, messege string, data any, err any) Response {
	return Response{
		Status_code: status_code,
		Messege:     messege,
		Data:        data,
		Error:       err,
	}
}
