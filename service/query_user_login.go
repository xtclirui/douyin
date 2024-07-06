package service

type QLoginResponse struct {
	Userid int64  `json:"userid"`
	Token  string `json:"token"`
}

type Service_QLoginResponse struct {
	username string
	password string
	data     *QLoginResponse
	userid   int64
	token    string
}

func New_Service_QLoginResponse(username string, password string) *Service_QLoginResponse {
	return &Service_QLoginResponse{username: username, password: password}
}

func (sq *Service_QLoginResponse) Do() (*QLoginResponse, error) {
	return nil, nil
}

func QueryUserLogin(username string, password string) (*QLoginResponse, error) {
	s_qLoginResponse := New_Service_QLoginResponse(username, password)
	return s_qLoginResponse.Do()
}
