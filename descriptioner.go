package renum

// Descriptioner allows types to describe themselves in more detail when asked, without having
// to embed this information in otherwise unstructured ways like fmt.Errorf("foo info: %v", err).
type Descriptioner interface {
	Description() string
}
