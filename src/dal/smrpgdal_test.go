package dal

import (
	"fmt"
	"testing"
)

// SetUp

func TestCreateDatabaseSchema(t *testing.T) {
	// Test schema creation or migration logic
}

// CRUD Operations
func TestAddUser(t *testing.T) {
	Connect()
	u1 := User{"Lovley", "nicky@gmail.com"}
	if AddUser(&u1) == false {
		t.Errorf("User %v should be able to be added", u1.Email)
	}
}

func TestAddCampaign(t *testing.T) {
	Connect()
	player_emails := [3]string{"user1@example.com", "user2@example.com", "user3@example.com"}

	if UpdateCampaign("TheBestCampaignEver", "test@gmail.com", player_emails) == false {
		t.Errorf("Campaign %v should be able to be added", "TheBestCampaignEver")
	}
}
func TestGetUser(t *testing.T) {
	Connect()
	u1 := GetUser("test@gmail.com")
	if u1 == nil {
		t.Errorf("User information not retrieved %s", "test@gmail.com")
	}

	if u1.Pwd != "Bert" || u1.Email != "test@gmail.com" {
		t.Errorf("User information not retrieved properly %s", u1.Email)
	}

}

func TestGetPassword(t *testing.T) {
	Connect()
	pwd := GetPassword("test@gmail.com")
	if pwd == nil {
		t.Errorf("User password not retrieved %s", "test@gmail.com")
	}

	if *pwd != "Bert" {
		t.Errorf("User information not retrieved properly %s", "test@gmail.com")
	}

	fmt.Printf("PWD is %s\n", *pwd)
}

func TestGetAllHostCampaigns(t *testing.T) {
	test_strings := [3]string{"TheBestCampaignEver", "GnomishAdventure", "AnotherCampaign"}

	Connect()
	c := GetAllHostCampaigns("test@gmail.com")

	for i := 0; i < len(c); i++ {
		if c[i].Name != test_strings[i] {
			t.Errorf("%s != %s", c[i].Name, test_strings[i])
		}
	}
}

func TestGetAllJoinedCampaigns(t *testing.T) {
	test_strings := [2]string{"TheBestCampaignEver", "AnotherCampaign"}

	Connect()
	c := GetAllHostCampaigns("user2@example.com")
	for i := 0; i < len(c); i++ {

		if c[i].Name != test_strings[i] {
			t.Errorf("%s != %s", c[i].Name, test_strings[i])
		}
	}
}

func TestGetAllCampaignSaves(t *testing.T) {
	test_strings := [2]string{"34dc23", "3w465f23456"}

	Connect()
	s := GetAllCampaignSaves("36bb4cb8-73a7-4944-ae38-3394eb6c914f")

	for i := 0; i < len(s); i++ {

		if s[i].Hash != test_strings[i] {
			t.Errorf("%s != %s", s[i].Hash, test_strings[i])
		}
	}

}

func TestReadRecord(t *testing.T) {
	// Test read logic
}

func TestUpdateRecord(t *testing.T) {
	// Test update logic
}

func TestRemoveUser(t *testing.T) {
	Connect()
	u1 := User{"Lovley", "nicky@gmail.com"}

	if RemoveUser(u1.Email) == false {
		t.Errorf("User %v should be able to be removed", u1.Email)
	}
}

func TestRemoveCampaign(t *testing.T) {
	Connect()
	if RemoveCampaign(2, `test@gmail.com`) == false {
		t.Errorf("Campaign %v should be able to be removed", 2)
	}
}

// Search and Filters
func TestFindUser(t *testing.T) {
	Connect()
	b := FindUserEmail("test@gmail.com")

	if b != true {
		t.Errorf("Output %v : User %v does exist", false, "test@gmail.com")
	}

	if FindUserEmail("bsh93@gmail.com") != false {
		t.Errorf("Output %v : User %v does not exist", true, "bsh93@gmail.com")
	}
}
