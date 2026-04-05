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
	if err != nil {
		return DashboardData{}, err
	}

	cpuInfo, err := cpu.Info()
	if err != nil {
		return DashboardData{}, err
	}

	hardDriveInfo, err := disk.Usage("/")
	if err != nil {
		return DashboardData{}, err
	}

	data := DashboardData{
		Title: "Systemressources",
		Sections: []MetricSection{
			{
				Name: "RAM",
				Cards: []MetricData{
					{
						ID:    "ram-total",
						Label: "Total RAM",
						Value: fmt.Sprintf("%d MB", convertByteToMB(int(vmem.Total))),
						Color: "accent",
					},
					{
						ID:    "ram-available",
						Label: "RAM available",
						Value: fmt.Sprintf("%d MB", convertByteToMB(int(vmem.Available))),
						Color: "success",
					},
					{
						ID:    "ram-used",
						Label: "RAM used",
						Value: fmt.Sprintf("%d MB", convertByteToMB(int(vmem.Used))),
						Color: "warning",
					},
					{
						ID:    "ram-used-percent",
						Label: "RAM used percent",
						Value: fmt.Sprintf("%.2f %%", vmem.UsedPercent),
						Color: "warning",
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
						Color: "accent",
					},
					{
						ID:    "cpu-mhz",
						Label: "CPU mhz",
						Value: fmt.Sprintf("%v", cpuInfo[0].Mhz),
						Color: "accent",
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
						ID:    "hardDrive-free",
						Label: "Hard Drive available",
						Value: fmt.Sprintf("%d MB", convertByteToMB(int(hardDriveInfo.Free))),
						Color: "success",
					},
					{
						ID:    "hardDrive-used",
						Label: "Hard Drive used",
						Value: fmt.Sprintf("%d MB", convertByteToMB(int(hardDriveInfo.Used))),
						Color: "warning",
					},
					{
						ID:    "hardDrive-used-percent",
						Label: "Hard Drive percent used",
						Value: fmt.Sprintf("%2.f%%", hardDriveInfo.UsedPercent),
						Color: "warning",
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
