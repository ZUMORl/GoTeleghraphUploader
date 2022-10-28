package telegraph

import (
	"context"
	"fmt"

	"gitlab.com/toby3d/telegraph"

	"github.com/zumorl/go-tgupload/entities"
	"github.com/zumorl/go-tgupload/services"
)

const (
	telegraphUploadAPI = "https://telegra.ph/upload"
)

var (
	_ services.CDN          = (*Server)(nil)
	_ services.TelegraphAPI = (*Server)(nil)
)

func New() *Server {
	return &Server{}
}

type Server struct {
	account *telegraph.Account
}

func (s *Server) Login(ctx context.Context, acc entities.Account) error {
	tgAcc, err := telegraph.CreateAccount(telegraph.Account{
		AccessToken: acc.AccessToken,
		AuthorURL:   acc.AuthorURL,
		AuthorName:  acc.AuthorName,
		ShortName:   acc.AuthorShortName,
	})
	if err != nil {
		return fmt.Errorf("Failed to create telegraph account")
	}

	if len(acc.AccessToken) != 0 {
		tgAcc.AccessToken = acc.AccessToken
	}

	s.account = tgAcc

	return nil
}