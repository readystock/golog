package golog

import (
    "testing"
)

func TestLogger_Print(t *testing.T) {
    log := Default
    log.Level = VerboseLevel
    log.print(VerboseLevel, "test verbose", true)
}
