package main

type CsvHeader struct {
	Date     string `csv:"date"`
	Name     string `csv:"name"`
	Slug     string `csv:"slug"`
	Distance string `csv:"distance"`
	Target   string `csv:"target"`
}

type CsvShotData struct {
	NotUsed01   string  `csv:"none"`
	Time        string  `csv:"time"`
	Id          string  `csv:"id"`
	Tags        string  `csv:"tags"`
	Score       string  `csv:"score"`
	XposMm      float32 `csv:"x (mm)"`
	YposMm      float32 `csv:"y (mm)"`
	XposIn      float32 `csv:"x (inch)"`
	YposIn      float32 `csv:"y (inch)"`
	XposMOA     float32 `csv:"x (moa)"`
	YPosMOA     float32 `csv:"y (moa)"`
	XPosMil     float32 `csv:"x (mil)"`
	YPosMil     float32 `csv:"y (mil)"`
	VelocityMs  float32 `csv:"v (m/s)"`
	VelocityFps float32 `csv:"v (fps)"`
	Yaw         float32 `csv:"yaw (deg)"`
	Pitch       float32 `csv:" pitch (deg)"`
	Quality     float32 `csv:"quality"`
	NotUsed02   string  `csv:"-"`
}

type CsvSummaryData struct {
	NotUsed01   string  `csv:"-"`
	NotUsed02   string  `csv:"-"`
	NotUsed03   string  `csv:"-"`
	NotUsed04   string  `csv:"-"`
	Name        string  `csv:"name"`
	XposMm      float32 `csv:"x (mm)"`
	YposMm      float32 `csv:"y (mm)"`
	XposIn      float32 `csv:"x (inch)"`
	YposIn      float32 `csv:"y (inch)"`
	XposMOA     float32 `csv:"x (moa)"`
	YPosMOA     float32 `csv:"y (moa)"`
	XPosMil     float32 `csv:"x (mil)"`
	YPosMil     float32 `csv:"y (mil)"`
	VelocityMs  float32 `csv:"v (m/s)"`
	VelocityFps float32 `csv:"v (fps)"`
	Yaw         float32 `csv:"yaw (deg)"`
	Pitch       float32 `csv:"pitch (deg)"`
	Quality     float32 `csv:"quality"`
	NotUsed05   string  `csv:"-"`
}
