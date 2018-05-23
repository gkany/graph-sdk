package types

//go:generate ffjson $GOFILE

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"strconv"

	"github.com/btcsuite/btcd/btcec"

	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

type Memo struct {
	From    PublicKey `json:"from"`
	To      PublicKey `json:"to"`
	Nonce   UInt64    `json:"nonce"`
	Message Buffer    `json:"message"`
}

func (p Memo) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(p.From); err != nil {
		return errors.Annotate(err, "encode from")
	}

	if err := enc.Encode(p.To); err != nil {
		return errors.Annotate(err, "encode to")
	}

	if err := enc.Encode(p.Nonce); err != nil {
		return errors.Annotate(err, "encode nonce")
	}

	if err := enc.Encode(p.Message); err != nil {
		return errors.Annotate(err, "encode Message")
	}

	return nil
}

//Encrypt encodes the memo message
func (p *Memo) Encrypt(priv *btcec.PrivateKey, msg string) error {
	sec, err := util.SharedSecret(priv, p.To.ToECDSA(), 16, 16)
	if err != nil {
		return errors.Annotate(err, "SharedSecret")
	}

	iv, blk, err := p.cypherBlock(sec)
	if err != nil {
		return errors.Annotate(err, "cypherBlock")
	}

	mode := cipher.NewCBCEncrypter(blk, iv)

	buf := []byte(msg)
	digest := sha256.Sum256(buf)

	// checksum + msg
	raw := digest[:4]
	raw = append(raw, buf...)

	if len(raw)%16 != 0 {
		raw = pad(raw, 16)
	}

	dst := make([]byte, len(raw))
	mode.CryptBlocks(dst, raw)
	p.Message = dst

	return nil

}

//Decrypt decodes the memo message
func (p Memo) Decrypt(priv *btcec.PrivateKey) (string, error) {
	sec, err := util.SharedSecret(priv, p.To.ToECDSA(), 16, 16)
	if err != nil {
		return "", errors.Annotate(err, "SharedSecret")
	}

	iv, blk, err := p.cypherBlock(sec)
	if err != nil {
		return "", errors.Annotate(err, "cypherBlock")
	}

	mode := cipher.NewCBCDecrypter(blk, iv)
	dst := make([]byte, len(p.Message))
	mode.CryptBlocks(dst, p.Message)

	//verify checksum
	chk1 := dst[:4]
	msg := unpad(dst[4:])
	dig := sha256.Sum256(msg)
	chk2 := dig[:4]

	if bytes.Compare(chk1, chk2) != 0 {
		return "", ErrInvalidChecksum
	}

	return string(msg), nil
}

func (p Memo) cypherBlock(sec []byte) ([]byte, cipher.Block, error) {
	ss := sha512.Sum512(sec)

	var seed []byte
	seed = append(seed, []byte(strconv.FormatUint(uint64(p.Nonce), 10))...)
	seed = append(seed, []byte(hex.EncodeToString(ss[:]))...)

	sd := sha512.Sum512(seed)
	blk, err := aes.NewCipher(sd[0:32])
	if err != nil {
		return nil, nil, errors.Annotate(err, "NewCipher")
	}

	return sd[32:48], blk, nil
}

func unpad(buf []byte) []byte {
	b := buf[len(buf)-1:][0]
	cnt := int(b)
	l := len(buf) - cnt

	a := bytes.Repeat([]byte{b}, cnt)
	if bytes.Compare(a, buf[l:]) == 0 {
		return buf[:l]
	}

	return buf
}

func pad(buf []byte, length int) []byte {
	cnt := length - len(buf)%length
	buf = append(buf, bytes.Repeat([]byte{byte(cnt)}, cnt)...)
	return buf
}
