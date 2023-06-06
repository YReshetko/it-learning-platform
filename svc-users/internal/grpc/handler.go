package grpc

import (
	"context"
	"fmt"
	"github.com/YReshetko/it-academy-cources/svc-users/internal/mapper"
	"github.com/YReshetko/it-academy-cources/svc-users/internal/storage"
	"github.com/YReshetko/it-academy-cources/svc-users/pb/users"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

/*
Handler the GRPC requests handler
@Constructor
*/
type Handler struct {
	users.UnimplementedUserServiceServer // @Exclude
	storage                              storage.UserStorage
}

func (s Handler) CreateUser(_ context.Context, request *users.CreateUserRequest) (*users.CreateUserResponse, error) {
	user := request.GetUser()
	if user == nil {
		return &users.CreateUserResponse{}, status.Error(codes.InvalidArgument, "no user is sent")
	}
	dbUser := mapper.PbToDBUser(*user)
	dbUser.CreatedAt = time.Now()
	dbUser.UpdatedAt = time.Now()

	dbUser, err := s.storage.Create(dbUser)
	if err != nil {
		fmt.Println("unable to save user", err)
		return &users.CreateUserResponse{}, status.Error(codes.Internal, "unable to save user")
	}
	usr := mapper.DBToPbUser(dbUser)
	return &users.CreateUserResponse{
		User: &usr,
	}, nil
}

func (s Handler) FindUsers(ctx context.Context, request *users.FindUsersRequest) (*users.FindUsersResponse, error) {
	var IDs []uuid.UUID
	for _, ID := range request.Ids {
		u, _ := uuid.FromBytes([]byte(ID))
		IDs = append(IDs, u)
	}
	usrs, err := s.storage.FindByIDs(IDs)
	if err != nil {
		fmt.Println("unable to find users by IDs:", err)
		return &users.FindUsersResponse{}, status.Error(codes.Internal, "unable to find users by IDs")
	}
	pbUsers := make([]*users.User, len(usrs))
	for i, user := range usrs {
		u := mapper.DBToPbUser(user)
		pbUsers[i] = &u
	}
	return &users.FindUsersResponse{User: pbUsers}, nil
}

func (s Handler) UserInfo(ctx context.Context, request *users.UserInfoRequest) (*users.UserInfoResponse, error) {
	UUID, err := uuid.FromBytes([]byte(request.Id))
	if err != nil {
		fmt.Println("unable to find user by ID:", err)
		return &users.UserInfoResponse{}, status.Error(codes.InvalidArgument, "invalid user ID")
	}
	user, err := s.storage.FindByID(UUID)
	if err != nil {
		fmt.Println("unable to retrieve user by ID:", err)
		return &users.UserInfoResponse{}, status.Error(codes.NotFound, "unable to retrieve user by ID")
	}
	pbUser := mapper.DBToPbUser(user)
	return &users.UserInfoResponse{User: &pbUser}, nil
}

func (s Handler) FindUserByExternalID(ctx context.Context, request *users.FindUserByExternalIDRequest) (*users.FindUserByExternalIDResponse, error) {
	UUID, err := uuid.FromBytes([]byte(request.ExternalId))
	if err != nil {
		fmt.Println("unable to find user by ID:", err)
		return &users.FindUserByExternalIDResponse{}, status.Error(codes.InvalidArgument, "invalid user ExternalId")
	}
	user, err := s.storage.FindByExternalID(UUID)
	if err != nil {
		fmt.Println("unable to retrieve user by ExternalId:", err)
		return &users.FindUserByExternalIDResponse{}, status.Error(codes.NotFound, "unable to retrieve user by ExternalId")
	}
	pbUser := mapper.DBToPbUser(user)
	return &users.FindUserByExternalIDResponse{User: &pbUser}, nil
}

func (s Handler) UpdateUser(ctx context.Context, request *users.UpdateUserRequest) (*users.UpdateUserResponse, error) {
	UUID, err := uuid.FromBytes([]byte(request.GetUser().GetId()))
	if err != nil {
		fmt.Println("unable to find user by ID:", err)
		return &users.UpdateUserResponse{}, status.Error(codes.InvalidArgument, "invalid user ID")
	}
	user, err := s.storage.FindByID(UUID)
	if err != nil {
		fmt.Println("unable to retrieve user by ID:", err)
		return &users.UpdateUserResponse{}, status.Error(codes.NotFound, "unable to retrieve user by ID")
	}

	user.Active = request.GetUser().GetActive()
	user.LastName = request.GetUser().GetLastName()
	user.FirstName = request.GetUser().GetFirstName()

	externalId, err := uuid.FromBytes([]byte(request.GetUser().GetExternalId()))
	if err != nil {
		fmt.Println("unable to update user with invalid external ID:", err)
		return &users.UpdateUserResponse{}, status.Error(codes.InvalidArgument, "unable to update user with invalid external ID")
	}
	user.ExternalID = externalId

	err = s.storage.Update(user)
	if err != nil {
		fmt.Println("unable to update user:", err)
		return &users.UpdateUserResponse{}, status.Error(codes.Internal, "unable to update user")
	}
	pbUser := mapper.DBToPbUser(user)
	return &users.UpdateUserResponse{User: &pbUser}, nil
}
