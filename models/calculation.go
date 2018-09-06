package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Shikigami struct {
	ShikigamiCode string // Shikigami's Code
	ShikigamiName string // Shikigami' Name
	ShikigamiWin  int    // How much game did the Shikigami's win
	ShikigamiLose int    // How much game did the Shikigami's lose

}

func Test() {
	fmt.Println("yeah")
}

//基础查询
func CodeSearchInfo() string {
	sql := `SELECT * FROM base_shikigami WHERE shikigami_code = ?`
	return sql
}
