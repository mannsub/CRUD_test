package types

type APIResponse struct {
	Result      int64       `json:"result"`
	Description string      `json:"description"`
	ErrorCode   interface{} `json:"errCode"`
}

func NewAPIResponse(description string, result int64, errCode interface{}) *APIResponse {
	return &APIResponse{
		Result:      result,
		Description: description,
		ErrorCode:   errCode,
	}
}
