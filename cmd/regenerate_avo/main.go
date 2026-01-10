package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/fmstephe/simd_explorer/pkg/generate"
)

var (
	flagPath = flag.String("path", "", "Names a path to process, if it points to a demo file then only that file is processed, if a directory then all demo files in that directory are processed.")
)

func main() {
	flag.Parse()

	paths := getDemoPaths(*flagPath)
	for _, path := range paths {
		println(path)
		regenerateDemo(path)
	}
}

func regenerateDemo(demoFile string) {
	stubFile := strings.Replace(demoFile, "demo", "stub", 1)
	avoFile := strings.Replace(demoFile, "demo", "_generate/asm", 1)
	println(demoFile, stubFile, avoFile)

	directory, instruction, sizeClass, discriminator := extractInfoFromDemoFileName(demoFile)
	println(directory, instruction, sizeClass, discriminator)
	pkg, description := processDemoFile(demoFile)
	println(pkg, description)
	stubArgs := processStubFile(stubFile)
	println(stubArgs)
	avoArgs := processAvoFile(avoFile, instruction)
	println(avoArgs)
	renamedAvoArgs := renameAvoArgs(avoArgs, sizeClass)

	generate.GenerateDemoFiles(directory, instruction, discriminator, stubArgs, description, sizeClass, renamedAvoArgs)
}

func processDemoFile(demoFile string) (pkg, description string) {
	fset := token.NewFileSet()

	demoF, err := parser.ParseFile(fset, demoFile, nil, 0)
	if err != nil {
		panic(err)
	}

	dg := &DemoGrabber{fset: fset}
	ast.Walk(dg, demoF)

	return dg.pkg, dg.description
}

func processStubFile(stubFile string) (args string) {
	fset := token.NewFileSet()

	stubF, err := parser.ParseFile(fset, stubFile, nil, 0)
	if err != nil {
		panic(err)
	}

	ag := &ArgsGrabber{fset: fset}
	ast.Walk(ag, stubF)

	return ag.args
}

func processAvoFile(avoFile, instruction string) (args string) {
	fset := token.NewFileSet()

	avoF, err := parser.ParseFile(fset, avoFile, nil, 0)
	if err != nil {
		panic(err)
	}

	ag := &AvoGrabber{
		fset:        fset,
		instruction: instruction,
	}
	ast.Walk(ag, avoF)

	return ag.args
}

type DemoGrabber struct {
	fset          *token.FileSet
	description   string
	pkg           string
	instruction   string
	sizeClass     int
	discriminator string
}

func (g *DemoGrabber) Visit(node ast.Node) ast.Visitor {
	switch node := node.(type) {
	case *ast.FuncDecl:
		if node.Name.Name == "Description" {
			statements := node.Body.List
			for _, stmt := range statements {
				if ret, ok := stmt.(*ast.ReturnStmt); ok {
					g.description = getValFromReturn(ret, g.fset)
				}
			}
		}
	case *ast.GenDecl:
		if node.Tok == token.TYPE && len(node.Specs) == 1 {
			// We specifically expect a single type declaration per demo
			// So we assume this is the one we are looking for
			if typeSpec, ok := node.Specs[0].(*ast.TypeSpec); ok {
				g.instruction, g.sizeClass, g.discriminator = extractInfoFromTypeName(typeSpec.Name.Name)
			}
		}
	}

	return g
}

func getValFromReturn(ret *ast.ReturnStmt, fset *token.FileSet) string {
	results := ret.Results
	if len(results) != 1 {
		ast.Print(fset, ret)
		panic(fmt.Errorf("Need a single return value, found %d", len(results)))
	}

	exp, ok := results[0].(*ast.BasicLit)
	if !ok {
		ast.Print(fset, ret)
		panic(fmt.Errorf("Need a BasicLit return value, found %T", results[0]))
	}

	if exp.Kind != token.STRING {
		ast.Print(fset, ret)
		panic(fmt.Errorf("Need a String literal return value, found %s", exp.Kind))
	}

	return strings.TrimSuffix(strings.TrimPrefix(exp.Value, `"`), `"`)
}

