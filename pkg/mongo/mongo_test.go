package mongo_test

import (
	"log"
	"reflect"
	"testing"

	"github.com/jcasado94/gobuyright/pkg/entity"
	"github.com/jcasado94/gobuyright/pkg/mongo"
	"github.com/jcasado94/gobuyright/pkg/mongo/service"
)

const (
	mongoUrl       = "localhost:27017"
	dbName         = "testDb"
	collectionName = "testCol"
)

func TestServices(t *testing.T) {
	t.Run("IUserService", iUserService)
	t.Run("UsageService", usageService)
	t.Run("UsageSelectionService", usageSelectionService)
}

func iUserService(t *testing.T) {
	t.Run("CreateIUser", createIUser_should_insert_user_into_mongo)
}

func createIUser_should_insert_user_into_mongo(t *testing.T) {
	session, err := mongo.NewSession(mongoUrl)
	if err != nil {
		log.Fatalf("Unable to connect to mongo: %s", err)
	}
	defer finishTest(session)
	iUserService := service.NewIUserService(session.Copy(), dbName, collectionName)

	testId, testUsername := "1111", "super_username"
	user := entity.IUser{
		ID:       testId,
		Username: testUsername,
	}

	err = iUserService.CreateUser(&user)
	if err != nil {
		t.Errorf("Unable to create user: %s", err)
	}

	results := make([]entity.IUser, 0)
	session.GetCollection(dbName, collectionName).Find(nil).All(&results)

	count := len(results)
	if count != 1 {
		t.Errorf("Incorrect number of results. Expecting 1, got %d", count)
	}

	if results[0].Username != user.Username {
		t.Errorf("Wrong username. Expected %s, got %s", testUsername, results[0].Username)
	}
}

func usageService(t *testing.T) {
	t.Run("GetAllUsages", getAllUsages_should_return_all_usages_in_mongo)
}

func getAllUsages_should_return_all_usages_in_mongo(t *testing.T) {
	session := connect()
	defer finishTest(session)
	usageService := service.NewUsageService(session.Copy(), dbName, collectionName)

	usages, err := usageService.GetAllUsages()
	if err != nil {
		t.Error("Coulddd not retrieve usages.")
	}
	if len(usages) != 0 {
		t.Errorf("Retrieved wrong number of usages. Expected %d, but got %d", 0, len(usages))
	}

	testId1, testUsageName1 := "1", "this_usageName1"
	testId2, testUsageName2 := "2", "this_usageName2"
	usage1 := &entity.Usage{
		ID:   testId1,
		Name: testUsageName1,
	}
	usage2 := &entity.Usage{
		ID:   testId2,
		Name: testUsageName2,
	}
	usageService.CreateUsage(usage1)
	usageService.CreateUsage(usage2)

	usages, err = usageService.GetAllUsages()
	if err != nil {
		t.Error("Could not retreive usages.")
	}
	if len(usages) != 2 {
		t.Errorf("Retrieved wrong number of usages. Expected %d, but got %d", 0, len(usages))
	}
	if usages[0].Name != testUsageName1 {
		t.Errorf("Wrong usage retrieved. Expected %s, but got %s", testUsageName1, usages[0].Name)
	}
	if usages[1].Name != testUsageName2 {
		t.Errorf("Wrong usage retrieved. Expected %s, but got %s", testUsageName2, usages[1].Name)
	}

}

func usageSelectionService(t *testing.T) {
	t.Run("CreateUsageSelection", createUsageSelection_should_insert_tag_into_mongo)
	t.Run("GetByUsernameAndTags", getByUsernameAndTags_should_retrieve_the_right_UsageSelection)
}

func createUsageSelection_should_insert_tag_into_mongo(t *testing.T) {
	session := connect()
	defer finishTest(session)
	usService := service.NewUsageSelectionService(session.Copy(), dbName, collectionName)

	testId, testUsername, testTagIDs, testUsageIDs := "1111", "user", []string{"tag1", "tag2"}, []string{"usage1", "usage2", "usage3"}
	usageSelection := entity.UsageSelection{
		ID:       testId,
		Username: testUsername,
		TagIDs:   testTagIDs,
		UsageIDs: testUsageIDs,
	}

	err := usService.CreateUsageSelection(&usageSelection)
	if err != nil {
		t.Errorf("Unable to create usage selection: %s", err)
	}

	results := make([]entity.UsageSelection, 0)
	session.GetCollection(dbName, collectionName).Find(nil).All(&results)

	count := len(results)
	if count != 1 {
		t.Errorf("Incorrect number of results. Expecting 1, got %d", count)
	}

	if results[0].Username != usageSelection.Username {
		t.Errorf("Wrong username. Expected %s, got %s", testUsername, results[0].Username)
	}
	if !reflect.DeepEqual(results[0].TagIDs, usageSelection.TagIDs) {
		t.Errorf("Wrong tagIDs. Expected %s, got %s", testTagIDs, results[0].TagIDs)
	}
}

