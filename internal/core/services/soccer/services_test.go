package soccer_test

import (
	"core-api/cmd/apitest"
	"core-api/internal/core/entities"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
)

var confDir string

func TestMain(m *testing.M) {
	//fmt.Println("INI KEPANGGIL GA SIH?")
	str, _ := os.Getwd()
	confDir = fmt.Sprintf("%s/../../../../", str)
	// Run Test
	exitCode := m.Run()

	// Exit
	os.Exit(exitCode)
}

func TestAddSoccerTeam(t *testing.T) {

	input := entities.FormAddSoccerTeam{
		Name:      "BAMBANG FC",
		CoachName: "Bambang Sejati",
	}
	testApp := apitest.InitializeTestApp(t, confDir)
	defer testApp.MockController.Finish()

	cleaningDB(testApp.DB)
	t.Log("Test Add Soccer Team")
	err := testApp.Services.Soccer.AddSoccerTeam(input)
	if err != nil {
		t.Errorf("Failed to add soccer team: %s", err)
	}
	res, err := getNewlyAddSoccerTeam(testApp.DB)
	if err != nil {
		t.Errorf("Failed to add soccer team: %s", err)
	}
	if res.Name != input.Name {
		t.Errorf("Failed at add soccer team, expected team_name : %s, got : %s", res.Name, input.Name)
	}
	if res.CoachName != input.CoachName {
		t.Errorf("Failed at add soccer team, expected team_name : %s, got : %s", res.CoachName, input.CoachName)
	}

	t.Log("OK")
	cleaningDB(testApp.DB)
}

func TestAddSoccerPlayer(t *testing.T) {
	testDOB, _ := time.Parse("2006-01-02", "1992-10-10")
	input := entities.FormAddSoccerPlayer{
		Fullname:    "Bambang Aduhai",
		DOB:         testDOB,
		StringDOB:   "1992-10-10",
		Age:         28,
		Nationality: "Indonesia",
		Position:    "GK",
		ShirtNumber: "22",
		TeamID:      11,
	}
	testApp := apitest.InitializeTestApp(t, confDir)
	defer testApp.MockController.Finish()

	cleaningDB(testApp.DB)
	initializeDummyData(testApp.DB)
	t.Log("Test Add Soccer Player")
	err := testApp.Services.Soccer.AddSoccerPlayer(input)
	if err != nil {
		t.Errorf("Failed to add soccer player: %s", err)
	}
	res, err := getNewlyAddSoccerPlayer(testApp.DB)
	if err != nil {
		t.Errorf("Failed to add soccer player: %s", err)
	}
	if res.Fullname != input.Fullname {
		t.Errorf("Failed at add soccer team, expected fullname : %s, got : %s", res.Fullname, input.Fullname)
	}
	if res.TeamID != input.TeamID {
		t.Errorf("Failed at add soccer team, expected team_id : %d, got : %d", res.TeamID, input.TeamID)
	}
	if res.ShirtNumber != input.ShirtNumber {
		t.Errorf("Failed at add soccer team, expected shirt_number : %s, got : %s", res.ShirtNumber, input.ShirtNumber)
	}
	if res.Nationality != input.Nationality {
		t.Errorf("Failed at add soccer team, expected nationality : %s, got : %s", res.Nationality, input.Nationality)
	}
	if res.Position != input.Position {
		t.Errorf("Failed at add soccer team, expected position : %s, got : %s", res.Position, input.Position)
	}
	if res.DOB != input.StringDOB {
		t.Errorf("Failed at add soccer team, expected dob : %s, got : %s", res.DOB, input.StringDOB)
	}
	if res.Age != input.Age {
		t.Errorf("Failed at add soccer team, expected age : %d, got : %d", res.Age, input.Age)
	}

	t.Log("OK")
	cleaningDB(testApp.DB)
}

