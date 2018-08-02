//
// Author: Ayoub Chouak (@ntauth)
// File:   ntdef.go
// Brief:  NT Structures Definition
//
package nt

import (
	"unsafe"
	"reflect"
)

//
// Structures
//
type wchar uint16

// typedef struct _LIST_ENTRY
type ListEntry struct {
	Flink *ListEntry
	Blink *ListEntry
}

//
// @arg ty the containing record type
// @arg field the list entry field name
// @return pointer to the containing record
func (le *ListEntry) ContainingRecord(ty reflect.Type, field string) unsafe.Pointer {
	fs, ok := ty.FieldByName(field)
	var record unsafe.Pointer = nil

	if ok {
		record = unsafe.Pointer((uintptr)(unsafe.Pointer(le)) - fs.Offset)
	}

	return record
}

func TypeOf(nt interface{}) reflect.Type {
	return reflect.TypeOf(nt)
}

type ExceptionRegistrationRecordPtr unsafe.Pointer

// typedef struct _NT_TIB
type NtTib struct
{
	ExceptionList ExceptionRegistrationRecordPtr
	StackBase unsafe.Pointer
	StackLimit unsafe.Pointer
	SubSystemTib unsafe.Pointer
	// union
	// +0x0: PVOID FiberData
	// +0x0: ULONG Version
	FiberDataVersionUnion unsafe.Pointer
	ArbitraryUserPointer unsafe.Pointer
	Self *NtTib
}

func (tib *NtTib) FiberData() unsafe.Pointer {
	return tib.FiberDataVersionUnion
}

func (tib *NtTib) Version() unsafe.Pointer {
	return tib.FiberDataVersionUnion
}

// typedef struct _CLIENT_ID
type ClientId struct
{
	UniqueProcess unsafe.Pointer
	UniqueThread unsafe.Pointer
}

// typedef struct _GDI_TEB_BATCH
type GdiTebBatch struct
{
	Offset uint32
	HDC uint32
	Buffer [310]uint32
}

type RtlActivationContextStackFramePtr unsafe.Pointer

// typedef struct _ACTIVATION_CONTEXT_STACK
type ActivationContextStack struct
{
	ActiveFrame RtlActivationContextStackFramePtr
	FrameListCache ListEntry
	Flags uint32
	NextCookieSequenceNumber uint32
	StackId uint32
}

// typedef struct _UNICODE_STRING
type UnicodeString struct
{
	Length int16
	MaximumLength int16
	Buffer *int16
}

// typedef struct _GUID
type Guid struct
{
	Data1 uint32
	Data2 int16
	Data3 int16
	Data4 [8]uint16
}

// typedef TEB_ACTIVE_FRAME_CONTEXT *PTEB_ACTIVE_FRAME_CONTEXT
type TebActiveFrameContextPtr unsafe.Pointer

// typedef struct _TEB_ACTIVE_FRAME
type TebActiveFrame struct
{
	Flags uint32
	Previous *TebActiveFrame
	Context TebActiveFrameContextPtr
}

// typedef struct _LARGE_INTEGER
type LargeInteger struct
{
	// union
	// +0x0: LowPart
	// +0x4: HighPart
	// +0x0: QuadPart
	LowPart uint32
	HighPart int32
}

func (l *LargeInteger) QuadPart() int64 {
	return (int64(l.HighPart) << 32) | int64(l.LowPart)
}

// typedef struct _ULARGE_INTEGER
type LargeIntegerUnsigned struct
{
	// union
	// +0x0: LowPart
	// +0x4: HighPart
	// +0x0: QuadPart
	LowPart uint32
	HighPart uint32
}

func (l *LargeIntegerUnsigned) QuadPart() uint64 {
	return (uint64(l.HighPart) << 32) | uint64(l.LowPart)
}

