package pkg

import (
	"encoding/json"
	"io"
	"os"

	model "blixza.github.io/model/record"
)

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func Load(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
}

func Write(file *os.File, record model.Record) error {
	var records []model.Record

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if len(byteValue) > 0 {
		err := json.Unmarshal(byteValue, &records)
		if err != nil {
			return err
		}
	}

	records = append(records, record)

	updatedData, err := json.MarshalIndent(records, "", " ")
	if err != nil {
		return err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	err = file.Truncate(0)
	if err != nil {
		return err
	}

	_, err = file.Write(updatedData)
	return err
}