func TestGetSoccerTeamList(t *testing.T) {
	totalExpected := 3
	expectedData := []entities.SoccerTeam{
		{
			ID:               10,
			Name:             "ABC FC",
			CoachName:        "Bapak ABC",
			PlayerAgeAverage: 21,
			TotalPlayer:      5,
		},
		{
			ID:               11,
			Name:             "XYZ FC",
			CoachName:        "Bapak XYZ",
			PlayerAgeAverage: 0,
			TotalPlayer:      0,
		},
		{
			ID:               12,
			Name:             "DEF FC",
			CoachName:        "Bapak DEF",
			PlayerAgeAverage: 21,
			TotalPlayer:      3,
		},
	}

	testApp := apitest.InitializeTestApp(t, confDir)
	defer testApp.MockController.Finish()

	cleaningDB(testApp.DB)
	initializeDummyData(testApp.DB)
	t.Log("Test Get Soccer Team List")
	result, err := testApp.Services.Soccer.GetSoccerTeamList()
	if err != nil {
		t.Errorf("Failed to get soccer team list: %s", err)
	}
	if len(result) != totalExpected {
		t.Errorf("Failed to get soccer team list, expected total row : %d, got : %d", totalExpected, len(result))
	}

	for i := 0; i < totalExpected; i++ {
		if result[i].ID != expectedData[i].ID {
			t.Errorf("Failed to get soccer team list, expected 1st row id : %d, got : %d", expectedData[i].ID, result[i].ID)
		}
		if result[i].Name != expectedData[i].Name {
			t.Errorf("Failed to get soccer team list, expected 1st row team_name : %s, got : %s", expectedData[i].Name, result[i].Name)
		}
		if result[i].CoachName != expectedData[i].CoachName {
			t.Errorf("Failed to get soccer team list, expected 1st row coach_name : %s, got : %s", expectedData[i].CoachName, result[i].CoachName)
		}
		if result[i].PlayerAgeAverage != expectedData[i].PlayerAgeAverage {
			t.Errorf("Failed to get soccer team list, expected 1st row player_age_average : %d, got : %d", expectedData[i].PlayerAgeAverage, result[i].PlayerAgeAverage)
		}
		if result[i].TotalPlayer != expectedData[i].TotalPlayer {
			t.Errorf("Failed to get soccer team list, expected 1st row total_player : %d, got : %d", expectedData[i].TotalPlayer, result[i].TotalPlayer)
		}
	}

	t.Log("OK")
	cleaningDB(testApp.DB)
}

func TestGetSoccerTeamDetail(t *testing.T) {
	expectedSoccerTeam := entities.SoccerTeam{
		ID:               10,
		TotalPlayer:      5,
		Name:             "ABC FC",
		CoachName:        "Bapak ABC",
		PlayerAgeAverage: 21,
	}
	expectedPlayer := []entities.SoccerPlayer{
		{
			ID:          13,
			TeamID:      10,
			Fullname:    "Kojiro Hyuga",
			Age:         21,
			DOB:         "1998-10-10",
			Position:    "AMF",
			Nationality: "Japan",
			ShirtNumber: "9",
		},
		{
			ID:          14,
			TeamID:      10,
			Fullname:    "Tsubasa Ozora",
			Age:         21,
			DOB:         "1998-10-11",
			Position:    "CMF",
			Nationality: "Japan",
			ShirtNumber: "10",
		},
		{
			ID:          15,
			TeamID:      10,
			Fullname:    "Misaki",
			Age:         21,
			DOB:         "1998-10-12",
			Position:    "CMF",
			Nationality: "Japan",
			ShirtNumber: "11",
		},
		{
			ID:          16,
			TeamID:      10,
			Fullname:    "Wakabayashi",
			Age:         21,
			DOB:         "1998-10-13",
			Position:    "GK",
			Nationality: "Japan",
			ShirtNumber: "1",
		},
		{
			ID:          17,
			TeamID:      10,
			Fullname:    "Ishizaki",
			Age:         21,
			DOB:         "1998-10-14",
			Position:    "CWB",
			Nationality: "Japan",
			ShirtNumber: "14",
		},
	}
	testApp := apitest.InitializeTestApp(t, confDir)
	defer testApp.MockController.Finish()

	cleaningDB(testApp.DB)
	initializeDummyData(testApp.DB)
	t.Log("Test Get Soccer Team Detail")
	result, err := testApp.Services.Soccer.GetSoccerTeamDetail(10)
	if err != nil {
		t.Errorf("Failed to get soccer team detail : %s", err)
	}
	if result.SoccerTeam != expectedSoccerTeam {
		t.Errorf("Failed to get soccer team detail, expected soccer_team : %v, got : %v", expectedSoccerTeam, result.SoccerTeam)
	}
	if len(result.Player) != len(expectedPlayer) {
		t.Errorf("Failed to get soccer team detail, expected total_item in player : %d, got : %d", len(expectedPlayer), len(result.Player))
	}

	for i := 0; i < len(expectedPlayer); i++ {
		if result.Player[i] != expectedPlayer[i] {
			t.Errorf("Failed to get soccer team detail, expected player : %v, got : %v", expectedPlayer[i], result.Player[i])
		}
	}

	t.Log("OK")
	cleaningDB(testApp.DB)
}

