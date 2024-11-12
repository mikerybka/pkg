package encryptedfs

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/mikerybka/pkg/util"
)

type Server struct {
	Key string
	Dir string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Encrypt path
	path := util.ParsePath(r.URL.Path)
	for i, p := range path {
		path[i] = encrypt(s.Key, p)
	}
	encryptedPath := filepath.Join(s.Dir, strings.Join(path, "/"))

	fmt.Println(encryptedPath)

	if r.Method == "GET" {
		if util.IsDir(encryptedPath) {
			entries, err := os.ReadDir(encryptedPath)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, "<html><body><pre>\n")
			for _, e := range entries {
				decryptedName := decrypt(s.Key, e.Name())
				fmt.Fprintf(w, "<a href=\"%s\">%s</a>\n", decryptedName, decryptedName)
			}
			fmt.Fprintf(w, "</pre></body></html>\n")
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			return
		}

		http.ServeFile(w, r, encryptedPath)
		return
	}

	if r.Method == "PUT" {
		err := os.MkdirAll(filepath.Dir(encryptedPath), os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		d := encrypt(s.Key, string(b))
		err = os.WriteFile(encryptedPath, []byte(d), os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func encrypt(key, s string) string {
	// Convert the key and plaintext to byte slices.
	keyBytes := util.SHA256([]byte(key))
	plaintextBytes := []byte(s)

	// Create a new AES cipher block with the key.
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err)
	}

	// Create a GCM (Galois/Counter Mode) cipher from the AES block.
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	// Generate a nonce. It must be unique for each encryption.
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	// Encrypt the plaintext using AES-GCM with the nonce.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintextBytes, nil)

	// Encode the ciphertext to base64 for easy display or storage.
	return base64.URLEncoding.EncodeToString(ciphertext)
}

func decrypt(key, s string) string {
	// Convert the key and ciphertext to byte slices.
	keyBytes := util.SHA256([]byte(key))
	ciphertext, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}

	// Create a new AES cipher block with the key.
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err)
	}

	// Create a GCM cipher from the AES block.
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	// Extract the nonce from the beginning of the ciphertext.
	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt the ciphertext.
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	return string(plaintext)
}
