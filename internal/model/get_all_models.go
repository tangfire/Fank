package model

import "Fank/internal/model/account"

func GetAllModels() []interface{} {
	return []interface{}{
		&account.Account{},
	}
}
