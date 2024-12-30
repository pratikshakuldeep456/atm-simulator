package store

type Postgres struct {
}

// 	UpdateAmount(int) error
// 	CheckBalance() (int, error)
// 	CreditAmount(amount int) error
// 	DebitAmount(amount int) error

func NewPG() *Postgres {
	return &Postgres{}
}

func (pg *Postgres) CheckBalance() (int, error) {

	return 0, nil
}

func (pg *Postgres) UpdateAmount(amt int) error {

	return nil
}

func (pg *Postgres) CreditAmount(amt int) error {
	return nil
}

func (pg *Postgres) DebitAmount(amt int) error {
	return nil
}
