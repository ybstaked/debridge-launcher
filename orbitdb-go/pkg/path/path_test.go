//+build unit

package path

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExplode(t *testing.T) {
	samples := []struct {
		in  string
		out []string
	}{
		{"/hello/you.txt", []string{"/hello/", "you", "txt"}},
		{"/hello/world.txt.c", []string{"/hello/", "world.txt", "c"}},
		{"/world/txt", []string{"/world/", "txt", ""}},
		{"/world/.txt", []string{"/world/", ".txt", ""}},
		{"/world.txt", []string{"/", "world", "txt"}},
		{"/world", []string{"/", "world", ""}},
		{"world", []string{"", "world", ""}},
		{"", []string{"", "", ""}},
	}
	for _, sample := range samples {
		t.Run(sample.in, func(t *testing.T) {
			dir, name, ext := Explode(sample.in)
			assert.Equal(t, sample.out, []string{dir, name, ext})
		})
	}
}
