//
// Author: Ayoub Chouak (@ntauth)
// File:   godyn.go
// Brief:  Godyn library functions
//
package godyn

import (
	"unsafe"
	"godyn/nt"
)

func GetProcAddress(proc string) unsafe.Pointer {

	teb := nt.GetTeb()
	peb := teb.ProcessEnvironmentBlock

	// Get a pointer to the currently loaded modules in the address space
	// note: modules is a sentry (i.e. not an actual list element)
	modules := &peb.Ldr.InMemoryOrderModuleList

	// Traverse the module list
	for module := modules.Flink; module != modules; module = module.Flink {

		ldr := (*nt.LdrDataTableEntry)(module.ContainingRecord(nt.TypeOf(nt.LdrDataTableEntry{}), "InMemoryOrderLinks"))
		_ = ldr.BaseDllName
	}

	return nil
}