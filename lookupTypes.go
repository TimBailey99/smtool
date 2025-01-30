package main

type LookupRow struct {
	No         string `csv:"no"`
	UIN        string `csv:"uin"`
	Name       string `csv:"name"`
	Discipline string `csv:"discipline"`
	Calibre    string `csv:"calibre"`
	CalibreRaw string `csv:"calibre_raw"`
}
