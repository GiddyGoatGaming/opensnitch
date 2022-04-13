package exprs

import "github.com/google/nftables/expr"

// NewExprIface returns a new network interface expression
func NewExprIface(iface string, isOut bool, cmpOp expr.CmpOp) *[]expr.Any {
	keyDev := expr.MetaKeyIIFNAME
	if isOut {
		keyDev = expr.MetaKeyOIFNAME
	}
	return &[]expr.Any{
		&expr.Meta{Key: keyDev, Register: 1},
		&expr.Cmp{
			Op:       cmpOp,
			Register: 1,
			Data:     ifname(iface),
		},
	}
}

// https://github.com/google/nftables/blob/master/nftables_test.go#L81
func ifname(n string) []byte {
	b := make([]byte, 16)
	copy(b, []byte(n+"\x00"))
	return b
}
