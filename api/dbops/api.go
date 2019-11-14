package dbops
import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"video_server/api/defs"
)


func AddUserCredential(loginName string, pwd string) error{
	stmtIns, err :=dbConn.Prepare("INSERT INTO users(login_name, pwd) values(?,?)")
	if err !=nil{
		return err
	}
	_,err = stmtIns.Exec(loginName,pwd)

}