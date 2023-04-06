package samples

import (
	"encoding/json"
	"fmt"
)

type MyResponse struct {
	Key   string `json:"key"`
	Value int    `json:"value,omitempty"`
}

func (m MyResponse) UnmarshalJSON(b []byte) error {
	type T struct{ Name string }
	var t T
	err := json.Unmarshal(b, &t)
	t.Name = m.Key + "wao"
	return err
}

// func (m MyResponse) MarshalJSON() ([]byte, error) {
// 	v, err := json.Marshal(&struct {
// 		Name string
// 	}{
// 		Name: "hoge",
// 	})
// 	return v, err
// }

func CallMarshal() {
	by := []byte(`{"key": "item", "value": 20}`)

	p := new(MyResponse)
	er := json.Unmarshal(by, &p)

	fmt.Println(er, p.Key, p.Value)

	v, _ := json.Marshal(p)
	fmt.Println(string(v))
}
