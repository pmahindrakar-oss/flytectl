package get

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateGetCommand(t *testing.T) {
	getCommand := CreateGetCommand()
	assert.Equal(t, "get", getCommand.Use)
	assert.Equal(t, "Used for fetching various flyte resources including tasks/workflows/launchplans/executions/project.", getCommand.Short)
	fmt.Println(getCommand.Commands())
	assert.Equal(t, len(getCommand.Commands()), 5)
	cmdNouns := getCommand.Commands()
	// Sort by Use value.
	sort.Slice(cmdNouns, func(i, j int) bool {
		return cmdNouns[i].Use < cmdNouns[j].Use
	})
	assert.Equal(t, "execution", cmdNouns[0].Use)
	assert.Equal(t, []string{"executions"}, cmdNouns[0].Aliases)
	assert.Equal(t, "Gets execution resources", cmdNouns[0].Short)
	assert.Equal(t, "launchplan", cmdNouns[1].Use)
	assert.Equal(t, []string{"launchplans"}, cmdNouns[1].Aliases)
	assert.Equal(t, "Gets launch plan resources", cmdNouns[1].Short)
	assert.Equal(t, "project", cmdNouns[2].Use)
	assert.Equal(t, []string{"projects"}, cmdNouns[2].Aliases)
	assert.Equal(t, "Gets project resources", cmdNouns[2].Short)
	assert.Equal(t, "task", cmdNouns[3].Use)
	assert.Equal(t, []string{"tasks"}, cmdNouns[3].Aliases)
	assert.Equal(t, "Gets task resources", cmdNouns[3].Short)
	assert.Equal(t, "workflow", cmdNouns[4].Use)
	assert.Equal(t, []string{"workflows"}, cmdNouns[4].Aliases)
	assert.Equal(t, "Gets workflow resources", cmdNouns[4].Short)
}
