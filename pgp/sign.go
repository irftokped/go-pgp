package pgp

import (
	"bytes"
	"golang.org/x/crypto/openpgp"
)

func Sign(entity *openpgp.Entity, message []byte) ([]byte, error) {
	writer := new(bytes.Buffer)
	// reader := bytes.NewReader(message)
	w, err := openpgp.Sign(writer, entity, nil, nil)
	if err != nil {
		return []byte{}, err
	}
	_, err = w.Write(message)
	if err != nil {
		return []byte{}, err
	}
	err = w.Close()
	// err := openpgp.ArmoredDetachSign(writer, entity, reader, nil)
	if err != nil {
		return []byte{}, err
	}
	return writer.Bytes(), nil
}
