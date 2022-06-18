package gocolor

var _default string = "\u001b[0m"

func Default(text string) string {
	return "\u001b[0m" + text
}

func Red(text string) string {
	return "\u001b[31m" + text + _default
}

func Blue(text string) string {
	return "\u001b[34m" + text + _default
}

func Yellow(text string) string {
	return "\u001b[33m" + text + _default
}

func Green(text string) string {
	return "\u001b[32m" + text + _default
}

func Purple(text string) string {
	return "\u001b[35m" + text + _default
}

func Cyan(text string) string {
	return "\u001b[36m" + text + _default
}

func White(text string) string {
	return "\u001b[37m" + text + _default
}

func Black(text string) string {
	return "\u001b[30m" + text + _default
}
