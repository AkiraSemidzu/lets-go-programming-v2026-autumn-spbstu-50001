package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/AkiraSemidzu/task-3/internal/model"
)

const (
	dirPerm    os.FileMode = 0o755
	filePerm   os.FileMode = 0o644
	jsonPrefix             = ""
	jsonIndent             = "    "
)

// WriteJSON сохраняет валюты в JSON, создавая недостающие директории.
func WriteJSON(path string, valutes []model.Valute) error {
	if valutes == nil {
		valutes = []model.Valute{}
	}

	data, err := json.MarshalIndent(valutes, jsonPrefix, jsonIndent)
	if err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(path), dirPerm); err != nil {
		return fmt.Errorf("create output directory for %q: %w", path, err)
	}

	if err := os.WriteFile(path, data, filePerm); err != nil {
		return fmt.Errorf("write output file %q: %w", path, err)
	}

	return nil
}
