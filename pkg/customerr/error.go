package customerr

type CustomErr struct {
	Err error `json:"error"`
}

func NewCustomError(err error) *CustomErr {
	return &CustomErr{
		Err: err,
	}
}

func (c CustomErr) Error() string {
	return c.Err.Error()
}
