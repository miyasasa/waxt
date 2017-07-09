package env

import (
	"Waxt/server/store"
)

type Environment struct {
	CustomerBucket string
	Store          *store.Store
}

func InitEnvironment() *Environment {
	st := store.NewStore("waxt.db")
	st.Open()
	return &Environment{CustomerBucket: "Customer", Store: st}
}
