package model

type Role string

const (
	ADMIN   Role = "ADMIN"
	MANAGER Role = "MANAGER"
	TEACHER Role = "TEACHER"
	STUDENT Role = "STUDENT"
)

func AllRoles() []Role {
	return []Role{ADMIN, MANAGER, TEACHER, STUDENT}
}

func (r Role) ToString() string {
	return string(r)
}

func RoleFromString(role string) (Role, bool) {
	m := map[string]Role{
		"ADMIN":   ADMIN,
		"MANAGER": MANAGER,
		"TEACHER": TEACHER,
		"STUDENT": STUDENT,
	}
	r, ok := m[role]
	return r, ok
}

func RolesFromStrings(roles []string) ([]Role, bool) {
	out := make([]Role, len(roles))
	var ok bool
	for i, role := range roles {
		out[i], ok = RoleFromString(role)
		if !ok {
			return nil, false
		}
	}
	return out, true
}

func RolesToStrings(roles []Role) []string {
	out := make([]string, len(roles))
	for i, role := range roles {
		out[i] = role.ToString()
	}
	return out
}
