package user

import (
	"fmt"
	"vago/internal/domain"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	repo   domain.UserRepository
	tokens domain.TokenProvider
}

func NewService(repo domain.UserRepository, tokens domain.TokenProvider) *Service {
	return &Service{
		repo:   repo,
		tokens: tokens,
	}
}

func (s *Service) CreateUser(dto domain.DTO) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user := domain.New(dto.Login, string(hash), dto.Username, dto.Color, dto.Role)
	return s.repo.CreateUser(user)
}

func (s *Service) DeleteUser(id int64) error {
	return s.repo.DeleteUser(id)
}

func (s *Service) GetByID(id domain.UserID) (domain.User, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Login(login, password string) (domain.User, *domain.TokenPair, error) {
	u, errGetUser := s.repo.GetByLogin(login)

	if errGetUser != nil {
		return domain.User{}, nil, domain.ErrUserNotFound
	}

	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) != nil {
		return u, nil, domain.ErrIncorrectPassword
	}

	tokens, err := s.tokens.CreateTokenPair(u.ID, string(u.Role))
	if err != nil {
		return domain.User{}, nil, fmt.Errorf("error creating tokens: %s", err.Error())
	}

	return u, tokens, nil
}

func (s *Service) Refresh(token string) (domain.User, string, error) {
	claims, errParseToken := s.tokens.ParseToken(token)
	if errParseToken != nil {
		return domain.User{}, "", status.Error(codes.Unauthenticated, "token read error")
	}
	u, errGetUser := s.repo.GetByID(domain.UserID(claims.UserID()))
	if errGetUser != nil {
		return domain.User{}, "", domain.ErrUserNotFound
	}

	newToken, errToken := s.tokens.CreateToken(u.ID, string(u.Role), true)
	if errToken != nil {
		return domain.User{}, "", status.Error(codes.Unauthenticated, fmt.Sprintf("Error creating new token: %s", errToken.Error()))
	}

	return u, newToken, nil
}

func (s *Service) GetAll() ([]domain.User, error) {
	return s.repo.GetAll()
}
