package gorm

import (
	"errors"
	"fmt"
	"vago/internal/app"
	"vago/internal/domain"

	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	UniqueCode      = "23505"
	ValueToLong     = "22001"
	ConstraintLogin = "users_login_key"
)

type UserRepository struct {
	db  *gorm.DB
	log *zap.SugaredLogger
}

func NewUserRepo(ctx *app.Context) domain.UserRepository {
	return &UserRepository{
		db:  ctx.DB,
		log: ctx.Log,
	}
}

func (r *UserRepository) CreateUser(u domain.User) error {
	entity := toEntity(u)
	if err := r.db.Create(&entity).Error; err != nil {
		if pgErr := parsePgError(err); pgErr != nil {
			switch pgErr.Code {
			case UniqueCode:
				switch pgErr.ConstraintName {
				case ConstraintLogin:
					return domain.ErrLoginExists
				}
			case ValueToLong:
				return domain.ErrValueToLong
			}
		}
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *UserRepository) DeleteUser(id int64) error {
	res := r.db.Delete(&UserEntity{}, id)

	if res.Error != nil {
		return fmt.Errorf("ошибка удаления пользователя: %w", res.Error)
	}

	if res.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}

func (r *UserRepository) GetByLogin(login string) (domain.User, error) {
	var entity UserEntity
	if err := r.db.Where("login = ?", login).First(&entity).Error; err != nil {
		return domain.User{}, err
	}

	return toDomain(entity), nil
}

func (r *UserRepository) GetByID(id int64) (domain.User, error) {
	var entity UserEntity
	if err := r.db.First(&entity, id).Error; err != nil {
		return domain.User{}, err
	}

	return toDomain(entity), nil
}

func (r *UserRepository) GetByIDs(ids []int64) ([]domain.User, error) {
	if len(ids) == 0 {
		return []domain.User{}, nil
	}

	var entities []UserEntity

	if err := r.db.
		Where("id IN ?", ids).
		Find(&entities).Error; err != nil {
		return nil, err
	}

	users := make([]domain.User, 0, len(entities))
	for _, e := range entities {
		users = append(users, toDomain(e))
	}

	return users, nil
}

func (r *UserRepository) GetAll() ([]domain.User, error) {
	var entities []UserEntity
	err := r.db.Order("id ASC").Find(&entities).Error
	result := make([]domain.User, 0, len(entities))

	for _, entity := range entities {
		result = append(result, toDomain(entity))
	}

	return result, err
}

func toDomain(e UserEntity) domain.User {
	return domain.User{
		ID:        e.ID,
		Login:     e.Login,
		Username:  e.Username,
		Password:  e.Password,
		CreatedAt: e.CreatedAt,
		Role:      domain.Role(e.Role),
		Color:     e.Color,
	}
}

func toEntity(u domain.User) UserEntity {
	return UserEntity{
		ID:        u.ID,
		Login:     u.Login,
		Password:  u.Password,
		Username:  u.Username,
		Color:     u.Color,
		Role:      string(u.Role), // доменный тип → строка
		CreatedAt: u.CreatedAt,
	}
}

func parsePgError(err error) *pgconn.PgError {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr
	}
	return nil
}
