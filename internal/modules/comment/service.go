package comment

import (
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
)

type CommentService interface {
	SendComment(userId uint64, contentId uint64, text string) (*entity.Comment, error)
	updateComment(id uint64, text string) (*entity.Comment, error)
	GetAllComments(contentId uint64) (*[]entity.Comment, error)
	DeleteContent(id uint64) error
}

type commentService struct {
	commentRepository CommentRepository
}

func NewContentService(commentRepository CommentRepository) CommentService {
	return &commentService{
		commentRepository: commentRepository,
	}
}

func (s *commentService) SendComment(userId uint64, contentId uint64, text string) (*entity.Comment, error) {
	comment, err := s.commentRepository.Create(userId, contentId, text)

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *commentService) updateComment(id uint64, text string) (*entity.Comment, error) {
	comment, err := s.commentRepository.FindOne(id)

	if err != nil {
		return nil, err
	}

	comment.Comment = text

	comment, err = s.commentRepository.Update(comment)

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *commentService) GetAllComments(contentId uint64) (*[]entity.Comment, error) {
	comment, err := s.commentRepository.FindAll(contentId)

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *commentService) DeleteContent(id uint64) error {
	err := s.commentRepository.DeleteOne(id)

	return err
}
