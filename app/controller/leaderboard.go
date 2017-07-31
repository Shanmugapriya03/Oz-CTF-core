package controller

import (
	"net/http"

	"github.com/Shivakishore14/Oz-CTF-core/app/model"
)

func Leaderboard(w http.ResponseWriter, r *http.Request) {
	query := "select users.user_name, users.name, sum(points) as point_sum, max(timestamp) as max_time from submissions LEFT JOIN users on users.id=submissions.user_id group by submissions.user_id order by point_sum desc, max_time asc;"
	sdb := db.DB()
	rows, err := sdb.Query(query)
	if err != nil {
		webresponse("error", err, nil, w)
		return
	}
	list := make([]model.LeaderboardItem, 0, 100)
	for rows.Next() {
		var username, name, timestamp string
		var points uint
		rows.Scan(&username, &name, &points, &timestamp)
		//fmt.Println(username, points, timestamp)
		obj := model.LeaderboardItem{UserName: username, Name: name, Points: points, Timestamp: timestamp}
		list = append(list, obj)
	}
	webresponse("success", nil, list, w)
}

// SCHEMA
// for score based on userid
// for score based on Challenge
// userid, challengeid, score, timestamp
