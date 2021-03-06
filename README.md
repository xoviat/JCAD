# jcad

This project aims to simplify and clarify the process of ordering SMT assembly boards from JLCPCB.
The problem is as such: given a KiCAD .sch and .kicad_pcb file, what is the simpliest and fastest way
to order SMT assembled boards from JLCPCB? The current process is as follows:

1. Generate a BOM in EEschema
2. Generate a POS file in PCBnew
3. Generate Gerbers in PCBnew
4. Run [tooling](https://github.com/wokwi/kicad-jlcpcb-bom-plugin) to transform the BOM and pos file
5. Remove thorugh-hole components from the BOM
6. Upload the files to JLCPCB and order the SMT board
7. Order additional through-hole components from a distributor, and solder these components yourself

## What can't change

1. Though-hole components and components not available from JLCPCB must be be ordered and soldered yourself.
2. As of now, we can't access the JLCPCB API, so we still have to click through their wizard

## What can change

1. The input files to the JLCPCB website can be directly derived from the .kicad_pcb file
2. The components can be assigned interactively, before uploading to JLCPCB.
3. Calculations can be done offline, using the JLCPCB component database.
4. Footprints can be assigned offline, using the JLCPCB component database, directly editing the .sch file.

## What our envisionsed workflow looks like

1. After creating a .kicad_pcb file, directly generate the required outputs for JLCPCB and derive the component parts from
   the designator, commment, and footprint. 
2. Directly generate an additional BOM that is uploadable to Mouser for the additional though-hole parts required.

## Why this is possible

1. The .kicad_pcb file can be read using the pcbnew Python API. We can generate the required outputs and postprocess them ourselves.
2. The BOM and POS files are simple to read and should not pose too many difficulties.

## Libraries

1. I golang because it's performant, compiles quickly, and executes natively.
2. This applicadtion will require KiCad to be installed to access the pcbnew api. It will use python scripts to access the pcbnew API.

## Issues

Feel free to give ideas in the issues. This project solves a real problem, and therefore it might (hopefully) become popular.
