// +build amd64

TEXT ·Obfuscate(SB), $0-16
    MOVQ data+0(FP), AX
    MOVQ key+8(FP), BX
    XORQ BX, AX
    MOVQ AX, ret+16(FP)
    RET
