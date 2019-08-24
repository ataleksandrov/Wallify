package environment

// Validator should be implemented by types that need input validation check.
type Validator interface {
	Validate() error
}
