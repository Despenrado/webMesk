package teststore

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/Despenrado/webMesk/internal/model"
	"github.com/stretchr/testify/assert"
)

type testUser struct {
	usr      model.User
	expected model.User
}

var tests = []testUser{
	testUser{
		usr:      model.User{UserName: "1", Email: "1", Password: "1"},
		expected: model.User{}},
}

func helperNewStorage() *Storage {
	os.Setenv("TZ", "UTC")
	return NewStorage()
}

func TestCreate(t *testing.T) {
	ur := helperNewStorage().User()
	for i, tusr := range tests {
		usr, err := ur.Create(context.TODO(), &tusr.usr)
		assert.Nil(t, err)
		assert.NotNil(t, usr)
		fmt.Println(usr)
		tests[i].usr = *usr
		tests[i].expected = *usr
	}
}

func TestReadAll(t *testing.T) {
	ur := helperNewStorage().User()
	for _, tusr := range tests {
		ur.Create(context.TODO(), &tusr.usr)
	}
	usrs, err := ur.ReadAll(context.TODO(), 0, len(tests))
	assert.Nil(t, err)
	assert.NotNil(t, usrs)
	assert.Equal(t, len(usrs), len(tests))
}

func TestFindById(t *testing.T) {
	ur := helperNewStorage().User()
	for i, tusr := range tests {
		t, _ := ur.Create(context.TODO(), &tusr.usr)
		tests[i].usr = *t
		tests[i].expected = *t
	}
	for _, tusr := range tests {
		usr, err := ur.FindById(context.TODO(), tusr.usr.ID)
		assert.Nil(t, err)
		assert.NotNil(t, usr)
		assert.Equal(t, &tusr.expected, usr)
	}
}

func TestUpdate(t *testing.T) {
	ur := helperNewStorage().User()
	for i, tusr := range tests {
		t, _ := ur.Create(context.TODO(), &tusr.usr)
		tests[i].usr = *t
	}
	for _, tusr := range tests {
		tusr.usr.Email = "2"
		usr, err := ur.Update(context.TODO(), &tusr.usr)
		assert.Nil(t, err)
		assert.NotNil(t, usr)
		assert.NotEqual(t, usr, tusr.expected)
		assert.Equal(t, &tusr.usr, usr)
	}
}

func TestFindByEmail(t *testing.T) {
	ur := helperNewStorage().User()
	for i, tusr := range tests {
		t, _ := ur.Create(context.TODO(), &tusr.usr)
		tests[i].expected = *t
	}
	for _, tusr := range tests {
		usr, err := ur.FindByEmail(context.TODO(), tusr.usr.Email)
		assert.Nil(t, err)
		assert.NotNil(t, usr)
		assert.Equal(t, &tusr.expected, usr)
	}
}

func TestDelete(t *testing.T) {
	ur := helperNewStorage().User()
	for i, tusr := range tests {
		t, _ := ur.Create(context.TODO(), &tusr.usr)
		tests[i].usr = *t
	}
	for _, tusr := range tests {
		err := ur.Delete(context.TODO(), tusr.usr.ID)
		assert.Nil(t, err)
	}
}
