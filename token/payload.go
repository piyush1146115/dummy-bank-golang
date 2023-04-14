package token

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

// Different types of error returned by the verify token function
var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

//func (payload *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
//	return &jwt.NumericDate{Time: payload.ExpiredAt}, nil
//}
//
//func (payload *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
//	return &jwt.NumericDate{Time: payload.IssuedAt}, nil
//}
//
//func (payload *Payload) GetNotBefore() (*jwt.NumericDate, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (payload *Payload) GetIssuer() (string, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (payload *Payload) GetSubject() (string, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (payload *Payload) GetAudience() (jwt.ClaimStrings, error) {
//	//TODO implement me
//	panic("implement me")
//}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenId,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}
