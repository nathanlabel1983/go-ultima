package authentication

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/nathanlabel1983/go-ultima/pkg/shared"
)

// AuthService represents the authentication service
type AuthService struct {
	filepath string           // The path to the data folder inluding filename
	running  bool             // Is the service running?
	accounts []shared.Account // The accounts
}

func NewAuthService(filepath string) *AuthService {
	return &AuthService{
		filepath: filepath,
	}
}

func (s *AuthService) Start() error {
	if s.running {
		return fmt.Errorf("AuthService: already running")
	}
	jsonFile, err := os.Open(s.filepath)
	if err != nil {
		return fmt.Errorf("AuthService: unable to open file: %s", err.Error())
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &s.accounts)
	s.running = true
	return nil
}

func (s *AuthService) Stop() error {
	if !s.running {
		return fmt.Errorf("AuthService: already stopped")
	}
	s.running = false
	return nil
}

func (s *AuthService) AuthAccount(username, password string) bool {
	for _, account := range s.accounts {
		if account.Username == username && account.Password == password {
			return true
		}
	}
	return false
}
