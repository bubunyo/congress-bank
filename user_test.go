package congressbank

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserCreation(t *testing.T) {
	user, ok := CreateNewUser("Bubu")
	assert.True(t, ok)

	fmt.Printf("user details name=%s, user_id=%s\n", user.Name, user.Id)

	acc, ok := user.NewAccount(Currency{})
	require.True(t, ok)
	fmt.Printf("account details id=%s\n", acc.Num)

	acc, ok = user.NewAccount(Currency{})
	require.True(t, ok)
	fmt.Printf("account details 2 id=%s\n", acc.Num)

	fmt.Printf("%v", user)

	// for k, _ := range user.Accounts {
	// 	acc := accountBase.Get(k)
	// }

	/*
			  acc = 12345
				user.Acc.Get(12345)
				// list
				for k,_ =  us.acc{
				if k == 12345
				return true
				_, ok:= user.Acc[1234]
				return ok
				}
		HahSet<Object>


	*/
}
