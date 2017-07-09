package store

import (
	"github.com/boltdb/bolt"
	"os/user"
	"path/filepath"
	"os"
	"fmt"
	"time"
)

const dir = ".waxt"

type Store struct {
	path string
	Db   *bolt.DB
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

	db, err := bolt.Open(s.path, 0755, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		return err
	}

	s.Db = db

	// Initialize all the required buckets.
	if err := s.Db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("Customer"))
		return nil
	}); err != nil {
		s.Close()
		return err
	}

	return nil
}

func (s *Store) Close() error {
	if s.Db != nil {
		s.Db.Close()
	}
	return nil
}