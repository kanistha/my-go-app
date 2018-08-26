package mock

import "mygoapp"

type RepositoryMock struct {

}

func NewRepositoryMock() *RepositoryMock {
	return &RepositoryMock{}
}

func (r *RepositoryMock) Get(userId string) mygoapp.User{
	return mygoapp.User{
		 Id: userId,
		 Name:"Foo Boo",
		 Email:"foo.boo@mail.com",
	}
}
