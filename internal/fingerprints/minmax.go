package fingerprints

import "fmt"

func MinList(list []int64) int64 {
	minI64 := list[0]
	for _, v := range list {
		if v < minI64 {
			minI64 = v
		}
	}
	return minI64
}

func MaxList(list []int64) int64 {
	maxI64 := list[0]
	for _, v := range list {
		if v > maxI64 {
			maxI64 = v
		}
	}
	return maxI64

}

func MinMaxSeq(seqs []*SEQ) []*SEQ {
	MinSeq := &SEQ{}
	MaxSeq := &SEQ{}
	var spList []int64
	var gcdList []int64
	var isrList []int64
	var tsList []int64
	for _, seq := range seqs {
		spList = append(spList, seq.SP)
		gcdList = append(gcdList, seq.GCD)
		isrList = append(isrList, seq.ISR)
		tsList = append(tsList, seq.TS)
	}
	MinSeq.SP = MinList(spList)
	// fmt.Printf("MinSeq.SP: %d\n", MinSeq.SP)
	MaxSeq.SP = MaxList(spList)
	// fmt.Printf("MaxSeq.SP: %d\n", MaxSeq.SP)
	MinSeq.GCD = MinList(gcdList)
	MaxSeq.GCD = MaxList(gcdList)
	MinSeq.ISR = MinList(isrList)
	MaxSeq.ISR = MaxList(isrList)
	MinSeq.TS = MinList(tsList)
	MaxSeq.TS = MaxList(tsList)
	// fmt.Printf("MinSp: %d, MaxSp: %d, MinGcd: %d, MaxGcd: %d, MinIsr: %d, MaxIsr: %d, MinTs: %d, MaxTs: %d\n", MinSeq.SP, MaxSeq.SP, MinSeq.GCD, MaxSeq.GCD, MinSeq.ISR, MaxSeq.ISR, MinSeq.TS, MaxSeq.TS)
	return []*SEQ{MinSeq, MaxSeq}
}

func MinMaxWin(wins []*WIN) []*WIN {
	MinWin := &WIN{}
	MaxWin := &WIN{}
	var w1List []int64
	var w2List []int64
	var w3List []int64
	var w4List []int64
	var w5List []int64
	var w6List []int64
	for _, win := range wins {
		w1List = append(w1List, win.W1)
		w2List = append(w2List, win.W2)
		w3List = append(w3List, win.W3)
		w4List = append(w4List, win.W4)
		w5List = append(w5List, win.W5)
		w6List = append(w6List, win.W6)

	}
	MinWin.W1 = MinList(w1List)
	MaxWin.W1 = MaxList(w1List)
	MinWin.W2 = MinList(w2List)
	MaxWin.W2 = MaxList(w2List)
	MinWin.W3 = MinList(w3List)
	MaxWin.W3 = MaxList(w3List)
	MinWin.W4 = MinList(w4List)
	MaxWin.W4 = MaxList(w4List)
	MinWin.W5 = MinList(w5List)
	MaxWin.W5 = MaxList(w5List)
	MinWin.W6 = MinList(w6List)
	MaxWin.W6 = MaxList(w6List)
	return []*WIN{MinWin, MaxWin}
}

func MinMaxEcn(ecns []*ECN) []*ECN {
	MinEcn := &ECN{}
	MaxEcn := &ECN{}
	var tList []int64
	var wList []int64
	for _, ecn := range ecns {
		tList = append(tList, ecn.T)
		wList = append(wList, ecn.W)
	}
	MinEcn.T = MinList(tList)
	MaxEcn.T = MaxList(tList)
	MinEcn.W = MinList(wList)
	MaxEcn.W = MaxList(wList)
	return []*ECN{MinEcn, MaxEcn}
}

func MinMaxT1(t1s []*T1) []*T1 {
	MinT1 := &T1{}
	MaxT1 := &T1{}
	var tList []int64
	var rdList []int64
	for _, t1 := range t1s {
		tList = append(tList, t1.T)
		rdList = append(rdList, t1.RD)
	}
	MinT1.T = MinList(tList)
	MaxT1.T = MaxList(tList)
	MinT1.RD = MinList(rdList)
	MaxT1.RD = MaxList(rdList)

	return []*T1{MinT1, MaxT1}
}

