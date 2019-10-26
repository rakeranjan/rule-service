package serializer

type JSONResponse struct {
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Error *APIError `json:"error"`
}

type APIError struct {
	Status interface{} `json:"status"`
	Title  interface{} `json:"title"`
}
