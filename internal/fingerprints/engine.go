package fingerprints

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func parseInt64(s string) int64 {
	// parse int from hex string, using strconv
	n, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return 0
	}
	return n
}

func parseInt64toStr(n int64) string {
	return fmt.Sprintf("%X", n)
}

func DeserializeSeq(data string) (*SEQ, error) {
	// seq data like "SEQ(SP=105%GCD=1%ISR=10A%TI=Z%CI=I%II=I%TS=7)"
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}
	seq := &SEQ{}
	data = strings.Replace(data, "SEQ(", "", -1)
	data = strings.Replace(data, ")", "", -1)
	items := strings.Split(data, "%")
	for _, item := range items {
		kv := strings.Split(item, "=")
		if len(kv) != 2 {
			continue
		}
		switch kv[0] {
		case "SP":
			seq.SP = parseInt64(kv[1])
		case "GCD":
			seq.GCD = parseInt64(kv[1])
		case "ISR":
			seq.ISR = parseInt64(kv[1])
		case "TI":
			seq.TI = kv[1]
		case "CI":
			seq.CI = kv[1]
		case "II":
			seq.II = kv[1]
		case "TS":
			seq.TS = parseInt64(kv[1])
		}
	}
	return seq, nil
}

func DeserializeOps(data string) (*OPS, error) {
	// ops data like "OPS(O1=M5B4ST11NW5%O2=M5B4ST11NW5%O3=M5B4NNT11NW5%O4=M5B4ST11NW5%O5=M5B4ST11NW5%O6=M5B4ST11)"
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}
	ops := &OPS{}
	data = strings.Replace(data, "OPS(", "", -1)
	data = strings.Replace(data, ")", "", -1)
	items := strings.Split(data, "%")
	for _, item := range items {
		kv := strings.Split(item, "=")
		if len(kv) != 2 {
			continue
		}
		switch kv[0] {
		case "O1":
			ops.O1 = kv[1]
		case "O2":
			ops.O2 = kv[1]
		case "O3":
			ops.O3 = kv[1]
		case "O4":
			ops.O4 = kv[1]
		case "O5":
			ops.O5 = kv[1]
		case "O6":
			ops.O6 = kv[1]
		}
	}
	return ops, nil
}

func DeserializeWin(data string) (*WIN, error) {
	// win data like "WIN(W1=7120%W2=7120%W3=7120%W4=7120%W5=7120%W6=7120)"
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}
	win := &WIN{}
	data = strings.Replace(data, "WIN(", "", -1)
	data = strings.Replace(data, ")", "", -1)
	items := strings.Split(data, "%")
	for _, item := range items {
		kv := strings.Split(item, "=")
		if len(kv) != 2 {
			continue
		}
		switch kv[0] {
		case "W1":
			win.W1 = parseInt64(kv[1])
		case "W2":
			win.W2 = parseInt64(kv[1])
		case "W3":
			win.W3 = parseInt64(kv[1])
		case "W4":
			win.W4 = parseInt64(kv[1])
		case "W5":
			win.W5 = parseInt64(kv[1])
		case "W6":
			win.W6 = parseInt64(kv[1])
		}
	}
	return win, nil
}

func DeserializeEcn(data string) (*ECN, error) {
	// ecn data like "ECN(R=Y%DF=Y%T=40%W=7210%O=M5B4NNSNW5%CC=N%Q=)"
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}
	ecn := &ECN{}
	data = strings.Replace(data, "ECN(", "", -1)
	data = strings.Replace(data, ")", "", -1)
	items := strings.Split(data, "%")
	for _, item := range items {
		kv := strings.Split(item, "=")
		if len(kv) != 2 {
			continue
		}
		switch kv[0] {
		case "R":
			ecn.R = kv[1]
		case "DF":
			ecn.DF = kv[1]
		case "T":
			ecn.T = parseInt64(kv[1])
		case "W":
			ecn.W = parseInt64(kv[1])
		case "O":
			ecn.O = kv[1]
		case "CC":
			ecn.CC = kv[1]
		case "Q":
			ecn.Q = kv[1]
		}
	}
	return ecn, nil
}

func DeserializeT1(data string) (*T1, error) {
	// t1 data like "T1(R=Y%DF=Y%T=40%S=O%A=S+%F=AS%RD=0%Q=)"
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}
	t1 := &T1{}
	data = strings.Replace(data, "T1(", "", -1)
	data = strings.Replace(data, ")", "", -1)
	items := strings.Split(data, "%")
	for _, item := range items {
		kv := strings.Split(item, "=")
		if len(kv) != 2 {
			continue
		}
		switch kv[0] {
		case "R":
			t1.R = kv[1]
		case "DF":
			t1.DF = kv[1]
		case "T":
			t1.T = parseInt64(kv[1])
		case "S":
			t1.S = kv[1]
		case "A":
			t1.A = kv[1]
		case "F":
			t1.F = kv[1]
		case "RD":
			t1.RD = parseInt64(kv[1])
		case "Q":
			t1.Q = kv[1]
		}
	}
	return t1, nil
}

