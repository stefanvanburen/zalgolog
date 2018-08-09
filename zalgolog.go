package zalgolog

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/kortschak/zalgo"
)

type Handler struct {
	mu   sync.Mutex
	z    *zalgo.Corrupter
	pain *bytes.Buffer
	h    log.Handler
}

var Default = New(os.Stderr)

func New(w io.Writer) *Handler {
	pain := bytes.NewBuffer(nil)
	z := zalgo.NewCorrupter(pain)

	z.Zalgo = func(n int, r rune, z *zalgo.Corrupter) bool {
		z.Up += 0.1
		z.Middle += complex(0.01, 0.01)
		z.Down += complex(real(z.Down)*0.1, 0)
		return false
	}

	return &Handler{
		pain: pain,
		z:    z,
		h:    text.New(w),
	}
}

func (h *Handler) HandleLog(e *log.Entry) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.pain.Reset()
	h.z.Up = complex(0, 0.2)
	h.z.Middle = complex(0, 0.2)
	h.z.Down = complex(0.001, 0.3)

	s := e.Message
	_, _ = fmt.Fprintf(h.z, s)
	e.Message = h.pain.String()

	return h.h.HandleLog(e)
}
