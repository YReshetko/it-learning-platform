package handlers

import (
	"github.com/YReshetko/it-learning-platform/api-app/internal/clients"
	"github.com/YReshetko/it-learning-platform/api-app/internal/http"
	"github.com/YReshetko/it-learning-platform/api-app/internal/http/models"
	"github.com/YReshetko/it-learning-platform/svc-auth/pb/auth"
	"github.com/sirupsen/logrus"
	rest "net/http"
)

/*
Self handler that returns user info to frontend
@Constructor
*/
type Self struct {
	client clients.AuthClient
	logger *logrus.Entry
}

func (s *Self) GetUserInfo(context http.Context, _ models.Empty) (models.SelfResponse, http.Status) {
	logger := s.logger.WithField("method", "GetUserInfo")
	userInfo, err := s.client.GetUserInfo(context.Context(), &auth.GetUserInfoRequest{
		Id: context.UserID.String(),
	})
	roles := make([]string, len(userInfo.GetUserInfo().GetRoles()))
	for i, role := range userInfo.GetUserInfo().GetRoles() {
		roles[i] = auth.UserRole_name[int32(role)]
	}
	if err != nil {
		logger.WithError(err).Error("Unable to get user info")
		return models.SelfResponse{}, http.Status{
			Error:      err,
			StatusCode: rest.StatusInternalServerError,
			Message:    "unable to get user info",
		}
	}

	return models.SelfResponse{
		ID:        context.UserID,
		FirstName: userInfo.GetUserInfo().GetFirstName(),
		LastName:  userInfo.GetUserInfo().GetLastName(),
		Roles:     roles,
	}, http.Status{StatusCode: rest.StatusOK}

}
