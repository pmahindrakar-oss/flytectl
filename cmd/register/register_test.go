package register

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterCommand(t *testing.T) {
	registerCommand := Command()
	assert.Equal(t, "register", registerCommand.Use)
	assert.Equal(t, "Registers tasks/workflows/launchplans from list of generated serialized files.", registerCommand.Short)
	fmt.Println(registerCommand.Commands())
	assert.Equal(t, len(registerCommand.Commands()), 1)
	cmdNouns := registerCommand.Commands()
	// Sort by Use value.
	sort.Slice(cmdNouns, func(i, j int) bool {
		return cmdNouns[i].Use < cmdNouns[j].Use
	})
	assert.Equal(t, "files", cmdNouns[0].Use)
	assert.Equal(t, []string{"file"}, cmdNouns[0].Aliases)
	assert.Equal(t, "Registers file resources", cmdNouns[0].Short)
}
