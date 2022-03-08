package timer

import "time"

// GetNowTime 提供当前时间
func GetNowTime() time.Time  {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
	
}

func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTimer.Add(duration), nil
}

