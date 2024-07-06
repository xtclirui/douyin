package models

type BaseResponse struct {
	StatusCode int32  `json:"statusCode"`
	StatusMsg  string `json:"statueMsg,omitempty"`
}
