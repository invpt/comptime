package comptime

import (
	"unsafe"
)

type ordinalMarker struct{}

type Ordinal0 = [0]ordinalMarker
type Ordinal1 = [1]ordinalMarker
type Ordinal2 = [2]ordinalMarker
type Ordinal3 = [3]ordinalMarker
type Ordinal4 = [4]ordinalMarker
type Ordinal5 = [5]ordinalMarker
type Ordinal6 = [6]ordinalMarker
type Ordinal7 = [7]ordinalMarker
type Ordinal8 = [8]ordinalMarker
type Ordinal9 = [9]ordinalMarker
type Ordinal10 = [10]ordinalMarker
type Ordinal11 = [11]ordinalMarker
type Ordinal12 = [12]ordinalMarker
type Ordinal13 = [13]ordinalMarker
type Ordinal14 = [14]ordinalMarker
type Ordinal15 = [15]ordinalMarker
type Ordinal16 = [16]ordinalMarker
type Ordinal17 = [17]ordinalMarker
type Ordinal18 = [18]ordinalMarker
type Ordinal19 = [19]ordinalMarker
type Ordinal20 = [20]ordinalMarker
type Ordinal21 = [21]ordinalMarker
type Ordinal22 = [22]ordinalMarker
type Ordinal23 = [23]ordinalMarker
type Ordinal24 = [24]ordinalMarker
type Ordinal25 = [25]ordinalMarker
type Ordinal26 = [26]ordinalMarker
type Ordinal27 = [27]ordinalMarker
type Ordinal28 = [28]ordinalMarker
type Ordinal29 = [29]ordinalMarker
type Ordinal30 = [30]ordinalMarker
type Ordinal31 = [31]ordinalMarker
type Ordinal32 = [32]ordinalMarker
type Ordinal33 = [33]ordinalMarker
type Ordinal34 = [34]ordinalMarker
type Ordinal35 = [35]ordinalMarker
type Ordinal36 = [36]ordinalMarker
type Ordinal37 = [37]ordinalMarker
type Ordinal38 = [38]ordinalMarker
type Ordinal39 = [39]ordinalMarker
type Ordinal40 = [40]ordinalMarker
type Ordinal41 = [41]ordinalMarker
type Ordinal42 = [42]ordinalMarker
type Ordinal43 = [43]ordinalMarker
type Ordinal44 = [44]ordinalMarker
type Ordinal45 = [45]ordinalMarker
type Ordinal46 = [46]ordinalMarker
type Ordinal47 = [47]ordinalMarker
type Ordinal48 = [48]ordinalMarker
type Ordinal49 = [49]ordinalMarker
type Ordinal50 = [50]ordinalMarker
type Ordinal51 = [51]ordinalMarker
type Ordinal52 = [52]ordinalMarker
type Ordinal53 = [53]ordinalMarker
type Ordinal54 = [54]ordinalMarker
type Ordinal55 = [55]ordinalMarker
type Ordinal56 = [56]ordinalMarker
type Ordinal57 = [57]ordinalMarker
type Ordinal58 = [58]ordinalMarker
type Ordinal59 = [59]ordinalMarker
type Ordinal60 = [60]ordinalMarker
type Ordinal61 = [61]ordinalMarker
type Ordinal62 = [62]ordinalMarker
type Ordinal63 = [63]ordinalMarker
type Ordinal64 = [64]ordinalMarker
type Ordinal65 = [65]ordinalMarker
type Ordinal66 = [66]ordinalMarker
type Ordinal67 = [67]ordinalMarker
type Ordinal68 = [68]ordinalMarker
type Ordinal69 = [69]ordinalMarker
type Ordinal70 = [70]ordinalMarker
type Ordinal71 = [71]ordinalMarker
type Ordinal72 = [72]ordinalMarker
type Ordinal73 = [73]ordinalMarker
type Ordinal74 = [74]ordinalMarker
type Ordinal75 = [75]ordinalMarker
type Ordinal76 = [76]ordinalMarker
type Ordinal77 = [77]ordinalMarker
type Ordinal78 = [78]ordinalMarker
type Ordinal79 = [79]ordinalMarker
type Ordinal80 = [80]ordinalMarker
type Ordinal81 = [81]ordinalMarker
type Ordinal82 = [82]ordinalMarker
type Ordinal83 = [83]ordinalMarker
type Ordinal84 = [84]ordinalMarker
type Ordinal85 = [85]ordinalMarker
type Ordinal86 = [86]ordinalMarker
type Ordinal87 = [87]ordinalMarker
type Ordinal88 = [88]ordinalMarker
type Ordinal89 = [89]ordinalMarker
type Ordinal90 = [90]ordinalMarker
type Ordinal91 = [91]ordinalMarker
type Ordinal92 = [92]ordinalMarker
type Ordinal93 = [93]ordinalMarker
type Ordinal94 = [94]ordinalMarker
type Ordinal95 = [95]ordinalMarker
type Ordinal96 = [96]ordinalMarker
type Ordinal97 = [97]ordinalMarker
type Ordinal98 = [98]ordinalMarker
type Ordinal99 = [99]ordinalMarker

