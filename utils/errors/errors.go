package errors

//인터페이스를 구현한 구조체를 변경해 유연하게 대응 가능
type UtilErr interface {
	GetMsg() string
	GetStatusCode() string
	GetDetail() string
}

type restErr struct {
	errMsg        string `json:"message"`     //사용자에게 출력되는 메세지
	errStatusCode string `json:"status_code"` //http 상태 코드
	errDetail     string `json:"detail"`      //세부 에러 내용
}

func (r restErr) GetMsg() string {
	return r.errMsg
}

func (r restErr) GetStatusCode() string {
	return r.errStatusCode
}

func (r restErr) GetDetail() string {
	return r.errDetail
}

func NewRestError(msg string, statusCode string, err string) UtilErr {
	return &restErr{
		errMsg:        msg,
		errStatusCode: statusCode,
		errDetail:     err,
	}
}
