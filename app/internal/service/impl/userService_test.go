package impl_test

import (
	"context"
	"testing"
	"time"

	"github.com/Despenrado/webMesk/internal/model"
	"gorm.io/gorm"

	"github.com/stretchr/testify/assert"
)

type testUser struct {
	user     *model.User
	expected *model.User
}

func TestCreateUser(t *testing.T) {
	defer storageInt.DB.Exec("DELETE FROM users")
	var testsData = []testUser{
		{
			user:     &model.User{UserName: "1", Email: "1", Password: "1"},
			expected: &model.User{UserName: "1", Email: "1", Password: "", LastOnline: time.Time{}},
		},
		{
			user:     &model.User{UserName: "2", Email: "2", Password: "1"},
			expected: &model.User{UserName: "2", Email: "2", Password: "", LastOnline: time.Time{}},
		},
	}

	ur := serviceInt.User()
	usr, err := ur.Create(context.TODO(), testsData[0].user)
	assert.Nil(t, err)
	assert.NotNil(t, usr)
	testsData[0].expected.ID = usr.ID
	compareUsers(t, usr, testsData[0].expected)
	usr, err = ur.Create(context.TODO(), testsData[0].user)
	assert.NotNil(t, err)
	assert.Nil(t, usr)
	usr, err = ur.Create(context.TODO(), testsData[1].user)
	assert.Nil(t, err)
	assert.NotNil(t, usr)
	testsData[1].expected.ID = usr.ID
	compareUsers(t, usr, testsData[1].expected)

}

func TestReadAllUser(t *testing.T) {
	defer storageInt.DB.Exec("DELETE FROM users")
	ur := serviceInt.User()
	var testsData = []testUser{
		{
			user:     &model.User{UserName: "1", Email: "2", Password: "1", LastOnline: time.Time{}},
			expected: &model.User{UserName: "1", Email: "2", Password: "", LastOnline: time.Time{}},
		},
	}
	for i, testData := range testsData {
		testData.user.BeforeCreate()
		usr, err := ur.Create(context.TODO(), testData.user)
		testData.expected.ID = usr.ID
		testsData[i].expected = testData.expected
		assert.Nil(t, err)
	}

	usrs, err := ur.ReadAll(context.TODO(), 0, 10)
	assert.Nil(t, err)
	assert.NotNil(t, usrs)
	assert.Equal(t, len(usrs), 1)
	for i, _ := range usrs {
		compareUsers(t, &usrs[i], testsData[i].expected)
	}
}

func TestFindByIdUser(t *testing.T) {
	defer storageInt.DB.Exec("DELETE FROM users")
	ur := serviceInt.User()
	var testsData = []testUser{
		{
			user:     &model.User{UserName: "1", Email: "1", Password: "1"},
			expected: &model.User{UserName: "1", Email: "1", Password: "", LastOnline: time.Time{}},
		},
	}
	for i, testData := range testsData {
		testData.user.BeforeCreate()
		usr, err := ur.Create(context.TODO(), testData.user)
		testData.expected.ID = usr.ID
		testsData[i].expected = testData.expected
		assert.Nil(t, err)
	}

	for _, testData := range testsData {
		usr, err := ur.FindById(context.TODO(), testData.expected.ID)
		assert.Nil(t, err)
		assert.NotNil(t, usr)
		compareUsers(t, usr, testData.expected)
	}
}

func TestUpdateUser(t *testing.T) {
	defer storageInt.DB.Exec("DELETE FROM users")
	ur := serviceInt.User()
	var testsData = []testUser{
		{
			user:     &model.User{UserName: "1", Email: "1", Password: "1"},
			expected: &model.User{UserName: "1", Email: "2", Password: ""},
		},
	}
	for i, testData := range testsData {
		usr, err := ur.Create(context.TODO(), testData.user)
		testData.expected.ID = usr.ID
		testsData[i].user = usr
		assert.Nil(t, err)
	}

	for _, testData := range testsData {
		testData.user.Email = "2"
		testData.user.Password = "1"
		usr, err := ur.Update(context.TODO(), testData.user)
		assert.Nil(t, err)
		assert.NotNil(t, usr)
		compareUsers(t, usr, testData.expected)
	}
}

func TestFindByEmailUser(t *testing.T) {
	defer storageInt.DB.Exec("DELETE FROM users")
	ur := serviceInt.User()
	var testsData = []testUser{
		{
			user:     &model.User{UserName: "1", Email: "1", Password: "1"},
			expected: &model.User{UserName: "1", Email: "1", Password: "", LastOnline: time.Time{}},
		},
	}
	for i, testData := range testsData {
		usr, err := ur.Create(context.TODO(), testData.user)
		testData.expected.ID = usr.ID
		testsData[i].expected = testData.expected
		assert.Nil(t, err)
	}

	for _, testData := range testsData {
		usr, err := ur.FindByEmail(context.TODO(), testData.user.Email)
		assert.Nil(t, err)
		assert.NotNil(t, usr)
		compareUsers(t, usr, testData.expected)
	}
}

func TestDeleteUser(t *testing.T) {
	defer storageInt.DB.Exec("DELETE FROM users")
	ur := serviceInt.User()
	var testsData = []testUser{
		{
			user:     &model.User{UserName: "1", Email: "1", Password: "1"},
			expected: &model.User{UserName: "1", Email: "1", Password: "", LastOnline: time.Time{}},
		},
	}
	for i, testData := range testsData {
		usr, err := ur.Create(context.TODO(), testData.user)
		testData.user.ID = usr.ID
		testsData[i].user = testData.user
		testData.expected.ID = usr.ID
		testsData[i].expected = testData.expected
		assert.Nil(t, err)
	}

	for _, testData := range testsData {
		err := ur.Delete(context.TODO(), testData.user.ID)
		assert.Nil(t, err)
		usr, err := ur.FindById(context.TODO(), testData.user.ID)
		assert.NotNil(t, err)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
		assert.Nil(t, usr)
	}
}

func TestGetUsersByFilterUser(t *testing.T) {
	defer storageInt.DB.Exec("DELETE FROM users")
	ur := serviceInt.User()
	var testsData = []testUser{
		{
			user:     &model.User{UserName: "1", Email: "1", Password: "1"},
			expected: &model.User{UserName: "1", Email: "1", Password: ""},
		},
		{
			user:     &model.User{UserName: "2", Email: "2", Password: "12"},
			expected: &model.User{UserName: "2", Email: "2", Password: ""},
		},
	}
	for _, testData := range testsData {
		usr, err := ur.Create(context.TODO(), testData.user)
		testData.expected.ID = usr.ID
		assert.Nil(t, err)
	}
	filter := &model.UserFilter{
		UserName: "1",
	}
	usrs, err := ur.FilterUser(context.TODO(), filter)
	assert.Nil(t, err)
	assert.NotNil(t, usrs)
	compareUsers(t, &usrs[0], testsData[0].expected)
}

func compareUsers(t *testing.T, actual *model.User, expected *model.User) {
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Email, actual.Email)
	assert.Equal(t, expected.UserName, actual.UserName)
	assert.Equal(t, expected.Password, actual.Password)
}
