package session

import (
	"encoding/json"
	"errors"

	"agussyahrilmubarok.github.io/web/internal/delivery/web/payload"
	"github.com/gin-contrib/sessions"
)

const userKey = "user"

var (
	ErrSessionNotFound = errors.New("session not found")
	ErrSessionInvalid  = errors.New("session invalid")
)

func SaveUser(s sessions.Session, user payload.UserResponse) error {
	bytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	s.Set(userKey, string(bytes))

	if err := s.Save(); err != nil {
		return err
	}
	return nil
}

func GetUser(s sessions.Session) (payload.UserResponse, error) {
	val := s.Get(userKey)
	if val == nil {
		return payload.UserResponse{}, ErrSessionNotFound
	}

	raw, ok := val.(string)
	if !ok {
		return payload.UserResponse{}, ErrSessionInvalid
	}

	var user payload.UserResponse
	if err := json.Unmarshal([]byte(raw), &user); err != nil {
		return payload.UserResponse{}, ErrSessionInvalid
	}
	return user, nil
}

func DeleteUser(s sessions.Session) error {
	s.Delete(userKey)

	if err := s.Save(); err != nil {
		return err
	}
	return nil
}
