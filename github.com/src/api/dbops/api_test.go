package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("add", testAddUser)
	t.Run("get", testGetUser)
	t.Run("delete", testDeleteUser)
	t.Run("get", testGetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("star", "123")
	if err != nil {
		t.Logf("add user error is %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("star")
	if err != nil {
		t.Logf("get user error is %v", err)
	}
	t.Logf("pwd is %s", pwd)
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("star", "123")
	if err != nil {
		t.Logf("delete user error is %v", err)
	}
}

func TestVideoWorkFlow(t *testing.T) {

}

var tempvid string

func testAddVideInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Eroor or AddVideoInfo: %v", err)
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Eroor or GetVideoInfo: %v", err)
	}
}

func TestDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Eroor or DeleteVideoInfo: %v", err)
	}
}

func TestCommentsFlow(t *testing.T) {

}

func testAddComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I like this video"

	err := AddNewComments(vid, aid, content)

	if err != nil {
		t.Errorf("Error of AddComments: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}
	for i, ele := range res {
		fmt.Printf("comment: %d, %v \n", i, ele)
	}
}
