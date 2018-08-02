//
// Author: Ayoub Chouak (@ntauth)
// File:   intrinsics_386.s
// Brief:  NT intrinsics for i386
//
#include "textflag.h"
#include "funcdata.h"

// func ReadFsDword(offset uint32) (uint32)
TEXT ·ReadFsDword(SB),$0-8
        MOVL offset+0(FP), AX
        // mov eax, dword ptr fs:[eax]
        BYTE $0x64; BYTE $0x8B; BYTE $0x00
        MOVL AX, ret+8(FP)
        RET
