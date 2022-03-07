package store

import (
	"github.com/Up2110H/e-wallet/pkg/model"
	"time"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, key, identified_at) VALUES ($1, $2, $3) RETURNING id",
		u.Email,
		u.Key,
		u.IdentifiedAt,
	).Scan(&u.ID); err != nil {
		return nil, err
	}

	wallet := &model.Wallet{
		UserId: u.ID,
		Amount: 0,
	}
	_, err := r.store.Wallet().Create(wallet)

	if err != nil {
		return nil, err
	}

	u.Wallet = wallet
	wallet.User = u

	return u, nil
}

func (r *UserRepository) GetById(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, key, identified_at FROM users WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.Email,
		&u.Key,
		&u.IdentifiedAt,
	); err != nil {
		return nil, err
	}

	wallet, err := r.store.Wallet().GetByUser(u)
	if err != nil {
		return nil, err
	}

	u.Wallet = wallet

	return u, nil
}

func (r *UserRepository) GetTransactionsInfo(user *model.User, startDate, endDate time.Time) (int, float64, error) {
	count, total := 0, 0.0
	var err error
	err = r.store.db.QueryRow(
		"SELECT count(*), sum(amount) FROM transactions WHERE wallet_id = $1 AND created_at >= $2 AND created_at < $3",
		user.Wallet.ID,
		startDate,
		endDate,
	).Scan(
		&count,
		&total,
	)

	return count, total, err
}
