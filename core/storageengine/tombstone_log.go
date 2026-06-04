package storageengine

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"os"
	"path/filepath"

	commons "github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/storageengine/utils"
	"github.com/hanami/tidets/core/storageengine/utils/codec"
	"github.com/hanami/tidets/core/storageengine/wal"
)

const (
	tombstoneLogFilename = "tombstones.log"
	tombstoneLogMagic    = 0x54535444 // "TSTD"
	tombstoneLogVersion  = 1

	tombstoneOpMarkRange uint8 = 1
)

type tombstoneLog struct {
	path     string
	file     *os.File
	w        *bufio.Writer
	syncMode wal.SyncMode
}

func openTombstoneLog(dir string, mode wal.SyncMode) (*tombstoneLog, error) {
	path := filepath.Join(dir, tombstoneLogFilename)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0o644)
	if err != nil {
		return nil, err
	}
	return &tombstoneLog{
		path:     path,
		file:     f,
		w:        bufio.NewWriter(f),
		syncMode: mode,
	}, nil
}

func loadTombstoneLog(dir string, apply func(keyStr string, start, end int64)) error {
	path := filepath.Join(dir, tombstoneLogFilename)
	f, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		keyStr, start, end, err := readTombstoneRecord(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		apply(keyStr, start, end)
	}
}

func (l *tombstoneLog) AppendMark(keyStr string, start, end int64) error {
	if l == nil || start > end {
		return nil
	}
	device, measurement := utils.SplitSeriesKey(keyStr)
	if measurement == "" {
		return commons.ErrStorageDeviceMeasurementRequired
	}
	if err := writeTombstoneRecord(l.w, device, measurement, start, end); err != nil {
		return err
	}
	return l.syncIfNeeded()
}

func (l *tombstoneLog) RewriteAll(ranges map[string][]timeRange) error {
	if l == nil {
		return nil
	}
	tmp := l.path + ".tmp"
	f, err := os.OpenFile(tmp, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0o644)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(f)
	for keyStr, items := range ranges {
		device, measurement := utils.SplitSeriesKey(keyStr)
		if measurement == "" {
			_ = f.Close()
			_ = os.Remove(tmp)
			return commons.ErrStorageDeviceMeasurementRequired
		}
		for _, r := range items {
			if err := writeTombstoneRecord(w, device, measurement, r.start, r.end); err != nil {
				_ = f.Close()
				_ = os.Remove(tmp)
				return err
			}
		}
	}
	if err := w.Flush(); err != nil {
		_ = f.Close()
		_ = os.Remove(tmp)
		return err
	}
	if l.syncMode == wal.SyncAlways {
		if err := f.Sync(); err != nil {
			_ = f.Close()
			_ = os.Remove(tmp)
			return err
		}
	}
	if err := f.Close(); err != nil {
		_ = os.Remove(tmp)
		return err
	}
	if l.file != nil {
		if err := l.Close(); err != nil {
			_ = os.Remove(tmp)
			return err
		}
	}
	if err := os.Rename(tmp, l.path); err != nil {
		return err
	}
	nl, err := openTombstoneLog(filepath.Dir(l.path), l.syncMode)
	if err != nil {
		return err
	}
	l.file = nl.file
	l.w = nl.w
	return nil
}

func (l *tombstoneLog) Sync() error {
	if l == nil {
		return nil
	}
	if l.w != nil {
		if err := l.w.Flush(); err != nil {
			return err
		}
	}
	if l.file != nil {
		return l.file.Sync()
	}
	return nil
}

func (l *tombstoneLog) Close() error {
	if l == nil {
		return nil
	}
	if l.w != nil {
		if err := l.w.Flush(); err != nil {
			return err
		}
	}
	if l.file != nil {
		err := l.file.Close()
		l.file = nil
		l.w = nil
		return err
	}
	return nil
}

func (l *tombstoneLog) syncIfNeeded() error {
	if l == nil {
		return nil
	}
	if l.w != nil {
		if err := l.w.Flush(); err != nil {
			return err
		}
	}
	if l.syncMode == wal.SyncAlways && l.file != nil {
		return l.file.Sync()
	}
	return nil
}

func writeTombstoneRecord(w io.Writer, device, measurement string, start, end int64) error {
	if err := binary.Write(w, binary.LittleEndian, uint32(tombstoneLogMagic)); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, uint32(tombstoneLogVersion)); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, tombstoneOpMarkRange); err != nil {
		return err
	}
	if err := codec.WriteString(w, device); err != nil {
		return err
	}
	if err := codec.WriteString(w, measurement); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, start); err != nil {
		return err
	}
	return binary.Write(w, binary.LittleEndian, end)
}

func readTombstoneRecord(r io.Reader) (string, int64, int64, error) {
	var magic uint32
	if err := binary.Read(r, binary.LittleEndian, &magic); err != nil {
		return "", 0, 0, err
	}
	if magic != tombstoneLogMagic {
		return "", 0, 0, commons.Wrap("storage", commons.CodeCorrupt, "read tombstone magic", commons.ErrWALCorruptRecord)
	}

	var version uint32
	if err := binary.Read(r, binary.LittleEndian, &version); err != nil {
		return "", 0, 0, err
	}
	if version != tombstoneLogVersion {
		return "", 0, 0, commons.Errorf("storage", commons.CodeInvalidArgument,
			"unsupported tombstone log version %d (want %d), remove data-dir/%s", version, tombstoneLogVersion, tombstoneLogFilename)
	}

	var op uint8
	if err := binary.Read(r, binary.LittleEndian, &op); err != nil {
		return "", 0, 0, err
	}
	if op != tombstoneOpMarkRange {
		return "", 0, 0, commons.ErrWALCorruptRecord
	}

	device, err := codec.ReadString(r)
	if err != nil {
		return "", 0, 0, err
	}
	measurement, err := codec.ReadString(r)
	if err != nil {
		return "", 0, 0, err
	}
	var start, end int64
	if err := binary.Read(r, binary.LittleEndian, &start); err != nil {
		return "", 0, 0, err
	}
	if err := binary.Read(r, binary.LittleEndian, &end); err != nil {
		return "", 0, 0, err
	}
	return device + "." + measurement, start, end, nil
}
