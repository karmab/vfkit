package cmdline

import "github.com/spf13/cobra"

type Options struct {
	Vcpus     uint
	MemoryMiB uint

	VmlinuzPath   string
	KernelCmdline string
	InitrdPath    string

	Bootloader stringSliceValue

	TimeSync string

	Devices []string
}

func AddFlags(cmd *cobra.Command, opts *Options) {
	cmd.Flags().StringVarP(&opts.VmlinuzPath, "kernel", "k", "", "path to the virtual machine linux kernel")
	cmd.Flags().StringVarP(&opts.KernelCmdline, "kernel-cmdline", "C", "", "linux kernel command line")
	cmd.Flags().StringVarP(&opts.InitrdPath, "initrd", "i", "", "path to the virtual machine initrd")

	cmd.Flags().VarP(&opts.Bootloader, "bootloader", "b", "bootloader configuration")

	cmd.MarkFlagsMutuallyExclusive("kernel", "bootloader")
	cmd.MarkFlagsMutuallyExclusive("initrd", "bootloader")
	cmd.MarkFlagsMutuallyExclusive("kernel-cmdline", "bootloader")
	cmd.MarkFlagsRequiredTogether("kernel", "initrd", "kernel-cmdline")

	cmd.Flags().UintVarP(&opts.Vcpus, "cpus", "c", 1, "number of virtual CPUs")
	// FIXME: use go-units for parsing
	cmd.Flags().UintVarP(&opts.MemoryMiB, "memory", "m", 512, "virtual machine RAM size in mibibytes")

	cmd.Flags().StringVarP(&opts.TimeSync, "timesync", "t", "", "sync guest time when host wakes up from sleep")

	cmd.Flags().StringArrayVarP(&opts.Devices, "device", "d", []string{}, "devices")
}