func MinMaxT4(t4s []*T4) []*T4 {
	MinT4 := &T4{}
	MaxT4 := &T4{}
	var tList []int64
	var wList []int64
	var rdList []int64
	for _, t4 := range t4s {
		tList = append(tList, t4.T)
		wList = append(wList, t4.W)
		rdList = append(rdList, t4.RD)
	}
	MinT4.T = MinList(tList)
	MaxT4.T = MaxList(tList)
	MinT4.W = MinList(wList)
	MaxT4.W = MaxList(wList)
	MinT4.RD = MinList(rdList)
	MaxT4.RD = MaxList(rdList)

	return []*T4{MinT4, MaxT4}
}

func MinMaxT5(t5s []*T5) []*T5 {
	MinT5 := &T5{}
	MaxT5 := &T5{}
	var tList []int64
	var wList []int64
	var rdList []int64
	for _, t5 := range t5s {
		tList = append(tList, t5.T)
		wList = append(wList, t5.W)
		rdList = append(rdList, t5.RD)
	}
	MinT5.T = MinList(tList)
	MaxT5.T = MaxList(tList)
	MinT5.W = MinList(wList)
	MaxT5.W = MaxList(wList)
	MinT5.RD = MinList(rdList)
	MaxT5.RD = MaxList(rdList)
	return []*T5{MinT5, MaxT5}
}

func MinMaxT6(t6s []*T6) []*T6 {
	MinT6 := &T6{}
	MaxT6 := &T6{}
	var tList []int64
	var wList []int64
	var rdList []int64
	for _, t6 := range t6s {
		tList = append(tList, t6.T)
		wList = append(wList, t6.W)
		rdList = append(rdList, t6.RD)
	}
	MinT6.T = MinList(tList)
	MaxT6.T = MaxList(tList)
	MinT6.W = MinList(wList)
	MaxT6.W = MaxList(wList)
	MinT6.RD = MinList(rdList)
	MaxT6.RD = MaxList(rdList)
	return []*T6{MinT6, MaxT6}
}

func MinMaxU1(u1s []*U1) []*U1 {
	MinU1 := &U1{}
	MaxU1 := &U1{}
	var tList []int64
	var iplList []int64
	var unList []int64
	for _, u1 := range u1s {
		tList = append(tList, u1.T)
		iplList = append(iplList, u1.IPL)
		unList = append(unList, u1.UN)

	}
	MinU1.T = MinList(tList)
	MaxU1.T = MaxList(tList)
	MinU1.IPL = MinList(iplList)
	MaxU1.IPL = MaxList(iplList)
	MinU1.UN = MinList(unList)
	MaxU1.UN = MaxList(unList)

	return []*U1{MinU1, MaxU1}
}

func MinMaxIe(ies []*IE) []*IE {
	MinIe := &IE{}
	MaxIe := &IE{}
	var tList []int64
	for _, ie := range ies {
		tList = append(tList, ie.T)
	}
	MinIe.T = MinList(tList)
	MaxIe.T = MaxList(tList)
	return []*IE{MinIe, MaxIe}
}

