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
	Win           bool    `json:"win"`
	Mutations     string  `json:"mutations"`
	Kills         int     `json:"kills"`
	Health        int     `json:"health"`
	SteamId       int     `json:"steamid"`
	Type          string  `json:"type"`
	Timestamp     int     `json:"timestamp"`
}

type ApiResponse struct {
	Current  *ApiResponseRun  `json:"current"`
	Previous *ApiResponseRun  `json:"previous"`
}

func NewApiResponseRun() *ApiResponseRun {
	arr := ApiResponseRun{
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		false,
		"",
		0,
		0,
		0,
		"",
		0,
	}

	return &arr
}

func NewApiResponse() *ApiResponse {
	ar := ApiResponse{
		NewApiResponseRun(),
		NewApiResponseRun(),
	}

	return &ar
}

func (ar *ApiResponse) ToRunData() *RunDataContainer {
	rdc := NewRunDataContainer()
	rdc.ReadFromApiResponse(ar)
	return rdc
}
