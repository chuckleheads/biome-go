package studio

import (
	"io"

	"github.com/docker/cli/cli/streams"
)

// Streams is an interface which exposes the standard input and output streams
type Streams interface {
	In() *streams.In
	Out() *streams.Out
	Err() io.Writer
}

// Cli represents the docker command line client.
type Cli interface {
	Out() *streams.Out
	Err() io.Writer
	In() *streams.In
	SetIn(in *streams.In)
}

type StudioCli struct {
	in  *streams.In
	out *streams.Out
	err io.Writer
}

// Out returns the writer used for stdout
func (cli *StudioCli) Out() *streams.Out {
	return cli.out
}

// Err returns the writer used for stderr
func (cli *StudioCli) Err() io.Writer {
	return cli.err
}

// SetIn sets the reader used for stdin
func (cli *StudioCli) SetIn(in *streams.In) {
	cli.in = in
}

// In returns the reader used for stdin
func (cli *StudioCli) In() *streams.In {
	return cli.in
}
