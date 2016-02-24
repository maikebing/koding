package repair

import (
	"errors"
	"fmt"
	"io"

	"koding/klientctl/ctlcli"
	"koding/klientctl/klient"

	"github.com/koding/logging"
)

type Options struct {
	MountName string
}

// Command implements the klientctl.Command interface for KD Repair
type Command struct {
	Options Options
	Stdout  io.Writer
	Stdin   io.Reader
	Log     logging.Logger

	// A collection of Repairers responsible for actually repairing a given mount.
	Repairers []Repairer

	// The klient instance this struct will use, mainly given to Repairers.
	//Klient interface {
	//}

	// The options to use if this struct needs to dial Klient.
	//
	// Note! These will be ignored if c.Klient is already defined before Run() is
	// called.
	KlientOptions klient.KlientOptions

	// the following vars exist primarily for mocking ability, and ensuring
	// an enclosed environment within the struct.

	// The ctlcli Helper. See the type docs for a better understanding of this.
	Helper ctlcli.Helper
}

// Help prints help to the caller.
func (c *Command) Help() {
	if c.Helper == nil {
		// Ugh, talk about a bad UX
		fmt.Fprintln(c.Stdout, "Error: Help was requested but command has no helper.")
		return
	}

	c.Helper(c.Stdout)
}

// printf is a helper function for printing to the internal writer.
func (c *Command) printfln(f string, i ...interface{}) {
	if c.Stdout == nil {
		return
	}

	fmt.Fprintf(c.Stdout, f+"\n", i...)
}

// Run the Mount command
func (c *Command) Run() (int, error) {
	if err := c.handleOptions(); err != nil {
		return 1, err
	}

	if err := c.initDefaultRepairers(); err != nil {
		return 2, err
	}

	if err := c.runRepairers(); err != nil {
		return 3, err
	}

	return 0, nil
}

func (c *Command) handleOptions() error {
	if c.Options.MountName == "" {
		c.printfln("MountName is a required option.")
		c.Help()
		return errors.New("Missing mountname option")
	}

	return nil
}

// initDefaultRepairers creates the repairers for this Command if the
// Command.Repairers field is *nil*. This allows a caller can specify their own
// repairers if desired.
func (c *Command) initDefaultRepairers() error {
	if c.Repairers != nil {
		return nil
	}

	return errors.New("Not implemented")
}

// runRepairers executes the given repairers. First running Statuses, and then
// Repair() on any of the Statuses that don't succeed. If any Repairs fail,
// the error is returned. It is the responsibility of the Repairer (usualy via the
// RetryRepairer) to repeat repair attempts on failures.
func (c *Command) runRepairers() error {
	// If there are no repairers to run, the core functionality of this command
	// is incapable of working. So, return an error.
	if len(c.Repairers) == 0 {
		return errors.New("Repair command has 0 repairers.")
	}

	for _, r := range c.Repairers {
		ok, err := r.Status()

		// If there is no problem from Status, we can just move onto the next Repairer.
		if ok {
			continue
		}

		// TODO: Improve this message, and move it to a package.
		c.printfln("Identified problem with the %s", r.Description())

		c.Log.Warning(
			"Repairer returned a non-ok status. Running its repair. repairer:%s, err:%s",
			r.Name(), err,
		)

		err = r.Repair()
		if err != nil {
			c.Log.Error("Repairer failed to repair. repairer:%s, err:%s", r.Name(), err)
			// TODO: Improve this message, and move it to a package.
			c.printfln("Unable to repair the %s", r.Description())
			return err
		}
	}

	return nil
}
