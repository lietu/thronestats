package thronestats

type ApiResponseRun struct {
	Character     int     `json:"char"`
	LastDamagedBy int     `json:"lasthit"`
	World         int     `json:"world"`
	Area          int     `json:"level"`
	Crown         int     `json:"crown"`
	Weapon1       int     `json:"wepA"`
	Weapon2       int     `json:"wepB"`
	BSkin         int     `json:"skin"`
	Ultra         int     `json:"ultra"`
	CharacterLvl  int     `json:"charlvl"`
	Loop          int     `json:"loops"`
	Win           int     `json:"win"`
	Mutations     string  `json:"mutations"`
	Kills         string  `json:"kills"`
	Health        string  `json:"health"`
	SteamId       int     `json:"steamid"`
	Type          string  `json:"type"`
	Timestamp     int     `json:"timestamp"`
}

type ApiResponse struct {
	Current  ApiResponseRun  `json:"current"`
	Previous ApiResponseRun  `json:"previous"`
}

func (ar ApiResponse) ToRunData() RunData {
	rd := RunData{}
	rd.ReadFromApiResponse(&ar)
	return rd
}
