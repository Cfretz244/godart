package dart

import (
  "github.com/cfretz244/godart/cdart"
)

type BooleanBuffer struct {
  native *cdart.Packet
  json string
  val, set bool
}

func boolFromPacket(pkt *cdart.Packet) *BooleanBuffer {
  if !pkt.IsBoolean() {
    panic("Native packet of unexpected type passed to BooleanBuffer converter")
  } else if !pkt.IsFinalized() {
    panic("Non-finalized boolean passed to BooleanBuffer converter")
  }
  return &BooleanBuffer{pkt, "", false, false}
}

func (num *BooleanBuffer) ctype() *cdart.Packet {
  return num.native
}

func (num *BooleanBuffer) IsObject() bool {
  return false
}

func (num *BooleanBuffer) IsArray() bool {
  return false
}

func (num *BooleanBuffer) IsString() bool {
  return false
}

func (num *BooleanBuffer) IsInteger() bool {
  return false
}

func (num *BooleanBuffer) IsDecimal() bool {
  return false
}

func (num *BooleanBuffer) IsBoolean() bool {
  return true
}

func (num *BooleanBuffer) IsNull() bool {
  return false
}

func (num *BooleanBuffer) IsFinalized() bool {
  return true
}

func (num *BooleanBuffer) GetType() int {
  return cdart.BooleanType
}

func (num *BooleanBuffer) Refcount() uint64 {
  return num.native.Refcount()
}

func (num *BooleanBuffer) equal(other wrapper) bool {
  return false
}

func (num *BooleanBuffer) ToJSON() string {
  if len(num.json) > 0 {
    // We've already generated JSON previously
    // Buffers are immutable, so just return it.
    return num.json
  } else if num.native != nil {
    // We haven't generated our JSON before, but we
    // have a native representation, so do it.
    json, err := num.native.ToJSON()
    errCheck(err, "boolean")

    num.json = json
    return num.json
  } else {
    // We're a default initialized numuct
    // Just return a static string
    return "false"
  }
}
