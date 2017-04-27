package openfpt

import (
	"github.com/nguyendangminh/openfpt/text2speech"
)

const TIMEOUT int64 = 10 // in seconds

func NewApplication(token string) (*Application, error) {
	if token == "" {
		return nil, ErrInvalidToken
	}
	
	return &Application{
		token: token,
		text2speechConfig: text2speech.DefaultConfig,
	}, nil
}