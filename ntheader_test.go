// Copyright 2021 Saferwall. All rights reserved.
// Use of this source code is governed by Apache v2 license
// license that can be found in the LICENSE file.

package pe

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestNtHeaderMachineType(t *testing.T) {

	tests := []struct {
		in  ImageFileHeaderMachineType
		out string
	}{
		{
			ImageFileHeaderMachineType(0x8664), "x64",
		},
		{
			ImageFileHeaderMachineType(0xffff), "?",
		},
	}

	for _, tt := range tests {
		name := "CaseNtHeaderMachineTypeEqualTo_" + strconv.Itoa(int(tt.in))
		t.Run(name, func(t *testing.T) {

			got := tt.in.String()
			if got != tt.out {
				t.Errorf("nt header machine type assertion failed, got %v, want %v",
					got, tt.out)
			}
		})
	}
}

func TestNtHeaderCharacteristicsType(t *testing.T) {

	tests := []struct {
		in  ImageFileHeaderCharacteristicsType
		out []string
	}{
		{
			ImageFileHeaderCharacteristicsType(0x0022), []string{"ExecutableImage", "LargeAddressAware"},
		},

	}

	for _, tt := range tests {
		name := "CaseNtHeaderCharacteristicsTypeEqualTo_" + strconv.Itoa(int(tt.in))
		t.Run(name, func(t *testing.T) {
			got := tt.in.String()
			sort.Strings(got)
			sort.Strings(tt.out)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("nt header Characteristics type assertion failed, got %v, want %v",
					got, tt.out)
			}
		})
	}
}

func TestOptionalHeaderSubsystemType(t *testing.T) {

	tests := []struct {
		in  ImageOptionalHeaderSubsystemType
		out string
	}{
		{
			ImageOptionalHeaderSubsystemType(0x2), "Windows GUI",
		},
		{
			ImageOptionalHeaderSubsystemType(0xff), "?",
		},
	}

	for _, tt := range tests {
		name := "CaseOptionalHeaderSubsystemTypeEqualTo_" + strconv.Itoa(int(tt.in))
		t.Run(name, func(t *testing.T) {
			got := tt.in.String()
			if got != tt.out {
				t.Errorf("optional header subsystem type assertion failed, got %v, want %v",
					got, tt.out)
			}
		})
	}
}

func TestPrettyDllCharacteristics(t *testing.T) {

	tests := []struct {
		in  string
		out []string
	}{
		{getAbsoluteFilePath("test/putty.exe"), []string{
			"DynamicBase", "HighEntropyVA", "NXCompact", "TerminalServiceAware"}},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			ops := Options{Fast: true}
			file, err := New(tt.in, &ops)
			if err != nil {
				t.Fatalf("New(%s) failed, reason: %v", tt.in, err)
			}
			err = file.Parse()
			if err != nil {
				t.Fatalf("Parse(%s) failed, reason: %v", tt.in, err)
			}

			dllCharacteristics := file.PrettyDllCharacteristics()
			sort.Strings(dllCharacteristics)
			if !reflect.DeepEqual(dllCharacteristics, tt.out) {
				t.Errorf("pretty dll characteristics type assertion failed, got %v, want %v",
					dllCharacteristics, tt.out)
			}

		})
	}
}
