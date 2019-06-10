package collectors

func value(source interface{}, parameters map[string]interface{}) interface{} {
	return source
}

func empty(source interface{}, parameters map[string]interface{}) interface{} {
	return ""
}
