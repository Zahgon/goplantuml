package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	goplantuml "github.com/jfeliu007/goplantuml/parser"
)

// RenderingOptionSlice will implements the sort interface
type RenderingOptionSlice []goplantuml.RenderingOption

// Len is the number of elements in the collection.
func (as RenderingOptionSlice) Len() int {
	_ = "STUB: not implemented"

	// Less reports whether the element with
	// index i should sort before the element with index j.
	return 0
}

func (as RenderingOptionSlice) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Swap swaps the elements with indexes i and j.
func (as RenderingOptionSlice) Swap(i, j int) { _ = "STUB: not implemented"; return }

func main() {
	recursive := flag.Bool("recursive", false, "walk all directories recursively")
	ignore := flag.String("ignore", "", "comma separated list of folders to ignore")
	maxDepth := flag.Int("max-depth", 0, "maximum nesting depth for packages (0 = unlimited)")
	showAggregations := flag.Bool("show-aggregations", false, "renders public aggregations even when -hide-connections is used (do not render by default)")
	hideFields := flag.Bool("hide-fields", false, "hides fields")
	hideMethods := flag.Bool("hide-methods", false, "hides methods")
	hideConnections := flag.Bool("hide-connections", false, "hides all connections in the diagram")
	showCompositions := flag.Bool("show-compositions", false, "Shows compositions even when -hide-connections is used")
	showImplementations := flag.Bool("show-implementations", false, "Shows implementations even when -hide-connections is used")
	showAliases := flag.Bool("show-aliases", false, "Shows aliases even when -hide-connections is used")
	showConnectionLabels := flag.Bool("show-connection-labels", false, "Shows labels in the connections to identify the connections types (e.g. extends, implements, aggregates, alias of")
	title := flag.String("title", "", "Title of the generated diagram")
	notes := flag.String("notes", "", "Comma separated list of notes to be added to the diagram")
	output := flag.String("output", "", "output file path. If omitted, then this will default to standard output")
	showOptionsAsNote := flag.Bool("show-options-as-note", false, "Show a note in the diagram with the none evident options ran with this CLI")
	aggregatePrivateMembers := flag.Bool("aggregate-private-members", false, "Show aggregations for private members. Ignored if -show-aggregations is not used.")
	hidePrivateMembers := flag.Bool("hide-private-members", false, "Hide private fields and methods")
	flag.Parse()
	renderingOptions := map[goplantuml.RenderingOption]interface{}{
		goplantuml.RenderConnectionLabels:  *showConnectionLabels,
		goplantuml.RenderFields:            !*hideFields,
		goplantuml.RenderMethods:           !*hideMethods,
		goplantuml.RenderAggregations:      *showAggregations,
		goplantuml.RenderTitle:             *title,
		goplantuml.AggregatePrivateMembers: *aggregatePrivateMembers,
		goplantuml.RenderPrivateMembers:    !*hidePrivateMembers,
	}
	if *hideConnections {
		renderingOptions[goplantuml.RenderAliases] = *showAliases
		renderingOptions[goplantuml.RenderCompositions] = *showCompositions
		renderingOptions[goplantuml.RenderImplementations] = *showImplementations

	}
	noteList := []string{}
	if *showOptionsAsNote {
		legend, err := getLegend(renderingOptions)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		noteList = append(noteList, legend)
	}
	if *notes != "" {
		noteList = append(noteList, "", "<b><u>Notes</u></b>")
	}
	split := strings.Split(*notes, ",")
	for _, note := range split {
		trimmed := strings.TrimSpace(note)
		if trimmed != "" {
			noteList = append(noteList, trimmed)
		}
	}
	renderingOptions[goplantuml.RenderNotes] = strings.Join(noteList, "\n")
	dirs, err := getDirectories()

	if err != nil {
		fmt.Println("usage:\ngoplantuml <DIR>\nDIR Must be a valid directory")
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	ignoredDirectories, err := getIgnoredDirectories(*ignore)
	if err != nil {

		fmt.Println("usage:\ngoplantuml [-ignore=<DIRLIST>]\nDIRLIST Must be a valid comma separated list of existing directories")
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	result, err := goplantuml.NewClassDiagramWithMaxDepth(dirs, ignoredDirectories, *recursive, *maxDepth)
	result.SetRenderingOptions(renderingOptions)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	rendered := result.Render()
	var writer io.Writer
	if *output != "" {
		writer, err = os.Create(*output)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	} else {
		writer = os.Stdout
	}
	fmt.Fprint(writer, rendered)
}

func getDirectories() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func getIgnoredDirectories(list string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getLegend(ro map[goplantuml.RenderingOption]interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
