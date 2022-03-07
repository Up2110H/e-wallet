package store

import (
	"github.com/Up2110H/e-wallet/pkg/model"
)

type TransactionRepository struct {
	store *Store
}

func (r *TransactionRepository) Create(t *model.Transaction) (*model.Transaction, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO transactions (wallet_id, amount) VALUES ($1, $2) RETURNING id, created_at",
		t.WalletId,
		t.Amount,
	).Scan(&t.ID, &t.CreatedAt); err != nil {
		return nil, err
	}

	_, err := r.store.db.Exec("UPDATE  wallets SET amount=amount+$1 WHERE id=$2", t.Amount, t.WalletId)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (r *TransactionRepository) GetByWallet(w *model.Wallet) ([]*model.Transaction, error) {
	transactions := make([]*model.Transaction, 0)
	rows, err := r.store.db.Query(
		"SELECT id, wallet_id, amount, created_at FROM transactions WHERE wallet_id = $1",
		w.ID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		transaction := &model.Transaction{}
		if err := rows.Scan(
			&transaction.ID,
			&transaction.WalletId,
			&transaction.Amount,
			&transaction.CreatedAt,
		); err != nil {
			return nil, err
		}

		transaction.Wallet = w

		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	w.Transactions = transactions

	return transactions, nil
}
