package repository

import (
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user entity.Users) (entity.Users, error)
	UpdateUser(user entity.Users) (entity.Users, error)
	UpdateUserPassword(user entity.Users) error
	FindByEmail(user entity.Users) (entity.Users, error)
	UserIsExist(id string) bool
	GetByUsername(username string) interface{}
	ProfileUser(user entity.Users) (entity.Users, error)
	EmailIsExist(email string) bool
	GetUserByRole(user entity.Users) ([]entity.Users, error)
	DeleteCreator(user entity.Users) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) InsertUser(user entity.Users) (entity.Users, error) {
	if user.Role == "" {
		user.Role = entity.Participant
	}
	err := r.db.Raw("INSERT INTO users (username, fullname, email, password, role) VALUES (@Username, @Fullname, @Email, @Password, @Role)", user).Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(user entity.Users) (entity.Users, error) {
	err := r.db.Exec("UPDATE users SET username=@Username fullname=@Fullname email=@Email role=@Role updated_at=@UpdatedAt WHERE id=@ID", user).Save(&user)
	if err.Error != nil {
		return user, err.Error
	}
	return user, nil
}

func (r *userRepository) UpdateUserPassword(user entity.Users) error {
	err := r.db.Exec("UPDATE users SET password=@Password WHERE id=@ID", user).Save(&user)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *userRepository) FindByEmail(user entity.Users) (entity.Users, error) {
	err := r.db.Raw("SELECT * FROM users WHERE email=@Email", user).Scan(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) UserIsExist(username string) bool {
	var user entity.Users
	res := r.db.Raw("SELECT * FROM users WHERE username=@Username", map[string]interface{}{"Username": username}).Take(&user)
	return res.Error == nil
}

func (r *userRepository) EmailIsExist(email string) bool {
	var user entity.Users
	res := r.db.Raw("SELECT * FROM users WHERE email=@Email", map[string]interface{}{"Username": email}).Take(&user)
	return res.Error == nil
}

func (r *userRepository) GetByUsername(username string) interface{} {
	user := entity.Users{}
	res := r.db.Raw("SELECT * FROM users WHERE username=@Username", map[string]interface{}{"Username": username}).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (r *userRepository) ProfileUser(user entity.Users) (entity.Users, error) {
	err := r.db.Raw("SELECT * FROM users WHERE id=@ID", user).Take(&user)
	if err != nil {
		return user, err.Error
	}
	return user, nil
}

func (r *userRepository) GetUserByRole(user entity.Users) ([]entity.Users, error) {
	var users []entity.Users
	err := r.db.Raw("SELECT * FROM users WHERE role = @Role", user).Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

func (r *userRepository) DeleteCreator(user entity.Users) error {
	err := r.db.Raw("SELECT * FROM users WHERE id = @ID", user).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
