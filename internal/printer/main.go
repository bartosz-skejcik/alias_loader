package printer

import (
	"fmt"
	"load_aliases/internal/utils"
	"os"

	"github.com/olekukonko/tablewriter"
)

func PrintAliasesTable(aliases map[string]map[string]string, withDescriptions bool) {
	if len(aliases) == 0 {
		fmt.Println("No aliases found")
		return
	}

	for _, aliasMap := range aliases {
		// title, err := utils.GetTableTitle(category)
		// if err != nil {
		// 	title = category
		// }

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader(getTableHeader(withDescriptions))
		// table.SetTitle(title)
		table.SetAutoWrapText(false)
		table.SetAutoFormatHeaders(true)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		// table.SetCenterSeparator("")
		// table.SetColumnSeparator("")
		// table.SetRowSeparator("")
		table.SetHeaderLine(true)
		table.SetBorder(true)
		table.SetTablePadding("\t")
		table.SetNoWhiteSpace(false)

		if withDescriptions {
			table.SetHeaderColor(
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
			)

			table.SetColumnColor(
				tablewriter.Colors{tablewriter.FgYellowColor},
				tablewriter.Colors{tablewriter.FgCyanColor},
				tablewriter.Colors{tablewriter.FgMagentaColor},
			)

		} else {
			table.SetHeaderColor(
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
			)

			table.SetColumnColor(
				tablewriter.Colors{tablewriter.FgYellowColor},
				tablewriter.Colors{tablewriter.FgCyanColor},
			)
		}

		for alias, cmd := range aliasMap {
			row := []string{alias, cmd}
			if withDescriptions {
				desc, err := utils.GetAliasDescription(fmt.Sprintf("%s='%s'", alias, cmd))
				if err != nil {
					desc = "No description available"
				}
				row = append(row, desc)
			}
			table.Append(row)
		}

		table.Render()
		fmt.Println() // Add a blank line between categories
	}
}

func getTableHeader(withDescriptions bool) []string {
	header := []string{"Alias", "Command"}
	if withDescriptions {
		header = append(header, "Description")
	}
	return header
}
