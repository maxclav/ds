package textbuffer

import "github.com/maxclav/ds/container"

// TextBuffer holds text you can edit anywhere, like an editor's buffer.
// Positions and Len count bytes, and String returns the whole text.
type TextBuffer interface {
	container.Container
	// Insert inserts s before position i.
	Insert(i int, s string)
	// Delete removes n bytes starting at position i.
	Delete(i, n int)
	// Slice returns the text from position i to j-1, like s[i:j].
	Slice(i, j int) string
}
