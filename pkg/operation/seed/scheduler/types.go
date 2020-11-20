// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package scheduler

import (
	"github.com/Masterminds/semver"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var (
	versionConstraintEqual1_18 *semver.Constraints
	versionConstraintEqual1_19 *semver.Constraints
	versionConstraintEqual1_20 *semver.Constraints
)

func init() {
	var err error

	versionConstraintEqual1_18, err = semver.NewConstraint("1.18.x")
	utilruntime.Must(err)

	versionConstraintEqual1_19, err = semver.NewConstraint("1.19.x")
	utilruntime.Must(err)

	versionConstraintEqual1_20, err = semver.NewConstraint("1.20.x")
	utilruntime.Must(err)
}
