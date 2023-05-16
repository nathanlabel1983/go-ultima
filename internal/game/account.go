package game

import "fmt"

type Account struct {
	id                       int    // The account ID
	connID                   int    // The connection ID
	username                 string // The account username
	seed                     string // The account seed
	major, minor, rev, proto int    // The account version
}

func NewAccount(id, connID int, username string) *Account {
	return &Account{
		id:       id,
		connID:   connID,
		username: username,
	}
}

func (a *Account) Seed(seed string, major, minor, rev, proto int) error {
	if a.IsSeeded() {
		return fmt.Errorf("Account: Account is already seeded")
	}
	a.seed = seed
	a.major = major
	a.minor = minor
	a.rev = rev
	a.proto = proto
	return nil
}

func (a *Account) IsSeeded() bool {
	if a.seed != "" && a.major != 0 && a.minor != 0 && a.rev != 0 && a.proto != 0 {
		return true
	}
	return false
}
