package controller

import (
	"encoding/json"
	"errors"

	"agussyahrilmubarok.github.io/web/internal/model"
	"github.com/gin-contrib/sessions"
)

const sessionUserKey = "user"

var (
	ErrSessionNotFound = errors.New("session not found")
	ErrSessionInvalid  = errors.New("session invalid")
)

func saveUserSession(s sessions.Session, user model.UserResponse) error {
	bytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	s.Set(sessionUserKey, string(bytes))

	if err := s.Save(); err != nil {
		return err
	}

	return nil
}

func GetUserSession(s sessions.Session) (model.UserResponse, error) {
	val := s.Get(sessionUserKey)
	if val == nil {
		return model.UserResponse{}, ErrSessionNotFound
	}

	raw, ok := val.(string)
	if !ok {
		return model.UserResponse{}, ErrSessionInvalid
	}

	var user model.UserResponse
	if err := json.Unmarshal([]byte(raw), &user); err != nil {
		return model.UserResponse{}, ErrSessionInvalid
	}

	return user, nil
}

func ClearUserSession(s sessions.Session) error {
	s.Delete(sessionUserKey)

	if err := s.Save(); err != nil {
		return err
	}

	return nil
}
