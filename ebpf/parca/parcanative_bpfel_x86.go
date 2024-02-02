// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package parca

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type ParcaNativeProcessInfoT struct {
	ShouldUseFpByDefault uint64
	IsJitCompiler        uint64
	InterpreterType      uint64
	Len                  uint64
	Mappings             [400]struct {
		LoadAddress  uint64
		Begin        uint64
		End          uint64
		ExecutableId uint64
		Type         uint64
	}
}

type ParcaNativeStackCountKeyT struct {
	Pid                int32
	Tgid               int32
	UserStackId        uint64
	KernelStackId      uint64
	InterpreterStackId uint64
}

type ParcaNativeStackTraceT struct {
	Len       uint64
	Addresses [127]uint64
}

type ParcaNativeStackUnwindTableT struct {
	Rows [250000]struct {
		Pc        uint64
		CfaType   uint8
		RbpType   uint8
		CfaOffset int16
		RbpOffset int16
	}
}

type ParcaNativeSymbolT struct {
	ClassName  [32]int8
	MethodName [64]int8
	Path       [128]int8
}

type ParcaNativeUnwindInfoChunksT struct {
	Chunks [30]struct {
		LowPc      uint64
		HighPc     uint64
		ShardIndex uint64
		LowIndex   uint64
		HighIndex  uint64
	}
}

type ParcaNativeUnwindStateT struct {
	Ip              uint64
	Sp              uint64
	Bp              uint64
	TailCalls       uint32
	_               [4]byte
	Stack           ParcaNativeStackTraceT
	UnwindingJit    bool
	UseFp           bool
	_               [6]byte
	InterpreterType uint64
	StackKey        ParcaNativeStackCountKeyT
}

type ParcaNativeUnwinderConfigT struct {
	FilterProcesses             bool
	VerboseLogging              bool
	MixedStackEnabled           bool
	PythonEnabled               bool
	RubyEnabled                 bool
	Padding1                    bool
	Padding2                    bool
	Padding3                    bool
	RateLimitUnwindInfo         uint32
	RateLimitProcessMappings    uint32
	RateLimitRefreshProcessInfo uint32
}

type ParcaNativeUnwinderStatsT struct {
	TotalRuns                          uint64
	TotalSamples                       uint64
	SuccessDwarf                       uint64
	ErrorTruncated                     uint64
	ErrorUnsupportedExpression         uint64
	ErrorUnsupportedFramePointerAction uint64
	ErrorUnsupportedCfaRegister        uint64
	ErrorCatchall                      uint64
	ErrorShouldNeverHappen             uint64
	ErrorPcNotCovered                  uint64
	ErrorPcNotCoveredJit               uint64
	ErrorJitUnupdatedMapping           uint64
	ErrorJitMixedModeDisabled          uint64
	SuccessJitFrame                    uint64
	SuccessJitToDwarf                  uint64
	SuccessDwarfToJit                  uint64
	SuccessDwarfReachBottom            uint64
	SuccessJitReachBottom              uint64
	EventRequestUnwindInformation      uint64
	EventRequestProcessMappings        uint64
	EventRequestRefreshProcessInfo     uint64
}

// LoadParcaNative returns the embedded CollectionSpec for ParcaNative.
func LoadParcaNative() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_ParcaNativeBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load ParcaNative: %w", err)
	}

	return spec, err
}

// LoadParcaNativeObjects loads ParcaNative and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*ParcaNativeObjects
//	*ParcaNativePrograms
//	*ParcaNativeMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadParcaNativeObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadParcaNative()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// ParcaNativeSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type ParcaNativeSpecs struct {
	ParcaNativeProgramSpecs
	ParcaNativeMapSpecs
}

// ParcaNativeSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type ParcaNativeProgramSpecs struct {
	Entrypoint   *ebpf.ProgramSpec `ebpf:"entrypoint"`
	NativeUnwind *ebpf.ProgramSpec `ebpf:"native_unwind"`
}

// ParcaNativeMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type ParcaNativeMapSpecs struct {
	DebugThreadsIds    *ebpf.MapSpec `ebpf:"debug_threads_ids"`
	Events             *ebpf.MapSpec `ebpf:"events"`
	EventsCount        *ebpf.MapSpec `ebpf:"events_count"`
	Heap               *ebpf.MapSpec `ebpf:"heap"`
	PercpuStats        *ebpf.MapSpec `ebpf:"percpu_stats"`
	ProcessInfo        *ebpf.MapSpec `ebpf:"process_info"`
	Programs           *ebpf.MapSpec `ebpf:"programs"`
	StackCounts        *ebpf.MapSpec `ebpf:"stack_counts"`
	StackTraces        *ebpf.MapSpec `ebpf:"stack_traces"`
	SymbolIndexStorage *ebpf.MapSpec `ebpf:"symbol_index_storage"`
	SymbolTable        *ebpf.MapSpec `ebpf:"symbol_table"`
	UnwindInfoChunks   *ebpf.MapSpec `ebpf:"unwind_info_chunks"`
	UnwindTables       *ebpf.MapSpec `ebpf:"unwind_tables"`
}

// ParcaNativeObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadParcaNativeObjects or ebpf.CollectionSpec.LoadAndAssign.
type ParcaNativeObjects struct {
	ParcaNativePrograms
	ParcaNativeMaps
}

func (o *ParcaNativeObjects) Close() error {
	return _ParcaNativeClose(
		&o.ParcaNativePrograms,
		&o.ParcaNativeMaps,
	)
}

// ParcaNativeMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadParcaNativeObjects or ebpf.CollectionSpec.LoadAndAssign.
type ParcaNativeMaps struct {
	DebugThreadsIds    *ebpf.Map `ebpf:"debug_threads_ids"`
	Events             *ebpf.Map `ebpf:"events"`
	EventsCount        *ebpf.Map `ebpf:"events_count"`
	Heap               *ebpf.Map `ebpf:"heap"`
	PercpuStats        *ebpf.Map `ebpf:"percpu_stats"`
	ProcessInfo        *ebpf.Map `ebpf:"process_info"`
	Programs           *ebpf.Map `ebpf:"programs"`
	StackCounts        *ebpf.Map `ebpf:"stack_counts"`
	StackTraces        *ebpf.Map `ebpf:"stack_traces"`
	SymbolIndexStorage *ebpf.Map `ebpf:"symbol_index_storage"`
	SymbolTable        *ebpf.Map `ebpf:"symbol_table"`
	UnwindInfoChunks   *ebpf.Map `ebpf:"unwind_info_chunks"`
	UnwindTables       *ebpf.Map `ebpf:"unwind_tables"`
}

func (m *ParcaNativeMaps) Close() error {
	return _ParcaNativeClose(
		m.DebugThreadsIds,
		m.Events,
		m.EventsCount,
		m.Heap,
		m.PercpuStats,
		m.ProcessInfo,
		m.Programs,
		m.StackCounts,
		m.StackTraces,
		m.SymbolIndexStorage,
		m.SymbolTable,
		m.UnwindInfoChunks,
		m.UnwindTables,
	)
}

// ParcaNativePrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadParcaNativeObjects or ebpf.CollectionSpec.LoadAndAssign.
type ParcaNativePrograms struct {
	Entrypoint   *ebpf.Program `ebpf:"entrypoint"`
	NativeUnwind *ebpf.Program `ebpf:"native_unwind"`
}

func (p *ParcaNativePrograms) Close() error {
	return _ParcaNativeClose(
		p.Entrypoint,
		p.NativeUnwind,
	)
}

func _ParcaNativeClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed parcanative_bpfel_x86.o
var _ParcaNativeBytes []byte
