package domain

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

type DepositRequest struct {
	Amount float64 `json:"amount"`
}

type WithdrawRequest struct {
	Amount float64 `json:"amount"`
}

type BalanceResponse struct {
	Balance float64 `json:"balance"`
}

type AccountCreationResponse struct {
	AccountID int `json:"account_id"`
}
