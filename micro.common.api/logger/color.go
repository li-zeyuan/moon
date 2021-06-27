package logger

import "fmt"

const (
	colorRed     = uint8(iota + 91)
	colorGreen   //	绿
	colorYellow  //	黄
	colorBlue    // 蓝
	colorMagenta //	洋红
)

func red(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorRed, s)
}

func green(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorGreen, s)
}

func yellow(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorYellow, s)
}

func blue(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorBlue, s)
}

func magenta(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorMagenta, s)
}
