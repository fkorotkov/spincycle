package cmd

import (
	"fmt"

	"github.com/square/spincycle/proto"
	"github.com/square/spincycle/spinc/app"
)

type Status struct {
	ctx   app.Context
	reqId string
}

func NewStatus(ctx app.Context) *Status {
	return &Status{
		ctx: ctx,
	}
}

func (c *Status) Prepare() error {
	if len(c.ctx.Command.Args) == 0 {
		return fmt.Errorf("Usage: spinc status <id>\n")
	}
	c.reqId = c.ctx.Command.Args[0]
	return nil
}

func (c *Status) Run() error {
	status, err := c.ctx.RMClient.RequestStatus(c.reqId)
	if err != nil {
		return err
	}
	if c.ctx.Options.Debug {
		app.Debug("status: %#v", status)
	}

	if c.ctx.Hooks.CommandRunResult != nil {
		c.ctx.Hooks.CommandRunResult(status, err)
		return nil
	}

	fmt.Printf("state:      %s\n", proto.StateName[status.State])
	fmt.Printf("jobs done:  %d\n", status.FinishedJobs)
	fmt.Printf("jobs total: %d\n", status.TotalJobs)
	fmt.Printf("created:    %s\n", status.CreatedAt)
	fmt.Printf("started:    %s\n", status.StartedAt)
	fmt.Printf("finished:   %s\n", status.FinishedAt)
	fmt.Printf("user:       %s\n", status.User)
	fmt.Printf("type:       %s\n", status.Type)
	fmt.Printf("id:         %s\n", status.Id)

	return nil
}
