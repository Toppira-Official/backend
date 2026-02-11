package usecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"github.com/Toppira-Official/backend/internal/configs"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleOauthRedirectURLUsecase interface {
	Execute(ctx context.Context) (redirectUrl, state string)
}

type googleOauthRedirectURLUsecase struct {
	envs configs.Environments
}

func NewGoogleOauthRedirectURLUsecase(envs configs.Environments) GoogleOauthRedirectURLUsecase {
	return &googleOauthRedirectURLUsecase{envs: envs}
}

func (uc *googleOauthRedirectURLUsecase) Execute(ctx context.Context) (redirectUrl, state string) {
	googleOauthConfig := &oauth2.Config{
		RedirectURL:  uc.envs.GOOGLE_REDIRECT_URL.String(),
		ClientID:     uc.envs.GOOGLE_CLIENT_ID.String(),
		ClientSecret: uc.envs.GOOGLE_CLIENT_SECRET.String(),
		Scopes: []string{
			"email",
			"openid",
			"profile",
		},
		Endpoint: google.Endpoint,
	}

	state = generateState()
	redirectUrl = googleOauthConfig.AuthCodeURL(state)

	return
}

func generateState() string {
	b := make([]byte, 16)
	rand.Read(b)

	return base64.URLEncoding.EncodeToString(b)
}
