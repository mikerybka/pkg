package backups

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
)

func Hash(src, dst string) error {
	fi, err := os.Stat(src)
	if err != nil {
		return err
	}
	if fi.IsDir() {
		err := os.MkdirAll(dst, os.ModePerm)
		if err != nil {
			return err
		}

		entries, err := os.ReadDir(src)
		if err != nil {
			return err
		}
		for _, e := range entries {
			err := Hash(
				filepath.Join(src, e.Name()),
				filepath.Join(dst, e.Name()),
			)
			if err != nil {
				return err
			}
		}
		return nil
	}

	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	hasher := sha256.New()
	_, err = io.Copy(hasher, f)
	if err != nil {
		return err
	}
	hash := hasher.Sum(nil)
	b := make([]byte, hex.EncodedLen(len(hash)))
	hex.Encode(b, hash)
	return os.WriteFile(dst, b, os.ModePerm)
}
