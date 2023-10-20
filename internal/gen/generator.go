package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Generator struct {
	plugin     *Plugin
	header     *bytes.Buffer
	output     *bytes.Buffer
	elements   int
	otherTypes int
	imports    map[string]bool
}

func generate(plugins map[string]*Plugin) error {
	for _, plugin := range plugins {
		if plugin.Filename == "gstlv2" {
			continue
		}

		g := &Generator{
			plugin: plugin,
			header: bytes.NewBuffer(nil),
			output: bytes.NewBuffer(nil),
			imports: map[string]bool{
				"gst": true,
			},
		}
		if err := g.generatePlugin(); err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) generatePlugin() error {
	for name, element := range g.plugin.Elements {
		element.Hierarchy[0] = toCamelCase(element.Hierarchy[0])
		if existingElements[element.Hierarchy[0]] == "" {
			g.generateElement(name, element)
			g.elements++
		}
	}

	for name, other := range g.plugin.OtherTypes {
		g.generateOther(toCamelCase(name), other)
		g.otherTypes++
	}

	if g.elements+g.otherTypes == 0 {
		return nil
	}

	g.generateHeader()
	filename := fmt.Sprintf("../../factory/%s.go", g.plugin.Filename)
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	if _, err = f.Write(g.header.Bytes()); err != nil {
		return err
	}
	if _, err = f.Write(g.output.Bytes()); err != nil {
		return err
	}

	return exec.Command("gofmt", "-w", filename).Run()
}

func (g *Generator) generateHeader() {
	g.H(`// source: `, g.plugin.Source)
	g.H()

	// package
	g.H(`package factory`)
	g.H()

	// imports
	g.H(`import (`)
	if g.imports["fmt"] {
		g.H(`	"fmt"`)
		g.H()
	}
	g.H(`	"github.com/go-gst/go-gst/gst"`)
	if g.imports["app"] {
		g.H(`	"github.com/go-gst/go-gst/gst/app"`)
	}
	if g.imports["base"] {
		g.H(`	"github.com/go-gst/go-gst/gst/base"`)
	}
	g.H(`)`)
	g.H()
}

func (g *Generator) generateElement(name string, element *Element) {
	g.generateDescription(element.Description)
	g.generateStructure(element.Hierarchy)

	constructOnly := false
	for _, property := range element.Properties {
		if property.ConstructOnly {
			constructOnly = true
		}
	}

	noPrefix := element.Hierarchy[0]
	if strings.HasPrefix(noPrefix, "Gst") {
		noPrefix = toCamelCase(noPrefix[3:])
	}
	for _, withName := range []bool{false, true} {
		if constructOnly {
			g.P(`func New`, noPrefix, `WithProperties(properties map[string]interface{}) (*`, element.Hierarchy[0], `, error) {`)
			g.P(`	e, err := gst.NewElementWithProperties("`, name, `", properties)`)
		} else if withName {
			g.P(`func New`, noPrefix, `WithName(name string) (*`, element.Hierarchy[0], `, error) {`)
			g.P(`	e, err := gst.NewElementWithName("`, name, `", name)`)
		} else {
			g.P(`func New`, noPrefix, `() (*`, element.Hierarchy[0], `, error) {`)
			g.P(`	e, err := gst.NewElement("`, name, `")`)
		}
		g.P(`	if err != nil {`)
		g.P(`		return nil, err`)
		g.P(`	}`)
		g.W(`	return &`)
		for i, h := range element.Hierarchy {
			if h == "GstElement" {
				g.W(`Element: e`)
				for j := 0; j < i; j++ {
					g.W(`}`)
				}
				g.P(`, nil`)
				break
			} else if e := existingElements[h]; e != "" {
				s := strings.Split(e, ".")
				g.imports[s[0]] = true
				g.W(s[1], `: &`, e, `{`)
			} else if i > 0 {
				g.W(h, `: &`, h, `{`)
			} else {
				g.W(h, `{`)
			}
		}
		g.P(`}`)
		g.P()
		if constructOnly {
			break
		}
	}

	for propertyName, property := range element.Properties {
		g.generateProperty(element.Hierarchy[0], propertyName, property)
	}

	for signalName, signal := range element.Signals {
		g.generateSignal(signalName, signal)
	}
}

func (g *Generator) generateStructure(hierarchy []string) {
	g.P(`type `, hierarchy[0], ` struct {`)
	parentElement := hierarchy[1]
	if e := existingElements[parentElement]; e != "" {
		parentElement = e
		g.imports[strings.Split(e, ".")[0]] = true
	}
	g.P(`	*`, parentElement)
	g.P(`}`)
	g.P()
}

func (g *Generator) generateProperty(elementName, propertyName string, prop *Property) {
	if !prop.Readable && !prop.Writable {
		return
	}

	g.generateDescription(prop.Blurb)
	if prop.Default != "" {
		g.W(`// Default: `, prop.Default)
		if prop.Min != "" {
			g.W(`, Min: `, prop.Min)
		}
		if prop.Max != "" {
			g.W(`, Max: `, prop.Max)
		}
	}
	g.P()

	goType := g.toGoType(prop.Type, g.plugin.OtherTypes)
	other := g.plugin.OtherTypes[goType]

	if prop.Writable {
		g.P(`func (e *`, elementName, `) Set`, toCamelCase(propertyName), `(value `, goType, `) error {`)
		if other != nil {
			if other.Kind == "enum" {
				g.P(`	e.Element.SetArg("`, propertyName, `", string(value))`)
			} else if other.Kind == "flags" {
				g.P(`	e.Element.SetArg("`, propertyName, `", fmt.Sprint(value))`)
			} else {
				// TODO
				fmt.Println("Set type", goType)
			}
			g.P(`	return nil`)
		} else {
			g.P(`	return e.Element.SetProperty("`, propertyName, `", value)`)
		}
		g.P(`}`)
		g.P()
	}

	if prop.Readable {
		g.P(`func (e *`, elementName, `) Get`, toCamelCase(propertyName), `() (`, goType, `, error) {`)
		if goType == "interface{}" {
			g.P(`	return e.Element.GetProperty("`, propertyName, `")`)
		} else {
			g.P(`	var value `, goType)
			g.P(`	var ok bool`)
			g.P(`	v, err := e.Element.GetProperty("`, propertyName, `")`)
			g.P(`	if err != nil {`)
			g.P(`		return value, err`)
			g.P(`	}`)
			g.P(`	value, ok = v.(`, goType, `)`)
			g.P(`	if !ok {`)
			g.imports["fmt"] = true
			g.P(`		return value, fmt.Errorf("could not cast value to `, goType, `")`)
			g.P(`	}`)
			g.P(`	return value, nil`)
		}
		g.P(`}`)
		g.P()
	}
}

func (g *Generator) generateSignal(name string, signal *Signal) {
	fmt.Printf("%s: %+v\n", name, signal)
}

func (g *Generator) generateOther(name string, other *OtherType) {
	switch other.Kind {
	case "enum":
		if other.IgnoreEnumMembers {
			return
		}
		g.P(`type `, name, ` string`)
		g.P()
		g.P(`const (`)
		for _, value := range other.Values {
			valueName := value.Name
			if strings.HasPrefix(value.Name, "-") {
				valueName = "Minus" + valueName[1:]
			}
			valueName = name + toCamelCase(valueName)
			g.P(`	`, valueName, ` `, name, ` = "`, value.Name, `" // `, value.Desc)
		}
		g.P(`)`)
		g.P()

	case "flags":
		g.P(`type `, name, ` int`)
		g.P()
		g.P(`const (`)
		for _, value := range other.Values {
			valueName := value.Name
			if strings.HasPrefix(value.Name, "-") {
				valueName = "Minus" + valueName[1:]
			}
			valueName = name + toCamelCase(valueName)
			g.P(`	`, valueName, ` `, name, ` = `, value.Value, ` // `, value.Desc)
		}
		g.P(`)`)
		g.P()

	case "object":
		other.Hierarchy[0] = toCamelCase(other.Hierarchy[0])
		for _, element := range g.plugin.Elements {
			if element.Hierarchy[0] == other.Hierarchy[0] {
				return
			}
		}

		for _, h := range other.Hierarchy {
			if h == "GstElement" {
				g.generateStructure(other.Hierarchy)

				for propertyName, property := range other.Properties {
					g.generateProperty(other.Hierarchy[0], propertyName, property)
				}

				for signalName, signal := range other.Signals {
					g.generateSignal(signalName, signal)
				}
			}
		}
	}
}

func (g *Generator) generateDescription(description string) {
	if len(description) > 0 {
		desc := strings.Split(description, "\n")
		for _, d := range desc {
			g.P(`// `, d)
		}
	}
}

func (g *Generator) W(args ...string) {
	for _, v := range args {
		g.output.WriteString(v)
	}
}

func (g *Generator) P(args ...string) {
	p(g.output, args...)
}

func (g *Generator) H(args ...string) {
	p(g.header, args...)
}

func p(out *bytes.Buffer, args ...string) {
	for _, v := range args {
		out.WriteString(v)
	}
	out.WriteByte('\n')
}

func (g *Generator) toGoType(gType string, otherTypes map[string]*OtherType) string {
	switch gType {
	case "gboolean":
		return "bool"
	case "gint":
		return "int"
	case "guint":
		return "uint"
	case "gint8":
		return "int8"
	case "guint8":
		return "uint8"
	case "gint16":
		return "int16"
	case "guint16":
		return "uint16"
	case "gint32":
		return "int32"
	case "guint32":
		return "uint32"
	case "gint64":
		return "int64"
	case "guint64":
		return "uint64"
	case "gfloat":
		return "float32"
	case "gdouble":
		return "float64"
	case "gchararray":
		return "string"
	default:
		if e := existingElements[gType]; e != "" {
			g.imports[strings.Split(e, ".")[0]] = true
			return "*" + e
		}
		if otherTypes[gType] != nil {
			return toCamelCase(gType)
		}
		return "interface{}"
	}
}

func toCamelCase(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	n := strings.Builder{}
	n.Grow(len(s))
	capNext := true
	for i, v := range []byte(s) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if capNext {
			if vIsLow {
				v += 'A'
				v -= 'a'
			}
		} else if i == 0 {
			if vIsCap {
				v += 'a'
				v -= 'A'
			}
		}

		if vIsCap || vIsLow {
			n.WriteByte(v)
		} else if vIsNum := v >= '0' && v <= '9'; vIsNum {
			n.WriteByte(v)
		}
		capNext = v == '_' || v == ' ' || v == '-' || v == '.'
	}
	return n.String()
}
