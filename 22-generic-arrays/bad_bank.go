package main

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

type Account struct {
	Name    string
	Balance float64
}

func NewBalanceFor(acc Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransaction,
		acc,
	)
}

func applyTransaction(acc Account, transaction Transaction) Account {
	if transaction.From == acc.Name {
		acc.Balance -= transaction.Sum
	}
	if transaction.To == acc.Name {
		acc.Balance += transaction.Sum
	}

	return acc
}
