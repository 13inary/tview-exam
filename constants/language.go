package constants

const (
	Installing = "installing ..."
)

var langMap = map[string][]string{
	Installing: {"安装中 ……"},
}

func Ch(text string) string {
	return langMap[text][0]
}
