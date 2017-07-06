package env

import (
	"Waxt/server/store"
)

type Environment struct {
	Db *store.Store
}

func InitEnvironment() *Environment {
	st := store.NewStore("waxt.db")
	return &Environment{Db: st}
}
