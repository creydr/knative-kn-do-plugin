package k8s

type ArgumentName string

type Arguments map[ArgumentName]interface{}

func (a Arguments) get(key ArgumentName, defaultValue string) string {
	val, ok := a[key].(string)
	if !ok {
		return defaultValue
	}

	return val
}
