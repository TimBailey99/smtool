package main

type CsvHeader struct {
	Date        string `csv:"date"`
	Name        string `csv:"name"`
	FiringPoint string `csv:"fp"`
	Distance    string `csv:"distance"`
	Target      string `csv:"target"`
	NotUsed01   string `csv:"none"`
	Stage       int
	LookupRow   LookupRow
}

type CsvShotData struct {
	NotUsed01 string  `csv:"none"`
	Time      string  `csv:"time"`
	Tags      string  `csv:"tags"`
	Id        string  `csv:"id"`
	Score     string  `csv:"score"`
	TempC     float32 `csv:"temp C"`
	XposMm    float32 `csv:"x mm"`
	YposMm    float32 `csv:"y mm"`
	Velocity  float32 `csv:"v fps"`
	Yaw       float32 `csv:"yaw deg"`
	Pitch     float32 `csv:" pitch deg"`
	Quality   float32 `csv:"quality"`
	XYError   string  `csv:"xy_err"`
}
