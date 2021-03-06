package lib

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	referred to as 'bcomponent'
*/
type BoardComponent struct {
	Designator string
	Comment    string
	Package    string
	X          string
	Y          string
	Rotation   string
	Layer      string
}

func (bc *BoardComponent) Key() []byte {
	key, _ := Marshal([]string{
		re1.ReplaceAllString(bc.Designator, ""),
		bc.Comment,
		bc.Package,
	})

	return key
}

func (bc *BoardComponent) Rotate(drotation float64) error {
	rotation, err := strconv.ParseFloat(bc.Rotation, 64)
	if err != nil {
		return fmt.Errorf("failed to parse board component rotation: %s\n", bc.Rotation)
	}

	rotation += drotation
	bc.Rotation = fmt.Sprintf("%.1f", rotation)

	return nil
}

type BOMEntry struct {
	Comment     string
	Designators []string
	Component   *LibraryComponent
}

func ReadCPL(src string) []*BoardComponent {
	fp, err := os.Open(src)
	if err != nil {
		return []*BoardComponent{}
	}
	defer fp.Close()

	components := []*BoardComponent{}
	reader := csv.NewReader(fp)
	for line, _ := reader.Read(); len(line) > 0; line, _ = reader.Read() {
		components = append(components, &BoardComponent{
			Designator: line[0],
			Comment:    line[1],
			Package:    line[2],
			X:          line[3],
			Y:          line[4],
			Rotation:   line[5],
			Layer:      line[6],
		})

		line = []string{}
	}

	return components
}

func WriteCPL(dst string, components []*BoardComponent) {
	fp, err := os.Create(dst)
	if err != nil {
		return
	}
	defer fp.Close()

	writer := csv.NewWriter(fp)
	writer.Write([]string{"Designator", "Mid X", "Mid Y", "Layer", "Rotation"})
	for _, component := range components {
		writer.Write([]string{
			component.Designator,
			component.X,
			component.Y,
			component.Layer,
			component.Rotation,
		})
	}

	writer.Flush()
}

func WriteBOM(dst string, entries []*BOMEntry) {
	fp, err := os.Create(dst)
	if err != nil {
		return
	}
	defer fp.Close()

	writer := csv.NewWriter(fp)
	writer.Write([]string{"Comment", "Designator", "Footprint", "LCSC Part #"})
	for _, entry := range entries {
		writer.Write([]string{
			entry.Comment,
			strings.Join(entry.Designators, ","),
			entry.Component.Package,
			entry.Component.CID(),
		})
	}

	writer.Flush()
}
