package lib

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

/*
	Given a source cpl file, generate a BOM and a result cpl file
*/
func GenerateOutputs(src, bom, cpl string, library *Library) {
	fp, err := os.Open(src)
	if err != nil {
		return
	}
	defer fp.Close()

	fpb, err := os.Create(bom)
	if err != nil {
		return
	}
	defer fpb.Close()

	fpc, err := os.Create(bom)
	if err != nil {
		return
	}
	defer fpc.Close()

	unpack := func(s []string) (string, string, string, string, string, string, string) {
		return s[0], s[1], s[2], s[3], s[4], s[5], s[6]
	}

	/*
		Map component numbers to designators
	*/
	dmap := make(map[string][]string)
	cmap := make(map[string]*LibraryComponent)
	//read data into multi-dimentional array of strings
	reader := csv.NewReader(fp)
	for line, _ := reader.Read(); len(line) > 0; line, _ = reader.Read() {
		// G***,LOGO,F4Silkscreen,57.2,-56.9,0.0,top
		// C1,10u,C_0805_2012Metric,30.0,-42.9,90.0,top
		designator, comment, footprint, x, y, rotation, layer := unpack(line)

		/*
			First, find the component for this designator
		*/
		component := library.FindMatching(comment, footprint)
		if component == nil {
			continue
		}

		/*
			Then, add it to the designator map
		*/
		if _, ok := dmap[component.ID]; !ok {
			dmap[component.ID] = []string{}
			cmap[component.ID] = component
		}
		dmap[component.ID] = append(dmap[component.ID], designator)

		/*
			Write the component to the position file, since we're keeping it
		*/
		fmt.Fprintf(fpc, "%s,%s,%s,%s,%s", designator, x, y, layer, rotation)

		line = []string{}
	}

	for ID, designators := range dmap {
		component := cmap[ID]
		designator := strings.Join(designators, ",")

		fmt.Fprintf(fpb, "%s,%s,%s,%s", component.Part, designator, component.Package, ID)
	}
}
