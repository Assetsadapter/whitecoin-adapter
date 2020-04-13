// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: proposalcreateoperation.go

package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/blocktree/whitecoin-adapter/libs/types"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *ProposalCreateOperation) MarshalJSON() ([]byte, error) {
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
func (j *ProposalCreateOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "expiration_time":`)

	{

		obj, err = j.ExpirationTime.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"extensions":`)

	{

		obj, err = j.Extensions.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"fee_paying_account":`)

	{

		obj, err = j.FeePayingAccount.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteByte(',')
	if j.ReviewPeriodSeconds != nil {
		if true {
			buf.WriteString(`"review_period_seconds":`)
			fflib.FormatBits2(buf, uint64(*j.ReviewPeriodSeconds), 10, false)
			buf.WriteByte(',')
		}
	}
	buf.WriteString(`"proposed_ops":`)
	if j.ProposedOps != nil {
		buf.WriteString(`[`)
		for i, v := range j.ProposedOps {
			if i != 0 {
				buf.WriteString(`,`)
			}
			/* Struct fall back. type=types.OperationEnvelopeHolder kind=struct */
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
	ffjtProposalCreateOperationbase = iota
	ffjtProposalCreateOperationnosuchkey

	ffjtProposalCreateOperationExpirationTime

	ffjtProposalCreateOperationExtensions

	ffjtProposalCreateOperationFeePayingAccount

	ffjtProposalCreateOperationReviewPeriodSeconds

	ffjtProposalCreateOperationProposedOps

	ffjtProposalCreateOperationFee
)

var ffjKeyProposalCreateOperationExpirationTime = []byte("expiration_time")

var ffjKeyProposalCreateOperationExtensions = []byte("extensions")

var ffjKeyProposalCreateOperationFeePayingAccount = []byte("fee_paying_account")

var ffjKeyProposalCreateOperationReviewPeriodSeconds = []byte("review_period_seconds")

var ffjKeyProposalCreateOperationProposedOps = []byte("proposed_ops")

var ffjKeyProposalCreateOperationFee = []byte("fee")

// UnmarshalJSON umarshall json - template of ffjson
func (j *ProposalCreateOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *ProposalCreateOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtProposalCreateOperationbase
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
				currentKey = ffjtProposalCreateOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'e':

					if bytes.Equal(ffjKeyProposalCreateOperationExpirationTime, kn) {
						currentKey = ffjtProposalCreateOperationExpirationTime
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyProposalCreateOperationExtensions, kn) {
						currentKey = ffjtProposalCreateOperationExtensions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyProposalCreateOperationFeePayingAccount, kn) {
						currentKey = ffjtProposalCreateOperationFeePayingAccount
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyProposalCreateOperationFee, kn) {
						currentKey = ffjtProposalCreateOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffjKeyProposalCreateOperationProposedOps, kn) {
						currentKey = ffjtProposalCreateOperationProposedOps
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffjKeyProposalCreateOperationReviewPeriodSeconds, kn) {
						currentKey = ffjtProposalCreateOperationReviewPeriodSeconds
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyProposalCreateOperationFee, kn) {
					currentKey = ffjtProposalCreateOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyProposalCreateOperationProposedOps, kn) {
					currentKey = ffjtProposalCreateOperationProposedOps
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyProposalCreateOperationReviewPeriodSeconds, kn) {
					currentKey = ffjtProposalCreateOperationReviewPeriodSeconds
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyProposalCreateOperationFeePayingAccount, kn) {
					currentKey = ffjtProposalCreateOperationFeePayingAccount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyProposalCreateOperationExtensions, kn) {
					currentKey = ffjtProposalCreateOperationExtensions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyProposalCreateOperationExpirationTime, kn) {
					currentKey = ffjtProposalCreateOperationExpirationTime
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtProposalCreateOperationnosuchkey
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

				case ffjtProposalCreateOperationExpirationTime:
					goto handle_ExpirationTime

				case ffjtProposalCreateOperationExtensions:
					goto handle_Extensions

				case ffjtProposalCreateOperationFeePayingAccount:
					goto handle_FeePayingAccount

				case ffjtProposalCreateOperationReviewPeriodSeconds:
					goto handle_ReviewPeriodSeconds

				case ffjtProposalCreateOperationProposedOps:
					goto handle_ProposedOps

				case ffjtProposalCreateOperationFee:
					goto handle_Fee

				case ffjtProposalCreateOperationnosuchkey:
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

handle_ExpirationTime:

	/* handler: j.ExpirationTime type=types.Time kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.ExpirationTime.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Extensions:

	/* handler: j.Extensions type=types.Extensions kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Extensions.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_FeePayingAccount:

	/* handler: j.FeePayingAccount type=types.AccountID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.FeePayingAccount.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ReviewPeriodSeconds:

	/* handler: j.ReviewPeriodSeconds type=types.UInt32 kind=uint32 quoted=false*/

	{
		if tok == fflib.FFTok_null {

			j.ReviewPeriodSeconds = nil

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			if j.ReviewPeriodSeconds == nil {
				j.ReviewPeriodSeconds = new(types.UInt32)
			}

			err = j.ReviewPeriodSeconds.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ProposedOps:

	/* handler: j.ProposedOps type=types.OperationEnvelopeHolders kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for OperationEnvelopeHolders", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.ProposedOps = nil
		} else {

			j.ProposedOps = []types.OperationEnvelopeHolder{}

			wantVal := true

			for {

				var tmpJProposedOps types.OperationEnvelopeHolder

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

				/* handler: tmpJProposedOps type=types.OperationEnvelopeHolder kind=struct quoted=false*/

				{
					/* Falling back. type=types.OperationEnvelopeHolder kind=struct */
					tbuf, err := fs.CaptureField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}

					err = json.Unmarshal(tbuf, &tmpJProposedOps)
					if err != nil {
						return fs.WrapErr(err)
					}
				}

				j.ProposedOps = append(j.ProposedOps, tmpJProposedOps)

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