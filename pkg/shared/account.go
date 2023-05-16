package shared

// Account represents a user account
type Account struct {
	ID        string `json:"ID"`
	Username  string `json:"Username"`
	Password  string `json:"Password"`
	LastLogon string `json:"LastLogon"`
}

type GameData struct {
	ShardName string // The name of the Server
	IPAddress string // The IP Address of the Server
}
