package csv_parser

import (
	"encoding/csv"
	"reflect"
	"strconv"
)

type FieldMismatch struct {
	expected, found int
}

func (e *FieldMismatch) Error() string {
	return "CSV line fields mismatch. Expected " + strconv.Itoa(e.expected) + " found " + strconv.Itoa(e.found)
}

type UnsupportedType struct {
	Type string
}

func (e *UnsupportedType) Error() string {
	return "Unsupported type: " + e.Type
}


// https://play.golang.org/p/kwc32A5mJf
func UnmarshalCsvRow(reader *csv.Reader, v interface{}) error {
    record, err := reader.Read()
    if err != nil {
        return err
    }
    s := reflect.ValueOf(v).Elem()
    if s.NumField() != len(record) {
        return &FieldMismatch{s.NumField(), len(record)}
    }
    for i := 0; i < s.NumField(); i++ {
        f := s.Field(i)
        switch f.Type().String() {
        case "string":
            f.SetString(record[i])
        case "int":
            ival, err := strconv.ParseInt(record[i], 10, 0)
            if err != nil {
                return err
            }
            f.SetInt(ival)
		case "float64":
			fval, err := strconv.ParseFloat(record[i], 64)
			if err != nil {
                return err
            }
            f.SetFloat(fval)
        default:
            return &UnsupportedType{f.Type().String()}
        }
    }
    return nil
}
