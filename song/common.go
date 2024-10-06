package song

import "time"

func timeFromString(strTime string) (*time.Time, error) {
	time, err := time.Parse("02.01.2006", strTime)
	if err != nil {
		return nil, err
	}
	return &time, nil
}

func likeStr(str string) string {
	return "%" + str + "%"
}
