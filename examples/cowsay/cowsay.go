package cowsay

import (
	"fmt"
	"strings"
)

type Cow struct{}

func (c *Cow) Say(message string) string {
	border := strings.Repeat("-", len(message)+2)
	return fmt.Sprintf(` %s
< %s >
 %s
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
`, border, message, border)
}
