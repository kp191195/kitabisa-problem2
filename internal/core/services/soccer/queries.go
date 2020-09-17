package soccer

const (
	getSoccerTeamList       = "SELECT a.id, a.name, a.coach_name, COUNT(b.id) as total_player, CAST(COALESCE(AVG(b.age),0) as bigint) as player_age_average FROM soccer_team a LEFT JOIN soccer_player b ON a.id = b.team_id GROUP BY a.id, a.name, a.coach_name ORDER BY a.id"
	getSoccerTeamById       = "SELECT a.id, a.name, a.coach_name, COUNT(b.id) as total_player, CAST(COALESCE(AVG(b.age),0) as bigint) as player_age_average FROM soccer_team a LEFT JOIN soccer_player b ON a.id = b.team_id WHERE a.id = $1 GROUP BY a.id, a.name, a.coach_name"
	getSoccerPlayerByTeamId = "SELECT id, position, shirt_number, team_id, fullname, CAST(dob as text) dob, age, nationality FROM soccer_player WHERE team_id = $1 ORDER BY id"
	addSoccerTeam           = "INSERT INTO soccer_team (name, coach_name, created_at, updated_at) VALUES (:name,:coach_name,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP)"
	addSoccerPlayer         = "INSERT INTO soccer_player (team_id, fullname, dob, age, nationality, position, shirt_number, created_at, updated_at) VALUES (:team_id, :fullname, :dob, :age, :nationality, :position, :shirt_number, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
)
