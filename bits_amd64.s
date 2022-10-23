//go:build amd64

TEXT ·_bitsToBytes64(SB), $0-24

    MOVQ res+0(FP), DI
    MOVQ data+8(FP), SI

    WORD $0x8b48; BYTE $0x0e     // mov    rcx, qword [rsi]
    WORD $0xc031                 // xor    eax, eax
    WORD $0x8548; BYTE $0xc9     // test    rcx, rcx
	JE LBB0_2
LBB0_1:
    LONG $0xbc0f48f3; BYTE $0xd1 // rep        bsf    rdx, rcx
    WORD $0x1488; BYTE $0x07     // mov    byte [rdi + rax], dl
    LONG $0xff518d48             // lea    rdx, [rcx - 1]
    WORD $0xff48; BYTE $0xc0     // inc    rax
    WORD $0x2148; BYTE $0xd1     // and    rcx, rdx
	JNE LBB0_1
LBB0_2:
    LONG $0x084e8b48             // mov    rcx, qword [rsi + 8]
    WORD $0x8548; BYTE $0xc9     // test    rcx, rcx
	JE LBB0_4
LBB0_3:
    LONG $0xbc0f48f3; BYTE $0xd1 // rep        bsf    rdx, rcx
    WORD $0xca80; BYTE $0x40     // or    dl, 64
    WORD $0x1488; BYTE $0x07     // mov    byte [rdi + rax], dl
    LONG $0xff518d48             // lea    rdx, [rcx - 1]
    WORD $0xff48; BYTE $0xc0     // inc    rax
    WORD $0x2148; BYTE $0xd1     // and    rcx, rdx
	JNE LBB0_3
LBB0_4:
    LONG $0x104e8b48             // mov    rcx, qword [rsi + 16]
    WORD $0x8548; BYTE $0xc9     // test    rcx, rcx
	JE LBB0_6
LBB0_5:
    LONG $0xbc0f48f3; BYTE $0xd1 // rep        bsf    rdx, rcx
    WORD $0xca80; BYTE $0x80     // or    dl, -128
    WORD $0x1488; BYTE $0x07     // mov    byte [rdi + rax], dl
    LONG $0xff518d48             // lea    rdx, [rcx - 1]
    WORD $0xff48; BYTE $0xc0     // inc    rax
    WORD $0x2148; BYTE $0xd1     // and    rcx, rdx
	JNE LBB0_5
LBB0_6:
    LONG $0x184e8b48             // mov    rcx, qword [rsi + 24]
    WORD $0x8548; BYTE $0xc9     // test    rcx, rcx
	JE LBB0_8
LBB0_7:
    LONG $0xbc0f48f3; BYTE $0xd1 // rep        bsf    rdx, rcx
    WORD $0xca80; BYTE $0xc0     // or    dl, -64
    WORD $0x1488; BYTE $0x07     // mov    byte [rdi + rax], dl
    LONG $0xff518d48             // lea    rdx, [rcx - 1]
    WORD $0xff48; BYTE $0xc0     // inc    rax
    WORD $0x2148; BYTE $0xd1     // and    rcx, rdx
	JNE LBB0_7
LBB0_8:
    MOVQ AX, int+16(FP)
    RET
