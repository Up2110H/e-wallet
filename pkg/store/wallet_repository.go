package store

import (
	"github.com/Up2110H/e-wallet/pkg/model"
)

type WalletRepository struct {
	store *Store
}

func (r *WalletRepository) Create(w *model.Wallet) (*model.Wallet, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO wallets (user_id, amount) VALUES ($1, $2) RETURNING id",
		w.UserId,
		w.Amount,
	).Scan(&w.ID); err != nil {
		return nil, err
	}

	return w, nil
}

func (r *WalletRepository) GetByUser(u *model.User) (*model.Wallet, error) {
	w := &model.Wallet{}
	if err := r.store.db.QueryRow(
		"SELECT id, user_id, amount FROM wallets WHERE id = $1",
		u.ID,
	).Scan(
		&w.ID,
		&w.UserId,
		&w.Amount,
	); err != nil {
		return nil, err
	}

	w.User = u

	return w, nil
}
