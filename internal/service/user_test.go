package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"grpcAvito/internal/entity"
	mock_usecase "grpcAvito/internal/usecase/mocks"
	"testing"
)

func TestService_GetBalance(t *testing.T) {
	type mockBehavior func(s *mock_usecase.MockUser, user entity.User)
	ctx := context.TODO()
	tests := []struct {
		name                 string
		inputBody            string
		inputUser            entity.User
		mockBehavior         mockBehavior
		expectedSuccess      bool
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"id": 1}`,
			inputUser: entity.User{
				ID: 1,
			},
			mockBehavior: func(s *mock_usecase.MockUser, user entity.User) {
				s.EXPECT().GetBalance(ctx, user).Return(true, nil)
			},
			expectedSuccess:      true,
			expectedResponseBody: `{"id":1,"balance":24,"success":true}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)

			getBalance := mock_usecase.NewMockUser(c)
			test.mockBehavior(getBalance, test.inputUser)

		})
	}
}
