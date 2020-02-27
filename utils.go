package config

func isZeroValue(value interface{}) bool {
	switch v := value.(type) {
	case bool:
		return v == *new(bool)
	case int:
		return v == *new(int)
	case string:
		return v == *new(string)
	case uint:
		return v == *new(uint)
	case float64:
		return v == *new(float64)
	default:
		return v == nil
	}
}