func GenFinalFingerprint(fps []*Fingerprint) string {
	var seqs []*SEQ
	var ops []*OPS
	var wins []*WIN
	var ecns []*ECN
	var t1s []*T1
	var t4s []*T4
	var t5s []*T5
	var t6s []*T6
	var t7s []*T7
	var u1s []*U1
	var ies []*IE
	for _, f := range fps {
		for _, seq := range f.SEQs {
			seqs = append(seqs, &seq)
		}
		ops = append(ops, &f.OPS)
		wins = append(wins, &f.WIN)
		ecns = append(ecns, &f.ECN)
		t1s = append(t1s, &f.T1)
		t4s = append(t4s, &f.T4)
		t5s = append(t5s, &f.T5)
		t6s = append(t6s, &f.T6)
		t7s = append(t7s, &f.T7)
		u1s = append(u1s, &f.U1)
		ies = append(ies, &f.IE)
	}
	MinSeq := MinMaxSeq(seqs)[0]
	MaxSeq := MinMaxSeq(seqs)[1]
	MinWin := MinMaxWin(wins)[0]
	MaxWin := MinMaxWin(wins)[1]
	MinEcn := MinMaxEcn(ecns)[0]
	MaxEcn := MinMaxEcn(ecns)[1]
	MinT1 := MinMaxT1(t1s)[0]
	MaxT1 := MinMaxT1(t1s)[1]
	MinT4 := MinMaxT4(t4s)[0]
	MaxT4 := MinMaxT4(t4s)[1]
	MinT5 := MinMaxT5(t5s)[0]
	MaxT5 := MinMaxT5(t5s)[1]
	MinT6 := MinMaxT6(t6s)[0]
	MaxT6 := MinMaxT6(t6s)[1]
	MinU1 := MinMaxU1(u1s)[0]
	MaxU1 := MinMaxU1(u1s)[1]
	MinIe := MinMaxIe(ies)[0]
	MaxIe := MinMaxIe(ies)[1]

	seqStr := GenSeqString(MinSeq.SP, MaxSeq.SP, MinSeq.GCD, MaxSeq.GCD, MinSeq.ISR, MaxSeq.ISR, seqs[0].TI, seqs[0].CI, seqs[0].II, MinSeq.TS, MaxSeq.TS)
	opsStr := GenOpsString(ops[0].O1, ops[0].O2, ops[0].O3, ops[0].O4, ops[0].O5, ops[0].O6)
	winStr := GenWinString(MinWin.W1, MaxWin.W1, MinWin.W2, MaxWin.W2, MinWin.W3, MaxWin.W3, MinWin.W4, MaxWin.W4, MinWin.W5, MaxWin.W5, MinWin.W6, MaxWin.W6)
	ecnStr := GenEcnString(ecns[0].R, ecns[0].DF, MinEcn.T, MaxEcn.T, MinEcn.W, MaxEcn.W, ecns[0].O, ecns[0].CC, ecns[0].Q)
	t1Str := GenT1String(t1s[0].R, t1s[0].DF, MinT1.T, MaxT1.T, t1s[0].S, t1s[0].A, t1s[0].F, MinT1.RD, MaxT1.RD, t1s[0].Q)
	t2Str := GenT2String(fps[0].T2.R)
	t3Str := GenT3String(fps[0].T3.R)
	t4Str := GenT4String(t4s[0].R, t4s[0].DF, MinT4.T, MaxT4.T, MinT4.W, MaxT4.W, t4s[0].S, t4s[0].A, t4s[0].F, t4s[0].O, MinT4.RD, MaxT4.RD, t4s[0].Q)
	t5Str := GenT5String(t5s[0].R, t5s[0].DF, MinT5.T, MaxT5.T, MinT5.W, MaxT5.W, t5s[0].S, t5s[0].A, t5s[0].F, t5s[0].O, MinT5.RD, MaxT5.RD, t5s[0].Q)
	t6Str := GenT6String(t6s[0].R, t6s[0].DF, MinT6.T, MaxT6.T, MinT6.W, MaxT6.W, t6s[0].S, t6s[0].A, t6s[0].F, t6s[0].O, MinT6.RD, MaxT6.RD, t6s[0].Q)
	t7Str := GenT7String(fps[0].T7.R)
	u1Str := GenU1String(u1s[0].R, u1s[0].DF, MinU1.T, MaxU1.T, MinU1.IPL, MaxU1.IPL, MinU1.UN, MaxU1.UN, u1s[0].RIPL, u1s[0].RID, u1s[0].RIPCK, u1s[0].RUCK, u1s[0].RUD)
	ieStr := GenIeString(ies[0].R, ies[0].DFI, MinIe.T, MaxIe.T, ies[0].CD)
	result := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n", seqStr, opsStr, winStr, ecnStr, t1Str, t2Str, t3Str, t4Str, t5Str, t6Str, t7Str, u1Str, ieStr)
	return result

}

