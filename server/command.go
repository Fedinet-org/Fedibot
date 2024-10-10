package main

import (
	"strings"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"
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

func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {

	arg := strings.Fields(args.Command)
	if len(arg) < 2 {
		return &model.CommandResponse{
			ResponseType: model.CommandResponseTypeEphemeral,
			Text:         "Usage: /create_team [team name] [channel-1] [channel-2] [channel-3] [channel-n]",
		}, nil
	}

	if args.Command != CommandTrigger {
		return &model.CommandResponse{}, nil
	}

	teamName := arg[1]
	channels := arg[2:]

	team := &model.Team{
		Name: teamName,
	}

	team, err := p.API.CreateTeam(team)
	if err != nil {
		return &model.CommandResponse{
			ResponseType: model.CommandResponseTypeEphemeral,
			Text:         "Failed to create team",
		}, nil
	}

	for _, channelName := range channels {
		channel := &model.Channel{
			TeamId: team.Id,
			Name:   channelName,
		}

		channel, err := p.API.CreateChannel(channel)
		if err != nil {
			return &model.CommandResponse{
				ResponseType: model.CommandResponseTypeEphemeral,
				Text:         "Failed to create channel",
			}, nil
		}
	}
}

// Parse the command arguments
// args.CommandArgs contains the arguments passed to the command
// args.CommandArgs.Command contains the command trigger
// args.CommandArgs.Command contains the arguments passed to the command
// args.CommandArgs.UserId contains the user ID of the user who executed the command
// args.CommandArgs.ChannelId contains the channel ID where the command was executed
// args.CommandArgs.TeamId contains the team ID where the command was executed
// args.CommandArgs.RootId contains the root ID of the post that triggered the command
// args.CommandArgs.ParentId contains the parent ID of the post that triggered the command
// args.CommandArgs.CommandArgs contains the arguments passed to the command
