// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package generator

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/livekit/gstplugins/internal/types"
)

const license = `// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.`

type Generator struct {
	output             *bytes.Buffer
	generatedConstants map[string]bool
}

func Generate(elements map[string]*types.Element) error {
	g := &Generator{}

	for _, element := range elements {
		if err := g.generate(element); err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) generate(element *types.Element) error {
	g.output = bytes.NewBuffer(nil)
	g.addHeader(element)

	addUnderscore := false
	if strings.HasPrefix(element.Name, "_") {
		addUnderscore = true
	}
	structName := strcase.ToCamel(element.Name)
	if addUnderscore {
		structName = "_" + structName
	}

	g.generateStruct(element, structName)

	if len(element.Properties) > 0 {
		g.P(`// ----- Properties -----`)
		g.P()
		for _, property := range element.Properties {
			g.generateProperty(structName, property, element.Constants)
		}
	}

	if len(element.Constants) > 0 {
		g.P(`// ----- Constants -----`)
		g.P()
		for _, constant := range element.Constants {
			g.generateConstant(constant)
		}
	}

	filename := fmt.Sprintf("%s.go", element.Name)
	if element.IsBase {
		filename = "base/" + filename
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	_, err = f.Write(g.output.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) addHeader(element *types.Element) {
	// license
	g.P(license)

	// package
	g.P()
	if element.IsBase {
		g.P(`package base`)
	} else {
		g.P(`package gstplugins`)
	}
	g.P()

	requireFmt := false
	for _, prop := range element.Properties {
		if prop.Readable {
			isInterface := false
			if strings.HasPrefix(prop.DataType, "Gst") {
				if _, isConstant := element.Constants[prop.DataType]; !isConstant {
					isInterface = true
				}
			}
			if !isInterface {
				requireFmt = true
				break
			}
		}
	}

	// imports
	g.P(`import (`)
	if requireFmt {
		g.P(`	"fmt"`)
		g.P()
	}
	g.P(`	"github.com/tinyzimmer/go-gst/gst"`)
	if element.BaseClass != "" {
		g.P()
		g.P(`	"github.com/livekit/gstplugins/base"`)
	}
	g.P(`)`)
	g.P()
}

func (g *Generator) generateStruct(element *types.Element, structName string) {
	g.P(`type `, structName, ` struct {`)
	if element.BaseClass != "" {
		g.P(`	*base.`, element.BaseClass)
	} else {
		g.P(`	Element *gst.Element`)
	}
	g.P(`}`)
	g.P()

	if !element.IsBase {
		g.P(`func New`, structName, `() (*`, structName, `, error) {`)
		g.P(`	e, err := gst.NewElement("`, element.FactoryName, `")`)
		g.P(`	if err != nil {`)
		g.P(`		return nil, err`)
		g.P(`	}`)
		if element.BaseClass != "" {
			g.P(`	return &`, structName, `{`, element.BaseClass, `: &base.`, element.BaseClass, `{Element: e}}, nil`)
		} else {
			g.P(`	return &`, structName, `{Element: e}, nil`)
		}
		g.P(`}`)
		g.P()
		g.P(`func New`, structName, `WithName(name string) (*`, structName, `, error) {`)
		g.P(`	e, err := gst.NewElementWithName("`, element.FactoryName, `", name)`)
		g.P(`	if err != nil {`)
		g.P(`		return nil, err`)
		g.P(`	}`)
		if element.BaseClass != "" {
			g.P(`	return &`, structName, `{`, element.BaseClass, `: &base.`, element.BaseClass, `{Element: e}}, nil`)
		} else {
			g.P(`	return &`, structName, `{Element: e}, nil`)
		}
		g.P(`}`)
		g.P()
	}
}

func (g *Generator) generateProperty(structName string, prop *types.Property, constants map[string]*types.Constant) {
	goType := prop.DataType
	var isConstant bool
	if strings.HasPrefix(prop.DataType, "Gst") {
		switch prop.DataType {
		case "GstSample":
			goType = "*gst.Sample"
		case "GstCaps":
			goType = "*gst.Caps"
		default:
			if constants[prop.DataType] != nil {
				goType = prop.DataType
				isConstant = true
			} else {
				goType = "interface{}"
			}
		}
	}

	g.P(`// `, prop.Name, ` (`, prop.DataType, `)`)
	if prop.Description != "" {
		g.P(`//`)
		desc := strings.Split(prop.Description, "\n")
		for _, line := range desc {
			g.P(`// `, line)
		}
	}
	g.P()

	if prop.Readable {
		g.P(`func (e *`, structName, `) Get`, strcase.ToCamel(prop.Name), `() (`, goType, `, error) {`)
		if goType == "interface{}" {
			g.P(`	return e.Element.GetProperty("`, prop.Name, `")`)
		} else {
			g.P(`	var value `, goType)
			g.P(`	var ok bool`)
			g.P(`	v, err := e.Element.GetProperty("`, prop.Name, `")`)
			g.P(`	if err != nil {`)
			g.P(`		return value, err`)
			g.P(`	}`)
			g.P(`	value, ok = v.(`, goType, `)`)
			g.P(`	if !ok {`)
			g.P(`		return value, fmt.Errorf("could not cast value to `, goType, `")`)
			g.P(`	}`)
			g.P(`	return value, nil`)
		}
		g.P(`}`)
		g.P()
	}

	if prop.Writable {
		g.P(`func (e *`, structName, `) Set`, strcase.ToCamel(prop.Name), `(value `, goType, `) error {`)
		if isConstant {
			g.P(`	e.Element.SetArg("`, prop.Name, `", string(value))`)
			g.P(`	return nil`)
		} else {
			g.P(`	return e.Element.SetProperty("`, prop.Name, `", value)`)
		}
		g.P(`}`)
		g.P()
	}
}

func (g *Generator) generateConstant(constant *types.Constant) {
	g.P(`type `, constant.Name, ` string`)
	g.P()
	g.P(`const (`)
	existing := make(map[string]int)
	for i := 0; i < len(constant.Members); i++ {
		name := constant.Name + strcase.ToCamel(constant.Members[i])
		if existing[name] > 0 {
			name += fmt.Sprint(existing[name])
		}
		existing[name]++
		value := constant.Members[i]
		desc := constant.Descriptions[i]
		g.P(`	`, name, ` `, constant.Name, ` = "`, value, `" // `, desc)
	}
	g.P(`)`)
	g.P()
}

// P forwards to g.gen.P, which prints output.
func (g *Generator) P(args ...string) {
	for _, v := range args {
		g.output.WriteString(v)
	}
	g.output.WriteByte('\n')
}
