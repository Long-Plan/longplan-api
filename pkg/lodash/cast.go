package lodash

import (
	"github.com/TylerBrock/colorjson"
	"github.com/bytedance/sonic"
)

func Recast(from, to interface{}) error {
	switch v := from.(type) {
	case []byte:
		return sonic.Unmarshal(v, to)
	default:
		buf, err := sonic.Marshal(from)
		if err != nil {
			return err
		}

		return sonic.Unmarshal(buf, to)
	}
}

func PrettyJson(from interface{}) string {
	obj := map[string]interface{}{}
	_ = Recast(from, &obj)
	f := colorjson.NewFormatter()
	buff, _ := f.Marshal(obj)
	return string(buff)
}
