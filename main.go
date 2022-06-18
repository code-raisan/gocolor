package gocolor

var _default string = "\x1b[0m"

func Default(text string) string {
	return "\x1b[0m" + text
}

func Red(text string) string {
	return "\x1b[31m" + text + _default
}

func Blue(text string) string {
	return "\x1b[34m" + text + _default
}

func Yellow(text string) string {
	return "\x1b[33m" + text + _default
}

func Green(text string) string {
	return "\x1b[32m" + text + _default
}

func Purple(text string) string {
	return "\x1b[35m" + text + _default
}

func Cyan(text string) string {
	return "\x1b[36m" + text + _default
}

func White(text string) string {
	return "\x1b[37m" + text + _default
}

func Black(text string) string {
	return "\x1b[30m" + text + _default
}
