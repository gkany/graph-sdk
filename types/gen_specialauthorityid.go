// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package types

import (
	"fmt"

	"github.com/gkany/cocos-go/logging"
	"github.com/gkany/cocos-go/util"
	"github.com/juju/errors"
)

type SpecialAuthorityID struct {
	ObjectID
}

func (p SpecialAuthorityID) Marshal(enc *util.TypeEncoder) error {
	n, err := enc.EncodeUVarintByByte(uint64(p.Instance()))
	if err != nil {
		return errors.Annotate(err, "encode instance")
	}

	for i := 0; i < 8-n; i++ {
		if err := enc.EncodeUVarint(uint64(0)); err != nil {
			return errors.Annotate(err, "encode zero")
		}
	}

	return nil
}

func (p *SpecialAuthorityID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}

	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeSpecialAuthority) << 48) | instance)
	return nil
}

type SpecialAuthorityIDs []SpecialAuthorityID

func (p SpecialAuthorityIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode SpecialAuthorityID")
		}
	}

	return nil
}

func SpecialAuthorityIDFromObject(ob GrapheneObject) SpecialAuthorityID {
	id, ok := ob.(*SpecialAuthorityID)
	if ok {
		return *id
	}

	p := SpecialAuthorityID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeSpecialAuthority {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeSpecialAuthority'", p.ID()))
	}

	return p
}

//NewSpecialAuthorityID creates an new SpecialAuthorityID object
func NewSpecialAuthorityID(id string) GrapheneObject {
	gid := new(SpecialAuthorityID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf(
			"SpecialAuthorityID parser error %v",
			errors.Annotate(err, "Parse"),
		)
		return nil
	}

	if gid.ObjectType() != ObjectTypeSpecialAuthority {
		logging.Errorf(
			"SpecialAuthorityID parser error %s",
			fmt.Sprintf("%q has no ObjectType 'ObjectTypeSpecialAuthority'", id),
		)
		return nil
	}

	return gid
}
