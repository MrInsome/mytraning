package internal

import (
	"strconv"
	"sync"
	"time"
)

type Account struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    time.Time
	AccountID    int
	Integration  Integration
}

type Integration struct {
	SecretKey        string
	ClientID         int
	RedirectURL      string
	AuthorizationURL string
}

type Database struct {
	accounts map[string]Account
	mutex    sync.RWMutex
}

func NewDatabase() *Database {
	return &Database{
		accounts: make(map[string]Account),
	}
}

func (db *Database) AddAccount(account Account) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.accounts[strconv.Itoa(account.AccountID)] = account
}

func (db *Database) RemoveAccount(accountID string) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	delete(db.accounts, accountID)
}

func (db *Database) GetAccount(accountID string) (Account, bool) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	account, ok := db.accounts[accountID]
	return account, ok
}
