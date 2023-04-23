package util

import (
	"bufio"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

func InitConsole() {
	scanner := bufio.NewScanner(os.Stdin)

	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.InfoLevel

	source := Source{log: log}

	go func() {
		for scanner.Scan() {
			if t := strings.TrimSpace(scanner.Text()); len(t) > 0 {
				name := strings.Split(t, " ")[0]
				if c, ok := cmd.ByAlias(name); ok {
					c.Execute(strings.TrimPrefix(strings.TrimPrefix(t, name), " "), source)
				} else {
					output := &cmd.Output{}
					output.Errorf("Unknown command '%s'", name)
					source.SendCommandOutput(output)
				}
			}
		}
	}()
}

type Source struct {
	log *logrus.Logger
}

func (Source) Position() mgl64.Vec3 {
	return mgl64.Vec3{}
}

func (s Source) SendCommandOutput(o *cmd.Output) {
	for _, e := range o.Errors() {
		s.log.Error(text.ANSI(e))
	}
	for _, m := range o.Messages() {
		s.log.Info(text.ANSI(m))
	}
}

func (Source) World() *world.World {
	return nil
}
