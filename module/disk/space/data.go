package space

import (
	"fmt"
	"time"

	sigar "github.com/elastic/gosigar"
	"github.com/elastic/beats/libbeat/common"
)

type FileSystemStats struct {
	FileSystem string `json:"filesystem"`
	Size       uint64 `json:"size"`
	Used       uint64 `json:"used"`
	Available  uint64 `json:"available"`
	Use        string `json:"use"`
	Mounted    string `json:"mounted"`
	ctime      time.Time
}

func eventMapping() (events []common.MapStr) {
	diskStats, _ := GetFilesystemStatList()


	// Iterate over stats and send an event per volume
	for _, stat := range diskStats {
		event := common.MapStr{
			"timestamp": common.Time(time.Now()),
			"type":      "disk",
			"disk":      stat,
		}
		events = append(events, event)

	}

	return events

}

func (f *FileSystemStats) String() string {

	output_format := "%-15s %4s %4s %5s %4s %-15s\n"
	// TODO: Implement
	return fmt.Sprintf(output_format, "Filesystem", "Size", "Used", "Avail", "Use%", "Mounted on")
}

func GetFilesystemStatList() ([]FileSystemStats, error) {
	fslist := sigar.FileSystemList{}
	fslist.Get()

	fss := make([]FileSystemStats, len(fslist.List))

	n := 0

	for _, fs := range fslist.List {
		disk, _ := GetFileSystemStat(fs)
		fss[n] = disk
		n = n + 1
	}

	return fss, nil
}

// Fetches the file system stats from sigar and puts it into the FileSystemStats object
func GetFileSystemStat(fs sigar.FileSystem) (FileSystemStats, error) {

	dir_name := fs.DirName
	usage := sigar.FileSystemUsage{}
	usage.Get(dir_name)

	filesystem := FileSystemStats{
		FileSystem: fs.DevName,
		Size:       usage.Total,
		Used:       usage.Used,
		Available:  usage.Avail,
		Use:        sigar.FormatPercent(usage.UsePercent()),
		Mounted:    dir_name,
		ctime:      time.Now(),
	}

	return filesystem, nil
}
