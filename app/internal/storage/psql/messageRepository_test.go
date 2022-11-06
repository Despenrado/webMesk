package psql_test

// type testMessage struct {
// 	message  model.Message
// 	expected model.Message
// }

// func TestCreate(t *testing.T) {
// 	mr := storageInt.Message()
// 	ur := storageInt.User()
// 	cr := storageInt.Chat()
// 	usr, err := ur.Create(context.TODO(), &model.User{UserName: "1", Email: "1", Password: "1"})
// 	assert.Nil(t, err)
// 	err := cr.Create(context.TODO(), &model.Chat{UserName: "1", Email: "1", Password: "1"})
// 	assert.Nil(t, err)
// 	var testsData = []testMessage{
// 		{
// 			message:  model.Message{UserID: usr.ID, ChatID: 1, MessageData: utils.JSONB{"msg": "test"}},
// 			expected: model.Message{UserID: usr.ID, ChatID: 1, MessageData: utils.JSONB{"msg": "test"}},
// 		},
// 	}
// 	for i, testData := range testsData {
// 		msg, err := mr.Create(context.TODO(), &testData.message)
// 		log.Println(msg)
// 		testData.expected.ID = msg.ID
// 		testsData[i].expected = testData.expected
// 		assert.Nil(t, err)
// 		assert.NotNil(t, msg)
// 		assert.Equal(t, testData.expected, *msg)
// 	}
// 	defer storageInt.DB.Exec("DELETE FROM messages")
// 	defer storageInt.DB.Exec("DELETE FROM users")
// }

// func TestReadAll(t *testing.T) {
// 	ur := storageInt.User()
// 	var testsData = []testUser{
// 		{
// 			user:     model.User{UserName: "1", Email: "2", Password: "1"},
// 			expected: model.User{UserName: "1", Email: "2", Password: "1", LastOnline: time.Time{}},
// 		},
// 	}
// 	for i, testData := range testsData {
// 		usr, err := ur.Create(context.TODO(), &testData.user)
// 		log.Println(usr)
// 		testData.expected.ID = usr.ID
// 		testsData[i].expected = testData.expected
// 		assert.Nil(t, err)
// 	}
// 	defer storageInt.DB.Exec("DELETE FROM users")

// 	usrs, err := ur.ReadAll(context.TODO(), 0, 10)
// 	log.Println(usrs)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, usrs)
// 	assert.Equal(t, len(usrs), 1)
// 	assert.Equal(t, []model.User{testsData[0].expected}, usrs)
// }

// func TestFindById(t *testing.T) {
// 	ur := storageInt.User()
// 	var testsData = []testUser{
// 		{
// 			user:     model.User{UserName: "1", Email: "1", Password: "1"},
// 			expected: model.User{UserName: "1", Email: "1", Password: "1", LastOnline: time.Time{}},
// 		},
// 	}
// 	for i, testData := range testsData {
// 		usr, err := ur.Create(context.TODO(), &testData.user)
// 		testData.expected.ID = usr.ID
// 		testsData[i].expected = testData.expected
// 		assert.Nil(t, err)
// 	}
// 	defer storageInt.DB.Exec("DELETE FROM users")

// 	for _, testData := range testsData {
// 		usr, err := ur.FindById(context.TODO(), testData.expected.ID)
// 		assert.Nil(t, err)
// 		assert.NotNil(t, usr)
// 		assert.Equal(t, testData.expected, *usr)
// 	}
// }

// func TestUpdate(t *testing.T) {
// 	ur := storageInt.User()
// 	var testsData = []testUser{
// 		{
// 			user:     model.User{UserName: "1", Email: "1", Password: "1"},
// 			expected: model.User{UserName: "1", Email: "1", Password: "1", LastOnline: time.Time{}},
// 		},
// 	}
// 	for i, testData := range testsData {
// 		usr, err := ur.Create(context.TODO(), &testData.user)
// 		testData.expected.ID = usr.ID
// 		testsData[i].expected = testData.expected
// 		assert.Nil(t, err)
// 	}
// 	defer storageInt.DB.Exec("DELETE FROM users")

// 	for _, testData := range testsData {
// 		testData.user.Email = "2"
// 		usr, err := ur.Update(context.TODO(), &testData.user)
// 		assert.Nil(t, err)
// 		assert.NotNil(t, usr)
// 		assert.NotEqual(t, usr, testData.expected)
// 		assert.Equal(t, &testData.user, usr)
// 	}
// }