func DeserializeT2(data string) (*T2, error) {
	// t2 data like "T2(R=N)"
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}
	t2 := &T2{}
	data = strings.Replace(data, "T2(", "", -1)
	data = strings.Replace(data, ")", "", -1)
	items := strings.Split(data, "%")
	for _, item := range items {
		kv := strings.Split(item, "=")
		if len(kv) != 2 {
			continue
		}
		switch kv[0] {
		case "R":
			t2.R = kv[1]
		}
	}
	return t2, nil
}

func DeserializeT3(data string) (*T3, error) {
	// t3 data like "T3(R=N)"
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}
	t3 := &T3{}
	data = strings.Replace(data, "T3(", "", -1)
	data = strings.Replace(data, ")", "", -1)
	items := strings.Split(data, "%")
	for _, item := range items {
		kv := strings.Split(item, "=")
		if len(kv) != 2 {
			continue
		}
		switch kv[0] {
		case "R":
			t3.R = kv[1]
		}
	}
	return t3, nil
}

func DeserializeT4(data string) (*T4, error) {
	// t4 data like "T4(R=Y%DF=Y%T=40%W=0%S=A%A=Z%F=R%O=%RD=0%Q=)"
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}
	t4 := &T4{}
	data = strings.Replace(data, "T4(", "", -1)
	data = strings.Replace(data, ")", "", -1)
	items := strings.Split(data, "%")
	for _, item := range items {
		kv := strings.Split(item, "=")
		if len(kv) != 2 {
			continue
		}
		switch kv[0] {
		case "R":
			t4.R = kv[1]
		case "DF":
			t4.DF = kv[1]
		case "T":
			t4.T = parseInt64(kv[1])
		case "W":
			t4.W = parseInt64(kv[1])
		case "S":
			t4.S = kv[1]
		case "A":
			t4.A = kv[1]
		case "F":
			t4.F = kv[1]
		case "O":
			t4.O = kv[1]
		case "RD":
			t4.RD = parseInt64(kv[1])
		case "Q":
			t4.Q = kv[1]
		}
	}
	return t4, nil
}

func DeserializeT5(data string) (*T5, error) {
	// t5 data like "T5(R=Y%DF=Y%T=40%W=0%S=Z%A=S+%F=AR%O=%RD=0%Q=)"
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}
	t5 := &T5{}
	data = strings.Replace(data, "T5(", "", -1)
	data = strings.Replace(data, ")", "", -1)
	items := strings.Split(data, "%")
	for _, item := range items {
		kv := strings.Split(item, "=")
		if len(kv) != 2 {
			continue
		}
		switch kv[0] {
		case "R":
			t5.R = kv[1]
		case "DF":
			t5.DF = kv[1]
		case "T":
			t5.T = parseInt64(kv[1])
		case "W":
			t5.W = parseInt64(kv[1])
		case "S":
			t5.S = kv[1]
		case "A":
			t5.A = kv[1]
		case "F":
			t5.F = kv[1]
		case "O":
			t5.O = kv[1]
		case "RD":
			t5.RD = parseInt64(kv[1])
		case "Q":
			t5.Q = kv[1]
		}
	}
	return t5, nil
}

func DeserializeT6(data string) (*T6, error) {
	// t6 data like "T6(R=Y%DF=Y%T=40%W=0%S=A%A=Z%F=R%O=%RD=0%Q=)"
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}
	t6 := &T6{}
	data = strings.Replace(data, "T6(", "", -1)
	data = strings.Replace(data, ")", "", -1)
	items := strings.Split(data, "%")
	for _, item := range items {
		kv := strings.Split(item, "=")
		if len(kv) != 2 {
			continue
		}
		switch kv[0] {
		case "R":
			t6.R = kv[1]
		case "DF":
			t6.DF = kv[1]
		case "T":
			t6.T = parseInt64(kv[1])
		case "W":
			t6.W = parseInt64(kv[1])
		case "S":
			t6.S = kv[1]
		case "A":
			t6.A = kv[1]
		case "F":
			t6.F = kv[1]
		case "O":
			t6.O = kv[1]
		case "RD":
			t6.RD = parseInt64(kv[1])
		case "Q":
			t6.Q = kv[1]
		}
	}
	return t6, nil
}

