package go_nik_parser

import "encoding/json"

func marshaler(a interface{}) (map[string]interface{}, error) {
	marshal, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	var b map[string]interface{}
	err = json.Unmarshal(marshal, &b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func convBulan(b int) string {
	switch b {
	case 1:
		return "January"
	case 2:
		return "February"
	case 3:
		return "March"
	case 4:
		return "April"
	case 5:
		return "May"
	case 6:
		return "June"
	case 7:
		return "July"
	case 8:
		return "August"
	case 9:
		return "Sepember"
	case 10:
		return "October"
	case 11:
		return "November"
	case 12:
		return "December"
	default:
		return ""
	}
}
