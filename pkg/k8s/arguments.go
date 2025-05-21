package k8s

type ArgumentName string

type Arguments map[ArgumentName]interface{}

func (a Arguments) get(key ArgumentName) interface{} {
	return a[key]
}
