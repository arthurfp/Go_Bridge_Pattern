package user

import "sync"

// UserManager manages multiple user profiles.
type UserManager struct {
	profiles map[string]*UserProfile
	lock     sync.RWMutex
}

// NewUserManager initializes a new user profile manager.
func NewUserManager() *UserManager {
	return &UserManager{
		profiles: make(map[string]*UserProfile),
	}
}

// AddUser adds a new user profile to the manager.
func (um *UserManager) AddUser(name string) *UserProfile {
	um.lock.Lock()
	defer um.lock.Unlock()

	profile := NewUserProfile(name)
	um.profiles[name] = profile
	return profile
}

// GetUser retrieves a user profile by name.
func (um *UserManager) GetUser(name string) (*UserProfile, bool) {
	um.lock.RLock()
	defer um.lock.RUnlock()

	profile, exists := um.profiles[name]
	return profile, exists
}

// UpdateUser updates settings for a user profile.
func (um *UserManager) UpdateUser(name string, volume, channel int) (*UserProfile, bool) {
	um.lock.Lock()
	defer um.lock.Unlock()

	if profile, exists := um.profiles[name]; exists {
		profile.Volume = volume
		profile.Channel = channel
		return profile, true
	}
	return nil, false
}
