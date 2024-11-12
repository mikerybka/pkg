package encryptedfs

import (
	"os"
	"path/filepath"
)

func Encrypt(key, src, dst string) error {
	err := os.MkdirAll(dst, os.ModePerm)
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, e := range entries {
		nextSrc := filepath.Join(src, e.Name())
		nextDst := filepath.Join(dst, encrypt(key, e.Name()))

		if e.IsDir() {
			err = Encrypt(key, nextSrc, nextDst)
			if err != nil {
				return err
			}
			continue
		}

		b, err := os.ReadFile(nextSrc)
		if err != nil {
			return err
		}

		err = os.WriteFile(nextDst, []byte(encrypt(key, string(b))), os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