func DeserializeT7(data string) (*T7, error) {
	// t7 data like "T7(R=Y%DF=Y%T=40%W=0%S=Z%A=S+%F=AR%O=%RD=0%Q=)"
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}
	t7 := &T7{}
	data = strings.Replace(data, "T7(", "", -1)
	data = strings.Replace(data, ")", "", -1)
	items := strings.Split(data, "%")
	for _, item := range items {
		kv := strings.Split(item, "=")
		if len(kv) != 2 {
			continue
		}
		switch kv[0] {
		case "R":
			t7.R = kv[1]
		}
	}
	return t7, nil
}

func DeserializeU1(data string) (*U1, error) {
	// u1 data like "U1(R=Y%DF=N%T=40%IPL=164%UN=0%RIPL=G%RID=G%RIPCK=G%RUCK=G%RUD=G)"
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}
	u1 := &U1{}
	data = strings.Replace(data, "U1(", "", -1)
	data = strings.Replace(data, ")", "", -1)
	items := strings.Split(data, "%")
	for _, item := range items {
		kv := strings.Split(item, "=")
		if len(kv) != 2 {
			continue
		}
		switch kv[0] {
		case "R":
			u1.R = kv[1]
		case "DF":
			u1.DF = kv[1]
		case "T":
			u1.T = parseInt64(kv[1])
		case "IPL":
			u1.IPL = parseInt64(kv[1])
		case "UN":
			u1.UN = parseInt64(kv[1])
		case "RIPL":
			u1.RIPL = kv[1]
		case "RID":
			u1.RID = kv[1]
		case "RIPCK":
			u1.RIPCK = kv[1]
		case "RUCK":
			u1.RUCK = kv[1]
		case "RUD":
			u1.RUD = kv[1]
		}
	}
	return u1, nil
}

func DeserializeIE(data string) (*IE, error) {
	// ie data like "IE(R=Y%DF=Y%T=40%W=7210%O=M5B4NNSNW5%CC=N%Q=)"
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}
	ie := &IE{}
	data = strings.Replace(data, "IE(", "", -1)
	data = strings.Replace(data, ")", "", -1)
	items := strings.Split(data, "%")
	for _, item := range items {
		kv := strings.Split(item, "=")
		if len(kv) != 2 {
			continue
		}
		switch kv[0] {
		case "R":
			ie.R = kv[1]
		case "DFI":
			ie.DFI = kv[1]
		case "T":
			ie.T = parseInt64(kv[1])
		case "CD":
			ie.CD = kv[1]
		}
	}
	return ie, nil
}

func DeserializeFp(data string) (*Fingerprint, error) {
	fp := &Fingerprint{}
	seqs := make([]SEQ, 0)
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "SEQ") {
			seq, err := DeserializeSeq(line)
			if err != nil {
				return nil, err
			}
			seqs = append(seqs, *seq)
			continue
		}
		if strings.HasPrefix(line, "OPS") {
			ops, err := DeserializeOps(line)
			if err != nil {
				return nil, err
			}
			fp.OPS = *ops
			continue
		}
		if strings.HasPrefix(line, "WIN") {
			win, err := DeserializeWin(line)
			if err != nil {
				return nil, err
			}
			fp.WIN = *win
			continue
		}
		if strings.HasPrefix(line, "ECN") {
			ecn, err := DeserializeEcn(line)
			if err != nil {
				return nil, err
			}
			fp.ECN = *ecn
			continue
		}
		if strings.HasPrefix(line, "T1") {
			t1, err := DeserializeT1(line)
			if err != nil {
				return nil, err
			}
			fp.T1 = *t1
			continue
		}
		if strings.HasPrefix(line, "T2") {
			t2, err := DeserializeT2(line)
			if err != nil {
				return nil, err
			}
			fp.T2 = *t2
			continue
		}
		if strings.HasPrefix(line, "T3") {
			t3, err := DeserializeT3(line)
			if err != nil {
				return nil, err
			}
			fp.T3 = *t3
			continue
		}
		if strings.HasPrefix(line, "T4") {
			t4, err := DeserializeT4(line)
			if err != nil {
				return nil, err
			}
			fp.T4 = *t4
			continue
		}
		if strings.HasPrefix(line, "T5") {
			t5, err := DeserializeT5(line)
			if err != nil {
				return nil, err
			}
			fp.T5 = *t5
			continue
		}
		if strings.HasPrefix(line, "T6") {
			t6, err := DeserializeT6(line)
			if err != nil {
				return nil, err
			}
			fp.T6 = *t6
			continue
		}
		if strings.HasPrefix(line, "T7") {
			t7, err := DeserializeT7(line)
			if err != nil {
				return nil, err
			}
			fp.T7 = *t7
			continue
		}
		if strings.HasPrefix(line, "U1") {
			u1, err := DeserializeU1(line)
			if err != nil {
				return nil, err
			}
			fp.U1 = *u1
			continue
		}
		if strings.HasPrefix(line, "IE") {
			ie, err := DeserializeIE(line)
			if err != nil {
				return nil, err
			}
			fp.IE = *ie
			continue
		}

	}
	fp.SEQs = seqs
	return fp, nil
}

