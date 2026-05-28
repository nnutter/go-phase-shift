package a

func UnannotatedMutation(p *int) {
	*p = 1
}

// Other annotation
func OtherAnnotationMutation(p *int) {
	*p = 1
}

//phase:nonmutating
func MutatesPointer(p *int) {
	*p = 1 // want "//phase:nonmutating function mutates pointer parameter p"
}

//phase:nonmutating
func ReadsPointer(p *int) int {
	return *p
}

//phase:nonmutating
func ReassignsParameter(p *int) {
	p = nil
}

//phase:nonmutating
func MutatesPointerAlias(p *int) {
	q := p
	*q = 1 // want "//phase:nonmutating function mutates pointer parameter p"
}

//phase:nonmutating
func MutatesSlice(s []int) {
	s[0] = 1 // want "//phase:nonmutating function mutates parameter s"
}

//phase:nonmutating
func MutatesSliceAlias(s []int) {
	t := s
	t[0] = 1 // want "//phase:nonmutating function mutates parameter s"
}

//phase:nonmutating
func MutatesMap(m map[string]int) {
	m["x"] = 1 // want "//phase:nonmutating function mutates parameter m"
}

//phase:nonmutating
func DeletesMapEntry(m map[string]int) {
	delete(m, "x") // want "//phase:nonmutating function deletes from map parameter m"
}

//phase:nonmutating
func MutatesLocal() int {
	x := 0
	x++
	return x
}

//phase:nonmutating
func MutatesLocalPointer() int {
	x := 0
	p := &x
	*p = 1
	return x
}

type Counter struct{ n int }

//phase:nonmutating
func (c *Counter) Inc() {
	c.n++ // want "//phase:nonmutating method mutates receiver c"
}

type Count struct{ n int }

//phase:nonmutating
func (c Count) Increment() Count {
	c.n++
	return c
}

type Buffer struct{ data []byte }

//phase:nonmutating
func (b Buffer) ClearFirstByte() {
	b.data[0] = 0 // want "//phase:nonmutating method mutates receiver b"
}

type Holder struct{ p *int }

//phase:nonmutating
func (h Holder) MutatesPointedValue() {
	*h.p = 1 // want "//phase:nonmutating method mutates receiver h"
}

type MHolder struct{ m map[string]int }

//phase:nonmutating
func (h MHolder) MutatesMapField() {
	h.m["x"] = 1 // want "//phase:nonmutating method mutates receiver h"
}

//phase:nonmutating
func (h Holder) ReplacesPointerInCopy() {
	h.p = nil
}

//phase:nonmutating
func MutatesNestedSlice(ss [][]int) {
	ss[0][0] = 1 // want "//phase:nonmutating function mutates parameter ss"
}

//phase:nonmutating
func DeletesLocalMap() {
	m := map[string]int{}
	delete(m, "x")
}

//phase:nonmutating
func MutatesDoublePointer(pp **int) {
	**pp = 1 // want "//phase:nonmutating function mutates pointer parameter pp"
}
