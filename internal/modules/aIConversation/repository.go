package aIconversation

import (
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"gorm.io/gorm"
)

type AIConversationRepository interface {
	Create(userId uint) (*entity.AIConversation, error)
	FindOne(userId uint) (*entity.AIConversation, error)
	DeleteOne(id uint) error
}

type aIConversationRepository struct {
	db *gorm.DB
}

func NewAIConversationRepository(db *gorm.DB) AIConversationRepository {
	return &aIConversationRepository{db: db}
}

func (r *aIConversationRepository) Create(userId uint) (*entity.AIConversation, error) {
	message := entity.AIConversation{
		UserID: userId,
	}

	err := r.db.Create(&message).Error

	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (r *aIConversationRepository) FindOne(userId uint) (*entity.AIConversation, error) {

	var DirectMessage entity.AIConversation

	err := r.db.Preload("Messages").Where("user_id = ?", userId).Take(&DirectMessage).Error

	if err != nil {
		return nil, err
	}

	return &DirectMessage, nil
}

func (r *aIConversationRepository) DeleteOne(id uint) error {

	err := r.db.Delete(&entity.AIConversation{}, id).Error

	if err != nil {
		return err
	}

	return nil
}