func DeserializeFpsFromStr(strFps []string) ([]*Fingerprint, error) {
	fps := make([]*Fingerprint, 0)
	for _, data := range strFps {
		fp, err := DeserializeFp(data)
		if err != nil {
			return nil, err
		}
		fps = append(fps, fp)
	}
	return fps, nil
}

func SerializeFp(fp *Fingerprint) string {
	data := ""
	for _, seq := range fp.SEQs {
		data += fmt.Sprintf("SEQ(SP=%X%%GCD=%X%%ISR=%X%%TI=%s%%CI=%s%%II=%s%%TS=%X)\n", seq.SP, seq.GCD, seq.ISR, seq.TI, seq.CI, seq.II, seq.TS)
	}
	data += fmt.Sprintf("OPS(O1=%s%%O2=%s%%O3=%s%%O4=%s%%O5=%s%%O6=%s)\n", fp.OPS.O1, fp.OPS.O2, fp.OPS.O3, fp.OPS.O4, fp.OPS.O5, fp.OPS.O6)
	data += fmt.Sprintf("WIN(W1=%X%%W2=%X%%W3=%X%%W4=%X%%W5=%X%%W6=%X)\n", fp.WIN.W1, fp.WIN.W2, fp.WIN.W3, fp.WIN.W4, fp.WIN.W5, fp.WIN.W6)
	data += fmt.Sprintf("ECN(R=%s%%DF=%s%%T=%X%%W=%X%%O=%s%%CC=%s%%Q=%s)\n", fp.ECN.R, fp.ECN.DF, fp.ECN.T, fp.ECN.W, fp.ECN.O, fp.ECN.CC, fp.ECN.Q)
	data += fmt.Sprintf("T1(R=%s%%DF=%s%%T=%X%%S=%s%%A=%s%%F=%s%%RD=%X%%Q=%s)\n", fp.T1.R, fp.T1.DF, fp.T1.T, fp.T1.S, fp.T1.A, fp.T1.F, fp.T1.RD, fp.T1.Q)
	data += fmt.Sprintf("T2(R=%s)\n", fp.T2.R)
	data += fmt.Sprintf("T3(R=%s)\n", fp.T3.R)
	data += fmt.Sprintf("T4(R=%s%%DF=%s%%T=%X%%W=%X%%S=%s%%A=%s%%F=%s%%O=%s%%RD=%X%%Q=%s)\n", fp.T4.R, fp.T4.DF, fp.T4.T, fp.T4.W, fp.T4.S, fp.T4.A, fp.T4.F, fp.T4.O, fp.T4.RD, fp.T4.Q)
	data += fmt.Sprintf("T5(R=%s%%DF=%s%%T=%X%%W=%X%%S=%s%%A=%s%%F=%s%%O=%s%%RD=%X%%Q=%s)\n", fp.T5.R, fp.T5.DF, fp.T5.T, fp.T5.W, fp.T5.S, fp.T5.A, fp.T5.F, fp.T5.O, fp.T5.RD, fp.T5.Q)
	data += fmt.Sprintf("T6(R=%s%%DF=%s%%T=%X%%W=%X%%S=%s%%A=%s%%F=%s%%O=%s%%RD=%X%%Q=%s)\n", fp.T6.R, fp.T6.DF, fp.T6.T, fp.T6.W, fp.T6.S, fp.T6.A, fp.T6.F, fp.T6.O, fp.T6.RD, fp.T6.Q)
	data += fmt.Sprintf("T7(R=%s)\n", fp.T7.R)
	data += fmt.Sprintf("U1(R=%s%%DF=%s%%T=%X%%IPL=%X%%UN=%X%%RIPL=%s%%RID=%s%%RIPCK=%s%%RUCK=%s%%RUD=%s)\n", fp.U1.R, fp.U1.DF, fp.U1.T, fp.U1.IPL, fp.U1.UN, fp.U1.RIPL, fp.U1.RID, fp.U1.RIPCK, fp.U1.RUCK, fp.U1.RUD)
	data += fmt.Sprintf("IE(R=%s%%DFI=%s%%T=%X%%CD=%s)\n", fp.IE.R, fp.IE.DFI, fp.IE.T, fp.IE.CD)
	return data
}
