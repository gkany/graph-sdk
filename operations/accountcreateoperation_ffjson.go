// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: accountcreateoperation.go

package operations

import (
	"bytes"
	"fmt"
	"github.com/denkhaus/bitshares/types"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *AccountCreateOperation) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *AccountCreateOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "registrar":`)

	{

		obj, err = j.Registrar.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"referrer":`)

	{

		obj, err = j.Referrer.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"referrer_percent":`)
	fflib.FormatBits2(buf, uint64(j.ReferrerPercent), 10, false)
	buf.WriteString(`,"owner":`)

	{

		err = j.Owner.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteString(`,"active":`)

	{

		err = j.Active.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteString(`,"name":`)

	{

		obj, err = j.Name.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"extensions":`)

	{

		err = j.Extensions.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteString(`,"options":`)

	{

		err = j.Options.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteByte(',')
	if j.Fee != nil {
		if true {
			buf.WriteString(`"fee":`)

			{

				err = j.Fee.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtAccountCreateOperationbase = iota
	ffjtAccountCreateOperationnosuchkey

	ffjtAccountCreateOperationRegistrar

	ffjtAccountCreateOperationReferrer

	ffjtAccountCreateOperationReferrerPercent

	ffjtAccountCreateOperationOwner

	ffjtAccountCreateOperationActive

	ffjtAccountCreateOperationName

	ffjtAccountCreateOperationExtensions

	ffjtAccountCreateOperationOptions

	ffjtAccountCreateOperationFee
)

var ffjKeyAccountCreateOperationRegistrar = []byte("registrar")

var ffjKeyAccountCreateOperationReferrer = []byte("referrer")

var ffjKeyAccountCreateOperationReferrerPercent = []byte("referrer_percent")

var ffjKeyAccountCreateOperationOwner = []byte("owner")

var ffjKeyAccountCreateOperationActive = []byte("active")

var ffjKeyAccountCreateOperationName = []byte("name")

var ffjKeyAccountCreateOperationExtensions = []byte("extensions")

var ffjKeyAccountCreateOperationOptions = []byte("options")

var ffjKeyAccountCreateOperationFee = []byte("fee")

// UnmarshalJSON umarshall json - template of ffjson
func (j *AccountCreateOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *AccountCreateOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtAccountCreateOperationbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtAccountCreateOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyAccountCreateOperationActive, kn) {
						currentKey = ffjtAccountCreateOperationActive
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffjKeyAccountCreateOperationExtensions, kn) {
						currentKey = ffjtAccountCreateOperationExtensions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyAccountCreateOperationFee, kn) {
						currentKey = ffjtAccountCreateOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'n':

					if bytes.Equal(ffjKeyAccountCreateOperationName, kn) {
						currentKey = ffjtAccountCreateOperationName
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'o':

					if bytes.Equal(ffjKeyAccountCreateOperationOwner, kn) {
						currentKey = ffjtAccountCreateOperationOwner
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountCreateOperationOptions, kn) {
						currentKey = ffjtAccountCreateOperationOptions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffjKeyAccountCreateOperationRegistrar, kn) {
						currentKey = ffjtAccountCreateOperationRegistrar
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountCreateOperationReferrer, kn) {
						currentKey = ffjtAccountCreateOperationReferrer
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountCreateOperationReferrerPercent, kn) {
						currentKey = ffjtAccountCreateOperationReferrerPercent
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyAccountCreateOperationFee, kn) {
					currentKey = ffjtAccountCreateOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountCreateOperationOptions, kn) {
					currentKey = ffjtAccountCreateOperationOptions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountCreateOperationExtensions, kn) {
					currentKey = ffjtAccountCreateOperationExtensions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAccountCreateOperationName, kn) {
					currentKey = ffjtAccountCreateOperationName
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAccountCreateOperationActive, kn) {
					currentKey = ffjtAccountCreateOperationActive
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAccountCreateOperationOwner, kn) {
					currentKey = ffjtAccountCreateOperationOwner
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyAccountCreateOperationReferrerPercent, kn) {
					currentKey = ffjtAccountCreateOperationReferrerPercent
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAccountCreateOperationReferrer, kn) {
					currentKey = ffjtAccountCreateOperationReferrer
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountCreateOperationRegistrar, kn) {
					currentKey = ffjtAccountCreateOperationRegistrar
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtAccountCreateOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtAccountCreateOperationRegistrar:
					goto handle_Registrar

				case ffjtAccountCreateOperationReferrer:
					goto handle_Referrer

				case ffjtAccountCreateOperationReferrerPercent:
					goto handle_ReferrerPercent

				case ffjtAccountCreateOperationOwner:
					goto handle_Owner

				case ffjtAccountCreateOperationActive:
					goto handle_Active

				case ffjtAccountCreateOperationName:
					goto handle_Name

				case ffjtAccountCreateOperationExtensions:
					goto handle_Extensions

				case ffjtAccountCreateOperationOptions:
					goto handle_Options

				case ffjtAccountCreateOperationFee:
					goto handle_Fee

				case ffjtAccountCreateOperationnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Registrar:

	/* handler: j.Registrar type=types.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Registrar.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Referrer:

	/* handler: j.Referrer type=types.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Referrer.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ReferrerPercent:

	/* handler: j.ReferrerPercent type=types.UInt16 kind=uint16 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.ReferrerPercent.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Owner:

	/* handler: j.Owner type=types.Authority kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			err = j.Owner.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Active:

	/* handler: j.Active type=types.Authority kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			err = j.Active.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Name:

	/* handler: j.Name type=types.String kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Name.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Extensions:

	/* handler: j.Extensions type=types.AccountCreateExtensions kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			err = j.Extensions.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Options:

	/* handler: j.Options type=types.AccountOptions kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			err = j.Options.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Fee:

	/* handler: j.Fee type=types.AssetAmount kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			j.Fee = nil

		} else {

			if j.Fee == nil {
				j.Fee = new(types.AssetAmount)
			}

			err = j.Fee.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
