package user

// UserProfile stores settings specific to a user.
type UserProfile struct {
	Name    string
	Volume  int
	Channel int
}

// NewUserProfile creates a new user profile with default settings.
func NewUserProfile(name string) *UserProfile {
	return &UserProfile{
		Name:    name,
		Volume:  50, // Default volume
		Channel: 1,  // Default channel
	}
}
