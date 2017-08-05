package model

import "github.com/jinzhu/gorm"

type Challenge struct {
	gorm.Model
	Name    string `json:"name"`
	Points  int    `json:"score"`
	Content string `json:"content"`
	Flag    string
	Solved  bool   `gorm:"-"`
	Hints   []Hint `json:"hint"`
}
type Hint struct {
	gorm.Model
	ChallengeId int
	Name        string
	Content     string
	IsActivated string
}

func (challenge Challenge) GetHints(db *gorm.DB) []Hint {
	hints := make([]Hint, 0, 100)
	sqlDb := db.DB()
	rows, err := sqlDb.Query("select name, content, is_activated, challenge_id from hints where challenge_id = ?", challenge.ID)
	if err == nil {
		for rows.Next() {
			var name, content string
			var challengeId int
			var isActivated string
			rows.Scan(&name, &content, &challengeId, &isActivated)
			hint := Hint{Name: name, ChallengeId: challengeId}
			if len(isActivated) != 0 {
				hint.Content = content
			}
			hints = append(hints, hint)
		}
	}
	//fmt.Print(err)
	return hints
}
