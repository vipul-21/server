package apis

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Auth(credentials Credential) string {
	identity := authenticate(credentials)
	if identity == false {
		return "invalid credential"
	}
	return "valid credetials"
}

func authenticate(c Credential) bool {
	if c.Username == "demo" && c.Password == "pass" {
		return true
	}
	return false
}
