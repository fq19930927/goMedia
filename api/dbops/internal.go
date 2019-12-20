package dbops

import (
	"../defs"
	"database/sql"
	"log"
	"strconv"
	"sync"
)

func InsertSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmt, err := dbConn.Prepare("INSERT INTO sessions (sessoin_id, TTL, login_name) " +
		"VALUES (?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}
	defer dbConn.Close()
	return nil
}

//根据sessionId检索session
func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmt, err := dbConn.Prepare("SELECT login_name, TTL FROM sessions WHERE session_id = ?")
	if err != nil {
		return nil, err
	}
	var uname string
	var ttl string
	stmt.QueryRow(sid).Scan(&uname, &ttl)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.Username = uname
	} else {
		return nil, err
	}
	defer stmt.Close()
	return ss, nil
}

//检索所有session
func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmt, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	for rows.Next() {
		var id string
		var ttlstr string
		var login_name string
		if err = rows.Scan(&id, &ttlstr, &login_name); err != nil {
			log.Printf("retrive sessions error: %s", err)
			break
		}
		if ttl, err := strconv.ParseInt(ttlstr, 10, 64); err == nil {
			ss := &defs.SimpleSession{Username: login_name, TTL: ttl}
			m.Store(id, ss)
			log.Printf("session id is: %s, ttl is: %d", id, ss.TTL)
		}
	}
	return m, nil
}

//根据sessionId删除session
func DeleteSession(sid string) error {
	stmt, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id = ?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	if _, err = stmt.Exec(sid); err != nil {
		log.Printf("%s", err)
		return err
	}
	return nil
}
