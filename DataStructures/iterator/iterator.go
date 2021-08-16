package iterator

type Iterator interface {
	Inc()
	Dec()
	Get() interface{}
	Set(interface{})
}