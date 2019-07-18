package stringutil

// ToStringSlice converts slice of interface{} to slice of string
func ToStringSlice(slice []interface{}) ([]string, bool) {
	if slice == nil {
		return nil, true
	}
	result := []string{}

	for _, item := range slice {
		if val, ok := item.(string); ok {
			result = append(result, val)
		} else {
			return nil, false
		}
	}
	return result, true
}
