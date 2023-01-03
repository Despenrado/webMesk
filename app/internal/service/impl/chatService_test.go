package impl_test

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
	actual   *model.Chat
	expected *model.Chat
}

func prepareData(t *testing.T) []testChat {
	ur := storageInt.User()
	usr1, err := ur.Create(context.TODO(), &model.User{UserName: "1", Email: "1", Password: "1"})
	assert.Nil(t, err)
	usr2, err := ur.Create(context.TODO(), &model.User{UserName: "1", Email: "2", Password: "1"})
	assert.Nil(t, err)
	usr3, err := ur.Create(context.TODO(), &model.User{UserName: "1", Email: "3", Password: "1"})
	assert.Nil(t, err)

	return []testChat{
		{
			actual: &model.Chat{
				ChatName: "test2", MemberList: []*model.User{
					usr1,
					usr2,
					usr3,
				},
			},
			expected: &model.Chat{
				ChatName: "test2", MemberList: []*model.User{
					usr1,
					usr2,
					usr3,
				},
			},
		},
	}

}

func TestCreateChat(t *testing.T) {
	defer func() {
		storageInt.DB.Exec("DELETE FROM user_chat")
		storageInt.DB.Exec("DELETE FROM chats")
		storageInt.DB.Exec("DELETE FROM users")
	}()
	cr := serviceInt.Chat()
	testsData := prepareData(t)

	for _, testData := range testsData {
		cht, err := cr.Create(context.TODO(), testData.actual)
		log.Println(cht)
		assert.Nil(t, err)
		assert.NotNil(t, cht)
		testData.expected.ID = cht.ID
		assert.Equal(t, testData.expected, cht)
	}

}

func TestReadAllChat(t *testing.T) {
	defer func() {
		storageInt.DB.Exec("DELETE FROM user_chat")
		storageInt.DB.Exec("DELETE FROM chats")
		storageInt.DB.Exec("DELETE FROM users")
	}()
	cr := serviceInt.Chat()
	testsData := prepareData(t)

	for _, testData := range testsData {
		actual, err := cr.Create(context.TODO(), testData.actual)
		log.Println(actual)
		testData.expected.ID = actual.ID
		assert.Nil(t, err)
	}

	actuals, err := cr.ReadAll(context.TODO(), 0, 10)
	assert.Nil(t, err)
	assert.NotNil(t, actuals)
	assert.Equal(t, len(actuals), len(testsData))
}

func TestFindByIdChat(t *testing.T) {
	defer func() {
		storageInt.DB.Exec("DELETE FROM user_chat")
		storageInt.DB.Exec("DELETE FROM chats")
		storageInt.DB.Exec("DELETE FROM users")
	}()
	cr := serviceInt.Chat()
	testsData := prepareData(t)

	for _, testData := range testsData {
		usr, err := cr.Create(context.TODO(), testData.actual)
		testData.expected.ID = usr.ID
		assert.Nil(t, err)
	}

	for _, testData := range testsData {
		usr, err := cr.FindById(context.TODO(), testData.expected.ID)
		assert.Nil(t, err)
		assert.NotNil(t, usr)
		compareChats(t, usr, testData.expected)
	}
}

func TestUpdateChat(t *testing.T) {
	defer func() {
		storageInt.DB.Exec("DELETE FROM user_chat")
		storageInt.DB.Exec("DELETE FROM chats")
		storageInt.DB.Exec("DELETE FROM users")
	}()
	cr := serviceInt.Chat()
	testsData := prepareData(t)

	for _, testData := range testsData {
		usr, err := cr.Create(context.TODO(), testData.actual)
		testData.expected.ID = usr.ID
		assert.Nil(t, err)
	}

	for _, testData := range testsData {
		testData.actual.MemberList = testData.actual.MemberList[:len(testData.actual.MemberList)-2]
		testData.expected.MemberList = testData.expected.MemberList[:len(testData.expected.MemberList)-2]
		actual, err := cr.Update(context.TODO(), testData.actual)
		assert.Nil(t, err)
		assert.NotNil(t, actual)
		compareChats(t, actual, testData.expected)
	}
}

func TestFindByUserIdChat(t *testing.T) {
	defer func() {
		storageInt.DB.Exec("DELETE FROM user_chat")
		storageInt.DB.Exec("DELETE FROM chats")
		storageInt.DB.Exec("DELETE FROM users")
	}()
	cr := serviceInt.Chat()
	testsData := prepareData(t)

	for _, testData := range testsData {
		usr, err := cr.Create(context.TODO(), testData.actual)
		testData.expected.ID = usr.ID
		assert.Nil(t, err)
	}

	for _, testData := range testsData {
		actual, err := cr.FindByUserId(context.TODO(), testData.expected.MemberList[0].ID)
		assert.Nil(t, err)
		assert.NotNil(t, actual)
		compareChats(t, &actual[0], testData.expected)
	}
}

func TestDeleteChat(t *testing.T) {
	defer func() {
		storageInt.DB.Exec("DELETE FROM user_chat")
		storageInt.DB.Exec("DELETE FROM chats")
		storageInt.DB.Exec("DELETE FROM users")
	}()
	cr := serviceInt.Chat()
	testsData := prepareData(t)

	for _, testData := range testsData {
		usr, err := cr.Create(context.TODO(), testData.actual)
		testData.expected.ID = usr.ID
		assert.Nil(t, err)
	}

	for _, testData := range testsData {
		err := cr.Delete(context.TODO(), testData.expected.ID)
		assert.Nil(t, err)
		usr, err := cr.FindById(context.TODO(), testData.expected.ID)
		assert.NotNil(t, err)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
		assert.Nil(t, usr)
	}
}

func TestGetUsersByFilterChat(t *testing.T) {
	defer func() {
		storageInt.DB.Exec("DELETE FROM user_chat")
		storageInt.DB.Exec("DELETE FROM chats")
		storageInt.DB.Exec("DELETE FROM users")
	}()
	cr := serviceInt.Chat()
	testsData := prepareData(t)

	for _, testData := range testsData {
		usr, err := cr.Create(context.TODO(), testData.actual)
		testData.expected.ID = usr.ID
		assert.Nil(t, err)
	}

	filter := &model.ChatFilter{
		ChatName: testsData[0].expected.ChatName,
	}
	actual, err := cr.FilterChat(context.TODO(), filter)
	assert.Nil(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, testsData[0].expected.ID, actual[0].ID)
	assert.Equal(t, testsData[0].expected.ChatName, actual[0].ChatName)
}

func compareChats(t *testing.T, actual *model.Chat, expected *model.Chat) {
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.ChatName, actual.ChatName)
	assert.Equal(t, len(expected.MemberList), len(actual.MemberList))
	for _, expMember := range expected.MemberList {
		notFound := true
		expMember.LastOnline = time.Time{}
		for _, actMember := range actual.MemberList {
			actMember.LastOnline = time.Time{}
			if *expMember == *actMember {
				notFound = false
			}
		}
		assert.False(t, notFound)
	}
}
