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
	ErrAccountNotFound = InteractorsErr{
		Message: "Akun tidak ditemukan",
	}
	ErrInsufficientBalance = InteractorsErr{
		Message: "Saldo tidak cukup ketika ditransfer",
	}
	ErrAmoutShouldNotBeNegative = InteractorsErr{
		Message: "Saldo tidak boleh bernilai negatif",
	}
	ErrSendToUserItself = InteractorsErr{
		Message: "Tidak diperbolehkan transfer ke akun pengirim sendiri",
	}
)

func InternalErrorCustom(msg string) InteractorsErr {
	return InteractorsErr{
		Message: msg,
	}
}
