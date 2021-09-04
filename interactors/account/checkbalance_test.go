package account

import (
	"linkaja-test/domain"
	"linkaja-test/interactors"
	"linkaja-test/platform"
	"linkaja-test/platform/mysql"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	mock_account "linkaja-test/mocks/platform/account"
)

func TestInteractors_CheckBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accountMock := mock_account.NewMockAccountAdapter(ctrl)

	type fields struct {
		accountRepo mysql.AccountAdapter
	}

	type args struct {
		accountNumber uint
	}
	tests := []struct {
		name              string
		fields            fields
		mockRepoFunc      func()
		args              args
		wantResp          ResponseCheckBalance
		wantInteractorErr interactors.Error
	}{
		{
			name: "account not found",
			fields: fields{
				accountRepo: accountMock,
			},
			mockRepoFunc: func() {
				accountMock.EXPECT().Get(gomock.Any()).Return(domain.Account{}, platform.ErrAccountNotFound)
			},
			args: args{
				accountNumber: 111,
			},
			wantResp:          ResponseCheckBalance{},
			wantInteractorErr: interactors.ErrAccountNotFound,
		},
		{
			name: "account found",
			fields: fields{
				accountRepo: accountMock,
			},
			mockRepoFunc: func() {
				accountMock.EXPECT().Get(gomock.Any()).Return(domain.Account{
					AccountNumber: 111,
					CustomerName:  "Dummy",
					Balance:       0,
				}, nil)
			},
			args: args{
				accountNumber: 111,
			},
			wantResp: ResponseCheckBalance{
				AccountNumber: 111,
				CustomerName:  "Dummy",
				Balance:       0,
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
			gotResp, gotInteractorErr := i.CheckBalance(tt.args.accountNumber)
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Interactors.CheckBalance() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if !reflect.DeepEqual(gotInteractorErr, tt.wantInteractorErr) {
				t.Errorf("Interactors.CheckBalance() gotInteractorErr = %v, want %v", gotInteractorErr, tt.wantInteractorErr)
			}
		})
	}
}
