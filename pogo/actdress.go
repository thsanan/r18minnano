package pogo

type Actress struct {
	ActId     int    `json:"act_id" db:"act_id"`
	MinnanoId string `json:"minnano_id" db:"minnano_id"`
	NameEn    string `json:"name_en" db:"name_en"`
	NameJp    string `json:"name_jp" db:"name_jp"`
	Birth     string `json:"birth" db:"birth"`
	Tall      int    `json:"tall" db:"tall"`
	Cup       string `json:"cup" db:"cup"`
	Waist     int    `json:"waist" db:"waist"`
	Hip       int    `json:"hip" db:"hip"`
	Retired   bool   `json:"retired" db:"retired"`
	Display   string `json:"display" db:"display"`
}

type ActTag struct {
	TagId    int    `json:"tag_id" db:"tag_id"`
	ActId    int    `json:"act_id" db:"act_id"`
	TagTitle string `json:"tag_title" db:"tag_title"`
}
