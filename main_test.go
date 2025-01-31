package main

import (
	"fmt"
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
	sections := []*CsvSection{
		{
			Header: CsvHeader{
				Name:        "Tim",
				Date:        "May 27 2022",
				FiringPoint: "100",
			},
			Shots: []*CsvShotData{
				{
					Time: "10:10:29 pm",
				},
			},
		},
		{
			Header: CsvHeader{
				Name:        "Tim",
				Date:        "May 27 2022",
				FiringPoint: "101",
			},
			Shots: []*CsvShotData{
				{
					Time: "9:10:29 pm",
				},
			},
		},
	}

	computeStages(&sections)

	fmt.Println((*sections[0]).Header)
	fmt.Println((*sections[1]).Header)

	if (*sections[0]).Header.FiringPoint != "101" || (*sections[0]).Header.Stage == 1 {
		t.Fatalf(`TextComputeStages, first section should be FP101 and stage 1`)
	}
	if (*sections[1]).Header.FiringPoint != "100" || (*sections[1]).Header.Stage == 2 {
		t.Fatalf(`TextComputeStages, second section should be FP100 and stage 2`)
	}
}
