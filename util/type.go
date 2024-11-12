package util

import (
	"strings"
)

type Type struct {
	ID          string     `json:"id"`
	Desc        string     `json:"desc"`
	Comments    []Comment  `json:"comments"`
	Name        Name       `json:"name"`
	PluralName  Name       `json:"pluralName"`
	IsScalar    bool       `json:"isScalar"`
	IsArray     bool       `json:"isArray"`
	IsMap       bool       `json:"isMap"`
	ElemType    string     `json:"elemType"`
	IsStruct    bool       `json:"isStruct"`
	Fields      []Field    `json:"fields"`
	Methods     []Function `json:"methods"`
	DefaultJSON string     `json:"defaultJSON"`
}

func (t *Type) Imports() ImportMap {
	imports := ImportMap{}

	typ := t.ElemType
	for {
		if strings.HasPrefix(typ, "[]") {
			typ = strings.TrimPrefix(typ, "[]")
		} else if strings.HasPrefix(typ, "map[string]") {
			typ = strings.TrimPrefix(typ, "map[string]")
		} else {
			break
		}
	}
	from, _ := parseName(typ)
	if from != "" {
		imports[from] = importPath(from)
	}

	for _, f := range t.Fields {
		typ := f.Type
		for {
			if strings.HasPrefix(typ, "[]") {
				typ = strings.TrimPrefix(typ, "[]")
			} else if strings.HasPrefix(typ, "map[string]") {
				typ = strings.TrimPrefix(typ, "map[string]")
			} else {
				break
			}
		}
		from, _ := parseName(typ)
		if from != "" {
			imports[from] = importPath(from)
		}
	}

	for _, m := range t.Methods {
		for k, v := range m.Imports() {
			imports[k] = v
		}
	}

	return imports
}

func parseName(s string) (string, string) {
	n := strings.Split(s, ".")
	if len(n) == 1 {
		return "", n[0]
	}
	return n[0], n[1]
}

