package commands

import (
	"context"
	"path/filepath"
	"strings"
	"time"

	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/photoprism"
	"github.com/photoprism/photoprism/internal/service"
	"github.com/photoprism/photoprism/pkg/fs"
	"github.com/photoprism/photoprism/pkg/txt"
	"github.com/urfave/cli"
)

// PurgeCommand registers the index cli command.
var PurgeCommand = cli.Command{
	Name:   "purge",
	Usage:  "Removes missing files from search results",
	Flags:  purgeFlags,
	Action: purgeAction,
}

var purgeFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "hard",
		Usage: "permanently delete from database",
	},
	cli.BoolFlag{
		Name:  "dry",
		Usage: "dry run, don't actually remove anything",
	},
}

// purgeAction removes missing files from search results
func purgeAction(ctx *cli.Context) error {
	start := time.Now()

	conf := config.NewConfig(ctx)
	service.SetConfig(conf)

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := conf.Init(); err != nil {
		return err
	}

	conf.InitDb()

	// get cli first argument
	subPath := strings.TrimSpace(ctx.Args().First())

	if subPath == "" {
		log.Infof("purge: removing missing files in %s", txt.Quote(filepath.Base(conf.OriginalsPath())))
	} else {
		log.Infof("purge: removing missing files in %s", txt.Quote(fs.RelName(filepath.Join(conf.OriginalsPath(), subPath), filepath.Dir(conf.OriginalsPath()))))
	}

	if conf.ReadOnly() {
		log.Infof("purge: read-only mode enabled")
	}

	w := service.Purge()

	opt := photoprism.PurgeOptions{
		Path: subPath,
		Dry:  ctx.Bool("dry"),
		Hard: ctx.Bool("hard"),
	}

	if files, photos, err := w.Start(opt); err != nil {
		return err
	} else {
		elapsed := time.Since(start)

		log.Infof("purge: removed %d files and %d photos in %s", len(files), len(photos), elapsed)
	}

	conf.Shutdown()

	return nil
}
