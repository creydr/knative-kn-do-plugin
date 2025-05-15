package k8s

type Arguments map[string]interface{}

func (a Arguments) get(key, defaultValue string) string {
	val, ok := a[key].(string)
	if !ok {
		return defaultValue
	}

	return val
}
