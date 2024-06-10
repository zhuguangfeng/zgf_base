package logger

var DEBUG = true

func Error(err error) Field {
	return Field{Key: "error", Val: err}
}
func Bool(key string, val bool) Field {
	return Field{Key: key, Val: val}
}
func String(key string, val string) Field {
	return Field{Key: key, Val: val}
}
func SafeString(key string, val string) Field {
	if DEBUG {
		return Field{Key: key, Val: val}
	} else {
		return Field{Key: key, Val: "*****"}
	}

}
func Int8(key string, val int8) Field {
	return Field{Key: key, Val: val}
}
func Int32(key string, val int32) Field {
	return Field{Key: key, Val: val}
}
func Int64(key string, val int64) Field {
	return Field{Key: key, Val: val}
}
