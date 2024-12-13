package user

import "fmt"

// Service contains business logic for user operations
type Service struct{}

// SignupService processes the signup logic
func (s *Service) SignupService(username, email, password string) string {
	return fmt.Sprintf("User %s with email %s has been registered.", username, email)
}

// SigninService processes the signin logic
func (s *Service) SigninService(username, password string) string {
	return fmt.Sprintf("Welcome back, %s!", username)
}
