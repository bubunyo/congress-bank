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

	acc1, ok := user.NewAccount(GhanaPesewas)
	require.True(t, ok)
	fmt.Printf("account details id=%s\n", acc1.Num)

	acc2, ok := user.NewAccount(GhanaPesewas)
	require.True(t, ok)
	fmt.Printf("account details 2 id=%s\n", acc2.Num)

	fmt.Printf("%v", user)
}
