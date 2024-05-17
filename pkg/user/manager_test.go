package user

import (
	"testing"
)

func TestUserProfileCreation(t *testing.T) {
	manager := NewUserManager()
	user := manager.AddUser("John")
	if user.Name != "John" {
		t.Errorf("Expected user name to be John, got %s", user.Name)
	}
}

func TestUpdateUserProfile(t *testing.T) {
	manager := NewUserManager()
	manager.AddUser("Jane")
	_, updated := manager.UpdateUser("Jane", 30, 15)
	if !updated {
		t.Errorf("User profile should be updated")
	}
	userProfile, _ := manager.GetUser("Jane")
	if userProfile.Volume != 30 || userProfile.Channel != 15 {
		t.Errorf("Expected Volume 30 and Channel 15, got Volume %d and Channel %d", userProfile.Volume, userProfile.Channel)
	}
}

func TestGetNonexistentUser(t *testing.T) {
	manager := NewUserManager()
	_, exists := manager.GetUser("DoesNotExist")
	if exists {
		t.Errorf("Should not be able to retrieve a non-existent user")
	}
}
