package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lukabrkovic/artemis/internal/config"
	"github.com/o1egl/paseto"
)

var (
	ErrInvalidToken     = errors.New("token is invalid")
	ErrExpiredToken     = errors.New("token has expired")
	ErrInvalidTokenType = errors.New("invalid token type")
)

type TokenType string

const (
	TokenTypeAccess  TokenType = "access"
	TokenTypeRefresh TokenType = "refresh"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Type      TokenType `json:"type"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(userID uuid.UUID, tokenType TokenType, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Payload{
		ID:        tokenID,
		UserID:    userID,
		Type:      tokenType,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

type Maker interface {
	CreateAccessToken(userID uuid.UUID) (string, *Payload, error)
	CreateRefreshToken(userID uuid.UUID) (string, *Payload, error)
	VerifyAccessToken(token string) (*Payload, error)
	VerifyRefreshToken(token string) (*Payload, error)
	Config() config.TokenConfig
}

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
	config       config.TokenConfig
}

func NewPasetoMaker(symmetricKey string, cfg config.TokenConfig) (Maker, error) {
	if len(symmetricKey) != 32 {
		return nil, errors.New("invalid key size: must be exactly 32 bytes")
	}

	return &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
		config:       cfg,
	}, nil
}

func (m *PasetoMaker) CreateAccessToken(userID uuid.UUID) (string, *Payload, error) {
	return m.createToken(userID, TokenTypeAccess, m.config.AccessTokenDuration)
}

func (m *PasetoMaker) CreateRefreshToken(userID uuid.UUID) (string, *Payload, error) {
	return m.createToken(userID, TokenTypeRefresh, m.config.RefreshTokenDuration)
}

func (m *PasetoMaker) createToken(userID uuid.UUID, tokenType TokenType, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(userID, tokenType, duration)
	if err != nil {
		return "", nil, err
	}

	token, err := m.paseto.Encrypt(m.symmetricKey, payload, nil)
	if err != nil {
		return "", nil, err
	}

	return token, payload, nil
}

func (m *PasetoMaker) VerifyAccessToken(token string) (*Payload, error) {
	return m.verifyToken(token, TokenTypeAccess)
}

func (m *PasetoMaker) VerifyRefreshToken(token string) (*Payload, error) {
	return m.verifyToken(token, TokenTypeRefresh)
}

func (m *PasetoMaker) verifyToken(token string, expectedType TokenType) (*Payload, error) {
	payload := &Payload{}

	err := m.paseto.Decrypt(token, m.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	if err := payload.Valid(); err != nil {
		return nil, err
	}

	if payload.Type != expectedType {
		return nil, ErrInvalidTokenType
	}

	return payload, nil
}

func (m *PasetoMaker) Config() config.TokenConfig {
	return m.config
}
