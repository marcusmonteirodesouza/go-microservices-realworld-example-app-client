package common

type ErrorResponse struct {
	Errors *errorResponseErrors `json:"errors"`
}

type errorResponseErrors struct {
	Body []string `json:"body"`
}
