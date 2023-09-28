package congressbank

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bubunyo/congress-bank/store"
)

var (
	src = rand.NewSource(time.Now().UnixNano())

	maxEffort = 50

	ub = store.NewStore[User]()
)

type User struct {
	Id, Name string
	Accounts *store.Store[struct{}]
}

func NewUserbase() store.Store[*User] {
	s := store.Store[*User]{}
	return s
}

func CreateNewUser(name string) (*User, bool) {
	u := &User{Name: name, Accounts: store.NewStore[struct{}]()}

	for i := 0; i < maxEffort; i++ {
		u.Id = fmt.Sprintf("u-%d", genId())
		if ok := ub.Insert(u.Id, *u); ok {
			return u, true
		}
	}

	return nil, false
}

func (u *User) NewAccount(curr CurrencyCode) (*Account, bool) {
	a := &Account{
		Curr:          curr,
		MoneyOrderIds: store.NewStore[struct{}](),
	}

	for i := 0; i < maxEffort; i++ {
		a.Num = fmt.Sprintf("%d", genId())
		if ok := accountBase.Insert(a.Num, *a); ok {
			_ = u.Accounts.Insert(a.Num, struct{}{})
			return a, true
		}
	}

	return nil, false
}
