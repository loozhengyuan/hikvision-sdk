package resource

// Resource is an interface for all API resources.
type Resource interface {
	Kind() string
	String() (string, error)
	StringIndent() (string, error)
}
