package buffer

import (
	"fmt"
)

type UpdatesService struct {
	client *Client
}

type Update struct {
	CreatedAt       int                    `json:"created_at,omitempty"`
	ID              string                 `json:"id,omitempty"`
	Day             string                 `json:"day,omitempty"`
	DueAt           int                    `json:"due_at,omitempty"`
	DueTime         string                 `json:"due_time,omitempty"`
	ProfileID       string                 `json:"profile_id,omitempty"`
	ProfileService  string                 `json:"profile_service,omitempty"`
	SentAt          int                    `json:"sent_at,omitempty"`
	ServiceUpdateID string                 `json:"service_update_id,omitempty"`
	Statistics      map[string]interface{} `json:"statistics,omitempty"`
	Status          string                 `json:"status,omitempty"`
	Text            string                 `json:"text,omitempty"`
	TextFormatted   string                 `json:"text_formatted,omitempty"`
	UserID          string                 `json:"user_id,omitempty"`
	Via             string                 `json:"via,omitempty"`
}

func (s *UpdatesService) Get(updateID string) (*Update, error) {
	u := fmt.Sprintf("/1/updates/%v.json", updateID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	update := new(Update)
	_, err = s.client.Do(req, update)

	return update, err
}
