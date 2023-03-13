package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/webmalc/vishleva-backend/models"
)

// ClientRepository is the repository.
type ClientRepository struct {
	db *gorm.DB
}

// GetOrCreate get or create a client.
func (s *ClientRepository) GetOrCreate(
	email, phone, name string,
) (*models.Client, error) {
	client := models.Client{}
	err := s.db.Where("phone = ?", phone).Or("email = ?", email).
		First(&client).Error
	if err == nil {
		return &client, nil
	}
	client.Name = name
	client.Comment = "automatically created client"
	client.Email = &email
	client.Phone = phone
	err = s.db.Create(&client).Error

	return &client, err
}

// NewClientRepository returns a new repository struct.
func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{db: db}
}
