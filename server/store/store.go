package store

import (
	"os/user"
	"path/filepath"
	"os"
	"fmt"
	"time"
	"github.com/asdine/storm"
	"github.com/boltdb/bolt"
)

const dir = ".waxt"

type Store struct {
	path string
	Db   *storm.DB
}

func NewStore(db string) *Store {
	return &Store{
		path: dbPath(db),
	}
}

func path() string {
	current, _ := user.Current()
	return current.HomeDir + string(filepath.Separator) + dir
}

func dbPath(dbName string) string {
	return path() + string(filepath.Separator) + dbName
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func (s *Store) Open() error {
	exist, err := exists(path())

	if err != nil {
		fmt.Print("failed to open db", err)
		os.Exit(1)
	}

	if !exist {
		os.Mkdir(path(), 0755)
	}

	db, err := storm.Open(s.path, storm.BoltOptions(0600, &bolt.Options{Timeout: 1 * time.Second}))

	if err != nil {
		return err
	}

	s.Db = db

	return nil
}

func (s *Store) Close() error {
	if s.Db != nil {
		s.Db.Close()
	}
	return nil
}