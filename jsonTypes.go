package main

type JsonRangeData struct {
	Location      string `json:"location"`       // "Hill Top"
	FiringPoint   string `json:"firingpoint"`    // "A5"
	FpNum         int    `json:"fp_num"`         // 5
	Range         string `json:"range"`          // "500"
	Units         string `json:"units"`          // "M"
	TargetType    string `json:"target_type"`    // "ICFRA"
	TargetTypeNum int    `json:"target_type_no"` // 0
}

type JsonShooterData struct {
	Name string `json:"name"` // "Barney NEILSEN"
	Club string `json:"club"` // "SHRC",
	UIN  string `json:"UIN"`  // "NEI0321",
	Num  string `json:"no"`   // "321"
}

type JsonDisplayData struct {
	Scalefactor  int    `json:"scalefactor"`  // 2,
	LastScotOnly int    `json:"lastshotonly"` // 0,
	Errorcode    int    `json:"errorcode"`    // 0,
	ErrorMessage string `json:"errormsg"`     // ""
}

type JsonShots struct {
	Discipline              string     `json:"discipline"`                  // "tr",
	CalibreRaw              string     `json:"calibre_raw"`                 // "1",
	Calibre                 string     `json:"calibre"`                     // "7.62mm",
	MaxBusShots             int        `json:"max_bus_shots"`               // 10,
	MaxSighters             int        `json:"max_sighters"`                // 2,
	SightersCut             int        `json:"sighters_cut"`                // 2,
	ShotsFired              int        `json:"shots_fired"`                 // 12,
	CountingShots           int        `json:"counting_shots"`              // 10,
	ElapsedTime             string     `json:"elapsed_time"`                // "20:22",
	ElaspsedTimeSecs        int        `json:"elapsed_time_secs"`           // 1222,
	AvgTimeBetweenShots     string     `json:"avg_time_between_shots"`      // "1:51",
	AvgTimeBetweenShotsSecs int        `json:"avg_time_between_shots_secs"` // 111,
	Comp                    []JsonComp `json:"comp"`
	Mpi                     []JsonMpi  `json:"mpi"`
	Stats                   JsonStats  `json:"stats"`
}

type JsonComp struct {
	No []JsonShotData `json:"no"`
}

type JsonShotData struct {
	ShotNo            int     `json:"shotno"`               // 0,
	XPos              float32 `json:"xpos"`                 // -6.9,
	YPos              float32 `json:"ypos"`                 // -158.5,
	Dfc               float32 `json:"dfc"`                  // 158.7,
	Temp              float32 `json:"temp"`                 // 150,
	Value             int     `json:"value"`                // 4,
	Status            int     `json:"status"`               // 1,
	Tof               float32 `json:"tof"`                  // 607,
	Qidx              float32 `json:"qidx"`                 // 3.6,
	Var               float32 `json:"var"`                  // 0.3,
	Time              int64   `json:"time"`                 // 47477784,
	TimeOfShot        string  `json:"time_of_shot"`         // "13:11:17",
	TimeSinceLastShot string  `json:"time_since_last_shot"` // "00:00:00"
}

type JsonMpi struct {
	XPos   float32 `json:"xpos"`   // 7.8,
	YPos   float32 `json:"ypos"`   // 6.1,
	Dia    float32 `json:"dia"`    // 212.1,
	Height float32 `json:"height"` // 101.3
}

type JsonStats struct {
	QiAvgGu float32 `json:"qiAVGu"` // 2.3,
	QiSdU   float32 `json:"qiSDu"`  // 1.8,
	QiAvgC  float32 `json:"qiAVGc"` // 0.3,
	QiSdC   float32 `json:"qiSDc"`  // 0.1,
	TofAvg  float32 `json:"tofAVG"` // 111,
	TofSd   float32 `json:"tofSD"`  // 298
}

type JsonData struct {
	Code        int             `json:"code"`     // 0,
	Format      int             `json:"format"`   // 2,
	Filename    string          `json:"filename"` // "NEI0321_0831_tr_0.json",
	Year        string          `json:"year"`     // "2024",
	Date        string          `json:"date"`     // "31 Aug 2024",
	Stage       string          `json:"stage"`    // "1",
	RangeData   JsonRangeData   `json:"range_data"`
	ShooterData JsonShooterData `json:"shooter_data"`
	DisplayData JsonDisplayData `json:"display_data"`
	Shots       JsonShots       `json:"shots"`
}
