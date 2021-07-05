package interactors

type Error struct {
	Code    int
	Message string
}

var (
	ErrInvalidInput = Error{
		Code:    400,
		Message: "Invalid Input",
	}
)

func InternalErrorCustom(msg string) Error {
	return Error{
		Code:    500,
		Message: msg,
	}
}
