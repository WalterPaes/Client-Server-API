package customerr

import "encoding/json"

type CustomErr struct {
	Err error `json:"error"`
}

func NewCustomError(err error) *CustomErr {
	return &CustomErr{
		Err: err,
	}
}

func (c CustomErr) Error() string {
	str, _ := json.Marshal(c)
	return string(str)
}
