// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || amd64p32 || arm || arm64 || mips64le || mips64p32le || mipsle || ppc64le || riscv64
// +build 386 amd64 amd64p32 arm arm64 mips64le mips64p32le mipsle ppc64le riscv64

package ebpf

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// loadCapture returns the embedded CollectionSpec for capture.
func loadCapture() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_CaptureBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load capture: %w", err)
	}

	return spec, err
}

// loadCaptureObjects loads capture and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*captureObjects
//	*capturePrograms
//	*captureMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadCaptureObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadCapture()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// captureSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type captureSpecs struct {
	captureProgramSpecs
	captureMapSpecs
}

// captureSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type captureProgramSpecs struct {
	Capture *ebpf.ProgramSpec `ebpf:"capture"`
}

// captureMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type captureMapSpecs struct {
	PerfMap *ebpf.MapSpec `ebpf:"perf_map"`
}

// captureObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadCaptureObjects or ebpf.CollectionSpec.LoadAndAssign.
type captureObjects struct {
	capturePrograms
	captureMaps
}

func (o *captureObjects) Close() error {
	return _CaptureClose(
		&o.capturePrograms,
		&o.captureMaps,
	)
}

// captureMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadCaptureObjects or ebpf.CollectionSpec.LoadAndAssign.
type captureMaps struct {
	PerfMap *ebpf.Map `ebpf:"perf_map"`
}

func (m *captureMaps) Close() error {
	return _CaptureClose(
		m.PerfMap,
	)
}

// capturePrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadCaptureObjects or ebpf.CollectionSpec.LoadAndAssign.
type capturePrograms struct {
	Capture *ebpf.Program `ebpf:"capture"`
}

func (p *capturePrograms) Close() error {
	return _CaptureClose(
		p.Capture,
	)
}

func _CaptureClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//go:embed capture_bpfel.o
var _CaptureBytes []byte
