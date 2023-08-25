/*******************************************************************************
 * Copyright (c) 2023 Contributors to the Eclipse Foundation
 *
 * See the NOTICE file(s) distributed with this work for additional
 * information regarding copyright ownership.
 *
 * This program and the accompanying materials are made available under the
 * terms of the Apache License, Version 2.0 which is available at
 * https://www.apache.org/licenses/LICENSE-2.0.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 ******************************************************************************/

package governance

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/eclipse-tractusx/tractusx-quality-checks/pkg/tractusx"
)

type CopyrightHeaderCheck struct {
	baseDir string
}

func NewCopyrightHeaderCheck(baseDir string) tractusx.QualityGuideline {
	return &CopyrightHeaderCheck{baseDir}
}

func (c *CopyrightHeaderCheck) Name() string {
	return "TRG 7.02 - License and Copyright header"
}

func (c *CopyrightHeaderCheck) Description() string {
	return "Where possible, all source code should contain appropriate copyright and license notices as well as information on each contribution. "
}

func (c *CopyrightHeaderCheck) ExternalDescription() string {
	return "https://eclipse-tractusx.github.io/docs/release/trg-7/trg-7-02"
}

func (c *CopyrightHeaderCheck) IsOptional() bool {
	return false
}

func (c *CopyrightHeaderCheck) Test() *tractusx.QualityResult {
	var filesNoHeader []string
	err := filepath.Walk(c.baseDir, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() && isAcceptedFileType(info.Name()) {
			if !hasValidCopywrightHeader(path) {
				filesNoHeader = append(filesNoHeader, path)
			}
		}
		return nil
	})

	if err != nil {
		return &tractusx.QualityResult{ErrorDescription: fmt.Sprintf("Can't files from %q: %v", c.baseDir, err)}
	}
	if len(filesNoHeader) > 0 {
		return &tractusx.QualityResult{ErrorDescription: fmt.Sprintf("Can't find copyright headers at:\n\t%s", strings.Join(filesNoHeader, "\n\t"))}
	}
	return &tractusx.QualityResult{Passed: true}
}

func isAcceptedFileType(filename string) bool {
	var validTypes = []string{".java", ".py", ".yaml", ".yml", ".js", ".sh", ".xml", ".sql", ".ts", ".cs", ".tsx", ".kt", ".tractusx", ".css", ".scss"}
	for _, v := range validTypes {
		if strings.HasSuffix(filename, v) {
			return true
		}
	}
	return strings.Contains(filename, "Dockerfile")
}

func hasValidCopywrightHeader(filepath string) bool {
	copyrightHeaderPart1 := "Apache-2.0"
	copyrightHeaderPart2 := "Contributors to the Eclipse Foundation"

	content, err := os.ReadFile(filepath)
	if err != nil {
		return false
	}
	s := string(content)
	if strings.Contains(s, copyrightHeaderPart1) && strings.Contains(s, copyrightHeaderPart2) {
		return true
	}
	return false
}
