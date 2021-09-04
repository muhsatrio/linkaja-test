package account

import (
	"linkaja-test/interactors"
	"linkaja-test/platform"
	"linkaja-test/platform/mysql"
	"reflect"
	"testing"

	mock_account "linkaja-test/mocks/platform/account"

	"github.com/golang/mock/gomock"
)

func TestInteractors_Transfer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accountMock := mock_account.NewMockAccountAdapter(ctrl)

	type fields struct {
		accountRepo mysql.AccountAdapter
	}
	type args struct {
		req RequestTransfer
	}
	tests := []struct {
		name              string
		mockRepoFunc      func()
		fields            fields
		args              args
		wantInteractorErr interactors.Error
	}{
		{
			name: "amount should not negative",
			fields: fields{
				accountRepo: accountMock,
			},
			mockRepoFunc: func() {
				accountMock.EXPECT().UpdateBalance(gomock.Any(), gomock.Any(), gomock.Any()).Return(platform.ErrInvalidInput)
			},
			args: args{
				req: RequestTransfer{
					SenderAccountNumber:   555001,
					ReceiverAccountNumber: 555002,
					Amount:                -1000,
				},
			},
			wantInteractorErr: interactors.ErrAmoutShouldNotBeNegative,
		},
		{
			name: "account not found",
			fields: fields{
				accountRepo: accountMock,
			},
			mockRepoFunc: func() {
				accountMock.EXPECT().UpdateBalance(gomock.Any(), gomock.Any(), gomock.Any()).Return(platform.ErrAccountNotFound)
			},
			args: args{
				req: RequestTransfer{
					SenderAccountNumber:   555001,
					ReceiverAccountNumber: 555002,
					Amount:                10000,
				},
			},
			wantInteractorErr: interactors.ErrAccountNotFound,
		},
		{
			name: "insufficient balance",
			fields: fields{
				accountRepo: accountMock,
			},
			mockRepoFunc: func() {
				accountMock.EXPECT().UpdateBalance(gomock.Any(), gomock.Any(), gomock.Any()).Return(platform.ErrInsufficientBalance)
			},
			args: args{
				req: RequestTransfer{
					SenderAccountNumber:   555001,
					ReceiverAccountNumber: 555002,
					Amount:                1000000,
				},
			},
			wantInteractorErr: interactors.ErrInsufficientBalance,
		},
		{
			name: "success transaction",
			fields: fields{
				accountRepo: accountMock,
			},
			mockRepoFunc: func() {
				accountMock.EXPECT().UpdateBalance(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			args: args{
				req: RequestTransfer{
					SenderAccountNumber:   555001,
					ReceiverAccountNumber: 555002,
					Amount:                1000000,
				},
			},
			wantInteractorErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Interactors{
				AccountRepo: tt.fields.accountRepo,
			}
			tt.mockRepoFunc()
			if gotInteractorErr := i.Transfer(tt.args.req); !reflect.DeepEqual(gotInteractorErr, tt.wantInteractorErr) {
				t.Errorf("Interactors.Transfer() = %v, want %v", gotInteractorErr, tt.wantInteractorErr)
			}
		})
	}
}
