package dbops

import "testing"

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
	err := AddUserCredential("fuqiang", "123")
	if err != nil {
		t.Logf("add user error is %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("fuqiang")
	if err != nil {
		t.Logf("get user error is %v", err)
	}
	t.Logf("pwd is %s", pwd)
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("fuqiang", "123")
	if err != nil {
		t.Logf("delete user error is %v", err)
	}
}
