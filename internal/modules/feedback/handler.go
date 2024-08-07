package feedback

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	feedbackService FeedbackService
}

type input struct {
	UserID  uint64   `json:"userId"`
	Message string `json:"message"`
}

func NewHandler(feedbackService FeedbackService) Handler {
	return Handler{feedbackService}
}

func (h *Handler) SendFeedback(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	feedback, _ := h.feedbackService.SendFeedback(input.UserID, input.Message)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(feedback); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAllFeedbacks(w http.ResponseWriter, r *http.Request) {

	feedbacks, _ := h.feedbackService.GetAllFeedbacks()

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(feedbacks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
