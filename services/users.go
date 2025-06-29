package services

// UsersService is the users service.
type UsersService struct {
	logger      InfoLogger
	userCreator UserCreator
}

// CreateAdmin create a new admin user.
func (s *UsersService) CreateAdmin(email, password string) error {
	user, err := s.userCreator.CreateUser(email, password)
	if err == nil {
		s.logger.Infof("admin user has been created %v", user.Email)
	}

	return err
}

// NewBookingService returns a new booking service.
func NewUsersService(
	logger InfoLogger,
	userCreator UserCreator,
) *UsersService {
	return &UsersService{
		logger:      logger,
		userCreator: userCreator,
	}
}
