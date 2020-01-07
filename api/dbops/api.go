package dbops
import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"video_server/api/defs"
)

func openConn()*sql.DB{

}

func AddUserCredential(loginName string, pwd string) error{
	stmtIns, err :=dbConn.Prepare("INSERT INTO users(login_name, pwd) values(?,?)")
	if err !=nil{
		return err
	}
	_,err = stmtIns.Exec(loginName,pwd)
    stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string)(string, error){
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err !=nil{
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err :=dbConn.Prepare("DELETE FROM users WHERE login_name=? AND pwd=?")
	if err != nil {
		log.Printf("Delete user error: %s", err)
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

func