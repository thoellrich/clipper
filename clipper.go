package clipper

type ClipboardMonitor struct{}

func NewClipboardMonitor() *ClipboardMonitor {
	return &ClipboardMonitor{}
}

func (c *ClipboardMonitor) Foo() int {
	return 42
}
