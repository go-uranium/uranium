package scan

type Scanner interface {
	Scan(...interface{}) error
}
