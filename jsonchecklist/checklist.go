package jsonchecklist

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Checklist struct {
	filename  string
	checklist map[string]bool
}

func (c *Checklist) Check(item string) {
	c.checklist[item] = true
}

func (c *Checklist) IsChecked(item string) bool {
	return c.checklist[item]
}

func (c *Checklist) Uncheck(item string) {
	c.checklist[item] = false
}

func New(filename string) (*Checklist, error) {
	checklist := &Checklist{
		filename:  filename,
		checklist: make(map[string]bool),
	}

	err := checklist.load(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to load checklist from file: %w", err)
	}

	return checklist, nil
}

func (c *Checklist) Flush() error {
	f, err := os.OpenFile(c.filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o0640)
	if err != nil {
		return fmt.Errorf("failed to open file for writing: %w", err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(c.checklist)
	if err != nil {
		return fmt.Errorf("failed to encode checklist: %w", err)
	}

	return nil
}

func (c *Checklist) load(filename string) error {
	f, err := os.Open(filename)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to open file: %w", err)
	}

	if f == nil {
		return nil
	}

	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&c.checklist)
	if err != nil {
		return fmt.Errorf("failed to decode checklist: %w", err)
	}

	return nil
}
