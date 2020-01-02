// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: dynamicglobalproperties.go

package types

import (
	"bytes"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *DynamicGlobalProperties) MarshalJSON() ([]byte, error) {
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
func (j *DynamicGlobalProperties) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"id":`)

	{

		obj, err = j.ID.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"head_block_number":`)
	fflib.FormatBits2(buf, uint64(j.HeadBlockNumber), 10, false)
	buf.WriteString(`,"head_block_id":`)

	{

		obj, err = j.HeadBlockID.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"time":`)

	{

		obj, err = j.Time.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"current_witness":`)

	{

		obj, err = j.CurrentWitness.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"current_transaction_count":`)
	fflib.FormatBits2(buf, uint64(j.CurrentTransactionCount), 10, false)
	buf.WriteString(`,"next_maintenance_time":`)

	{

		obj, err = j.NextMaintenanceTime.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"last_budget_time":`)

	{

		obj, err = j.LastBudgetTime.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"witness_budget":`)

	{

		obj, err = j.WitnessBudget.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"accounts_registered_this_interval":`)
	fflib.FormatBits2(buf, uint64(j.AccountsRegisteredThisInterval), 10, j.AccountsRegisteredThisInterval < 0)
	buf.WriteString(`,"recently_missed_count":`)
	fflib.FormatBits2(buf, uint64(j.RecentlyMissedCount), 10, j.RecentlyMissedCount < 0)
	buf.WriteString(`,"current_aslot":`)
	fflib.FormatBits2(buf, uint64(j.CurrentAslot), 10, j.CurrentAslot < 0)
	buf.WriteString(`,"recent_slots_filled":`)

	{

		obj, err = j.RecentSlotsFilled.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"dynamic_flags":`)
	fflib.FormatBits2(buf, uint64(j.DynamicFlags), 10, j.DynamicFlags < 0)
	buf.WriteString(`,"last_irreversible_block_num":`)
	fflib.FormatBits2(buf, uint64(j.LastIrreversibleBlockNum), 10, false)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtDynamicGlobalPropertiesbase = iota
	ffjtDynamicGlobalPropertiesnosuchkey

	ffjtDynamicGlobalPropertiesID

	ffjtDynamicGlobalPropertiesHeadBlockNumber

	ffjtDynamicGlobalPropertiesHeadBlockID

	ffjtDynamicGlobalPropertiesTime

	ffjtDynamicGlobalPropertiesCurrentWitness

	ffjtDynamicGlobalPropertiesCurrentTransactionCount

	ffjtDynamicGlobalPropertiesNextMaintenanceTime

	ffjtDynamicGlobalPropertiesLastBudgetTime

	ffjtDynamicGlobalPropertiesWitnessBudget

	ffjtDynamicGlobalPropertiesAccountsRegisteredThisInterval

	ffjtDynamicGlobalPropertiesRecentlyMissedCount

	ffjtDynamicGlobalPropertiesCurrentAslot

	ffjtDynamicGlobalPropertiesRecentSlotsFilled

	ffjtDynamicGlobalPropertiesDynamicFlags

	ffjtDynamicGlobalPropertiesLastIrreversibleBlockNum
)

var ffjKeyDynamicGlobalPropertiesID = []byte("id")

var ffjKeyDynamicGlobalPropertiesHeadBlockNumber = []byte("head_block_number")

var ffjKeyDynamicGlobalPropertiesHeadBlockID = []byte("head_block_id")

var ffjKeyDynamicGlobalPropertiesTime = []byte("time")

var ffjKeyDynamicGlobalPropertiesCurrentWitness = []byte("current_witness")

var ffjKeyDynamicGlobalPropertiesCurrentTransactionCount = []byte("current_transaction_count")

var ffjKeyDynamicGlobalPropertiesNextMaintenanceTime = []byte("next_maintenance_time")

var ffjKeyDynamicGlobalPropertiesLastBudgetTime = []byte("last_budget_time")

var ffjKeyDynamicGlobalPropertiesWitnessBudget = []byte("witness_budget")

var ffjKeyDynamicGlobalPropertiesAccountsRegisteredThisInterval = []byte("accounts_registered_this_interval")

var ffjKeyDynamicGlobalPropertiesRecentlyMissedCount = []byte("recently_missed_count")

var ffjKeyDynamicGlobalPropertiesCurrentAslot = []byte("current_aslot")

var ffjKeyDynamicGlobalPropertiesRecentSlotsFilled = []byte("recent_slots_filled")

var ffjKeyDynamicGlobalPropertiesDynamicFlags = []byte("dynamic_flags")

var ffjKeyDynamicGlobalPropertiesLastIrreversibleBlockNum = []byte("last_irreversible_block_num")

// UnmarshalJSON umarshall json - template of ffjson
func (j *DynamicGlobalProperties) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *DynamicGlobalProperties) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtDynamicGlobalPropertiesbase
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
				currentKey = ffjtDynamicGlobalPropertiesnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesAccountsRegisteredThisInterval, kn) {
						currentKey = ffjtDynamicGlobalPropertiesAccountsRegisteredThisInterval
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'c':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesCurrentWitness, kn) {
						currentKey = ffjtDynamicGlobalPropertiesCurrentWitness
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyDynamicGlobalPropertiesCurrentTransactionCount, kn) {
						currentKey = ffjtDynamicGlobalPropertiesCurrentTransactionCount
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyDynamicGlobalPropertiesCurrentAslot, kn) {
						currentKey = ffjtDynamicGlobalPropertiesCurrentAslot
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'd':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesDynamicFlags, kn) {
						currentKey = ffjtDynamicGlobalPropertiesDynamicFlags
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'h':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesHeadBlockNumber, kn) {
						currentKey = ffjtDynamicGlobalPropertiesHeadBlockNumber
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyDynamicGlobalPropertiesHeadBlockID, kn) {
						currentKey = ffjtDynamicGlobalPropertiesHeadBlockID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesID, kn) {
						currentKey = ffjtDynamicGlobalPropertiesID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'l':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesLastBudgetTime, kn) {
						currentKey = ffjtDynamicGlobalPropertiesLastBudgetTime
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyDynamicGlobalPropertiesLastIrreversibleBlockNum, kn) {
						currentKey = ffjtDynamicGlobalPropertiesLastIrreversibleBlockNum
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'n':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesNextMaintenanceTime, kn) {
						currentKey = ffjtDynamicGlobalPropertiesNextMaintenanceTime
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesRecentlyMissedCount, kn) {
						currentKey = ffjtDynamicGlobalPropertiesRecentlyMissedCount
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyDynamicGlobalPropertiesRecentSlotsFilled, kn) {
						currentKey = ffjtDynamicGlobalPropertiesRecentSlotsFilled
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesTime, kn) {
						currentKey = ffjtDynamicGlobalPropertiesTime
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffjKeyDynamicGlobalPropertiesWitnessBudget, kn) {
						currentKey = ffjtDynamicGlobalPropertiesWitnessBudget
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyDynamicGlobalPropertiesLastIrreversibleBlockNum, kn) {
					currentKey = ffjtDynamicGlobalPropertiesLastIrreversibleBlockNum
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyDynamicGlobalPropertiesDynamicFlags, kn) {
					currentKey = ffjtDynamicGlobalPropertiesDynamicFlags
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyDynamicGlobalPropertiesRecentSlotsFilled, kn) {
					currentKey = ffjtDynamicGlobalPropertiesRecentSlotsFilled
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyDynamicGlobalPropertiesCurrentAslot, kn) {
					currentKey = ffjtDynamicGlobalPropertiesCurrentAslot
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyDynamicGlobalPropertiesRecentlyMissedCount, kn) {
					currentKey = ffjtDynamicGlobalPropertiesRecentlyMissedCount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyDynamicGlobalPropertiesAccountsRegisteredThisInterval, kn) {
					currentKey = ffjtDynamicGlobalPropertiesAccountsRegisteredThisInterval
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyDynamicGlobalPropertiesWitnessBudget, kn) {
					currentKey = ffjtDynamicGlobalPropertiesWitnessBudget
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyDynamicGlobalPropertiesLastBudgetTime, kn) {
					currentKey = ffjtDynamicGlobalPropertiesLastBudgetTime
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyDynamicGlobalPropertiesNextMaintenanceTime, kn) {
					currentKey = ffjtDynamicGlobalPropertiesNextMaintenanceTime
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyDynamicGlobalPropertiesCurrentTransactionCount, kn) {
					currentKey = ffjtDynamicGlobalPropertiesCurrentTransactionCount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyDynamicGlobalPropertiesCurrentWitness, kn) {
					currentKey = ffjtDynamicGlobalPropertiesCurrentWitness
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyDynamicGlobalPropertiesTime, kn) {
					currentKey = ffjtDynamicGlobalPropertiesTime
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyDynamicGlobalPropertiesHeadBlockID, kn) {
					currentKey = ffjtDynamicGlobalPropertiesHeadBlockID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyDynamicGlobalPropertiesHeadBlockNumber, kn) {
					currentKey = ffjtDynamicGlobalPropertiesHeadBlockNumber
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyDynamicGlobalPropertiesID, kn) {
					currentKey = ffjtDynamicGlobalPropertiesID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtDynamicGlobalPropertiesnosuchkey
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

				case ffjtDynamicGlobalPropertiesID:
					goto handle_ID

				case ffjtDynamicGlobalPropertiesHeadBlockNumber:
					goto handle_HeadBlockNumber

				case ffjtDynamicGlobalPropertiesHeadBlockID:
					goto handle_HeadBlockID

				case ffjtDynamicGlobalPropertiesTime:
					goto handle_Time

				case ffjtDynamicGlobalPropertiesCurrentWitness:
					goto handle_CurrentWitness

				case ffjtDynamicGlobalPropertiesCurrentTransactionCount:
					goto handle_CurrentTransactionCount

				case ffjtDynamicGlobalPropertiesNextMaintenanceTime:
					goto handle_NextMaintenanceTime

				case ffjtDynamicGlobalPropertiesLastBudgetTime:
					goto handle_LastBudgetTime

				case ffjtDynamicGlobalPropertiesWitnessBudget:
					goto handle_WitnessBudget

				case ffjtDynamicGlobalPropertiesAccountsRegisteredThisInterval:
					goto handle_AccountsRegisteredThisInterval

				case ffjtDynamicGlobalPropertiesRecentlyMissedCount:
					goto handle_RecentlyMissedCount

				case ffjtDynamicGlobalPropertiesCurrentAslot:
					goto handle_CurrentAslot

				case ffjtDynamicGlobalPropertiesRecentSlotsFilled:
					goto handle_RecentSlotsFilled

				case ffjtDynamicGlobalPropertiesDynamicFlags:
					goto handle_DynamicFlags

				case ffjtDynamicGlobalPropertiesLastIrreversibleBlockNum:
					goto handle_LastIrreversibleBlockNum

				case ffjtDynamicGlobalPropertiesnosuchkey:
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

handle_ID:

	/* handler: j.ID type=types.DynamicGlobalPropertyID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.ID.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_HeadBlockNumber:

	/* handler: j.HeadBlockNumber type=types.UInt32 kind=uint32 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.HeadBlockNumber.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_HeadBlockID:

	/* handler: j.HeadBlockID type=types.String kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.HeadBlockID.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Time:

	/* handler: j.Time type=types.Time kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Time.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CurrentWitness:

	/* handler: j.CurrentWitness type=types.WitnessID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.CurrentWitness.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CurrentTransactionCount:

	/* handler: j.CurrentTransactionCount type=types.UInt32 kind=uint32 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.CurrentTransactionCount.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_NextMaintenanceTime:

	/* handler: j.NextMaintenanceTime type=types.Time kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.NextMaintenanceTime.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_LastBudgetTime:

	/* handler: j.LastBudgetTime type=types.Time kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.LastBudgetTime.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_WitnessBudget:

	/* handler: j.WitnessBudget type=types.String kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.WitnessBudget.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_AccountsRegisteredThisInterval:

	/* handler: j.AccountsRegisteredThisInterval type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.AccountsRegisteredThisInterval = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_RecentlyMissedCount:

	/* handler: j.RecentlyMissedCount type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.RecentlyMissedCount = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CurrentAslot:

	/* handler: j.CurrentAslot type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.CurrentAslot = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_RecentSlotsFilled:

	/* handler: j.RecentSlotsFilled type=types.String kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.RecentSlotsFilled.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_DynamicFlags:

	/* handler: j.DynamicFlags type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.DynamicFlags = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_LastIrreversibleBlockNum:

	/* handler: j.LastIrreversibleBlockNum type=types.UInt32 kind=uint32 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.LastIrreversibleBlockNum.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
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
