// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package types

import (
	"fmt"

	"github.com/blocktree/whitecoin-adapter/libs/util"
	"github.com/denkhaus/logging"
	"github.com/juju/errors"
)

type OperationHistoryID struct {
	ObjectID
}

func (p OperationHistoryID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.Instance())); err != nil {
		return errors.Annotate(err, "encode instance")
	}

	return nil
}

func (p *OperationHistoryID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}

	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeOperationHistory) << 48) | instance)
	return nil
}

type OperationHistoryIDs []OperationHistoryID

func (p OperationHistoryIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode OperationHistoryID")
		}
	}

	return nil
}

func OperationHistoryIDFromObject(ob GrapheneObject) OperationHistoryID {
	id, ok := ob.(*OperationHistoryID)
	if ok {
		return *id
	}

	p := OperationHistoryID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeOperationHistory {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeOperationHistory'", p.ID()))
	}

	return p
}

//NewOperationHistoryID creates an new OperationHistoryID object
func NewOperationHistoryID(id string) GrapheneObject {
	gid := new(OperationHistoryID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf(
			"OperationHistoryID parser error %v",
			errors.Annotate(err, "Parse"),
		)
		return nil
	}

	if gid.ObjectType() != ObjectTypeOperationHistory {
		logging.Errorf(
			"OperationHistoryID parser error %s",
			fmt.Sprintf("%q has no ObjectType 'ObjectTypeOperationHistory'", id),
		)
		return nil
	}

	return gid
}
