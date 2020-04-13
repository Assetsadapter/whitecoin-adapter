// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: balanceclaimoperation.go

package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *BalanceClaimOperation) MarshalJSON() ([]byte, error) {
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
func (j *BalanceClaimOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "balance_to_claim":`)

	{

		obj, err = j.BalanceToClaim.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"balance_owner_key":`)

	{

		obj, err = j.BalanceOwnerKey.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"deposit_to_account":`)

	{

		obj, err = j.DepositToAccount.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	/* Struct fall back. type=types.AssetAmount kind=struct */
	buf.WriteString(`,"total_claimed":`)
	err = buf.Encode(&j.TotalClaimed)
	if err != nil {
		return err
	}
	buf.WriteByte(',')
	if j.Fee != nil {
		if true {
			/* Struct fall back. type=types.AssetAmount kind=struct */
			buf.WriteString(`"fee":`)
			err = buf.Encode(j.Fee)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtBalanceClaimOperationbase = iota
	ffjtBalanceClaimOperationnosuchkey

	ffjtBalanceClaimOperationBalanceToClaim

	ffjtBalanceClaimOperationBalanceOwnerKey

	ffjtBalanceClaimOperationDepositToAccount

	ffjtBalanceClaimOperationTotalClaimed

	ffjtBalanceClaimOperationFee
)

var ffjKeyBalanceClaimOperationBalanceToClaim = []byte("balance_to_claim")

var ffjKeyBalanceClaimOperationBalanceOwnerKey = []byte("balance_owner_key")

var ffjKeyBalanceClaimOperationDepositToAccount = []byte("deposit_to_account")

var ffjKeyBalanceClaimOperationTotalClaimed = []byte("total_claimed")

var ffjKeyBalanceClaimOperationFee = []byte("fee")

// UnmarshalJSON umarshall json - template of ffjson
func (j *BalanceClaimOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *BalanceClaimOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtBalanceClaimOperationbase
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
				currentKey = ffjtBalanceClaimOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'b':

					if bytes.Equal(ffjKeyBalanceClaimOperationBalanceToClaim, kn) {
						currentKey = ffjtBalanceClaimOperationBalanceToClaim
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyBalanceClaimOperationBalanceOwnerKey, kn) {
						currentKey = ffjtBalanceClaimOperationBalanceOwnerKey
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'd':

					if bytes.Equal(ffjKeyBalanceClaimOperationDepositToAccount, kn) {
						currentKey = ffjtBalanceClaimOperationDepositToAccount
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyBalanceClaimOperationFee, kn) {
						currentKey = ffjtBalanceClaimOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffjKeyBalanceClaimOperationTotalClaimed, kn) {
						currentKey = ffjtBalanceClaimOperationTotalClaimed
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyBalanceClaimOperationFee, kn) {
					currentKey = ffjtBalanceClaimOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyBalanceClaimOperationTotalClaimed, kn) {
					currentKey = ffjtBalanceClaimOperationTotalClaimed
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBalanceClaimOperationDepositToAccount, kn) {
					currentKey = ffjtBalanceClaimOperationDepositToAccount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBalanceClaimOperationBalanceOwnerKey, kn) {
					currentKey = ffjtBalanceClaimOperationBalanceOwnerKey
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyBalanceClaimOperationBalanceToClaim, kn) {
					currentKey = ffjtBalanceClaimOperationBalanceToClaim
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtBalanceClaimOperationnosuchkey
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

				case ffjtBalanceClaimOperationBalanceToClaim:
					goto handle_BalanceToClaim

				case ffjtBalanceClaimOperationBalanceOwnerKey:
					goto handle_BalanceOwnerKey

				case ffjtBalanceClaimOperationDepositToAccount:
					goto handle_DepositToAccount

				case ffjtBalanceClaimOperationTotalClaimed:
					goto handle_TotalClaimed

				case ffjtBalanceClaimOperationFee:
					goto handle_Fee

				case ffjtBalanceClaimOperationnosuchkey:
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

handle_BalanceToClaim:

	/* handler: j.BalanceToClaim type=types.BalanceID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.BalanceToClaim.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BalanceOwnerKey:

	/* handler: j.BalanceOwnerKey type=types.PublicKey kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.BalanceOwnerKey.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_DepositToAccount:

	/* handler: j.DepositToAccount type=types.AccountID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.DepositToAccount.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_TotalClaimed:

	/* handler: j.TotalClaimed type=types.AssetAmount kind=struct quoted=false*/

	{
		/* Falling back. type=types.AssetAmount kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.TotalClaimed)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Fee:

	/* handler: j.Fee type=types.AssetAmount kind=struct quoted=false*/

	{
		/* Falling back. type=types.AssetAmount kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Fee)
		if err != nil {
			return fs.WrapErr(err)
		}
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