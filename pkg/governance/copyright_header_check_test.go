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
	"github.com/eclipse-tractusx/tractusx-quality-checks/internal/filesystem"
	"path"
	"testing"
)

func TestShouldPassIfFileContainsCopyrightHeader(t *testing.T) {
	listOfFilesWithCopyrightHeader := []string{"test1.java",
		"test2.py",
		"test3.yaml",
	}

	dir := t.TempDir()
	for _, file := range listOfFilesWithCopyrightHeader {
		err := filesystem.CopyFile(path.Join(dir, file), path.Join("test", file))
		if err != nil {
			t.Errorf(fmt.Sprintf("Unable to copy test file %v: %v", file, err))
		}
	}
	result := NewCopyrightHeaderCheck(dir).Test()
	if !result.Passed {
		t.Errorf("Test should pass, test files contain copyright headers.")
	}
}
