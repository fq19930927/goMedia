package dbops

import (
	"../defs"
	"../util"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

//创建用户
func AddUserCredential(loginName string, pwd string) error {
	stmt, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) values (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

//通过loginName获取密码
func GetUserCredential(loginName string) (string, error) {
	stmt, err := dbConn.Prepare("SELECT pwd from users where login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	err = stmt.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmt.Close()
	return pwd, nil
}

//删除用户
func DeleteUser(loginName string, pwd string) error {
	stmt, err := dbConn.Prepare("DELETE FROM users where login_name = ? and pwd = ?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	_, err = stmt.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

//上传视频
func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	//create uuid
	vid, err := util.NewUUID()
	if err != nil {
		return nil, err
	}
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:06")
	stmt, err := dbConn.Prepare("INSERT INTO video_info (id, author_id, name, display_name)" +
		"values (?,?,?,?)")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	defer stmt.Close()
	return res, nil
}

//根据id获取video
func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmt, err := dbConn.Prepare("SELECT * FROM video_info WHERE id = ?")
	var aid int
	var dct string
	var name string
	err = stmt.QueryRow(vid).Scan(&aid, &dct, &name)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer stmt.Close()
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}
	return res, nil
}

func ListVideoInfo(uname string, from, to int) ([]*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare(`SELECT video_info.id, video_info.author_id, video_info.name, video_info.display_ctime FROM video_info 
		INNER JOIN users ON video_info.author_id = users.id
		WHERE users.login_name = ? AND video_info.create_time > FROM_UNIXTIME(?) AND video_info.create_time <= FROM_UNIXTIME(?) 
		ORDER BY video_info.create_time DESC`)

	var res []*defs.VideoInfo

	if err != nil {
		return res, err
	}

	rows, err := stmtOut.Query(uname, from, to)
	if err != nil {
		log.Printf("%s", err)
		return res, err
	}

	for rows.Next() {
		var id, name, ctime string
		var aid int
		if err := rows.Scan(&id, &aid, &name, &ctime); err != nil {
			return res, err
		}

		vi := &defs.VideoInfo{Id: id, AuthorId: aid, Name: name, DisplayCtime: ctime}
		res = append(res, vi)
	}

	defer stmtOut.Close()

	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}