func GenSeqString(minSp int64, maxSp int64, minGcd int64, maxGcd int64, minIsr int64, maxIsr int64, ti string, ci string, ii string, minTs int64, maxTs int64) string {
	var spStr, gcdStr, isrStr, tsStr string
	if minSp == maxSp {
		spStr = fmt.Sprintf("%s", parseInt64toStr(minSp))
	} else {
		spStr = fmt.Sprintf("%s-%s", parseInt64toStr(minSp), parseInt64toStr(maxSp))
	}
	if minGcd == maxGcd {
		gcdStr = fmt.Sprintf("%s", parseInt64toStr(minGcd))
	} else {
		gcdStr = fmt.Sprintf("%s-%s", parseInt64toStr(minGcd), parseInt64toStr(maxGcd))

	}
	if minIsr == maxIsr {
		isrStr = fmt.Sprintf("%s", parseInt64toStr(minIsr))
	} else {
		isrStr = fmt.Sprintf("%s-%s", parseInt64toStr(minIsr), parseInt64toStr(maxIsr))
	}
	tiStr := fmt.Sprintf("%s", ti)
	ciStr := fmt.Sprintf("%s", ci)
	iiStr := fmt.Sprintf("%s", ii)
	if minTs == maxTs {
		tsStr = fmt.Sprintf("%s", parseInt64toStr(minTs))
	} else {
		tsStr = fmt.Sprintf("%s-%s", parseInt64toStr(minTs), parseInt64toStr(maxTs))

	}
	result := fmt.Sprintf("SEQ(SP=%s%%GCD=%s%%ISR=%s%%TI=%s%%CI=%s%%II=%s%%TS=%s)", spStr, gcdStr, isrStr, tiStr, ciStr, iiStr, tsStr)
	return result
}

func GenOpsString(o1 string, o2 string, o3 string, o4 string, o5 string, o6 string) string {
	result := fmt.Sprintf("OPS(O1=%s%%O2=%s%%O3=%s%%O4=%s%%O5=%s%%O6=%s)", o1, o2, o3, o4, o5, o6)
	return result
}

func GenWinString(minW1 int64, maxW1 int64, minW2 int64, maxW2 int64, minW3 int64, maxW3 int64, minW4 int64, maxW4 int64, minW5 int64, maxW5 int64, minW6 int64, maxW6 int64) string {
	var w1Str, w2Str, w3Str, w4Str, w5Str, w6Str string
	if minW1 == maxW1 {
		w1Str = fmt.Sprintf("%s", parseInt64toStr(minW1))
	} else {
		w1Str = fmt.Sprintf("%s-%s", parseInt64toStr(minW1), parseInt64toStr(maxW1))
	}
	if minW2 == maxW2 {
		w2Str = fmt.Sprintf("%s", parseInt64toStr(minW2))
	} else {
		w2Str = fmt.Sprintf("%s-%s", parseInt64toStr(minW2), parseInt64toStr(maxW2))
	}
	if minW3 == maxW3 {
		w3Str = fmt.Sprintf("%s", parseInt64toStr(minW3))
	} else {
		w3Str = fmt.Sprintf("%s-%s", parseInt64toStr(minW3), parseInt64toStr(maxW3))
	}
	if minW4 == maxW4 {
		w4Str = fmt.Sprintf("%s", parseInt64toStr(minW4))
	} else {
		w4Str = fmt.Sprintf("%s-%s", parseInt64toStr(minW4), parseInt64toStr(maxW4))
	}
	if minW5 == maxW5 {
		w5Str = fmt.Sprintf("%s", parseInt64toStr(minW5))
	} else {
		w5Str = fmt.Sprintf("%s-%s", parseInt64toStr(minW5), parseInt64toStr(maxW5))
	}
	if minW6 == maxW6 {
		w6Str = fmt.Sprintf("%s", parseInt64toStr(minW6))
	} else {
		w6Str = fmt.Sprintf("%s-%s", parseInt64toStr(minW6), parseInt64toStr(maxW6))
	}
	result := fmt.Sprintf("WIN(W1=%s%%W2=%s%%W3=%s%%W4=%s%%W5=%s%%W6=%s)", w1Str, w2Str, w3Str, w4Str, w5Str, w6Str)
	return result
}

