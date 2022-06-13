package jsondemo

type TestJsonObj struct {
	Id       int      `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	ListObj  []string `json:"listObj"'`
}
