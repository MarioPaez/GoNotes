package formatter

type Formatter interface {
	Export([]any) string
}
