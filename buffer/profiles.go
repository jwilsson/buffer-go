package buffer

import (
	"fmt"
)

type ProfilesService struct {
	client *Client
}

type Profiles struct {
	Items []Profile
}

type Profile struct {
	Avatar            string                 `json:"avatar,omitempty"`
	CreatedAt         int                    `json:"created_at,omitempty"`
	Default           bool                   `json:"default,omitempty"`
	FormattedUsername string                 `json:"formatted_username,omitempty"`
	ID                string                 `json:"id,omitempty"`
	Schedules         []ProfileSchedule      `json:"schedules,omitempty"`
	Service           string                 `json:"service,omitempty"`
	ServiceID         string                 `json:"service_id,omitempty"`
	ServiceUsername   string                 `json:"service_username,omitempty"`
	Statistics        map[string]interface{} `json:"statistics,omitempty"`
	TeamMembers       []string               `json:"team_members,omitempty"`
	Timezone          string                 `json:"timezone,omitempty"`
	UserID            string                 `json:"user_id,omitempty"`
}

type ProfileSchedule struct {
	Days  []string `json:"days,omitempty"`
	Times []string `json:"times,omitempty"`
}

func (s *ProfilesService) Get(profileID string) (*Profile, error) {
	u := fmt.Sprintf("/1/profiles/%v.json", profileID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	profile := new(Profile)
	_, err = s.client.Do(req, profile)

	return profile, err
}

func (s *ProfilesService) GetAll() (*Profiles, error) {
	u := "/1/profiles.json"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	profiles := new(Profiles)
	_, err = s.client.Do(req, profiles)

	return profiles, err
}

func (s *ProfilesService) GetSchedules(profileID string) (*ProfileSchedule, error) {
	u := fmt.Sprintf("/1/profiles/%v.json", profileID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	schedule := new(ProfileSchedule)
	_, err = s.client.Do(req, schedule)

	return schedule, err
}

func (s *ProfilesService) GetPendingUpdates(profileID string, params *PagingParams) (*Updates, error) {
	u := fmt.Sprintf("/1/profiles/%v/updates/pending.json", profileID)

	req, err := s.client.NewRequest("GET", u, params)
	if err != nil {
		return nil, err
	}

	updates := new(Updates)
	_, err = s.client.Do(req, updates)

	return updates, err
}