// An ordinal value known at compile-time. Since Go only allows up to 100 types in a union, only the
// ordinal values from 0-99 are available. The term "ordinal" is used instead of "integer" to make
// it clear that this is *not* a generalized compile-time integer.
type Ordinal interface {
	~Ordinal0 | ~Ordinal1 | ~Ordinal2 | ~Ordinal3 | ~Ordinal4 | ~Ordinal5 | ~Ordinal6 | ~Ordinal7 | ~Ordinal8 | ~Ordinal9 |
		~Ordinal10 | ~Ordinal11 | ~Ordinal12 | ~Ordinal13 | ~Ordinal14 | ~Ordinal15 | ~Ordinal16 | ~Ordinal17 | ~Ordinal18 | ~Ordinal19 |
		~Ordinal20 | ~Ordinal21 | ~Ordinal22 | ~Ordinal23 | ~Ordinal24 | ~Ordinal25 | ~Ordinal26 | ~Ordinal27 | ~Ordinal28 | ~Ordinal29 |
		~Ordinal30 | ~Ordinal31 | ~Ordinal32 | ~Ordinal33 | ~Ordinal34 | ~Ordinal35 | ~Ordinal36 | ~Ordinal37 | ~Ordinal38 | ~Ordinal39 |
		~Ordinal40 | ~Ordinal41 | ~Ordinal42 | ~Ordinal43 | ~Ordinal44 | ~Ordinal45 | ~Ordinal46 | ~Ordinal47 | ~Ordinal48 | ~Ordinal49 |
		~Ordinal50 | ~Ordinal51 | ~Ordinal52 | ~Ordinal53 | ~Ordinal54 | ~Ordinal55 | ~Ordinal56 | ~Ordinal57 | ~Ordinal58 | ~Ordinal59 |
		~Ordinal60 | ~Ordinal61 | ~Ordinal62 | ~Ordinal63 | ~Ordinal64 | ~Ordinal65 | ~Ordinal66 | ~Ordinal67 | ~Ordinal68 | ~Ordinal69 |
		~Ordinal70 | ~Ordinal71 | ~Ordinal72 | ~Ordinal73 | ~Ordinal74 | ~Ordinal75 | ~Ordinal76 | ~Ordinal77 | ~Ordinal78 | ~Ordinal79 |
		~Ordinal80 | ~Ordinal81 | ~Ordinal82 | ~Ordinal83 | ~Ordinal84 | ~Ordinal85 | ~Ordinal86 | ~Ordinal87 | ~Ordinal88 | ~Ordinal89 |
		~Ordinal90 | ~Ordinal91 | ~Ordinal92 | ~Ordinal93 | ~Ordinal94 | ~Ordinal95 | ~Ordinal96 | ~Ordinal97 | ~Ordinal98 | ~Ordinal99
}

// Gets the value of the compile-time ordinal constant `I` as a runtime uint8.
func GetOrdinal[I Ordinal]() uint8 {
	var i I
	return uint8(len(i))
}

type boolMarker struct{}

type False = [0]boolMarker
type True = [1]boolMarker

// A boolean value known at compile-time.
type Bool interface{ ~False | ~True }

// Gets the value of the compile-time boolean constant `B` as a runtime bool.
func GetBool[B Bool]() bool {
	var b B
	var t True
	return len(b) == len(t)
}

type unionItemMarker struct{}

type UnionItem1 = [0]unionItemMarker
type UnionItem2 = [1]unionItemMarker
type UnionItem interface{ ~UnionItem1 | ~UnionItem2 }

// A compile-time union type. It's like using `|` in a type constraint, except you can do compile-time
// specialization!
//
// The `unsafe` package is used for implementation, but the user-visible API is fully safe regardless
// of how it is used.
type Union[Tag UnionItem, T1, T2 any] struct {
	ptr unsafe.Pointer
}

// Creates an instance of [Union] with the first element.
func Union1[T1, T2 any](t1 *T1) Union[UnionItem1, T1, T2] {
	return Union[UnionItem1, T1, T2]{unsafe.Pointer(t1)}
}

// Creates an instance of [Union] with the second element.
func Union2[T1, T2 any](t2 *T2) Union[UnionItem2, T1, T2] {
	return Union[UnionItem2, T1, T2]{unsafe.Pointer(t2)}
}

// Gets a pointer to the value stored in the union.
//
// At most one of the two return values will be non-nil. It is possible that both are nil, which
// happens in the case that [Union1] or [Union2] was called with a nil pointer.
func (u Union[Tag, T1, T2]) Get() (*T1, *T2) {
	if u.Is1() {
		return (*T1)(u.ptr), nil
	} else {
		return nil, (*T2)(u.ptr)
	}
}

// Gets the T1 element of the union. Returns nil if this union actually T2 instead of T1.
//
// Also returns nil when the union was constructed by passing nil to [Union1].
func (u Union[Tag, T1, T2]) Get1() *T1 {
	if u.Is1() {
		return (*T1)(u.ptr)
	} else {
		return nil
	}
}

// Gets the T2 element of the union. Returns nil if this union actually T1 instead of T2.
//
// Also returns nil when the union was constructed by passing nil to [Union2].
func (u Union[Tag, T1, T2]) Get2() *T2 {
	if u.Is2() {
		return (*T2)(u.ptr)
	} else {
		return nil
	}
}

// Checks if this union contains T1.
func (u Union[Tag, T1, T2]) Is1() bool {
	var tag Tag
	return len(tag) == 0
}

// Checks if this union contains T2.
func (u Union[Tag, T1, T2]) Is2() bool {
	var tag Tag
	return len(tag) == 1
}
