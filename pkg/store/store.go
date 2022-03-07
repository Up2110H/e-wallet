package store

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

type Store struct {
	config                *Config
	db                    *sql.DB
	userRepository        *UserRepository
	walletRepository      *WalletRepository
	transactionRepository *TransactionRepository
}

func NewStore(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open(os.Getenv("DB_DRIVER"), s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

func (s *Store) Wallet() *WalletRepository {
	if s.walletRepository != nil {
		return s.walletRepository
	}

	s.walletRepository = &WalletRepository{
		store: s,
	}

	return s.walletRepository
}

func (s *Store) Transaction() *TransactionRepository {
	if s.transactionRepository != nil {
		return s.transactionRepository
	}

	s.transactionRepository = &TransactionRepository{
		store: s,
	}

	return s.transactionRepository
}
