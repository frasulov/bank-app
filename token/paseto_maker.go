package token

import (
	"fmt"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
	"time"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size. Must be %d characters", chacha20poly1305.KeySize)
	}
	maker := &PasetoMaker{
		symmetricKey: []byte(symmetricKey),
		paseto:       paseto.NewV2(),
	}
	return maker, nil
}

func (p *PasetoMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", nil, err
	}
	res, err := p.paseto.Encrypt(p.symmetricKey, payload, nil)
	return res, payload, err
}

func (p *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := p.paseto.Decrypt(token, p.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrTokenInvalid
	}
	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil
}
