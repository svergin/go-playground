package fit

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/tormoder/fit"
)

type recordCorrectionEntry struct {
	Speed       uint16
	SpeedNew    uint16
	Distance    uint32
	DistanceNew uint32
}

type recordCorrection struct {
	entries   map[time.Time]recordCorrectionEntry
	speed     uint16
	starttime time.Time
	endtime   time.Time
}

type recordSwitchState int

const (
	stateAddToSlice    recordSwitchState = iota
	stateAddNotToSlice                   = 1
	stateStartNewSlice                   = 2
	stateEndOfSlice                      = 3
)

func readFitFile() ([]recordCorrection, error) {
	// Read our FIT test file data
	testFile := filepath.Join("c:\\", "fit", "i26475458.fit")
	testData, err := os.ReadFile(testFile)
	if err != nil {
		return nil, err
	}

	// Decode the FIT file data
	fit, err := fit.Decode(bytes.NewReader(testData))
	if err != nil {
		return nil, err
	}

	// Get the actual activity
	activity, err := fit.Activity()
	if err != nil {
		return nil, err
	}

	result := make([]recordCorrection, 0)
	rc := recordCorrection{}
	rc.entries = make(map[time.Time]recordCorrectionEntry)
	for idx, rec := range activity.Records {

		var re recordCorrectionEntry
		switch state := getState(rec, idx, activity.Records); state {
		case stateStartNewSlice:
			rc = recordCorrection{}
			rc.entries = make(map[time.Time]recordCorrectionEntry)
			re = recordCorrectionEntry{}
			re.Speed = rec.Speed
			re.Distance = rec.Distance
			rc.entries[rec.Timestamp] = re
		case stateEndOfSlice:
			result = append(result, rc)
		case stateAddToSlice:
			re = recordCorrectionEntry{}
			re.Speed = rec.Speed
			re.Distance = rec.Distance
			rc.entries[rec.Timestamp] = re
		case stateAddNotToSlice:
			continue
		}

		// fmt.Println("time: ", record.Timestamp, " speed: ", record.GetSpeedScaled(), "speed_ori: ", record.Speed, " distance: ", record.GetDistanceScaled(), " distance_ori: ", record.Distance)
		// i++
		// if i > 30 {
		// 	break

		// }
	}
	return result, nil
}

func PrintResult() {
	result, err := readFitFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	count := 0
	for i, v := range result {
		fmt.Println("Slice nummer ", i)
		fmt.Println("Länge: ", len(v.entries))
		timestamps := make([]time.Time, 0, len(v.entries))

		for k, _ := range v.entries {
			timestamps = append(timestamps, k)
		}
		sort.Slice(timestamps, func(i, j int) bool {
			return timestamps[i].Before(timestamps[j])
		})
		for _, k := range timestamps {
			fmt.Println("time: ", k, " - Value: ", v.entries[k])
			count++
		}

	}
	fmt.Println("kaputte Datensätze: ", count)
	// fmt.Println("Länge: ", len(rc.entries))
	// fmt.Println("Startzeit: ", rc.starttime)
}

func Print() {
	// Read our FIT test file data
	testFile := filepath.Join("c:\\", "fit", "i26475458.fit")
	testData, err := os.ReadFile(testFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Decode the FIT file data
	fit, err := fit.Decode(bytes.NewReader(testData))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the actual activity
	activity, err := fit.Activity()
	if err != nil {
		fmt.Println(err)
		return
	}
	count := 0
	for _, rec := range activity.Records {
		if math.IsNaN(rec.GetSpeedScaled()) {
			fmt.Println("time: ", rec.Timestamp, " speed: ", rec.GetSpeedScaled(), "speed_ori: ", rec.Speed, " distance: ", rec.Distance)
			count++
		}
	}
	fmt.Println("kaputte Datensätze: ", count)
}

func getState(actRecMsg *fit.RecordMsg, idx int, recs []*fit.RecordMsg) recordSwitchState {
	if idx == 0 {
		if math.IsNaN(actRecMsg.GetSpeedScaled()) {
			return stateStartNewSlice
		} else {
			return stateAddToSlice
		}
	}
	prevRecMsg := recs[idx-1]
	if math.IsNaN(actRecMsg.GetSpeedScaled()) && !math.IsNaN(prevRecMsg.GetSpeedScaled()) {
		return stateStartNewSlice
	} else if !math.IsNaN(actRecMsg.GetSpeedScaled()) && math.IsNaN(prevRecMsg.GetSpeedScaled()) {
		return stateEndOfSlice
	} else if math.IsNaN(actRecMsg.GetSpeedScaled()) {
		return stateAddToSlice
	} else {
		return stateAddNotToSlice
	}
}
