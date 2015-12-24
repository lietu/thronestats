package thronestats

type ApiResponse struct {
	Health        string		`json:"heal"`
	BSkin         string        `json:"bskn"`
	Character     string        `json:"char"`
	Mutations     string        `json:"muts"`
	Area          string        `json:"area"`
	Crown         string        `json:"crow"`
	World         string        `json:"worl"`
	Loop          string        `json:"loop"`
	Weapon1       string        `json:"wep1"`
	Weapon2       string        `json:"wep2"`
	RunTime       string        `json:"time"`
	Kills         string        `json:"kill"`
	IsDaily       string        `json:"dail"`
	IsWeekly      string        `json:"week"`
	Ultra         string        `json:"ultr"`
	LastDamagedBy string    	`json:"dead"`
}

func (ar ApiResponse) ToRunData() RunData {
	rd := RunData{}
	rd.ReadFromApiResponse(&ar)
	return rd
}
