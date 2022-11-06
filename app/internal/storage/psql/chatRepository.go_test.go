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

type testChat struct {
	chat     model.Chat
	expected model.Chat
}

func TestCreateChat(t *testing.T) {
	ur := storageInt.User()
	usr1, err := ur.Create(context.TODO(), &model.User{UserName: "1", Email: "1", Password: "1"})
	assert.Nil(t, err)
	usr2, err := ur.Create(context.TODO(), &model.User{UserName: "1", Email: "2", Password: "1"})
	assert.Nil(t, err)
	usr3, err := ur.Create(context.TODO(), &model.User{UserName: "1", Email: "3", Password: "1"})
	assert.Nil(t, err)
	cr := storageInt.Chat()
	var testsData = []testChat{
		{
			chat: model.Chat{
				ChatName: "test2", MemberList: []*model.User{
					usr1,
					usr2,
					usr3,
				},
			},
			expected: model.Chat{
				ChatName: "test2", MemberList: []*model.User{
					usr1,
					usr2,
					usr3,
				},
			},
		},
	}

	for i, testData := range testsData {
		cht, err := cr.Create(context.TODO(), &testData.chat)
		log.Println(cht)
		assert.Nil(t, err)
		assert.NotNil(t, cht)
		testData.expected.ID = cht.ID
		testsData[i].expected = testData.expected
		assert.Equal(t, testData.expected, *cht)
	}
	defer func() {
		// storageInt.DB.Exec("DELETE FROM user_chat")
		// storageInt.DB.Exec("DELETE FROM chats")
		// storageInt.DB.Exec("DELETE FROM users")
	}()
}

func TestReadAllChat(t *testing.T) {
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

func TestFindByIdChat(t *testing.T) {
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

func TestUpdateChat(t *testing.T) {
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

func TestFindByEmailChat(t *testing.T) {
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

func TestFindByUserNameChat(t *testing.T) {
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

func TestDeleteChat(t *testing.T) {
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

func TestGetUsersByFilterChat(t *testing.T) {
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
