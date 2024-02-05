package datasources

import (
	"database/sql"

	"github.com/patrick0806/my-bank/internal/entities"
)

type AccountDatasource struct {
	DB *sql.DB
}

func (ad *AccountDatasource) FindByCPF(cpf string) (*entities.Account, error) {
	statement, err := ad.DB.Prepare("SELECT id, balance, name, birthdate, cpf, email, password, phone_number FROM accounts WHERE cpf = $1")
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	var account entities.Account

	err = statement.
		QueryRow(cpf).
		Scan(&account.Id, &account.Balance, &account.Name, &account.Birthdate, &account.CPF, &account.Email, &account.Password, &account.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &account, nil
}

func (ad *AccountDatasource) CreateAccount(account *entities.Account) error {
	statement, err := ad.DB.
		Prepare("INSERT INTO accounts (id, balance, name, birthdate, cpf, email, password, phone_number) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(account.Id, account.Balance, account.Name, account.Birthdate, account.CPF, account.Email, account.Password, account.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}
