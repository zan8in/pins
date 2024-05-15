package timeutil

import "time"

func Format(format string) string {
	return time.Now().Format(format)
}

// 使用time.Parse函数来尝试将输入日期字符串解析为指定的日期格式。
// 如果日期格式不正确，会返回一个错误，然后我们可以根据错误来判断日期格式是否正确。
func ValidateDateFormat(value, layout string) bool {
	_, err := time.Parse(layout, value)
	return err == nil
}

// 获取几天前日期
func GetAnyTimeAgo(days int) string {
	t := time.Now().AddDate(0, 0, -days)
	return t.Format("2006-01-02")
}
