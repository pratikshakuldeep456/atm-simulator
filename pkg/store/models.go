package store

const (
	postgres    = "postgres"
	fileStorage = "file_storage"
)

type Amount struct {
	Balance  int    `json:"balance"`
	Currency string `json:"currency"`
}
