package app

import (
	pb "account-management-service/gen/go/v1/proto"
	"account-management-service/internal/service"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccountServer struct {
	pb.UnimplementedAccountServiceServer
	service service.AccountService
}

func NewAccountServer(s service.AccountService) *AccountServer {
	return &AccountServer{service: s}
}

func (s AccountServer) CreateAccount(ctx context.Context, request *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	account, err := s.service.CreateAccount(ctx, service.CreateAccountDTO{
		FullName:             request.GetFullName(),
		Email:                request.GetEmail(),
		Username:             request.GetUsername(),
		Password:             request.GetPassword(),
		PasswordConfirmation: request.GetPasswordConfirmation(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateAccountResponse{
		Id:                account.ID.String(),
		FullName:          account.FullName,
		Email:             account.Email,
		Username:          account.Username,
		CreatedAt:         timestamppb.New(account.CreatedAt),
		PasswordUpdatedAt: timestamppb.New(account.PasswordUpdatedAt),
		EmailConfirmedAt:  nil,
	}, nil
}