func getNewlyAddSoccerTeam(db *sqlx.DB) (result entities.SoccerTeam, err error) {
	err = db.Get(&result, "SELECT a.id, a.name, a.coach_name, COUNT(b.id) as total_player, CAST(COALESCE(AVG(b.age),0) as bigint) as player_age_average FROM soccer_team a LEFT JOIN soccer_player b ON a.id = b.team_id WHERE a.id = $1 GROUP BY a.id, a.name, a.coach_name", 10)
	if err != nil {
		return
	}
	return
}

func getNewlyAddSoccerPlayer(db *sqlx.DB) (result entities.SoccerPlayer, err error) {
	err = db.Get(&result, "SELECT id, position, shirt_number, team_id, fullname, CAST(dob as text) dob, age, nationality FROM soccer_player ORDER BY id DESC LIMIT 1")
	if err != nil {
		return
	}
	return
}

func cleaningDB(db *sqlx.DB) {
	_, err := db.Exec("DELETE FROM soccer_player")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("DELETE FROM soccer_team")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("SELECT setval('soccer_player_id_seq',9)")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("SELECT setval('soccer_team_id_seq',9)")
	if err != nil {
		panic(err)
	}
}

func initializeDummyData(db *sqlx.DB) {
	query1 := "INSERT INTO soccer_team "
	query1 += "(name,coach_name,created_at,updated_at)"
	query1 += "VALUES "
	query1 += "('ABC FC','Bapak ABC',CURRENT_TIMESTAMP,CURRENT_TIMESTAMP),"
	query1 += "('XYZ FC','Bapak XYZ',CURRENT_TIMESTAMP,CURRENT_TIMESTAMP),"
	query1 += "('DEF FC','Bapak DEF',CURRENT_TIMESTAMP,CURRENT_TIMESTAMP)"

	_, err := db.Exec(query1)
	if err != nil {
		panic(err)
	}

	query2 := "INSERT INTO soccer_player "
	query2 += "(team_id, fullname, age, dob, nationality, shirt_number, position, created_at, updated_at)"
	query2 += "VALUES "
	query2 += "(12, 'John Pantau', 21, '1998-11-19', 'Italia', '10', 'AMF', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), "
	query2 += "(12, 'Johndy Fandy', 21, '1998-11-10', 'Italia', '11', 'CMF', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), "
	query2 += "(12, 'Kiansantang', 21, '1998-10-10', 'Italia', '14', 'CWB', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), "
	query2 += "(10, 'Kojiro Hyuga', 21, '1998-10-10', 'Japan', '9', 'AMF', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), "
	query2 += "(10, 'Tsubasa Ozora', 21, '1998-10-11', 'Japan', '10', 'CMF', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), "
	query2 += "(10, 'Misaki', 21, '1998-10-12', 'Japan', '11', 'CMF', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), "
	query2 += "(10, 'Wakabayashi', 21, '1998-10-13', 'Japan', '1', 'GK', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), "
	query2 += "(10, 'Ishizaki', 21, '1998-10-14', 'Japan', '14', 'CWB', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"

	_, err = db.Exec(query2)
	if err != nil {
		panic(err)
	}
}
