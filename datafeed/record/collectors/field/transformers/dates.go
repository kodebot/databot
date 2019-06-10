package transformers

import (
	"time"

	"github.com/golang/glog"
)

func formatDate(value interface{}, parameters map[string]interface{}) interface{} {
	if value == nil {
		return value
	}

	if valueTimePointer, ok := value.(*time.Time); ok {
		return valueTimePointer.String()
	}

	if valueTime, ok := value.(time.Time); ok {
		return valueTime.String()
	}

	glog.Errorf("formatDate is not allowed on non time.Time type")
	return value
}

func parseDate(value interface{}, parameters map[string]interface{}) interface{} {
	if value == nil {
		return value
	}

	layoutString := time.RFC3339
	if parameters != nil && parameters["layout"] != nil {
		layoutString = parameters["layout"].(string)
	}

	valueString := value.(string)
	result, err := time.Parse(layoutString, valueString)

	if err != nil {
		glog.Errorf("parsing date failed with layout %s", layoutString)
		return value
	}

	return result
}

func utcMidnightToNow(value interface{}, parameters map[string]interface{}) interface{} {
	if value == nil {
		return value
	}

	var inputTime *time.Time
	if valueTimePointer, ok := value.(*time.Time); ok {
		inputTime = valueTimePointer
	} else if valueTime, ok := value.(time.Time); ok {
		inputTime = &valueTime
	}

	if inputTime != nil && isUtcMidnight(inputTime) {
		return time.Now().UTC()
	}

	return value
}

func isUtcMidnight(datetime *time.Time) bool {
	glog.Infoln("checking if the date is midnight utc...")
	return (datetime.Hour() == 0 &&
		datetime.Minute() == 0 &&
		datetime.Second() == 0 &&
		datetime.Nanosecond() == 0 &&
		datetime.Location().String() == "UTC")

}
