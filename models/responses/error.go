package responses

type ErrorWrapper struct {
	Error     error  `json:"error"`
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
} //@name ErrorWrapper
