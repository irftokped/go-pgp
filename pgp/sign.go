package pgp

import (
	"bytes"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/clearsign"
)

func Sign(entity *openpgp.Entity, message []byte) ([]byte, error) {
	writer := new(bytes.Buffer)
	// reader := bytes.NewReader(message)
	dec, err := clearsign.Encode(writer, entity.PrivateKey, nil)
	if err != nil {
		return []byte{}, err
	}
	_, err = dec.Write(message)
	if err != nil {
		return []byte{}, err
	}
	err = dec.Close()
	if err != nil {
		return []byte{}, err
	}
	return writer.Bytes(), nil
}
