package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/samber/lo"
)

type CsvSection struct {
	Header  CsvHeader
	Shots   []*CsvShotData
	Summary []*CsvSummaryData
}

func exportOzScore(section CsvSection, outputFolder string) {

	// Create OZScore compatible JSON
	const shortForm = "Jan 02 2006"
	jDate, _ := time.Parse(shortForm, section.Header.Date)

	jsonFileName := fmt.Sprintf("%s_%s_tr_0.json", section.Header.Name, jDate.Format("0102"))

	j := JsonData{}
	j.Code = 0
	j.Format = 2
	j.Date = section.Header.Date
	j.Filename = jsonFileName

	j.Year = fmt.Sprintf("%v", jDate.Year())

	j.RangeData = JsonRangeData{
		Location:    "Hill Top",
		FiringPoint: section.Header.Slug,
		TargetType:  "ISSF",
		Range:       strings.ReplaceAll(strings.ToLower(section.Header.Distance), "m", ""),
		Units:       "M",
	}
	j.ShooterData = JsonShooterData{
		Name: section.Header.Name,
		Club: "SHRC",
		UIN:  "TBA",
		Num:  "TBA",
	}

	j.DisplayData.Scalefactor = 2

	j.Stage = "0"

	j.Shots = JsonShots{}
	j.Shots.Discipline = "TBA"
	j.Shots.Calibre = "TBA"
	j.Shots.CalibreRaw = "TBA"
	j.Shots.SightersCut = lo.Reduce(section.Shots, func(agg int, item *CsvShotData, _ int) int {
		if strings.ToLower(item.Tags) == "sighter" || strings.HasPrefix(strings.ToLower(item.Id), "s") {
			return agg + 1
		}
		return agg
	}, 0)
	j.Shots.ShotsFired = len(section.Shots)
	j.Shots.CountingShots = lo.Reduce(section.Shots, func(agg int, item *CsvShotData, _ int) int {
		if strings.ToLower(item.Tags) != "sighter" && !strings.HasPrefix(strings.ToLower(item.Id), "s") {
			return agg + 1
		}
		return agg
	}, 0)

	j.Shots.Mpi = []JsonMpi{}

	var minX float32 = 0
	var maxX float32 = 0
	var minY float32 = 0
	var maxY float32 = 0
	lo.ForEach(section.Shots, func(item *CsvShotData, index int) {
		if item.XposMm < minX {
			minX = item.XposMm
		}
		if item.XposMm > maxX {
			maxX = item.XposMm
		}
		if item.YposMm < minY {
			minY = item.YposMm
		}
		if item.YposMm > maxY {
			maxY = item.YposMm
		}
	})

	mpi := JsonMpi{
		Height: maxX - minX,
		Dia:    maxY - minY,
	}
	j.Shots.Mpi = append(j.Shots.Mpi, mpi)

	j.Shots.Comp = []JsonComp{}

	comp := JsonComp{}

	var prev time.Time
	comp.No = lo.Map(section.Shots, func(s *CsvShotData, index int) JsonShotData {
		value := 0
		if s.Score == "X" {
			value = 6
		} else if s.Score == "V" {
			value = 5
		} else {
			i, _ := strconv.Atoi(s.Score)
			value = i
		}

		const shortForm = "Jan 02 2006 3:04:05 pm"
		const TwentyFourHourForm = "15:04:05"
		t, _ := time.Parse(shortForm, j.Date+" "+s.Time)

		if prev.IsZero() {
			prev = t
		}

		since := int(t.Sub(prev).Seconds())
		mins := since / 60

		prev = t

		return JsonShotData{
			ShotNo:            index,
			XPos:              s.XposMm,
			YPos:              s.YposMm,
			Dfc:               float32(math.Sqrt(math.Pow(float64(s.XposMm), 2) + math.Pow(float64(s.YposMm), 2))),
			Value:             value,
			Temp:              150,
			Status:            1,
			TimeOfShot:        t.Format(TwentyFourHourForm),
			Time:              t.Unix(),
			TimeSinceLastShot: fmt.Sprintf("%d:%02d", mins, since-(mins*60)),
		}
	})

	j.Shots.Comp = append(j.Shots.Comp, comp)

	// Export object to JSON
	jb, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}

	os.WriteFile(filepath.Join(outputFolder, jsonFileName), jb, 0644)

}

