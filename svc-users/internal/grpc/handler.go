package grpc

import (
	"context"
	"github.com/YReshetko/it-learning-platform/svc-users/internal/mapper"
	"github.com/YReshetko/it-learning-platform/svc-users/internal/storage"
	"github.com/YReshetko/it-learning-platform/svc-users/pb/users"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
	logger                               *logrus.Entry
}

func (h *Handler) CreateUser(_ context.Context, request *users.CreateUserRequest) (*users.CreateUserResponse, error) {
	logger := h.logger.WithField("method", "CreateUser").WithField("request", request)
	user := request.GetUser()
	if user == nil {
		logger.Error("no user is sent")
		return &users.CreateUserResponse{}, status.Error(codes.InvalidArgument, "no user is sent")
	}
	dbUser, err := mapper.PbToDBUser(*user)
	if err != nil {
		logger.WithError(err).Error("unable to map proto to database user model")
		return &users.CreateUserResponse{}, status.Error(codes.Internal, "unable to map proto to database user model")
	}
	dbUser.CreatedAt = time.Now()
	dbUser.UpdatedAt = time.Now()

	dbUser, err = h.storage.Create(dbUser)
	if err != nil {
		logger.WithError(err).Error("unable to save user")
		return &users.CreateUserResponse{}, status.Error(codes.Internal, "unable to save user")
	}
	usr := mapper.DBToPbUser(dbUser)
	return &users.CreateUserResponse{
		User: &usr,
	}, nil
}

func (h *Handler) FindUsers(ctx context.Context, request *users.FindUsersRequest) (*users.FindUsersResponse, error) {
	logger := h.logger.WithField("method", "FindUsers").WithField("request", request)
	var IDs []uuid.UUID
	for _, ID := range request.Ids {
		u, _ := uuid.Parse(ID)
		IDs = append(IDs, u)
	}
	usrs, err := h.storage.FindByIDs(IDs)
	if err != nil {
		logger.WithError(err).Error("unable to find users by IDs")
		return &users.FindUsersResponse{}, status.Error(codes.Internal, "unable to find users by IDs")
	}
	pbUsers := make([]*users.User, len(usrs))
	for i, user := range usrs {
		u := mapper.DBToPbUser(user)
		pbUsers[i] = &u
	}
	return &users.FindUsersResponse{User: pbUsers}, nil
}

func (h *Handler) UserInfo(ctx context.Context, request *users.UserInfoRequest) (*users.UserInfoResponse, error) {
	logger := h.logger.WithField("method", "UserInfo").WithField("request", request)
	UUID, err := uuid.Parse(request.Id)
	if err != nil {
		logger.WithError(err).Error("unable to find user by ID")
		return &users.UserInfoResponse{}, status.Error(codes.InvalidArgument, "invalid user ID")
	}
	user, err := h.storage.FindByID(UUID)
	if err != nil {
		logger.WithError(err).Error("unable to retrieve user by ID")
		return &users.UserInfoResponse{}, status.Error(codes.NotFound, "unable to retrieve user by ID")
	}
	pbUser := mapper.DBToPbUser(user)
	return &users.UserInfoResponse{User: &pbUser}, nil
}

func (h *Handler) FindUserByExternalID(ctx context.Context, request *users.FindUserByExternalIDRequest) (*users.FindUserByExternalIDResponse, error) {
	logger := h.logger.WithField("method", "FindUserByExternalID").WithField("request", request)
	UUID, err := uuid.Parse(request.ExternalId)
	if err != nil {
		logger.WithError(err).Error("unable to find user by ID")
		return &users.FindUserByExternalIDResponse{}, status.Error(codes.InvalidArgument, "invalid user ExternalId")
	}
	user, err := h.storage.FindByExternalID(UUID)
	if err != nil {
		logger.WithError(err).Error("unable to retrieve user by ExternalId")
		return &users.FindUserByExternalIDResponse{}, status.Error(codes.NotFound, "unable to retrieve user by ExternalId")
	}
	pbUser := mapper.DBToPbUser(user)
	return &users.FindUserByExternalIDResponse{User: &pbUser}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, request *users.UpdateUserRequest) (*users.UpdateUserResponse, error) {
	logger := h.logger.WithField("method", "UpdateUser").WithField("request", request)
	UUID, err := uuid.Parse(request.GetUser().GetId())
	if err != nil {
		logger.WithError(err).Error("unable to find user by ID")
		return &users.UpdateUserResponse{}, status.Error(codes.InvalidArgument, "invalid user ID")
	}
	user, err := h.storage.FindByID(UUID)
	if err != nil {
		logger.WithError(err).Error("unable to retrieve user by ID")
		return &users.UpdateUserResponse{}, status.Error(codes.NotFound, "unable to retrieve user by ID")
	}

	user.Active = request.GetUser().GetActive()
	user.LastName = request.GetUser().GetLastName()
	user.FirstName = request.GetUser().GetFirstName()

	externalId, err := uuid.Parse(request.GetUser().GetExternalId())
	if err != nil {
		logger.WithError(err).Error("unable to update user with invalid external ID")
		return &users.UpdateUserResponse{}, status.Error(codes.InvalidArgument, "unable to update user with invalid external ID")
	}
	user.ExternalID = externalId

	err = h.storage.Update(user)
	if err != nil {
		logger.WithError(err).Error("unable to update user")
		return &users.UpdateUserResponse{}, status.Error(codes.Internal, "unable to update user")
	}
	pbUser := mapper.DBToPbUser(user)
	return &users.UpdateUserResponse{User: &pbUser}, nil
}
