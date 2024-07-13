package entity

import "time"

type AIConversation struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint
	Messages  []AIMessage `gorm:"foreignKey:ConversationID;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time   `gorm:"autoCreateTime"`
	UpdatedAt time.Time   `gorm:"autoUpdateTime"`
}