func parseCsv(fileName string, outputFolder string) []CsvSection {
	result := []CsvSection{}

	// Make sure output folder exists
	err := os.MkdirAll(outputFolder, os.ModePerm)
	if err != nil {
		panic(err)
	}

	state := string("Waiting")

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	header := []*CsvHeader{}
	headerCsv := []string{"date,name,slug,distance,target", ""}

	shots := []*CsvShotData{}
	shotsCsv := []string{}

	summary := []*CsvSummaryData{}
	summaryCsv := []string{}

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	groupNumber := 1
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		// Skip rows 1 & 2
		if lineNumber <= 2 {
			continue
		}

		// Blank lines
		if state == "Waiting" && len(strings.TrimSpace(line)) == 0 {
			continue
		}
		if state == "WaitingForData" && len(strings.TrimSpace(line)) == 0 {
			state = "ConsumingDataHeader"
			continue
		}
		if state == "ConsumingData" && len(strings.TrimSpace(line)) == 0 {
			// We've finished consuming the shots
			//fmt.Println(strings.Join(shotsCsv[:], "\n"))

			if err := gocsv.UnmarshalString(strings.Join(shotsCsv[:], "\n"), &shots); err != nil {
				panic(err)
			}

			state = "ConsumingSummary"
			continue
		}
		if state == "ConsumingSummary" && len(strings.TrimSpace(line)) == 0 {
			// We've finished consuming the summary
			if len(summaryCsv) > 0 {
				//fmt.Println("Summary----------------------------------------------")
				//fmt.Println(strings.Join(summaryCsv[:], "\n"))

				summaryCsv = append([]string{"none,none,none,none,name,x (mm),y (mm),x (inch),y (inch),x (moa),y (moa),x (mil),y (mil),v (m/s),v (fps),yaw (deg), pitch (deg),quality,none"}, summaryCsv...)

				if err := gocsv.UnmarshalString(strings.Join(summaryCsv[:], "\n"), &summary); err != nil {
					panic(err)
				}
			}

			data := CsvSection{}
			data.Header = *header[0]
			data.Shots = shots
			data.Summary = summary

			result = append(result, data)

			b, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				fmt.Println("error:", err)
			}
			// os.Stdout.Write(b)
			csvJsonFileName := fmt.Sprintf(filepath.Join(outputFolder, "csv_%03d.json"), groupNumber)
			os.WriteFile(csvJsonFileName, b, 0644)

			groupNumber++
			state = "Waiting"

			shotsCsv = []string{}
			summaryCsv = []string{}

			continue
		}

		if state == "Waiting" {
			headerCsv[1] = line

			if err := gocsv.UnmarshalString(strings.Join(headerCsv[:], "\n"), &header); err != nil { // Load clients from file
				panic(err)
			}

			//fmt.Println("Header----------------------------------------------")
			//fmt.Println(strings.Join(headerCsv[:], "\n"))
			//fmt.Println("----------------------------------------------------")

			state = "WaitingForData"
		} else if state == "ConsumingDataHeader" {
			shotsCsv = append(shotsCsv, "none"+line+",none")
			state = "ConsumingData"
		} else if state == "ConsumingData" {
			shotsCsv = append(shotsCsv, line)
		} else if state == "ConsumingSummary" {
			summaryCsv = append(summaryCsv, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func main() {
	fmt.Println("SHRC Shotmarker tool")

	exportCmd := flag.NewFlagSet("export", flag.ExitOnError)
	exportFile := exportCmd.String("f", "", "Filename of the csv file to export")
	exportFolder := exportCmd.String("o", "", "Destination folder for exported files")

	if len(os.Args) < 2 {
		fmt.Println("expected 'export' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "export":
		exportCmd.Parse(os.Args[2:])

		if len(strings.TrimSpace(*exportFile)) == 0 {
			fmt.Println("expected 'file' to have a valid filename")
			os.Exit(1)
		}

		csvResults := parseCsv(*exportFile, *exportFolder)

		for _, csvSection := range csvResults {
			exportOzScore(csvSection, *exportFolder)
		}

		fullPath, _ := filepath.Abs(*exportFolder)
		fmt.Printf("Exported %d sections from csv file '%s' into folder '%s'\n", len(csvResults), *exportFile, fullPath)

	default:
		fmt.Println("expected 'export' subcommands")
		os.Exit(1)
	}

}
