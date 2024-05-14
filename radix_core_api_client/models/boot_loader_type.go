package models

import (
	"errors"
)

type BootLoaderType int

const (
	KERNELBOOT_BOOTLOADERTYPE BootLoaderType = iota
	SYSTEMBOOT_BOOTLOADERTYPE
	VMBOOT_BOOTLOADERTYPE
)

func (i BootLoaderType) String() string {
	return []string{"KernelBoot", "SystemBoot", "VmBoot"}[i]
}

func ParseBootLoaderType(v string) (any, error) {
	result := KERNELBOOT_BOOTLOADERTYPE
	switch v {
	case "KernelBoot":
		result = KERNELBOOT_BOOTLOADERTYPE
	case "SystemBoot":
		result = SYSTEMBOOT_BOOTLOADERTYPE
	case "VmBoot":
		result = VMBOOT_BOOTLOADERTYPE
	default:
		return 0, errors.New("Unknown BootLoaderType value: " + v)
	}
	return &result, nil
}

func SerializeBootLoaderType(values []BootLoaderType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}

func (i BootLoaderType) isMultiValue() bool {
	return false
}
