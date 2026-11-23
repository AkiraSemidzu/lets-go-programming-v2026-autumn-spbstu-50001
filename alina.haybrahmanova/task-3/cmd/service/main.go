package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/AkiraSemidzu/task-3/internal/model"
	"github.com/AkiraSemidzu/task-3/internal/reader"
	"github.com/AkiraSemidzu/task-3/internal/settings"
	"github.com/AkiraSemidzu/task-3/internal/writer"
)

const (
	configFlagName    = "config"
	configFlagDefault = ""
	configFlagUsage   = "path to YAML config file"
)

func main() {
	configPath := flag.String(configFlagName, configFlagDefault, configFlagUsage)
	flag.Parse()

	if err := run(*configPath); err != nil {
		panic(err)
	}
}

func run(configPath string) error {
	if configPath == "" {
		return errors.New("flag -config is required")
	}

	cfg, err := settings.Load(configPath)
	if err != nil {
		return fmt.Errorf("load settings: %w", err)
	}

	valutes, err := reader.ReadValutes(cfg.InputFile)
	if err != nil {
		return fmt.Errorf("read valutes: %w", err)
	}

	model.SortByValueDesc(valutes)

	if err := writer.WriteJSON(cfg.OutputFile, valutes); err != nil {
		return fmt.Errorf("write valutes: %w", err)
	}

	return nil
}
