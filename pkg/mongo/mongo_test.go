package mongo_test

import (
	"gobuyright/pkg/entity"
	"gobuyright/pkg/mongo"
	"gobuyright/pkg/mongo/service"
	"log"
	"reflect"
	"testing"
)

const (
	mongoUrl       = "localhost:27017"
	dbName         = "testDb"
	collectionName = "testCol"
)

func TestServices(t *testing.T) {
	t.Run("IUserService", iUserService)
	t.Run("UsageService", usageService)
	t.Run("TagService", tagService)
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
	t.Run("CreateUsage", createUsage_should_insert_usage_into_mongo)
	t.Run("GetAllUsages", getAllUsages_should_return_all_usages_in_mongo)
}

func createUsage_should_insert_usage_into_mongo(t *testing.T) {
	session := connect()
	defer finishTest(session)
	usageService := service.NewUsageService(session.Copy(), dbName, collectionName)

	testId, testUsageID, testUsageName := "1111", "this_usageID", "this_usageName"
	usage := entity.Usage{
		ID:        testId,
		UsageID:   testUsageID,
		UsageName: testUsageName,
	}

	err := usageService.CreateUsage(&usage)
	if err != nil {
		t.Errorf("Unable to create usage: %s", err)
	}

	results := make([]entity.Usage, 0)
	session.GetCollection(dbName, collectionName).Find(nil).All(&results)

	count := len(results)
	if count != 1 {
		t.Errorf("Incorrect number of results. Expecting 1, got %d", count)
	}

	if results[0].UsageID != usage.UsageID {
		t.Errorf("Wrong usageID. Expected %s, got %s", testUsageID, results[0].UsageID)
	}
	if results[0].UsageName != usage.UsageName {
		t.Errorf("Wrong UsageName. Expected %s, got %s", testUsageName, results[0].UsageName)
	}

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

	testId1, testUsageID1, testUsageName1 := "1", "this_usageID1", "this_usageName1"
	testId2, testUsageID2, testUsageName2 := "2", "this_usageID2", "this_usageName2"
	usage1 := &entity.Usage{
		ID:        testId1,
		UsageID:   testUsageID1,
		UsageName: testUsageName1,
	}
	usage2 := &entity.Usage{
		ID:        testId2,
		UsageID:   testUsageID2,
		UsageName: testUsageName2,
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
	if usages[0].UsageID != testUsageID1 {
		t.Errorf("Wrong usage retrieved. Expected id %s, but got id %s", testUsageID1, usages[0].UsageID)
	}
	if usages[1].UsageID != testUsageID2 {
		t.Errorf("Wrong usage retrieved. Expected id %s, but got id %s", testUsageID2, usages[1].UsageID)
	}

}

func tagService(t *testing.T) {
	t.Run("CreateTag", createTag_should_insert_tag_into_mongo)
}

func createTag_should_insert_tag_into_mongo(t *testing.T) {
	session := connect()
	defer finishTest(session)
	tagService := service.NewTagService(session.Copy(), dbName, collectionName)

	testId, testTagID, testTagName := "1111", "this_tagID", "this_tagName"
	tag := entity.Tag{
		ID:      testId,
		TagID:   testTagID,
		TagName: testTagName,
	}

	err := tagService.CreateTag(&tag)
	if err != nil {
		t.Errorf("Unable to create tag: %s", err)
	}

	results := make([]entity.Tag, 0)
	session.GetCollection(dbName, collectionName).Find(nil).All(&results)

	count := len(results)
	if count != 1 {
		t.Errorf("Incorrect number of results. Expecting 1, got %d", count)
	}

	if results[0].TagID != tag.TagID {
		t.Errorf("Wrong tagID. Expected %s, got %s", testTagID, results[0].TagID)
	}
	if results[0].TagName != tag.TagName {
		t.Errorf("Wrong tagName. Expected %s, got %s", testTagName, results[0].TagID)
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
