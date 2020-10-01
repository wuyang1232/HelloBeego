package db_mysql

func QueryUse(name string)(int, error) {
	row := Db.QueryRow("select count(name) admin_num from user where name = ?", name)
	var admin_num int
	err := row.Scan(&admin_num)
	if err != nil {
		return 0, err
	}
	return admin_num, nil
}