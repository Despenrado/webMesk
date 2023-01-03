package impl_test

import (
	"context"
	"testing"
	"time"

	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/pkg/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type testMessage struct {
	actual   *model.Message
	expected *model.Message
}

func prepareDataMessage(t *testing.T) []testMessage {
	ur := storageInt.User()
	usr1, err := ur.Create(context.TODO(), &model.User{UserName: "1", Email: "1", Password: "1"})
	assert.Nil(t, err)
	usr2, err := ur.Create(context.TODO(), &model.User{UserName: "1", Email: "2", Password: "1"})
	assert.Nil(t, err)
	usr3, err := ur.Create(context.TODO(), &model.User{UserName: "1", Email: "3", Password: "1"})
	assert.Nil(t, err)

	chats := []*model.Chat{
		&model.Chat{
			ChatName: "test2", MemberList: []*model.User{
				usr1,
				usr2,
				usr3,
			},
		},
		&model.Chat{
			ChatName: "test3", MemberList: []*model.User{
				usr1,
				usr2,
			},
		},
	}

	cr := storageInt.Chat()
	for i, chat := range chats {
		chats[i], err = cr.Create(context.TODO(), chat)
		assert.Nil(t, err)
	}

	return []testMessage{
		{
			actual: &model.Message{
				UserID: usr1.ID,
				ChatID: chats[0].ID,
				MessageData: utils.JSONB{
					"text": "msg_1",
				},
			},
			expected: &model.Message{
				UserID: usr1.ID,
				ChatID: chats[0].ID,
				MessageData: utils.JSONB{
					"text": "msg_1",
				},
			},
		},
		{
			actual: &model.Message{
				UserID: usr1.ID,
				ChatID: chats[1].ID,
				MessageData: utils.JSONB{
					"text": "msg_2",
				},
			},
			expected: &model.Message{
				UserID: usr1.ID,
				ChatID: chats[1].ID,
				MessageData: utils.JSONB{
					"text": "msg_2",
				},
			},
		},
	}
}

func TestCreateMessage(t *testing.T) {
	defer func() {
		storageInt.DB.Exec("DELETE FROM user_chat")
		storageInt.DB.Exec("DELETE FROM chats")
		storageInt.DB.Exec("DELETE FROM users")
		storageInt.DB.Exec("DELETE FROM messages")
	}()
	mr := serviceInt.Message()
	testsData := prepareDataMessage(t)

	for _, testData := range testsData {
		msg, err := mr.Create(context.TODO(), testData.actual)
		assert.Nil(t, err)
		assert.NotNil(t, msg)
		msg.DateTime = time.Time{}
		testData.expected.ID = msg.ID
		assert.Equal(t, testData.expected, msg)
	}

}

func TestReadAllMessages(t *testing.T) {
	defer func() {
		storageInt.DB.Exec("DELETE FROM user_chat")
		storageInt.DB.Exec("DELETE FROM chats")
		storageInt.DB.Exec("DELETE FROM users")
		storageInt.DB.Exec("DELETE FROM messages")
	}()
	mr := serviceInt.Message()
	testsData := prepareDataMessage(t)

	for _, testData := range testsData {
		actual, err := mr.Create(context.TODO(), testData.actual)
		testData.expected.ID = actual.ID
		assert.Nil(t, err)
	}

	actuals, err := mr.ReadAll(context.TODO(), 0, 10)
	assert.Nil(t, err)
	assert.NotNil(t, actuals)
	assert.Equal(t, len(actuals), len(testsData))
}

func TestFindByIdMessages(t *testing.T) {
	defer func() {
		storageInt.DB.Exec("DELETE FROM user_chat")
		storageInt.DB.Exec("DELETE FROM chats")
		storageInt.DB.Exec("DELETE FROM users")
		storageInt.DB.Exec("DELETE FROM messages")
	}()
	mr := serviceInt.Message()
	testsData := prepareDataMessage(t)

	for _, testData := range testsData {
		usr, err := mr.Create(context.TODO(), testData.actual)
		testData.expected.ID = usr.ID
		assert.Nil(t, err)
	}

	for _, testData := range testsData {
		msg, err := mr.FindById(context.TODO(), testData.expected.ID)
		assert.Nil(t, err)
		assert.NotNil(t, msg)
		msg.DateTime = time.Time{}
		assert.Equal(t, testData.expected.ChatID, msg.ChatID)
		assert.Equal(t, testData.expected.ID, msg.ID)
		assert.Equal(t, testData.expected.ReadBy, msg.ReadBy)
		assert.Equal(t, testData.expected.UserID, msg.UserID)
		assert.Equal(t, testData.expected.MessageData, msg.MessageData)
	}
}

func TestUpdateMessage(t *testing.T) {
	defer func() {
		storageInt.DB.Exec("DELETE FROM user_chat")
		storageInt.DB.Exec("DELETE FROM chats")
		storageInt.DB.Exec("DELETE FROM users")
		storageInt.DB.Exec("DELETE FROM messages")
	}()
	mr := serviceInt.Message()
	testsData := prepareDataMessage(t)

	for _, testData := range testsData {
		usr, err := mr.Create(context.TODO(), testData.actual)
		testData.expected.ID = usr.ID
		assert.Nil(t, err)
	}

	for _, testData := range testsData {
		testData.actual.MessageData["text"] = "updated"
		testData.expected.MessageData["text"] = "updated"
		actual, err := mr.Update(context.TODO(), testData.actual)
		assert.Nil(t, err)
		assert.NotNil(t, actual)
		assert.Equal(t, testData.expected.ChatID, actual.ChatID)
		assert.Equal(t, testData.expected.ID, actual.ID)
		assert.Equal(t, testData.expected.ReadBy, actual.ReadBy)
		assert.Equal(t, testData.expected.UserID, actual.UserID)
		assert.Equal(t, testData.expected.MessageData, actual.MessageData)
	}
}

func TestDeleteMessage(t *testing.T) {
	defer func() {
		storageInt.DB.Exec("DELETE FROM user_chat")
		storageInt.DB.Exec("DELETE FROM chats")
		storageInt.DB.Exec("DELETE FROM users")
	}()
	mr := serviceInt.Message()
	testsData := prepareDataMessage(t)

	for _, testData := range testsData {
		usr, err := mr.Create(context.TODO(), testData.actual)
		testData.expected.ID = usr.ID
		assert.Nil(t, err)
	}

	for _, testData := range testsData {
		err := mr.Delete(context.TODO(), testData.expected)
		assert.Nil(t, err)
		usr, err := mr.FindById(context.TODO(), testData.expected.ID)
		assert.NotNil(t, err)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
		assert.Nil(t, usr)
	}
}

func TestGetUsersByFilterMessage(t *testing.T) {
	defer func() {
		storageInt.DB.Exec("DELETE FROM user_chat")
		storageInt.DB.Exec("DELETE FROM chats")
		storageInt.DB.Exec("DELETE FROM users")
	}()
	mr := serviceInt.Message()
	testsData := prepareDataMessage(t)

	for _, testData := range testsData {
		usr, err := mr.Create(context.TODO(), testData.actual)
		testData.expected.ID = usr.ID
		assert.Nil(t, err)
	}

	filter := &model.MessageFilter{
		UserID: testsData[0].expected.UserID,
	}
	actual, err := mr.FilterMessage(context.TODO(), filter)
	assert.Nil(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, testsData[0].expected.ChatID, actual[0].ChatID)
	assert.Equal(t, testsData[0].expected.ID, actual[0].ID)
	assert.Equal(t, testsData[0].expected.ReadBy, actual[0].ReadBy)
	assert.Equal(t, testsData[0].expected.UserID, actual[0].UserID)
	assert.Equal(t, testsData[0].expected.MessageData, actual[0].MessageData)
}

func TestMarkAsReadMessage(t *testing.T) {
	defer func() {
		storageInt.DB.Exec("DELETE FROM user_chat")
		storageInt.DB.Exec("DELETE FROM chats")
		storageInt.DB.Exec("DELETE FROM users")
	}()
	mr := serviceInt.Message()
	testsData := prepareDataMessage(t)

	for _, testData := range testsData {
		usr, err := mr.Create(context.TODO(), testData.actual)
		testData.expected.ID = usr.ID
		assert.Nil(t, err)
	}

	err := mr.MarkAsRead(context.TODO(), testsData[0].expected.ID, testsData[0].expected.UserID)
	assert.Nil(t, err)
	msg, err := mr.FindById(context.TODO(), testsData[0].expected.ID)
	assert.Nil(t, err)
	assert.NotNil(t, msg)
	assert.Equal(t, testsData[0].expected.ChatID, msg.ChatID)
	assert.Equal(t, testsData[0].expected.ID, msg.ID)
	assert.Equal(t, testsData[0].expected.UserID, msg.UserID)
	assert.Equal(t, testsData[0].expected.MessageData, msg.MessageData)
}
