package repository

import (
	"billing-service/src/handler/transactions/model"
	"database/sql"
	"errors"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) ProcessTransaction(trx model.TransactionRequest) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var balance int64

	err = tx.QueryRow(`
		SELECT balance FROM accounts 
		WHERE id = $1 FOR UPDATE
	`, trx.SenderID).Scan(&balance)

	if err != nil {
		return err
	}

	if balance < trx.Amount {
		return errors.New("insufficient balance")
	}

	_, err = tx.Exec(`
		UPDATE accounts SET balance = balance - $1 WHERE id = $2
	`, trx.Amount, trx.SenderID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		UPDATE accounts SET balance = balance + $1 WHERE id = $2
	`, trx.Amount, trx.ReceiverID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO ledger (account_id, amount, type, reference_id)
		VALUES 
		($1, $2, 'debit', $3),
		($4, $2, 'credit', $3)
	`, trx.SenderID, trx.Amount, trx.ReferenceID, trx.ReceiverID)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *TransactionRepository) GetTransactionSender(userId string, page, limit int) (int, int64, []map[string]interface{}, error) {
	var balance int64

	err := r.DB.QueryRow(`
		SELECT balance FROM accounts WHERE id = $1
	`, userId).Scan(&balance)
	if err != nil {
		return 0, 0, nil, err
	}

	var total int
	err = r.DB.QueryRow(`
		SELECT COUNT(*) FROM ledger WHERE account_id = $1
	`, userId).Scan(&total)
	if err != nil {
		return 0, 0, nil, err
	}

	offset := (page - 1) * limit

	rows, err := r.DB.Query(`
		SELECT amount, type, reference_id
		FROM ledger
		WHERE account_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`, userId, limit, offset)
	if err != nil {
		return 0, 0, nil, err
	}
	defer rows.Close()

	var data []map[string]interface{}

	for rows.Next() {
		var amount int64
		var ttype string
		var ref string

		err := rows.Scan(&amount, &ttype, &ref)
		if err != nil {
			return 0, 0, nil, err
		}

		data = append(data, map[string]interface{}{
			"amount":       amount,
			"type":         ttype,
			"reference_id": ref,
		})
	}

	return total, balance, data, nil
}
