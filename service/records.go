package service

import (
	"ferdyrurka/architektura-systemow/util"
	"fmt"
)

func Print(stringRecords map[string]uint16, boolRecords map[string]byte) {
	fmt.Println("############### String records ###############")
	for index, element := range stringRecords {
		fmt.Println(
			util.BLUE,
			"Record:",
			util.NORMAL_TEXT,
			index,
			util.BLUE,
			"Value uint16:",
			util.NORMAL_TEXT,
			element)
	}

	fmt.Println("############### Bool records ###############")
	for index, element := range boolRecords {
		fmt.Println(
			util.BLUE,
			"Record:",
			util.NORMAL_TEXT,
			index,
			util.BLUE,
			"Value binary:",
			util.NORMAL_TEXT,
			element)
	}
}
