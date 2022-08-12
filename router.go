package nipper

type Route interface {
	Route() string
	Get(interface{}) Route
	Post(interface{}) Route
	Put(interface{}) Route
	Delete(interface{}) Route
	Patch(interface{}) Route
	Head(interface{}) Route
	Options(interface{}) Route
}

type Router interface {
	Route(string) Route
}
