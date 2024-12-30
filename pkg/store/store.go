package store

type StoreSerivce interface {
	//GetAmount() (int, error)
	UpdateAmount(int) error
	CheckBalance() (int, error)
	CreditAmount(int) error
	DebitAmount(int) error
}

// type Store struct {
// 	FS *FileStorage
// 	PG *Postgres
// }

// func NewStore(fs *FileStorage, pg *Postgres) *Store {
// 	return &Store{FS: fs, PG: pg}
// }

// factory method
func GetStore(str string, filePath string) StoreSerivce {

	pg := NewPG()
	fs := NewFileStorage(filePath)

	if str == postgres {
		return pg
	} else {
		return fs
	}
}
