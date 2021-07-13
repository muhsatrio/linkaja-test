package auth

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/muhsatrio/golang-boilerplate/interactors"
	mock_jwt "github.com/muhsatrio/golang-boilerplate/mocks/platform/jwt"
)

func TestService_Login(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	jwtMock := mock_jwt.NewMockJwtAdapter(mockCtrl)

	svc := Service{
		JwtRepo: jwtMock,
	}

	type args struct {
		req RequestLogin
	}
	tests := []struct {
		name              string
		mock              func()
		args              args
		wantResp          ResponseLogin
		wantInteractorErr interactors.Error
	}{
		{
			name: "login success",
			mock: func() {
				jwtMock.EXPECT().GenerateToken(gomock.Any()).Return("12345", nil)
			},
			args: args{
				req: RequestLogin{
					Username: "muhsatrio",
					Password: "12345",
				},
			},
			wantResp: ResponseLogin{
				Token: "12345",
			},
			wantInteractorErr: nil,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			gotResp, gotInteractorErr := svc.Login(tt.args.req)
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Service.Login() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if !reflect.DeepEqual(gotInteractorErr, tt.wantInteractorErr) {
				t.Errorf("Service.Login() gotInteractorErr = %v, want %v", gotInteractorErr, tt.wantInteractorErr)
			}
		})
	}
}
