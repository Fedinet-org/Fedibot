package main

import (
	"github.com/mattermost/mattermost/server/public/model"
	"github.com/pkg/errors"
)

const CommandTrigger = "createteam"

func (p *Plugin) registerCommand() error {
	if err := p.API.RegisterCommand(&model.Command{
		Trigger:          CommandTrigger,
		AutoComplete:     true,
		AutoCompleteHint: "[team name] [channel-1] [channel-2] [channel-3] [channel-n]",
		AutoCompleteDesc: "Create a team with predefined channels",
	}); err != nil {
		return errors.Wrap(err, "failed to register command")
	}

	return nil
}