type ArgsGrabber struct {
	fset *token.FileSet
	args string
}

// NB: We expect the file being processed to have a single function
// If this assumption stops working we will have to make this visitor fussier
func (g *ArgsGrabber) Visit(node ast.Node) ast.Visitor {
	switch node := node.(type) {
	case *ast.FuncDecl:
		g.args = formatArgs(node, g.fset)
	}

	return g
}

type AvoGrabber struct {
	fset        *token.FileSet
	instruction string
	args        string
}

func (g *AvoGrabber) Visit(node ast.Node) ast.Visitor {
	switch node := node.(type) {
	case *ast.CallExpr:
		switch call := node.Fun.(type) {
		case *ast.Ident:
			if strings.ToLower(call.Name) == strings.ToLower(g.instruction) {
				g.args = formatArgs(node, g.fset)
			}
		}
	}

	return g
}

// Captures args from both a functiona call and a function declaration
var argsCapture = regexp.MustCompile(`(?:func )?\w+\((.*)\)`)

// This feels like a bit of a hack - but we can't directly print _just_ the args using printer.Fprint(...)
// So we print the function declaration and then use a regex to extract the args.
func formatArgs(funcDecl ast.Node, fset *token.FileSet) string {
	buff := &bytes.Buffer{}
	printer.Fprint(buff, fset, funcDecl)
	matches := argsCapture.FindStringSubmatch(buff.String())
	if len(matches) != 2 {
		panic(fmt.Errorf("Regex failed to capture function args for: %q", buff.String()))
	}
	return matches[1]
}

var instructionCaptureTypeName = regexp.MustCompile(`([A-Z]+)(\d\d\d)(.*)`)

func extractInfoFromTypeName(typeName string) (instruction string, sizeClass int, discriminator string) {
	matches := instructionCaptureTypeName.FindStringSubmatch(typeName)
	if len(matches) != 4 {
		panic(fmt.Errorf("Regex failed to capture instruction for: %q %d", typeName, len(matches)))
	}
	instruction = matches[1]

	sizeClass, err := strconv.Atoi(matches[2])
	if err != nil {
		panic(err)
	}

	discriminator = matches[3]

	return instruction, sizeClass, discriminator
}

var instructionCaptureDemoFileName = regexp.MustCompile(`^demo_([^_]+)_(\d+)(?:_(.+))?\.go$`)

func extractInfoFromDemoFileName(filePath string) (directory, instruction string, sizeClass int, discriminator string) {
	dir := filepath.Dir(filePath)
	if dir == "." {
		dir = ""
	}
	name := filepath.Base(filePath)
	matches := instructionCaptureDemoFileName.FindStringSubmatch(name)
	if len(matches) != 4 {
		panic(fmt.Errorf("Regex failed to capture info for: %q %d", filePath, len(matches)))
	}
	directory = dir

	instruction = matches[1]

	sizeClass, err := strconv.Atoi(matches[2])
	if err != nil {
		panic(err)
	}

	discriminator = matches[3]

	return directory, instruction, sizeClass, discriminator
}

func getDemoPaths(demoPath string) []string {
	if demoPath == "" {
		panic(fmt.Errorf("Empty demo path provided"))
	}
	// Strip out any unnecessary parts of path
	demoPath = filepath.Clean(demoPath)

	fi, err := os.Stat(demoPath)
	if err != nil {
		panic(err)
	}

	// If path points to a single file - then return just that file path
	if !fi.IsDir() {
		name := filepath.Base(demoPath)
		if !instructionCaptureDemoFileName.MatchString(name) {
			panic(fmt.Errorf("bad demo file path %q (%q)", demoPath, name))
		}
		return []string{demoPath}
	}

	entries, err := os.ReadDir(demoPath)
	if err != nil {
		panic(err)
	}

	paths := make([]string, 0, len(entries))
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if instructionCaptureDemoFileName.MatchString(name) {
			paths = append(paths, filepath.Join(demoPath, name))
		}
	}

	sort.Strings(paths)
	return paths
}
