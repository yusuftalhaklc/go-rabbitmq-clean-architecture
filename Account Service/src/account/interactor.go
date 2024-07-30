package account

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UseCase interface {
	Register(request Register) error
	Verify(request VerifyCode) error
}

type interactor struct {
	services Services
}

func NewInteractor(service Services) UseCase {
	return &interactor{services: service}
}

func (i *interactor) Register(request Register) error {
	id := uuid.New()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.PasswordText), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	err = i.services.AccountRepo.Create(Account{
		Id:           id.String(),
		FullName:     request.FullName,
		Email:        request.Email,
		PasswordHash: string(hashedPassword),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		return err
	}

	verificationCode, err := generateRandomString(40)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("verification:%s", request.Email)
	err = i.services.RedisRepo.SetEmailVerificationCode(key, verificationCode)
	if err != nil {
		return err
	}

	link := fmt.Sprintf("http://localhost:3000/account?code=%s&email=%s", verificationCode, request.Email)

	template := EmailTemplate{
		To:   request.Email,
		Link: link,
		Type: Verification,
	}

	body, err := json.Marshal(template)
	if err != nil {
		return err
	}

	err = i.services.RabbitMqProducer.Publish(body)
	if err != nil {
		return err
	}

	return nil
}

func (i *interactor) Verify(request VerifyCode) error {
	key := fmt.Sprintf("verification:%s", request.Email)
	value := i.services.RedisRepo.GetEmailVerification(key)
	if value != request.Code {
		return fmt.Errorf("cannot verified")
	}
	
	err := i.services.AccountRepo.Update(request.Email, true)
	if err != nil {
		return err
	}
	return nil
}

func generateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result = make([]byte, length)
	bytes := make([]byte, length+(length/4))
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	for i := 0; i < length; i++ {
		result[i] = charset[bytes[i]%byte(len(charset))]
	}

	return string(result), nil
}