// typedef struct _TEB
type Teb struct
{
	NtTib NtTib
	EnvironmentPointer unsafe.Pointer
	ClientId ClientId
	ActiveRpcHandle unsafe.Pointer
	ThreadLocalStoragePointer unsafe.Pointer
	ProcessEnvironmentBlock *Peb
	LastErrorValue uint32
	CountOfOwnedCriticalSections uint32
	CsrClientThread unsafe.Pointer
	Win32ThreadInfo unsafe.Pointer
	User32Reserved [26]uint32
	UserReserved [5]uint32
	WOW32Reserved unsafe.Pointer
	CurrentLocale uint32
	FpSoftwareStatusRegister uint32
	SystemReserved1 [54]unsafe.Pointer
	ExceptionCode int32
	ActivationContextStackPointer *ActivationContextStack
	SpareBytes1 [36]byte
	TxFsContext uint32
	GdiTebBatch GdiTebBatch
	RealClientId ClientId
	GdiCachedProcessHandle unsafe.Pointer
	GdiClientPID uint32
	GdiClientTID uint32
	GdiThreadLocalInfo unsafe.Pointer
	Win32ClientInfo [62]uint32
	glDispatchTable [233]unsafe.Pointer
	glReserved1 [29]uint32
	glReserved2 unsafe.Pointer
	glSectionInfo unsafe.Pointer
	glSection unsafe.Pointer
	glTable unsafe.Pointer
	glCurrentRC unsafe.Pointer
	glContext unsafe.Pointer
	LastStatusValue uint32
	StaticUnicodeString UnicodeString
	StaticUnicodeBuffer [261]wchar
	DeallocationStack unsafe.Pointer
	TlsSlots [64]unsafe.Pointer
	TlsLinks ListEntry
	Vdm unsafe.Pointer
	ReservedForNtRpc unsafe.Pointer
	DbgSsReserved [2]unsafe.Pointer
	HardErrorMode uint32
	Instrumentation [9]unsafe.Pointer
	ActivityId Guid
	SubProcessTag unsafe.Pointer
	EtwLocalData unsafe.Pointer
	EtwTraceData unsafe.Pointer
	WinSockData unsafe.Pointer
	GdiBatchCount uint32
	SpareBool0 byte
	SpareBool1 byte
	SpareBool2 byte
	IdealProcessor byte
	GuaranteedStackBytes uint32
	ReservedForPerf unsafe.Pointer
	ReservedForOle unsafe.Pointer
	WaitingOnLoaderLock uint32
	SavedPriorityState unsafe.Pointer
	SoftPatchPtr1 uint32
	ThreadPoolData unsafe.Pointer
	TlsExpansionSlots unsafe.Pointer // Double Ref
	ImpersonationLocale uint32
	IsImpersonating uint32
	NlsCache unsafe.Pointer
	pShimData unsafe.Pointer
	HeapVirtualAffinity uint32
	CurrentTransactionHandle unsafe.Pointer
	ActiveFrame *TebActiveFrame
	FlsData unsafe.Pointer
	PreferredLanguages unsafe.Pointer
	UserPrefLanguages unsafe.Pointer
	MergedPrefLanguages unsafe.Pointer
	MuiImpersonation uint32
	CrossTebFlags int16
	SpareCrossTebBits uint16 // SpareCrossTebBits: 16
	SameTebFlags int16
	// Bitfield
	DbgBitfield byte
	//ULONG DbgSafeThunkCall: 1;
	//ULONG DbgInDebugPrint: 1;
	//ULONG DbgHasFiberData: 1;
	//ULONG DbgSkipThreadAttach: 1;
	//ULONG DbgWerInShipAssertCode: 1;
	//ULONG DbgRanProcessInit: 1;
	//ULONG DbgClonedThread: 1;
	//ULONG DbgSuppressDebugMsg: 1;
	SpareSameTebBits byte // ULONG SpareSameTebBits: 8;
	TxnScopeEnterCallback unsafe.Pointer
	TxnScopeExitCallback unsafe.Pointer
	TxnScopeContext unsafe.Pointer
	LockCount uint32
	ProcessRundown uint32
	LastSwitchTime uint64
	TotalSwitchOutTime uint64
	WaitReasonBitMap LargeInteger
}

// typedef PEB_LDR_DATA *PPEB_LDR_DATA
type PebLdrDataPtr *PebLdrData

// typedef RTL_USER_PROCESS_PARAMETERS *PRTL_USER_PROCESS_PARAMETERS
type RtlUserProcessParametersPtr unsafe.Pointer

// typedef RTL_CRITICAL_SECTION *PRTL_CRITICAL_SECTION
type RtlCriticalSectionPtr unsafe.Pointer

// typedef PEB_FREE_BLOCK *PPEB_FREE_BLOCK
type PebFreeBlockPtr unsafe.Pointer

// typedef _ACTIVATION_CONTEXT_DATA *PACTIVATION_CONTEXT_DATA
type ActivationContextDataPtr unsafe.Pointer

// typedef _ASSEMBLY_STORAGE_MAP *PASSEMBLY_STORAGE_MAP
type AssemblyStorageMapPtr unsafe.Pointer

// typedef _FLS_CALLBACK_INFO *PFLS_CALLBACK_INFO
type FlsCallbackInfoPtr unsafe.Pointer


