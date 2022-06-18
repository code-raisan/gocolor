package gocolor

var _default string = "\u001b[0m"

func Default(text string) string {
	result := "\u001b[0m" + text
	return result
}

func Red(text string) string {
	result := "\u001b[31m" + text + _default
	return result
}

func Blue(text string) string {
	result := "\u001b[34m" + text + _default
	return result
}

func Yellow(text string) string {
	result := "\u001b[33m" + text + _default
	return result
}

func Green(text string) string {
	result := "\u001b[32m" + text + _default
	return result
}

func Purple(text string) string {
	result := "\u001b[35m" + text + _default
	return result
}

func Cyan(text string) string {
	result := "\u001b[36m" + text + _default
	return result
}

func White(text string) string {
	result := "\u001b[37m" + text + _default
	return result
}

func Black(text string) string {
	result := "\u001b[30m" + text + _default
	return result
}
