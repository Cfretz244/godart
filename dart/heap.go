package dart

import "strings"

type Heap struct {
  impl wrapper
}

type HeapIterator struct {
  val *Heap
  impl func () *Heap
}

var nullHeap *Heap

func init() {
  nullHeap = &Heap{&NullHeap{}}
}

func (pkt *Heap) IsObject() bool {
  return pkt.impl.IsObject()
}

func (pkt *Heap) IsArray() bool {
  return pkt.impl.IsArray()
}

func (pkt *Heap) IsString() bool {
  return pkt.impl.IsString()
}

func (pkt *Heap) IsInteger() bool {
  return pkt.impl.IsInteger()
}

func (pkt *Heap) IsDecimal() bool {
  return pkt.impl.IsDecimal()
}

func (pkt *Heap) IsBoolean() bool {
  return pkt.impl.IsBoolean()
}

func (pkt *Heap) IsNull() bool {
  return pkt.impl.IsNull()
}

func (pkt *Heap) IsFinalized() bool {
  return false
}

func (pkt *Heap) GetType() int {
  return pkt.impl.GetType()
}

func (pkt *Heap) Refcount() uint64 {
  return pkt.impl.Refcount()
}

func (pkt *Heap) toJSON(out *strings.Builder) {
  pkt.impl.toJSON(out)
}

func (pkt *Heap) ToJSON() string {
  return pkt.impl.ToJSON()
}

func (it *HeapIterator) Next() bool {
  if it.impl != nil {
    it.val = it.impl()
    return it.val != nil
  } else {
    return false
  }
}

func (it *HeapIterator) Value() *Heap {
  if it.val != nil {
    return it.val
  } else {
    return nullHeap
  }
}