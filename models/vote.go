package models

import (
	"fmt"
	"log"
)

type Tuse struct {
	TuseHouse_id        int64  `json:"tuse_house_id"`
	Tuse_id             int64  `json:"tuse_id"`
	Thouse_id           int64  `json:"thouse_id"`
	TuseHouse_sumAmount string `json:"tuse_house_sum_amount"`
	TuseHouse_ZJfentan  string `json:"tuse_house_zj_fen_tan"`
	TuseHouse_XJfentan  string `json:"tuse_house_xj_fen_tan"`
	TuseHouse_voteValue int64  `json:"tuse_house_vote_value"`
	Tuse_content        string `json:"tuse_content"`
	Tuse_hezhunAmount   string `json:"tuse_hezhun_amount"`
	Tuse_fentanHouse    string `json:"tuse_fentan_house"`
	Taccount_name       string `json:"taccount_name"`
}

func GetTuses(mobile string) []*Tuse {
	queryString := `select TuseHouse.TuseHouse_id , TuseHouse.Tuse_id, TuseHouse.Thouse_id,
       TuseHouse.TuseHouse_sumAmount,TuseHouse.TuseHouse_ZJfentan,TuseHouse.TuseHouse_XJfentan,
       TuseHouse.TuseHouse_voteValue,Tuse.Tuse_content,Tuse.Tuse_hezhunAmount,Tuse.Tuse_fentanHouse,
Taccount.Taccount_name from TuseHouse
left join Tuse on Tuse.Tuse_id=TuseHouse.Tuse_id
left join Taccount on Taccount.Thouse_id =TuseHouse.Thouse_id
where Tuse.Tuse_voteStatus=1 and Taccount. Taccount_tel = '` + mobile + `'
and TuseHouse.TuseHouse_voteValue<1`
	fmt.Println(queryString)
	rows, err := db.Query(queryString)
	if err != nil {
		log.Fatal("Query failed:", err.Error())
	}
	defer rows.Close()

	var rowsData []*Tuse
	for rows.Next() {
		var row = new(Tuse)
		rows.Scan(
			&row.TuseHouse_id,
			&row.Tuse_id,
			&row.Thouse_id,
			&row.TuseHouse_sumAmount,
			&row.TuseHouse_ZJfentan,
			&row.TuseHouse_XJfentan,
			&row.TuseHouse_voteValue,
			&row.Tuse_content,
			&row.Tuse_hezhunAmount,
			&row.Tuse_fentanHouse,
			&row.Taccount_name,
		)
		rowsData = append(rowsData, row)
	}
	for _, val := range rowsData {
		fmt.Println(val)
	}
	return rowsData
}
