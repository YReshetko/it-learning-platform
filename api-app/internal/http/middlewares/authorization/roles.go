package authorization

// TODO duplication in svc-auth
type Role string

const (
	ADMIN   Role = "ADMIN"
	MANAGER Role = "MANAGER"
	TEACHER Role = "TEACHER"
	STUDENT Role = "STUDENT"
)
