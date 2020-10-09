package db_mysql

import ("HelloBeego190604/models"
		_"大一下学期/github.com/go-sql-driver/mysql"
)

func QueryUse(name string)(int, error) {
	row := Db.QueryRow("select count(name) admin_num from user where name = ?", name)
	var admin_num int
	err := row.Scan(&admin_num)
	if err != nil {
		return 0, err
	}
	return admin_num, nil
}
func QueryAllHero(nam string)([]models.User,error){
	rows,err := Db.Query("select * from user where name = ?",nam)
	if err != nil{
		return nil,err
	}
	users := make([]models.User,0)
	for rows.Next(){
		var user models.User
		err = rows.Scan(&user.Name,&user.Birthday,&user.Address,&user.Password)
		if err != nil{
			return nil,err
		}
		users = append(users,user)
	}
	return users,nil
}
