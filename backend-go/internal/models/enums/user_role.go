package enums

// UserRole represents the role of a user in the system
type UserRole string

const (
	RoleSuperAdmin UserRole = "super_admin"
	RoleMaveric    UserRole = "maveric"
)

// ValidUserRoles returns all valid user roles
func ValidUserRoles() []UserRole {
	return []UserRole{
		RoleSuperAdmin,
		RoleMaveric,
	}
}

// IsValid checks if a user role is valid
func (r UserRole) IsValid() bool {
	for _, validRole := range ValidUserRoles() {
		if r == validRole {
			return true
		}
	}
	return false
}

// String returns the string representation of the user role
func (r UserRole) String() string {
	return string(r)
}
