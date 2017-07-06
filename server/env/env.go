package env

import (
	"Waxt/server/store"
)

type Environment struct {
	Store *store.Store
}

func InitEnvironment() *Environment {
	st := store.NewStore("waxt.db")
	st.Open()
	return &Environment{Store: st}
}
