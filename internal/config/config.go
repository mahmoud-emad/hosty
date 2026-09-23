package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
)

type Config struct {
	filePath string
}

func NewConfig(filePath string) *Config {
	return &Config{
		filePath: filePath,
	}
}

func (c *Config) Load(target any) error {
	file, err := os.Open(c.filePath)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}

	defer func() {
		_ = file.Close()
	}()

	err = json.NewDecoder(file).Decode(target)
	if errors.Is(err, io.EOF) {
		return nil
	}

	return err
}

func (c *Config) Save(source any) error {
	if err := os.MkdirAll(filepath.Dir(c.filePath), 0o700); err != nil {
		return err
	}

	file, err := os.Create(c.filePath)
	if err != nil {
		return err
	}

	defer func() {
		if closeErr := file.Close(); err == nil {
			err = closeErr
		}
	}()

	return json.NewEncoder(file).Encode(source)
}
