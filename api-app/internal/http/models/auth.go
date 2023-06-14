package models

import "github.com/YReshetko/it-academy-cources/svc-auth/pb/auth"

type Roles []string

func (r Roles) ToAuthProtoRoles() []auth.UserRole {
	if len(r) == 0 {
		return nil
	}
	out := make([]auth.UserRole, len(r))
	for i, role := range r {
		out[i] = auth.UserRole(auth.UserRole_value[role])
	}
	return out
}

type AuthUser struct {
	Login     string `json:"login"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Roles     Roles  `json:"roles"`
}

type AuthResponse struct {
}

type AuthUsers struct {
	Users []AuthUser `json:"users"`
}
