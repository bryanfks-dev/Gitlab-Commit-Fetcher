package sharedview

import (
	"fmt"

	"github.com/alexeyco/simpletable"
)

func (v *sharedView) NewRegisteredProjectTable() string {
	registeredProjects := v.getRegisteredProjects()

	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignLeft, Text: "Name"},
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignLeft, Text: "Project URL"},
		},
	}
	for index, project := range registeredProjects {
		cell := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", index+1)},
			{Align: simpletable.AlignLeft, Text: project.Name},
			{
				Align: simpletable.AlignRight,
				Text:  fmt.Sprintf("%d", project.ID),
			},
			{Align: simpletable.AlignLeft, Text: project.WebURL},
		}
		table.Body.Cells = append(table.Body.Cells, cell)
	}
	tableStr := table.String()
	tableStr += fmt.Sprintf("\nTotal: %d", len(registeredProjects))
	return tableStr
}
