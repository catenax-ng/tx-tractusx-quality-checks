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

import "github.com/eclipse-tractusx/tractusx-quality-checks/pkg/tractusx"

type CopyrightHeaderCheck struct {
	baseDir string
}

func NewCopyrightHeaderCheck(baseDir string) tractusx.QualityGuideline {
	return &CopyrightHeaderCheck{baseDir}
}

func (r *CopyrightHeaderCheck) Name() string {
	return "TRG 7.02 - License and Copyright header"
}

func (r *CopyrightHeaderCheck) Description() string {
	return "Where possible, all source code should contain appropriate copyright and license notices as well as information on each contribution. "
}

func (r *CopyrightHeaderCheck) ExternalDescription() string {
	return "https://eclipse-tractusx.github.io/docs/release/trg-7/trg-7-02"
}

func (r *CopyrightHeaderCheck) IsOptional() bool {
	return false
}

func (r *CopyrightHeaderCheck) Test() *tractusx.QualityResult {

}