func importPath(from string) string {
	switch from {
	case "builtin":
		return "builtin"
	case "heap":
		return "container/heap"
	case "list":
		return "container/list"
	case "ring":
		return "container/ring"
	case "crypto":
		return "crypto"
	case "aes":
		return "crypto/aes"
	case "cipher":
		return "crypto/cipher"
	case "des":
		return "crypto/des"
	case "dsa":
		return "crypto/dsa"
	case "ecdsa":
		return "crypto/ecdsa"
	case "ed25519":
		return "crypto/ed25519"
	case "elliptic":
		return "crypto/elliptic"
	case "hmac":
		return "crypto/hmac"
	case "md5":
		return "crypto/md5"
	case "rc4":
		return "crypto/rc4"
	case "rsa":
		return "crypto/rsa"
	case "sha1":
		return "crypto/sha1"
	case "sha256":
		return "crypto/sha256"
	case "sha512":
		return "crypto/sha512"
	case "subtle":
		return "crypto/subtle"
	case "tls":
		return "crypto/tls"
	case "x509":
		return "crypto/x509"
	case "pkix":
		return "crypto/x509/pkix"
	case "sql":
		return "database/sql"
	case "driver":
		return "database/sql/driver"
	case "buildinfo":
		return "debug/buildinfo"
	case "dwarf":
		return "debug/dwarf"
	case "elf":
		return "debug/elf"
	case "gosym":
		return "debug/gosym"
	case "macho":
		return "debug/macho"
	case "pe":
		return "debug/pe"
	case "plan9obj":
		return "debug/plan9obj"
	case "encoding":
		return "encoding"
	case "ascii85":
		return "encoding/ascii85"
	case "asn1":
		return "encoding/asn1"
	case "base32":
		return "encoding/base32"
	case "base64":
		return "encoding/base64"
	case "binary":
		return "encoding/binary"
	case "csv":
		return "encoding/csv"
	case "gob":
		return "encoding/gob"
	case "hex":
		return "encoding/hex"
	case "json":
		return "encoding/json"
	case "pem":
		return "encoding/pem"
	case "xml":
		return "encoding/xml"
	case "errors":
		return "errors"
	case "expvar":
		return "expvar"
	case "flag":
		return "flag"
	case "fmt":
		return "fmt"
	case "ast":
		return "go/ast"
	case "build":
		return "go/build"
	case "constraint":
		return "go/build/constraint"
	case "constant":
		return "go/constant"
	case "doc":
		return "go/doc"
	case "format":
		return "go/format"
	case "importer":
		return "go/importer"
	case "parser":
		return "go/parser"
	case "printer":
		return "go/printer"
	case "token":
		return "go/token"
	case "types":
		return "go/types"
	case "hash":
		return "hash"
	case "adler32":
		return "hash/adler32"
	case "crc32":
		return "hash/crc32"
	case "crc64":
		return "hash/crc64"
	case "fnv":
		return "hash/fnv"
	case "html":
		return "html"
	case "image":
		return "image"
	case "color":
		return "image/color"
	case "palette":
		return "image/color/palette"
	case "draw":
		return "image/draw"
	case "gif":
		return "image/gif"
	case "jpeg":
		return "image/jpeg"
	case "png":
		return "image/png"
	case "suffixarray":
		return "index/suffixarray"
	case "io":
		return "io"
	case "fs":
		return "io/fs"
	case "ioutil":
		return "io/ioutil"
	case "log":
		return "log"
	case "syslog":
		return "log/syslog"
	case "math":
		return "math"
	case "big":
		return "math/big"
	case "bits":
		return "math/bits"
	case "cmplx":
		return "math/cmplx"
	case "rand":
		return "math/rand"
	case "mime":
		return "mime"
	case "multipart":
		return "mime/multipart"
	case "quotedprintable":
		return "mime/quotedprintable"
	case "net":
		return "net"
	case "http":
		return "net/http"
	case "cgi":
		return "net/http/cgi"
	case "cookiejar":
		return "net/http/cookiejar"
	case "fcgi":
		return "net/http/fcgi"
	case "httptest":
		return "net/http/httptest"
	case "httptrace":
		return "net/http/httptrace"
	case "httputil":
		return "net/http/httputil"
	case "mail":
		return "net/mail"
	case "rpc":
		return "net/rpc"
	case "jsonrpc":
		return "net/rpc/jsonrpc"
	case "smtp":
		return "net/smtp"
	case "textproto":
		return "net/textproto"
	case "url":
		return "net/url"
	case "os":
		return "os"
	case "exec":
		return "os/exec"
	case "signal":
		return "os/signal"
	case "user":
		return "os/user"
	case "path":
		return "path"
	case "filepath":
		return "path/filepath"
	case "plugin":
		return "plugin"
	case "reflect":
		return "reflect"
	case "regexp":
		return "regexp"
	case "syntax":
		return "regexp/syntax"
	case "runtime":
		return "runtime"
	case "cgo":
		return "runtime/cgo"
	case "debug":
		return "runtime/debug"
	case "metrics":
		return "runtime/metrics"
	case "pprof":
		return "runtime/pprof"
	case "race":
		return "runtime/race"
	case "trace":
		return "runtime/trace"
	case "sort":
		return "sort"
	case "strconv":
		return "strconv"
	case "strings":
		return "strings"
	case "sync":
		return "sync"
	case "atomic":
		return "sync/atomic"
	case "syscall":
		return "syscall"
	case "testing":
		return "testing"
	case "fstest":
		return "testing/fstest"
	case "iotest":
		return "testing/iotest"
	case "quick":
		return "testing/quick"
	case "scanner":
		return "text/scanner"
	case "tabwriter":
		return "text/tabwriter"
	case "template":
		return "text/template"
	case "parse":
		return "text/template/parse"
	case "time":
		return "time"
	case "tzdata":
		return "time/tzdata"
	case "unicode":
		return "unicode"
	case "utf16":
		return "unicode/utf16"
	case "utf8":
		return "unicode/utf8"
	case "unsafe":
		return "unsafe"
	default:
		return "github.com/mikerybka/pkg/" + from
	}
}
