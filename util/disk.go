package util

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/disk"
)

// Get disk usage for root
func GetDiskUsage() string {
	return getRootDiskUsage()
}

// Extract the disk usage for the root path
// We will iterate over partitions and find the one
// that has the mountpoint set to `/`
func getRootDiskUsage() string {
	partitions, err := disk.Partitions(false)
	if err != nil {
		fmt.Println("error occurred while finding partitions, ", err)
		return ""
	}

	for _, partition := range partitions {
		if partition.Mountpoint == "/" {
			// Get the usage
			diskUsage, err := disk.Usage(partition.Mountpoint)
			if err != nil {
				fmt.Println("error occurred while getting disk usage for root, ", err)
				return ""
			}

			total := humanize.Bytes(diskUsage.Total)
			used := humanize.Bytes(diskUsage.Used)

			return fmt.Sprint(used, " / ", total)
		}
	}

	return ""
}
