package interactors

type Error interface {
	error
}

type InteractorsErr struct {
	Message string
}

var _ Error = InteractorsErr{}

func (s InteractorsErr) Error() string {
	return s.Message
}

var (
	ErrInvalidInput = InteractorsErr{
		Message: "Invalid input",
	}
	ErrRequiredFieldEmpty = InteractorsErr{
		Message: "Required field is empty",
	}
	ErrUnauthorized = InteractorsErr{
		Message: "Unauthorized",
	}
	ErrForbiddenAccess = InteractorsErr{
		Message: "Forbidden access",
	}
	ErrDataNotFound = InteractorsErr{
		Message: "Data not found",
	}
	ErrDuplicateDataAdd = InteractorsErr{
		Message: "Can not add data with duplicate identifier",
	}
)

func InternalErrorCustom(msg string) InteractorsErr {
	return InteractorsErr{
		Message: msg,
	}
}
