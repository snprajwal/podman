package crit

import (
	"errors"
	"path/filepath"

	"github.com/checkpoint-restore/go-criu/v6/crit/images"
)

// Helper function to load stats file into Go struct
func getStats(path string) (*images.StatsEntry, error) {
	c := New(path, "", "", false, false)
	statsImg, err := c.Decode()
	if err != nil {
		return nil, err
	}

	stats, ok := statsImg.Entries[0].Message.(*images.StatsEntry)
	if !ok {
		return nil, errors.New("Failed to type assert stats image")
	}

	return stats, nil
}

// GetDumpStats returns the dump statistics of a checkpoint.
// dir is the path to the directory with the checkpoint images.
func GetDumpStats(dir string) (*images.DumpStatsEntry, error) {
	stats, err := getStats(filepath.Join(dir, "stats-dump"))
	if err != nil {
		return nil, err
	}

	return stats.GetDump(), nil
}

// GetRestoreStats returns the restore statistics of a checkpoint.
// dir is the path to the directory with the checkpoint images.
func GetRestoreStats(dir string) (*images.RestoreStatsEntry, error) {
	stats, err := getStats(filepath.Join(dir, "stats-restore"))
	if err != nil {
		return nil, err
	}

	return stats.GetRestore(), nil
}
