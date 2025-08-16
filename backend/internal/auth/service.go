package auth

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"net/mail"
	"time"

	"github.com/resend/resend-go/v2"
)

type Service struct {
	storage      *Storage
	resendClient *resend.Client
	jwtSecret    string
}

func NewService(storage *Storage, resendCLient *resend.Client, jwtSecret string) *Service {
	return &Service{
		storage:      storage,
		resendClient: resendCLient,
		jwtSecret:    jwtSecret,
	}
}

var ErrInvalidEmailFormat = errors.New("invalid email format")

func validateEmailFormat(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func generateOTP() (string, error) {
	max := big.NewInt(1000000)
	n, err := rand.Int(rand.Reader, max)

	if err != nil {
		return "", err
	}

	otp := fmt.Sprintf("%06d", n)

	return otp, nil
}

func (s *Service) RequestOTP(ctx context.Context, email string) error {
	if !validateEmailFormat(email) {
		return ErrInvalidEmailFormat
	}

	otp, err := generateOTP()

	if err != nil {
		return err
	}

	expiresAt := time.Now().Add(5 * time.Minute)
	if err := s.storage.SaveOTP(ctx, email, otp, expiresAt); err != nil {
		return err
	}

	params := &resend.SendEmailRequest{
		From:    "Likely <no-reply@likely.ro>",
		To:      []string{email},
		Html:    "<strong>" + otp + "</strong>",
		Subject: "Cod OTP",
	}

	sent, err := s.resendClient.Emails.Send(params)

	if err != nil {
		return err
	}

	fmt.Println(sent)

	return nil
}
