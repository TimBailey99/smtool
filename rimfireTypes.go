package main

type RimfireRow struct {
	Name       string `csv:"Name"`
	Six        string `csv:"6"`
	SixTotal   string `csv:"-"`
	Five       string `csv:"5"`
	FiveTotal  string `csv:"-"`
	Four       string `csv:"4"`
	FourTotal  string `csv:"-"`
	Three      string `csv:"3"`
	ThreeTotal string `csv:"-"`
	Two        string `csv:"2"`
	TwoTotal   string `csv:"-"`
	One        string `csv:"1"`
	OneTotal   string `csv:"-"`
	Blank      string `csv:"-"`
	Total      string `csv:"TOTAL"`
	Vs         string `csv:"Vs"`
	Shots      string `csv:"Shots"`
}
