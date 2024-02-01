package katas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_toHumanReadableTime_NullSekunden(t *testing.T) {
	res, _ := toHumanReadableTime(0)
	assert.Equal(t, "00:00:00", res)
}

func Test_toHumanReadableTime_59Sekunden(t *testing.T) {
	res, _ := toHumanReadableTime(59)
	assert.Equal(t, "00:00:59", res)
}
func Test_toHumanReadableTime_60Sekunden(t *testing.T) {
	res, _ := toHumanReadableTime(60)
	assert.Equal(t, "00:01:00", res)
}

func Test_toHumanReadableTime_61Sekunden(t *testing.T) {
	res, _ := toHumanReadableTime(61)
	assert.Equal(t, "00:01:01", res)
}

func Test_toHumanReadableTime_59Minuten0Sekunden(t *testing.T) {
	res, _ := toHumanReadableTime(59 * 60)
	assert.Equal(t, "00:59:00", res)
}

func Test_toHumanReadableTime_1Stunde0Minuten0Sekunden(t *testing.T) {
	res, _ := toHumanReadableTime(60 * 60)
	assert.Equal(t, "01:00:00", res)
}

func Test_toHumanReadableTime_1Stunde0Minuten1Sekunde(t *testing.T) {
	res, _ := toHumanReadableTime(60*60 + 1)
	assert.Equal(t, "01:00:01", res)
}

func Test_toHumanReadableTime_1Stunde1Minute1Sekunde(t *testing.T) {
	res, _ := toHumanReadableTime(60*60 + 61)
	assert.Equal(t, "01:01:01", res)
}

func Test_toHumanReadableTime_1Stunde59Minuten1Sekunde(t *testing.T) {
	res, _ := toHumanReadableTime(60*60 + 60*59 + 1)
	assert.Equal(t, "01:59:01", res)
}

func Test_toHumanReadableTime_1Stunde59Minuten59Sekunden(t *testing.T) {
	res, _ := toHumanReadableTime(60*60 + 60*59 + 59)
	assert.Equal(t, "01:59:59", res)
}
func Test_toHumanReadableTime_25Stunden59Minuten59Sekunden(t *testing.T) {
	res, _ := toHumanReadableTime(60*60*25 + 60*48 + 59)
	assert.Equal(t, "25:48:59", res)
}
func Test_toHumanReadableTime_maxValue(t *testing.T) {
	res, _ := toHumanReadableTime(359999)
	assert.Equal(t, "99:59:59", res)
}
func Test_toHumanReadableTime_ErrorOverMaxValue(t *testing.T) {
	_, err := toHumanReadableTime(360000)
	assert.Error(t, err)
}