func GenEcnString(r string, df string, minT int64, maxT int64, minW int64, maxW int64, o string, cc string, q string) string {
	var tStr, wStr string
	if minT == maxT {
		tStr = fmt.Sprintf("%s", parseInt64toStr(minT))
	} else {
		tStr = fmt.Sprintf("%s-%s", parseInt64toStr(minT), parseInt64toStr(maxT))
	}
	if minW == maxW {
		wStr = fmt.Sprintf("%s", parseInt64toStr(minW))
	} else {
		wStr = fmt.Sprintf("%s-%s", parseInt64toStr(minW), parseInt64toStr(maxW))
	}
	result := fmt.Sprintf("ECN(R=%s%%DF=%s%%T=%s%%W=%s%%O=%s%%CC=%s%%Q=%s)", r, df, tStr, wStr, o, cc, q)
	return result
}

func GenT1String(r string, df string, minT int64, maxT int64, s string, a string, f string, minRD int64, maxRD int64, q string) string {
	var tStr, rdStr string
	if minT == maxT {
		tStr = fmt.Sprintf("%s", parseInt64toStr(minT))
	} else {
		tStr = fmt.Sprintf("%s-%s", parseInt64toStr(minT), parseInt64toStr(maxT))
	}
	if minRD == maxRD {
		rdStr = fmt.Sprintf("%s", parseInt64toStr(minRD))
	} else {
		rdStr = fmt.Sprintf("%s-%s", parseInt64toStr(minRD), parseInt64toStr(maxRD))
	}
	//T1(R=Y%DF=Y%T=40%S=O%A=S+%F=AS%RD=0%Q=)
	result := fmt.Sprintf("T1(R=%s%%DF=%s%%T=%s%%S=%s%%A=%s%%F=%s%%RD=%s%%Q=%s)", r, df, tStr, s, a, f, rdStr, q)
	return result
}

func GenT2String(r string) string {
	result := fmt.Sprintf("T2(R=%s)", r)
	return result
}

func GenT3String(r string) string {
	result := fmt.Sprintf("T3(R=%s)", r)
	return result
}

func GenT4String(r string, df string, minT int64, maxT int64, minW int64, maxW int64, s string, a string, f string, o string, minRD int64, maxRD int64, q string) string {
	var tStr, wStr, rdStr string
	if minT == maxT {
		tStr = fmt.Sprintf("%s", parseInt64toStr(minT))
	} else {
		tStr = fmt.Sprintf("%s-%s", parseInt64toStr(minT), parseInt64toStr(maxT))
	}
	if minW == maxW {
		wStr = fmt.Sprintf("%s", parseInt64toStr(minW))
	} else {
		wStr = fmt.Sprintf("%s-%s", parseInt64toStr(minW), parseInt64toStr(maxW))
	}
	if minRD == maxRD {
		rdStr = fmt.Sprintf("%s", parseInt64toStr(minRD))
	} else {
		rdStr = fmt.Sprintf("%s-%s", parseInt64toStr(minRD), parseInt64toStr(maxRD))
	}
	// T4(R=Y%DF=Y%T=40%W=0%S=A%A=Z%F=R%O=%RD=0%Q=)
	result := fmt.Sprintf("T4(R=%s%%DF=%s%%T=%s%%W=%s%%S=%s%%A=%s%%F=%s%%O=%s%%RD=%s%%Q=%s)", r, df, tStr, wStr, s, a, f, o, rdStr, q)
	return result
}

func GenT5String(r string, df string, minT int64, maxT int64, minW int64, maxW int64, s string, a string, f string, o string, minRD int64, maxRD int64, q string) string {
	var tStr, wStr, rdStr string
	if minT == maxT {
		tStr = fmt.Sprintf("%s", parseInt64toStr(minT))
	} else {
		tStr = fmt.Sprintf("%s-%s", parseInt64toStr(minT), parseInt64toStr(maxT))
	}
	if minW == maxW {
		wStr = fmt.Sprintf("%s", parseInt64toStr(minW))
	} else {
		wStr = fmt.Sprintf("%s-%s", parseInt64toStr(minW), parseInt64toStr(maxW))
	}
	if minRD == maxRD {
		rdStr = fmt.Sprintf("%s", parseInt64toStr(minRD))
	} else {
		rdStr = fmt.Sprintf("%s-%s", parseInt64toStr(minRD), parseInt64toStr(maxRD))
	}
	// T5(R=Y%DF=Y%T=40%W=0%S=Z%A=S+%F=AR%O=%RD=0%Q=)
	result := fmt.Sprintf("T5(R=%s%%DF=%s%%T=%s%%W=%s%%S=%s%%A=%s%%F=%s%%O=%s%%RD=%s%%Q=%s)", r, df, tStr, wStr, s, a, f, o, rdStr, q)
	return result
}

