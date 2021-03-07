// Copyright 2021 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package embedded defines the internal templates that Hugo provides.
package embedded

import (
	"embed"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const templateFolder = "templates"

//go:embed templates/*
var templatesFS embed.FS

// LoadTemplates loads all templates that have been embedded using the embed directive.
func LoadTemplates() ([][2]string, error) {
	var embeddedTemplates [][2]string
	err := fs.WalkDir(templatesFS, templateFolder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		// This isn't needed necessarily as skipping dotfiles is an implementation
		// detail of the embed directive.
		if strings.HasPrefix(d.Name(), ".") {
			return nil
		}
		templateName := filepath.ToSlash(strings.TrimPrefix(path, templateFolder+string(os.PathSeparator)))
		templateContent, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		embeddedTemplates = append(embeddedTemplates, [2]string{templateName, string(templateContent)})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return embeddedTemplates, nil
}
