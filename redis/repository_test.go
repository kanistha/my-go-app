package redis

import (
	"testing"
	"mygoapp/asserts"
)

func TestGet_success(t *testing.T) {

	userId := "123"
	repo := NewRepository()

	user,_ := repo.Get(userId)

	asserts.Equals(t, true, userId == user.Id )



}