// typedef struct _PEB
type Peb struct
{
	InheritedAddressSpace byte
	ReadImageFileExecOptions byte
	BeingDebugged byte
	BitField byte
	// Bitfield
	ProcessFeatureBitfield byte
	//ULONG ImageUsesLargePages: 1;
	//ULONG IsProtectedProcess: 1;
	//ULONG IsLegacyProcess: 1;
	//ULONG IsImageDynamicallyRelocated: 1;
	//ULONG SpareBits: 4;
	Mutant unsafe.Pointer
	ImageBaseAddress unsafe.Pointer
	Ldr PebLdrDataPtr
	ProcessParameters RtlUserProcessParametersPtr
	SubSystemData unsafe.Pointer
	ProcessHeap unsafe.Pointer
	FastPebLock RtlCriticalSectionPtr
	AtlThunkSListPtr unsafe.Pointer
	IFEOKey unsafe.Pointer
	CrossProcessFlags uint32

	// Bitfield
	ProcessStatusBitfield uint32
	//ULONG ProcessInJob: 1;
	//ULONG ProcessInitializing: 1;
	//ULONG ReservedBits0: 30;

	// Union
	ProcessPtrUnion0 unsafe.Pointer
	// PVOID KernelCallbackTable;
	// PVOID UserSharedInfoPtr;

	SystemReserved [1]uint32
	SpareUlong uint32
	FreeList PebFreeBlockPtr
	TlsExpansionCounter uint32
	TlsBitmap unsafe.Pointer
	TlsBitmapBits [2]uint32
	ReadOnlySharedMemoryBase unsafe.Pointer
	HotpatchInformation unsafe.Pointer
	ReadOnlyStaticServerData unsafe.Pointer // Double Ref
	AnsiCodePageData unsafe.Pointer
	OemCodePageData unsafe.Pointer
	UnicodeCaseTableData unsafe.Pointer
	NumberOfProcessors uint32
	NtGlobalFlag uint32
	CriticalSectionTimeout LargeInteger
	HeapSegmentReserve uint32
	HeapSegmentCommit uint32
	HeapDeCommitTotalFreeThreshold uint32
	HeapDeCommitFreeBlockThreshold uint32
	NumberOfHeaps uint32
	MaximumNumberOfHeaps uint32
	ProcessHeaps unsafe.Pointer // Double Ref
	GdiSharedHandleTable unsafe.Pointer
	ProcessStarterHelper unsafe.Pointer
	GdiDCAttributeList uint32
	LoaderLock RtlCriticalSectionPtr
	OSMajorVersion uint32
	OSMinorVersion uint32
	OSBuildNumber uint16
	OSCSDVersion uint16
	OSPlatformId uint32
	ImageSubsystem uint32
	ImageSubsystemMajorVersion uint32
	ImageSubsystemMinorVersion uint32
	ImageProcessAffinityMask uint32
	GdiHandleBuffer [34]uint32
	PostProcessInitRoutine unsafe.Pointer
	TlsExpansionBitmap unsafe.Pointer
	TlsExpansionBitmapBits [32]uint32
	SessionId uint32
	AppCompatFlags LargeIntegerUnsigned
	AppCompatFlagsUser LargeIntegerUnsigned
	pShimData unsafe.Pointer
	AppCompatInfo unsafe.Pointer
	CSDVersion UnicodeString
	ActivationContextData ActivationContextDataPtr
	ProcessAssemblyStorageMap AssemblyStorageMapPtr
	SystemDefaultActivationContextData ActivationContextDataPtr
	SystemAssemblyStorageMap AssemblyStorageMapPtr
	MinimumStackCommit uint32
	FlsCallback FlsCallbackInfoPtr
	FlsListHead ListEntry
	FlsBitmap unsafe.Pointer
	FlsBitmapBits [4]uint32
	FlsHighIndex uint32
	WerRegistrationData unsafe.Pointer
	WerShipAssertPtr unsafe.Pointer
}

// typedef struct _PEB_LDR_DATA
type PebLdrData struct
{
	Length uint32
	Initialized byte
	SsHandle unsafe.Pointer
	InLoadOrderModuleList ListEntry
	InMemoryOrderModuleList ListEntry
	InInitializationOrderModuleList ListEntry
	EntryInProgress unsafe.Pointer
	ShutdownInProgress byte
	ShutdownThreadId unsafe.Pointer
}

// typedef _ACTIVATION_CONTEXT *PACTIVATION_CONTEXT
type ActivationContextPtr unsafe.Pointer

// typedef struct _LDR_DATA_TABLE_ENTRY
type LdrDataTableEntry struct
{
	InMemoryOrderLinks ListEntry
	InInitializationOrderLinks ListEntry
	DllBase unsafe.Pointer
	EntryPoint unsafe.Pointer
	SizeOfImage uint32
	FullDllName UnicodeString
	BaseDllName UnicodeString
	Flags uint32
	LoadCount int16
	TlsIndex int16

	// Union
	Union0 uint64
	// LIST_ENTRY HashLinks;
	// struct
	// {
	//	 PVOID SectionPointer;
	//	 ULONG CheckSum;
	// };

	// Union
	Union1 uintptr
	// ULONG TimeDateStamp;
	// PVOID LoadedImports;

	EntryPointActivationContext ActivationContextPtr
	PatchInformation unsafe.Pointer
	ForwarderLinks ListEntry
	ServiceTagLinks ListEntry
	StaticLinks ListEntry
}

// PE Header Structures


//
// Methods
//
func GetTeb() (teb *Teb) {
	teb = (*Teb)(unsafe.Pointer(uintptr(ReadFsDword(NtTibTebOffset))))
	return
}