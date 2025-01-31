package main

import (
	"testing"
)

func TestComputeStages(t *testing.T) {
	sections := []*CsvSection{
		{
			Header: CsvHeader{
				Name:        "Tim",
				Date:        "May 27 2022",
				FiringPoint: "100",
			},
		},
	}

	computeStages(&sections)

	if sections[0].Header.Stage != 1 {
		t.Fatalf(`TextComputeStages, single section should be stage 1`)
	}
}

func TestComputeStagesWithTwoGroupMembers(t *testing.T) {
	section100 := CsvSection{
		Header: CsvHeader{
			Name:        "Tim",
			Date:        "May 27 2022",
			FiringPoint: "100",
		},
		Shots: []*CsvShotData{
			{
				Time: "10:10:29 am",
			},
		},
	}
	section101 := CsvSection{
		Header: CsvHeader{
			Name:        "Tim",
			Date:        "May 27 2022",
			FiringPoint: "101",
		},
		Shots: []*CsvShotData{
			{
				Time: "9:10:29 am",
			},
		},
	}
	sections := []*CsvSection{
		&section100,
		&section101,
	}

	computeStages(&sections)

	if section100.Header.Stage != 2 {
		t.Fatalf(`TextComputeStages, section100 should be stage 2`)
	}
	if section101.Header.Stage != 1 {
		t.Fatalf(`TextComputeStages, section101 should be stage 1`)
	}
}
