// +build !windows

package ios

import (
	"errors"
	"os"
)

func (s *IOStreams) EnableVirtualTerminalProcessing() error {
	return nil
}

func enableVirtualTerminalProcessing(f *os.File) error {
	return errors.New("not implemented")
}
