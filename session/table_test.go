package session

type User struct {
	Name string `orm:"PRIMARY KEY"`
	Age  int
}

// func TestSession_CreateTable(t *testing.T) {
// 	s := NewSession().Model(&User{})
// 	_ = s.DropTable()
// 	_ = s.CreateTable()
// 	if !s.HasTable() {
// 		t.Fatal("Failed to create table User")
// 	}
// }
