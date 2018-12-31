// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: bidcolatteraloperation.go

package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/denkhaus/bitshares/types"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *BidCollateralOperation) MarshalJSON() ([]byte, error) {
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
func (j *BidCollateralOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "additional_collateral":`)

	{

		err = j.AdditionalCollateral.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteString(`,"bidder":`)

	{

		obj, err = j.Bidder.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"debt_covered":`)

	{

		err = j.DebtCovered.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteString(`,"extensions":`)
	if j.Extensions != nil {
		buf.WriteString(`[`)
		for i, v := range j.Extensions {
			if i != 0 {
				buf.WriteString(`,`)
			}
			/* Interface types must use runtime reflection. type=interface {} kind=interface */
			err = buf.Encode(v)
			if err != nil {
				return err
			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
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
	ffjtBidCollateralOperationbase = iota
	ffjtBidCollateralOperationnosuchkey

	ffjtBidCollateralOperationAdditionalCollateral

	ffjtBidCollateralOperationBidder

	ffjtBidCollateralOperationDebtCovered

	ffjtBidCollateralOperationExtensions

	ffjtBidCollateralOperationFee
)

var ffjKeyBidCollateralOperationAdditionalCollateral = []byte("additional_collateral")

var ffjKeyBidCollateralOperationBidder = []byte("bidder")

var ffjKeyBidCollateralOperationDebtCovered = []byte("debt_covered")

var ffjKeyBidCollateralOperationExtensions = []byte("extensions")

var ffjKeyBidCollateralOperationFee = []byte("fee")

// UnmarshalJSON umarshall json - template of ffjson
func (j *BidCollateralOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *BidCollateralOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtBidCollateralOperationbase
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
				currentKey = ffjtBidCollateralOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyBidCollateralOperationAdditionalCollateral, kn) {
						currentKey = ffjtBidCollateralOperationAdditionalCollateral
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'b':

					if bytes.Equal(ffjKeyBidCollateralOperationBidder, kn) {
						currentKey = ffjtBidCollateralOperationBidder
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'd':

					if bytes.Equal(ffjKeyBidCollateralOperationDebtCovered, kn) {
						currentKey = ffjtBidCollateralOperationDebtCovered
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffjKeyBidCollateralOperationExtensions, kn) {
						currentKey = ffjtBidCollateralOperationExtensions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyBidCollateralOperationFee, kn) {
						currentKey = ffjtBidCollateralOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyBidCollateralOperationFee, kn) {
					currentKey = ffjtBidCollateralOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBidCollateralOperationExtensions, kn) {
					currentKey = ffjtBidCollateralOperationExtensions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyBidCollateralOperationDebtCovered, kn) {
					currentKey = ffjtBidCollateralOperationDebtCovered
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyBidCollateralOperationBidder, kn) {
					currentKey = ffjtBidCollateralOperationBidder
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyBidCollateralOperationAdditionalCollateral, kn) {
					currentKey = ffjtBidCollateralOperationAdditionalCollateral
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtBidCollateralOperationnosuchkey
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

				case ffjtBidCollateralOperationAdditionalCollateral:
					goto handle_AdditionalCollateral

				case ffjtBidCollateralOperationBidder:
					goto handle_Bidder

				case ffjtBidCollateralOperationDebtCovered:
					goto handle_DebtCovered

				case ffjtBidCollateralOperationExtensions:
					goto handle_Extensions

				case ffjtBidCollateralOperationFee:
					goto handle_Fee

				case ffjtBidCollateralOperationnosuchkey:
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

handle_AdditionalCollateral:

	/* handler: j.AdditionalCollateral type=types.AssetAmount kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			err = j.AdditionalCollateral.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Bidder:

	/* handler: j.Bidder type=types.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Bidder.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_DebtCovered:

	/* handler: j.DebtCovered type=types.AssetAmount kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			err = j.DebtCovered.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Extensions:

	/* handler: j.Extensions type=types.Extensions kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for Extensions", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Extensions = nil
		} else {

			j.Extensions = []interface{}{}

			wantVal := true

			for {

				var tmpJExtensions interface{}

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJExtensions type=interface {} kind=interface quoted=false*/

				{
					/* Falling back. type=interface {} kind=interface */
					tbuf, err := fs.CaptureField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}

					err = json.Unmarshal(tbuf, &tmpJExtensions)
					if err != nil {
						return fs.WrapErr(err)
					}
				}

				j.Extensions = append(j.Extensions, tmpJExtensions)

				wantVal = false
			}
		}
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
