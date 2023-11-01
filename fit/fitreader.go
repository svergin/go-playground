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

type recordResult struct {
	recordCorrections []recordCorrection
}

func ReadFitFile() {
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

	rc := recordCorrection{}
	rc.entries = make(map[time.Time]recordCorrectionEntry)
	for _, record := range activity.Records {
		if math.IsNaN(record.GetSpeedScaled()) {
			re := recordCorrectionEntry{}
			re.Speed = record.Speed
			re.Distance = record.Distance
			rc.entries[record.Timestamp] = re
		}
		// fmt.Println("time: ", record.Timestamp, " speed: ", record.GetSpeedScaled(), "speed_ori: ", record.Speed, " distance: ", record.GetDistanceScaled(), " distance_ori: ", record.Distance)
		// i++
		// if i > 30 {
		// 	break

		// }
	}
	timestamps := make([]time.Time, 0, len(rc.entries))

	for k, _ := range rc.entries {
		timestamps = append(timestamps, k)
	}
	sort.Slice(timestamps, func(i, j int) bool {
		return timestamps[i].Before(timestamps[j])
	})
	for _, k := range timestamps {
		fmt.Println("time: ", k, " - Value: ", rc.entries[k])
	}
	fmt.Println("Länge: ", len(rc.entries))

}
