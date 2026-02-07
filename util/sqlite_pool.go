package util

import (
	"database/sql"
	"log"
	"sync"
)

type SqlitePoolConn struct {
	Db *sql.DB

	inPool bool
	prev   *SqlitePoolConn
}

type SqlitePool struct {
	list  *SqlitePoolConn
	path  string
	mutex sync.Mutex
}

func SqlitePoolNew(path string, initSize int) (*SqlitePool, error) {
	pool := new(SqlitePool)
	for range initSize {
		db, err := openSqliteConn(path)
		if err != nil {
			pool.closeAll()
			return nil, err
		}

		conn := new(SqlitePoolConn)
		conn.inPool = true
		conn.Db = db
		pool.push(conn)
	}

	pool.path = path
	return pool, nil
}

func (s *SqlitePool) Destory() {
	s.closeAll()
}

func (s *SqlitePool) GetConn() *SqlitePoolConn {
	s.mutex.Lock()
	conn := s.pop()
	s.mutex.Unlock()

	if conn == nil {
		conn = new(SqlitePoolConn)
		db, err := openSqliteConn(s.path)
		if err != nil {
			return nil
		}

		conn.Db = db
	}

	return conn
}

func (s *SqlitePool) PutConn(conn *SqlitePoolConn) {
	if !conn.inPool {
		conn.Db.Close()
		return
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.push(conn)
}

func (s *SqlitePool) push(conn *SqlitePoolConn) {
	conn.prev = s.list
	s.list = conn
	log.Printf("putting db conn: %p -> %+v", conn, conn)
}

func (s *SqlitePool) pop() *SqlitePoolConn {
	item := s.list
	if item == nil {
		return nil
	}

	log.Printf("getting db conn: %p -> %+v", item, item)
	s.list = item.prev
	return item
}

func (s *SqlitePool) closeAll() {
	for {
		conn := s.pop()
		if conn == nil {
			return
		}

		log.Println("closing db conn:", conn)
		conn.Db.Close()
	}
}

func openSqliteConn(path string) (*sql.DB, error) {
	ret, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	query := "PRAGMA journal_mode=WAL"
	_, err = ret.Exec(query)
	if err != nil {
		ret.Close()
		return nil, err
	}

	return ret, nil
}