func GenT6String(r string, df string, minT int64, maxT int64, minW int64, maxW int64, s string, a string, f string, o string, minRD int64, maxRD int64, q string) string {
	var tStr, wStr, rdStr string
	if minT == maxT {
		tStr = fmt.Sprintf("%s", parseInt64toStr(minT))
	} else {
		tStr = fmt.Sprintf("%s-%s", parseInt64toStr(minT), parseInt64toStr(maxT))
	}
	if minW == maxW {
		wStr = fmt.Sprintf("%s", parseInt64toStr(minW))
	} else {
		wStr = fmt.Sprintf("%s-%s", parseInt64toStr(minW), parseInt64toStr(maxW))
	}
	if minRD == maxRD {
		rdStr = fmt.Sprintf("%s", parseInt64toStr(minRD))
	} else {
		rdStr = fmt.Sprintf("%s-%s", parseInt64toStr(minRD), parseInt64toStr(maxRD))
	}
	// T6(R=Y%DF=Y%T=40%W=0%S=A%A=Z%F=R%O=%RD=0%Q=)
	result := fmt.Sprintf("T6(R=%s%%DF=%s%%T=%s%%W=%s%%S=%s%%A=%s%%F=%s%%O=%s%%RD=%s%%Q=%s)", r, df, tStr, wStr, s, a, f, o, rdStr, q)
	return result
}

func GenT7String(r string) string {
	result := fmt.Sprintf("T7(R=%s)", r)
	return result
}

func GenU1String(r string, df string, minT int64, maxT int64, minIPL int64, maxIPL int64, minUN int64, maxUN int64, ripl string, rid string, ripck string, ruck string, rud string) string {
	var tStr, iplStr, unStr string
	if minT == maxT {
		tStr = fmt.Sprintf("%s", parseInt64toStr(minT))
	} else {
		tStr = fmt.Sprintf("%s-%s", parseInt64toStr(minT), parseInt64toStr(maxT))
	}
	if minIPL == maxIPL {
		iplStr = fmt.Sprintf("%s", parseInt64toStr(minIPL))
	} else {
		iplStr = fmt.Sprintf("%s-%s", parseInt64toStr(minIPL), parseInt64toStr(maxIPL))
	}
	if minUN == maxUN {
		unStr = fmt.Sprintf("%s", parseInt64toStr(minUN))
	} else {
		unStr = fmt.Sprintf("%s-%s", parseInt64toStr(minUN), parseInt64toStr(maxUN))
	}
	// U1(R=Y%DF=N%T=40%IPL=164%UN=0%RIPL=G%RID=G%RIPCK=G%RUCK=G%RUD=G)
	result := fmt.Sprintf("U1(R=%s%%DF=%s%%T=%s%%IPL=%s%%UN=%s%%RIPL=%s%%RID=%s%%RIPCK=%s%%RUCK=%s%%RUD=%s)", r, df, tStr, iplStr, unStr, ripl, rid, ripck, ruck, rud)
	return result
}

func GenIeString(r string, dfi string, minT int64, maxT int64, cd string) string {
	var tStr string
	if minT == maxT {
		tStr = fmt.Sprintf("%s", parseInt64toStr(minT))
	} else {
		tStr = fmt.Sprintf("%s-%s", parseInt64toStr(minT), parseInt64toStr(maxT))
	}
	// IE(R=Y%DFI=N%T=40%CD=S)
	result := fmt.Sprintf("IE(R=%s%%DFI=%s%%T=%s%%CD=%s)", r, dfi, tStr, cd)
	return result
}
