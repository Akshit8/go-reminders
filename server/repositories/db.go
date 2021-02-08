package repositories

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/Akshit8/go-reminders/server/models"
)

// dbConfig represents the config which is used when DB is initialized
type dbConfig struct {
	ID       int    `json:"id"`
	CheckSum string `json:"check_sum"`
}

// DB represents the application database (json file)
type DB struct {
	dbPath    string
	dbCfgPath string
	cfg       dbConfig
	db        []byte
}

// // NewDB creates a new instance of application file DB
func NewDB(dbPath string, dbCfgPath string) *DB {
	return &DB{dbPath: dbPath, dbCfgPath: dbCfgPath}
}

func (d *DB) Start() error {
	bs, err := d.read(d.dbCfgPath)
	if err != nil {
		return models.WrapError("could not read db config contents", err)
	}
	var cfg dbConfig
	if len(bs) == 0 {
		bs = []byte("{}")
	}
	err = json.Unmarshal(bs, &cfg)
	if err != nil {
		return models.WrapError("could not unmarshal db config", err)
	}

	bs, err = d.read(d.dbPath)
	if err != nil {
		return models.WrapError("could not read db contents", err)
	}

	d.db = bs
	if d.cfg.CheckSum == "" {
		checksum, err := genChecksum(bytes.NewReader(bs))
		if err != nil {
			return err
		}
		cfg.CheckSum = checksum
	}
	d.cfg = cfg
	return nil
}

// Read fetches a list of reminders by given ids
func (d *DB) Read(bs []byte) (int, error) {
	n, err := bytes.NewReader(d.db).Read(bs)
	if err != nil && err != io.EOF {
		return 0, models.WrapError("could not read db file bytes", err)
	}
	return n, nil
}

// genCheckSum generates check sum for a reader
func genChecksum(r io.Reader) (string, error) {
	hash := sha256.New()
	if _, err := io.Copy(hash, r); err != nil {
		return "", models.WrapError("could not copy db contents", err)
	}
	sum := hash.Sum(nil)
	return fmt.Sprintf("%x", sum), nil
}

// reads the contents of a db file
func (d *DB) read(path string) ([]byte, error) {
	dbFile, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if errors.Is(err, os.ErrNotExist) {
		dbFile, err = os.Create(path)
	}
	if err != nil {
		return nil, models.WrapError("could not open or create db file", err)
	}
	return ioutil.ReadAll(dbFile)
}
