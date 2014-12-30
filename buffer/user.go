package buffer

type UserService struct {
    client *Client
}

type User struct {
    ActivityAt int    `json:"activity_at,omitempty"`
    CreatedAt  int    `json:"created_at,omitempty"`
    ID         string `json:"id,omitempty"`
    Plan       string `json:"plan,omitempty"`
    Timezone   string `json:"timezone,omitempty"`
}

func (s *UserService) Get() (*User, error) {
    u := "/1/user.json"

    req, err := s.client.NewRequest("GET", u, nil)
    if err != nil {
        return nil, err
    }

    user := new(User)
    _, err = s.client.Do(req, user)

    return user, err
}