func getByUsernameAndTags_should_retrieve_the_right_UsageSelection(t *testing.T) {
	session := connect()
	defer finishTest(session)
	usService := service.NewUsageSelectionService(session.Copy(), dbName, collectionName)

	testUsername1, testTagIDs1, testUsageIDs1 := "user1", []string{"tag1", "tag2"}, []string{"usage1", "usage2", "usage3"}
	testUsername2, testTagIDs2, testUsageIDs2 := "user1", []string{"tag3", "tag4"}, []string{"usage2", "usage3"}
	testUsername3, testTagIDs3, testUsageIDs3 := "user2", []string{"tag3", "tag4"}, []string{"usage2", "usage1"}
	usageSelection1, usageSelection2, usageSelection3 :=
		entity.UsageSelection{
			Username: testUsername1,
			TagIDs:   testTagIDs1,
			UsageIDs: testUsageIDs1,
		}, entity.UsageSelection{
			Username: testUsername2,
			TagIDs:   testTagIDs2,
			UsageIDs: testUsageIDs2,
		}, entity.UsageSelection{
			Username: testUsername3,
			TagIDs:   testTagIDs3,
			UsageIDs: testUsageIDs3,
		}

	err := usService.CreateUsageSelection(&usageSelection1)
	err = usService.CreateUsageSelection(&usageSelection2)
	err = usService.CreateUsageSelection(&usageSelection3)
	if err != nil {
		t.Errorf("Unable to create usage selections: %s", err)
	}

	result1, err := usService.GetByUsernameAndTags("user1", []string{"tag3", "tag4"})
	if err != nil {
		t.Errorf("Error while querying. %s", err)
	}
	result2, err := usService.GetByUsernameAndTags("user1", []string{"tag1", "tag2"})
	if err != nil {
		t.Errorf("Error while querying. %s", err)
	}

	usageSelection2.ID, usageSelection1.ID = result1.ID, result2.ID
	if !reflect.DeepEqual(*result1, usageSelection2) {
		t.Errorf("First query failed. Queried [%s, [%s, %s]], obtained %v", "user1", "tag3", "tag4", result1)
	}
	if !reflect.DeepEqual(*result2, usageSelection1) {
		t.Errorf("First query failed. Queried [%s, [%s, %s]], obtained %v", "user1", "tag1", "tag2", result2)
	}

}

func itemService(t *testing.T) {
	t.Run("GetAllItems", getAllItems_should_return_all_items_in_mongo)
}

func getAllItems_should_return_all_items_in_mongo(t *testing.T) {
	session := connect()
	defer finishTest(session)
	itemService := service.NewItemService(session.Copy(), dbName, collectionName)

	items, err := itemService.GetAllItems()
	if err != nil {
		t.Error("Coulddd not retrieve items.")
	}
	if len(items) != 0 {
		t.Errorf("Retrieved wrong number of items. Expected %d, but got %d", 0, len(items))
	}

	testId1, testItemName1, testItemImg1 := "1", "itemName1", "http://google.com/kewtImg.jpg"
	testId2, testItemName2, testItemImg2 := "2", "itemName2", "http://google.com/kewtImg.jpg"
	item1 := &entity.Item{
		ID:   testId1,
		Name: testItemName1,
		Img:  testItemImg1,
	}
	item2 := &entity.Item{
		ID:   testId2,
		Name: testItemName1,
		Img:  testItemImg2,
	}
	itemService.CreateItem(item1)
	itemService.CreateItem(item2)

	items, err = itemService.GetAllItems()
	if err != nil {
		t.Error("Could not retreive items.")
	}
	if len(items) != 2 {
		t.Errorf("Retrieved wrong number of items. Expected %d, but got %d", 0, len(items))
	}
	if items[0].Name != testItemName1 {
		t.Errorf("Wrong usage retrieved. Expected %s, but got %s", testItemName1, items[0].Name)
	}
	if items[1].Name != testItemName2 {
		t.Errorf("Wrong usage retrieved. Expected %s, but got %s", testItemName2, items[1].Name)
	}

}

func connect() *mongo.Session {
	session, err := mongo.NewSession(mongoUrl)
	if err != nil {
		log.Fatalf("Unable to connect to mongo %s", err)
	}
	return session
}

func finishTest(s *mongo.Session) {
	s.DropDatabase(dbName)
	s.Close()
}