// func TestFindByEmail(t *testing.T) {
// 	ur := storageInt.User()
// 	var testsData = []testUser{
// 		{
// 			user:     model.User{UserName: "1", Email: "1", Password: "1"},
// 			expected: model.User{UserName: "1", Email: "1", Password: "1", LastOnline: time.Time{}},
// 		},
// 	}
// 	for i, testData := range testsData {
// 		usr, err := ur.Create(context.TODO(), &testData.user)
// 		testData.expected.ID = usr.ID
// 		testsData[i].expected = testData.expected
// 		assert.Nil(t, err)
// 	}
// 	defer storageInt.DB.Exec("DELETE FROM users")

// 	for _, testData := range testsData {
// 		usr, err := ur.FindByEmail(context.TODO(), testData.user.Email)
// 		assert.Nil(t, err)
// 		assert.NotNil(t, usr)
// 		assert.Equal(t, &testData.expected, usr)
// 	}
// }

// func TestFindByUserName(t *testing.T) {
// 	ur := storageInt.User()
// 	var testsData = []testUser{
// 		{
// 			user:     model.User{UserName: "1", Email: "1", Password: "1"},
// 			expected: model.User{UserName: "1", Email: "1", Password: "1", LastOnline: time.Time{}},
// 		},
// 	}
// 	for i, testData := range testsData {
// 		usr, err := ur.Create(context.TODO(), &testData.user)
// 		testData.expected.ID = usr.ID
// 		testsData[i].expected = testData.expected
// 		assert.Nil(t, err)
// 	}
// 	defer storageInt.DB.Exec("DELETE FROM users")

// 	for _, testData := range testsData {
// 		usr, err := ur.FindByUserName(context.TODO(), testData.user.UserName)
// 		assert.Nil(t, err)
// 		assert.NotNil(t, usr)
// 		assert.Equal(t, &testData.expected, usr)
// 	}
// }

// func TestDelete(t *testing.T) {
// 	ur := storageInt.User()
// 	var testsData = []testUser{
// 		{
// 			user:     model.User{UserName: "1", Email: "1", Password: "1"},
// 			expected: model.User{UserName: "1", Email: "1", Password: "1", LastOnline: time.Time{}},
// 		},
// 	}
// 	for i, testData := range testsData {
// 		usr, err := ur.Create(context.TODO(), &testData.user)
// 		testData.user.ID = usr.ID
// 		testsData[i].user = testData.user
// 		testData.expected.ID = usr.ID
// 		testsData[i].expected = testData.expected
// 		assert.Nil(t, err)
// 	}
// 	// defer storageInt.db.Exec("DELETE FROM users")

// 	for _, testData := range testsData {
// 		err := ur.Delete(context.TODO(), testData.user.ID)
// 		assert.Nil(t, err)
// 		usr, err := ur.FindById(context.TODO(), testData.user.ID)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, gorm.ErrRecordNotFound, err)
// 		assert.Nil(t, usr)
// 	}
// }

// func TestGetUsersByFilter(t *testing.T) {
// 	ur := storageInt.User()
// 	var testsData = []testUser{
// 		{
// 			user:     model.User{UserName: "1", Email: "1", Password: "1"},
// 			expected: model.User{UserName: "1", Email: "1", Password: "1", LastOnline: time.Time{}},
// 		},
// 		{
// 			user:     model.User{UserName: "2", Email: "2", Password: "12"},
// 			expected: model.User{UserName: "2", Email: "2", Password: "12", LastOnline: time.Time{}},
// 		},
// 	}
// 	for i, testData := range testsData {
// 		usr, err := ur.Create(context.TODO(), &testData.user)
// 		testData.expected.ID = usr.ID
// 		testsData[i].expected = testData.expected
// 		assert.Nil(t, err)
// 	}
// 	defer storageInt.DB.Exec("DELETE FROM users")
// 	filter := &model.UserFilter{
// 		UserName: "1",
// 	}
// 	usrs, err := ur.GetUsersByFilter(context.TODO(), filter)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, usrs)
// 	assert.Equal(t, []model.User{testsData[0].expected}, usrs)
// }
