package fingerprints

type SEQ struct {
	SP  int64
	GCD int64
	ISR int64
	TI  string
	CI  string
	II  string
	TS  int64
}

type OPS struct {
	O1 string
	O2 string
	O3 string
	O4 string
	O5 string
	O6 string
}

type WIN struct {
	W1 int64
	W2 int64
	W3 int64
	W4 int64
	W5 int64
	W6 int64
}

type ECN struct {
	R  string
	DF string
	T  int64
	TG int64
	W  int64
	O  string
	CC string
	Q  string
}

type T1 struct {
	R  string
	DF string
	T  int64
	TG int64
	S  string
	A  string
	F  string
	RD int64
	Q  string
}

type T2 struct {
	R string
}

type T3 struct {
	R string
}

type T4 struct {
	R  string
	DF string
	T  int64
	TG int64
	W  int64
	S  string
	A  string
	F  string
	O  string
	RD int64
	Q  string
}

type T5 struct {
	R  string
	DF string
	T  int64
	TG int64
	W  int64
	S  string
	A  string
	F  string
	O  string
	RD int64
	Q  string
}

type T6 struct {
	R  string
	DF string
	T  int64
	TG int64
	W  int64
	S  string
	A  string
	F  string
	O  string
	RD int64
	Q  string
}

type T7 struct {
	R string
}

type U1 struct {
	R     string
	DF    string
	T     int64
	IPL   int64
	UN    int64
	RIPL  string
	RID   string
	RIPCK string
	RUCK  string
	RUD   string
}

type IE struct {
	R   string
	DFI string
	T   int64
	CD  string
}

type Fingerprint struct {
	SEQs []SEQ
	OPS  OPS
	WIN  WIN
	ECN  ECN
	T1   T1
	T2   T2
	T3   T3
	T4   T4
	T5   T5
	T6   T6
	T7   T7
	U1   U1
	IE   IE
}
