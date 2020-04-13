// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: transferfromblindoperation.go

package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Assetsadapter/whitecoin-adapter/libs/types"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *TransferFromBlindOperation) MarshalJSON() ([]byte, error) {
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
func (j *TransferFromBlindOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	/* Struct fall back. type=types.AssetAmount kind=struct */
	buf.WriteString(`{ "amount":`)
	err = buf.Encode(&j.Amount)
	if err != nil {
		return err
	}
	buf.WriteString(`,"to":`)

	{

		obj, err = j.To.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"blinding_factor":`)

	{

		obj, err = j.BlindFactor.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"inputs":`)
	if j.BlindInputs != nil {
		buf.WriteString(`[`)
		for i, v := range j.BlindInputs {
			if i != 0 {
				buf.WriteString(`,`)
			}
			/* Struct fall back. type=types.BlindInput kind=struct */
			err = buf.Encode(&v)
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
	ffjtTransferFromBlindOperationbase = iota
	ffjtTransferFromBlindOperationnosuchkey

	ffjtTransferFromBlindOperationAmount

	ffjtTransferFromBlindOperationTo

	ffjtTransferFromBlindOperationBlindFactor

	ffjtTransferFromBlindOperationBlindInputs

	ffjtTransferFromBlindOperationFee
)

var ffjKeyTransferFromBlindOperationAmount = []byte("amount")

var ffjKeyTransferFromBlindOperationTo = []byte("to")

var ffjKeyTransferFromBlindOperationBlindFactor = []byte("blinding_factor")

var ffjKeyTransferFromBlindOperationBlindInputs = []byte("inputs")

var ffjKeyTransferFromBlindOperationFee = []byte("fee")

// UnmarshalJSON umarshall json - template of ffjson
func (j *TransferFromBlindOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *TransferFromBlindOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtTransferFromBlindOperationbase
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
				currentKey = ffjtTransferFromBlindOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyTransferFromBlindOperationAmount, kn) {
						currentKey = ffjtTransferFromBlindOperationAmount
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'b':

					if bytes.Equal(ffjKeyTransferFromBlindOperationBlindFactor, kn) {
						currentKey = ffjtTransferFromBlindOperationBlindFactor
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyTransferFromBlindOperationFee, kn) {
						currentKey = ffjtTransferFromBlindOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyTransferFromBlindOperationBlindInputs, kn) {
						currentKey = ffjtTransferFromBlindOperationBlindInputs
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffjKeyTransferFromBlindOperationTo, kn) {
						currentKey = ffjtTransferFromBlindOperationTo
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyTransferFromBlindOperationFee, kn) {
					currentKey = ffjtTransferFromBlindOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyTransferFromBlindOperationBlindInputs, kn) {
					currentKey = ffjtTransferFromBlindOperationBlindInputs
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyTransferFromBlindOperationBlindFactor, kn) {
					currentKey = ffjtTransferFromBlindOperationBlindFactor
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyTransferFromBlindOperationTo, kn) {
					currentKey = ffjtTransferFromBlindOperationTo
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyTransferFromBlindOperationAmount, kn) {
					currentKey = ffjtTransferFromBlindOperationAmount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtTransferFromBlindOperationnosuchkey
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

				case ffjtTransferFromBlindOperationAmount:
					goto handle_Amount

				case ffjtTransferFromBlindOperationTo:
					goto handle_To

				case ffjtTransferFromBlindOperationBlindFactor:
					goto handle_BlindFactor

				case ffjtTransferFromBlindOperationBlindInputs:
					goto handle_BlindInputs

				case ffjtTransferFromBlindOperationFee:
					goto handle_Fee

				case ffjtTransferFromBlindOperationnosuchkey:
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

handle_Amount:

	/* handler: j.Amount type=types.AssetAmount kind=struct quoted=false*/

	{
		/* Falling back. type=types.AssetAmount kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Amount)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_To:

	/* handler: j.To type=types.AccountID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.To.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BlindFactor:

	/* handler: j.BlindFactor type=types.FixedBuffer kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.BlindFactor.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BlindInputs:

	/* handler: j.BlindInputs type=types.BlindInputs kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for BlindInputs", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.BlindInputs = nil
		} else {

			j.BlindInputs = []types.BlindInput{}

			wantVal := true

			for {

				var tmpJBlindInputs types.BlindInput

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

				/* handler: tmpJBlindInputs type=types.BlindInput kind=struct quoted=false*/

				{
					/* Falling back. type=types.BlindInput kind=struct */
					tbuf, err := fs.CaptureField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}

					err = json.Unmarshal(tbuf, &tmpJBlindInputs)
					if err != nil {
						return fs.WrapErr(err)
					}
				}

				j.BlindInputs = append(j.BlindInputs, tmpJBlindInputs)

				wantVal = false
			}
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
