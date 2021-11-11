package track

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

// View 追跡状況を表示
func (t *Track) View() error {
	data := t.CreateTableData()

	if len(data) == 0 {
		return fmt.Errorf("tracking number or shipping carrier is incorrect")
	}

	ShowTable(&data)

	return nil
}

// ShowTable テーブルを表示
func ShowTable(data *[][]string) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"運送業者", "追跡番号", "メモ", "配達状況", "日時", "営業所"})
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetAutoMergeCells(true)
	table.AppendBulk(*data)
	table.Render()
}
