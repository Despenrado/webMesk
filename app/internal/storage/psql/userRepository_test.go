package psql_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/Despenrado/webMesk/internal/model"
	"gorm.io/gorm"

	"github.com/stretchr/testify/assert"
)

type testUser struct {
	user     model.User
	expected model.User
}

func TestCreateUser(t *testing.T) {
	var testsData = []testUser{
		{
			user:     model.User{UserName: "1", Email: "1", Password: "1"},
			expected: model.User{UserName: "1", Email: "1", Password: "1", LastOnline: time.Time{}},
		},
		{
			user:     model.User{UserName: "1", Email: "2", Password: "1"},
			expected: model.User{UserName: "1", Email: "2", Password: "1", LastOnline: time.Time{}},
		},
	}

	ur := storageInt.User()
	usr, err := ur.Create(context.TODO(), &testsData[0].user)
	assert.Nil(t, err)
	assert.NotNil(t, usr)
	testsData[0].expected.ID = usr.ID
	usr, err = ur.Create(context.TODO(), &testsData[0].user)
	assert.NotNil(t, err)
	assert.Nil(t, usr)
	usr, err = ur.Create(context.TODO(), &testsData[1].user)
	assert.Nil(t, err)
	assert.NotNil(t, usr)
	testsData[1].expected.ID = usr.ID
	assert.Equal(t, testsData[1].expected, *usr)
	defer storageInt.DB.Exec("DELETE FROM users")
}

func TestReadAllUser(t *testing.T) {
	ur := storageInt.User()
	var testsData = []testUser{
		{
			user:     model.User{UserName: "1", Email: "2", Password: "1"},
			expected: model.User{UserName: "1", Email: "2", Password: "1", LastOnline: time.Time{}},
		},
	}
	for i, testData := range testsData {
		usr, err := ur.Create(context.TODO(), &testData.user)
		log.Println(usr)
		testData.expected.ID = usr.ID
		testsData[i].expected = testData.expected
		assert.Nil(t, err)
	}
	defer storageInt.DB.Exec("DELETE FROM users")

	usrs, err := ur.ReadAll(context.TODO(), 0, 10)
	log.Println(usrs)
	assert.Nil(t, err)
	assert.NotNil(t, usrs)
	assert.Equal(t, len(usrs), 1)
	assert.Equal(t, []model.User{testsData[0].expected}, usrs)
}

func TestFindByIdUser(t *testing.T) {
	ur := storageInt.User()
	var testsData = []testUser{
		{
			user:     model.User{UserName: "1", Email: "1", Password: "1"},
			expected: model.User{UserName: "1", Email: "1", Password: "1", LastOnline: time.Time{}},
		},
	}
	for i, testData := range testsData {
		usr, err := ur.Create(context.TODO(), &testData.user)
		testData.expected.ID = usr.ID
		testsData[i].expected = testData.expected
		assert.Nil(t, err)
	}
	defer storageInt.DB.Exec("DELETE FROM users")

	for _, testData := range testsData {
		usr, err := ur.FindById(context.TODO(), testData.expected.ID)
		assert.Nil(t, err)
		assert.NotNil(t, usr)
		assert.Equal(t, testData.expected, *usr)
	}
}

func TestUpdateUser(t *testing.T) {
	ur := storageInt.User()
	var testsData = []testUser{
		{
			user:     model.User{UserName: "1", Email: "1", Password: "1"},
			expected: model.User{UserName: "1", Email: "1", Password: "1", LastOnline: time.Time{}},
		},
	}
	for i, testData := range testsData {
		usr, err := ur.Create(context.TODO(), &testData.user)
		testData.expected.ID = usr.ID
		testsData[i].expected = testData.expected
		assert.Nil(t, err)
	}
	defer storageInt.DB.Exec("DELETE FROM users")

	for _, testData := range testsData {
		testData.user.Email = "2"
		usr, err := ur.Update(context.TODO(), &testData.user)
		assert.Nil(t, err)
		assert.NotNil(t, usr)
		assert.NotEqual(t, usr, testData.expected)
		assert.Equal(t, &testData.user, usr)
	}
}

func TestFindByEmailUser(t *testing.T) {
	ur := storageInt.User()
	var testsData = []testUser{
		{
			user:     model.User{UserName: "1", Email: "1", Password: "1"},
			expected: model.User{UserName: "1", Email: "1", Password: "1", LastOnline: time.Time{}},
		},
	}
	for i, testData := range testsData {
		usr, err := ur.Create(context.TODO(), &testData.user)
		testData.expected.ID = usr.ID
		testsData[i].expected = testData.expected
		assert.Nil(t, err)
	}
	defer storageInt.DB.Exec("DELETE FROM users")

	for _, testData := range testsData {
		usr, err := ur.FindByEmail(context.TODO(), testData.user.Email)
		assert.Nil(t, err)
		assert.NotNil(t, usr)
		assert.Equal(t, &testData.expected, usr)
	}
}

func TestFindByUserNameUser(t *testing.T) {
	ur := storageInt.User()
	var testsData = []testUser{
		{
			user:     model.User{UserName: "1", Email: "1", Password: "1"},
			expected: model.User{UserName: "1", Email: "1", Password: "1", LastOnline: time.Time{}},
		},
	}
	for i, testData := range testsData {
		usr, err := ur.Create(context.TODO(), &testData.user)
		testData.expected.ID = usr.ID
		testsData[i].expected = testData.expected
		assert.Nil(t, err)
	}
	defer storageInt.DB.Exec("DELETE FROM users")

	for _, testData := range testsData {
		usr, err := ur.FindByUserName(context.TODO(), testData.user.UserName)
		assert.Nil(t, err)
		assert.NotNil(t, usr)
		assert.Equal(t, &testData.expected, usr)
	}
}

func TestDeleteUser(t *testing.T) {
	ur := storageInt.User()
	var testsData = []testUser{
		{
			user:     model.User{UserName: "1", Email: "1", Password: "1"},
			expected: model.User{UserName: "1", Email: "1", Password: "1", LastOnline: time.Time{}},
		},
	}
	for i, testData := range testsData {
		usr, err := ur.Create(context.TODO(), &testData.user)
		testData.user.ID = usr.ID
		testsData[i].user = testData.user
		testData.expected.ID = usr.ID
		testsData[i].expected = testData.expected
		assert.Nil(t, err)
	}
	// defer storageInt.db.Exec("DELETE FROM users")

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
	ur := storageInt.User()
	var testsData = []testUser{
		{
			user:     model.User{UserName: "1", Email: "1", Password: "1"},
			expected: model.User{UserName: "1", Email: "1", Password: "1", LastOnline: time.Time{}},
		},
		{
			user:     model.User{UserName: "2", Email: "2", Password: "12"},
			expected: model.User{UserName: "2", Email: "2", Password: "12", LastOnline: time.Time{}},
		},
	}
	for i, testData := range testsData {
		usr, err := ur.Create(context.TODO(), &testData.user)
		testData.expected.ID = usr.ID
		testsData[i].expected = testData.expected
		assert.Nil(t, err)
	}
	defer storageInt.DB.Exec("DELETE FROM users")
	filter := &model.UserFilter{
		UserName: "1",
	}
	usrs, err := ur.GetUsersByFilter(context.TODO(), filter)
	assert.Nil(t, err)
	assert.NotNil(t, usrs)
	assert.Equal(t, []model.User{testsData[0].expected}, usrs)
}
