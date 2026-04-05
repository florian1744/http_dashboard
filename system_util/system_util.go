package systemutil

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

type MetricData struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Value string `json:"value"`
	Color string `json:"color"`
}

type MetricSection struct {
	Name  string       `json:"name"`
	Cards []MetricData `json:"cards"`
}

type DashboardData struct {
	Title    string          `json:"title"`
	Sections []MetricSection `json:"sections"`
}

func BuildDashboardData() (DashboardData, error) {
	vmem, err := mem.VirtualMemory()
	cpuInfo, err := cpu.Info()
	hardDriveInfo, err := disk.Usage("/")

	if err != nil {
		return DashboardData{}, err
	}

	data := DashboardData{
		Title: "Systemressourcen",
		Sections: []MetricSection{
			{
				Name: "RAM",
				Cards: []MetricData{
					{
						ID:    "ram-total",
						Label: "Gesamtspeicher",
						Value: fmt.Sprintf("%d MB", convertByteToMB(int(vmem.Total))),
						Color: "accent",
					},
					{
						ID:    "ram-available",
						Label: "Verfügbar",
						Value: fmt.Sprintf("%d MB", convertByteToMB(int(vmem.Available))),
						Color: "success",
					},
					{
						ID:    "ram-used",
						Label: "Belegt",
						Value: fmt.Sprintf("%d MB", convertByteToMB(int(vmem.Used))),
						Color: "warning",
					},
					{
						ID:    "ram-used-percent",
						Label: "Auslastung",
						Value: fmt.Sprintf("%.2f %%", vmem.UsedPercent),
						Color: "",
					},
				},
			},

			{
				Name: "CPU",
				Cards: []MetricData{

					{
						ID:    "cpu-model-name",
						Label: "CPU Name",
						Value: fmt.Sprintf("%v", cpuInfo[0].ModelName),
						Color: "accent",
					},
					{
						ID:    "cpu-cores",
						Label: "CPU cores",
						Value: fmt.Sprintf("%v", cpuInfo[0].Cores),
						Color: "",
					},
					{
						ID:    "cpu-mhz",
						Label: "CPU mhz",
						Value: fmt.Sprintf("%v", cpuInfo[0].Mhz),
						Color: "",
					},
				},
			},
			{
				Name: "Harddrive",
				Cards: []MetricData{
					{
						ID:    "hardDrive-total",
						Label: "Hard Drive total",
						Value: fmt.Sprintf("%d MB", convertByteToMB(int(hardDriveInfo.Total))),
						Color: "",
					},
					{
						ID:    "hardDrive-used",
						Label: "Hard Drive used",
						Value: fmt.Sprintf("%d MB", convertByteToMB(int(hardDriveInfo.Used))),
						Color: "",
					},
					{
						ID:    "hardDrive-used-percent",
						Label: "Hard Drive percent",
						Value: fmt.Sprintf("%d MB", hardDriveInfo.UsedPercent),
						Color: "warning",
					},
					{
						ID:    "hardDrive-free",
						Label: "Hard Drive free",
						Value: fmt.Sprintf("%d MB", convertByteToMB(int(hardDriveInfo.Free))),
						Color: "success",
					},
				},
			},
		},
	}
	return data, nil
}

func convertByteToMB(input int) int {
	return input / 1024 / 1024
}
