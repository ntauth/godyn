//
// Author: Ayoub Chouak (@ntauth)
// File:   intrinsics_amd64.s
// Brief:  NT intrinsics for amd64
//
#include "textflag.h"
#include "funcdata.h"

// func ReadFsDword(offset uint32) (uint32)
TEXT ·ReadFsDword(SB),$0-8
        MOVL offset+0(FP), AX
        // mov eax, dword ptr gs:[eax]
        BYTE $0x65; BYTE $0x8B; BYTE $0x00
        MOVL AX, ret+8(FP)
        RET
