package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Plugin struct {
	Description string              `json:"description"`
	Elements    map[string]*Element `json:"elements"`
	Filename    string              `json:"filename"`
	// License     string                `json:"license"`
	OtherTypes map[string]*OtherType `json:"other-types"`
	// Package     string                `json:"package"`
	Source string `json:"source"`
	// Url         string                `json:"url"`
}

type Element struct {
	// Author       string                  `json:"author"`
	Description string   `json:"description"`
	Hierarchy   []string `json:"hierarchy"`
	Interfaces  []string `json:"interfaces"`
	// Klass        string                  `json:"klass"`
	// LongName     string                  `json:"long-name"`
	// PadTemplates map[string]*PadTemplate `json:"pad-templates"`
	Properties map[string]*Property `json:"properties"`
	// Rank       string               `json:"rank"`
	Signals map[string]*Signal `json:"signals"`
}

// type PadTemplate struct {
// 	Caps      string `json:"caps"`
// 	Direction string `json:"direction"`
// 	Presence  string `json:"presence"`
// }

type Property struct {
	Blurb string `json:"blurb"`
	// ConditionallyAvailable bool   `json:"conditionally-available"`
	// Construct              bool   `json:"construct"`
	ConstructOnly bool `json:"construct-only"`
	// Controllable           bool   `json:"controllable"`
	Default string `json:"default"`
	Max     string `json:"max"`
	Min     string `json:"min"`
	// Mutable                string `json:"mutable"`
	Readable bool   `json:"readable"`
	Type     string `json:"type"`
	Writable bool   `json:"writable"`
}

type Signal struct {
	Args       []*SignalArg `json:"args"`
	ReturnType string       `json:"return-type"`
	When       string       `json:"when"`
}

type SignalArg struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type OtherType struct {
	Hierarchy         []string             `json:"hierarchy"`
	Interfaces        []string             `json:"interfaces"`
	IgnoreEnumMembers bool                 `json:"ignore-enum-members"`
	Kind              string               `json:"kind"`
	Properties        map[string]*Property `json:"properties"`
	Signals           map[string]*Signal   `json:"signals"`
	Values            []*Value             `json:"values"`
}

type Value struct {
	Desc  string `json:"desc"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func loadPlugins() (map[string]*Plugin, error) {
	// pluginFiles, err := findPlugins()
	// if err != nil {
	// 	return nil, err
	// }

	pluginFiles := []string{
		"../gstreamer/subprojects/gstreamer/docs/plugins/gst_plugins_cache.json",
		"../gstreamer/subprojects/gst-plugins-base/docs/plugins/gst_plugins_cache.json",
		"../gstreamer/subprojects/gst-plugins-good/docs/gst_plugins_cache.json",
		"../gstreamer/subprojects/gst-plugins-bad/docs/plugins/gst_plugins_cache.json",
		// "../gstreamer/subprojects/gst-plugins-ugly/docs/gst_plugins_cache.json",
		// "../gstreamer/subprojects/gst-libav/docs/gst_plugins_cache.json",
		// "../gstreamer/subprojects/gstreamer-vaapi/docs/gst_plugins_cache.json",
		// "../gstreamer/subprojects/gst-rtsp-server/docs/plugins/gst_plugins_cache.json",
	}

	plugins := make(map[string]*Plugin)
	for _, pluginFile := range pluginFiles {
		if pluginFile == "../gstreamer/subprojects/gst-editing-services/docs/gst_plugins_cache.json" {
			// deprecated
			continue
		}

		fmt.Println(pluginFile)

		b, err := os.ReadFile(pluginFile)
		if err != nil {
			return nil, err
		}

		if err = json.Unmarshal(b, &plugins); err != nil {
			return nil, err
		}
	}

	return plugins, nil
}

func findPlugins() ([]string, error) {
	cmd := exec.Command("find", "../gstreamer/subprojects", "-name", "gst_plugins_cache.json")
	b := bytes.NewBuffer(nil)
	cmd.Stdout = b
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return strings.Split(strings.TrimSpace(b.String()), "\n"), nil
}
