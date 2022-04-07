package map01

func ModifyUser(user map[string]map[string]string, name string) {
	if user[name] != nil {
		//存在用户
		user[name]["n"] = name
		user[name]["P"] = "888888"
	} else {
		user[name] = make(map[string]string)
		user[name]["n"] = name
		user[name]["P"] = "123456"

	}

}
