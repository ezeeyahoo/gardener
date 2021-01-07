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

package v20_test

import (
	"testing"

	v20 "github.com/gardener/gardener/pkg/operation/seed/scheduler/v20"
	"k8s.io/kube-scheduler/v20/v1beta1"

	"k8s.io/utils/pointer"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConfigurator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Configurator component Suite")
}

var _ = Describe("NewConfigurator", func() {
	It("should not return nil", func() {
		c, err := v20.NewConfigurator("test", &v1beta1.KubeSchedulerConfiguration{})

		Expect(err).NotTo(HaveOccurred())
		Expect(c).NotTo(BeNil())
	})
})

var _ = Describe("Config", func() {
	var output, sha256 string

	JustBeforeEach(func() {
		c, err := v20.NewConfigurator("test", &v1beta1.KubeSchedulerConfiguration{
			Profiles: []v1beta1.KubeSchedulerProfile{
				{
					SchedulerName: pointer.StringPtr("test"),
				},
			},
		})

		Expect(err).NotTo(HaveOccurred())
		Expect(c).NotTo(BeNil())

		output, sha256, err = c.Config()
		Expect(err).NotTo(HaveOccurred())
	})

	It("returns correct config", func() {
		Expect(output).To(Equal(`apiVersion: kubescheduler.config.k8s.io/v1beta1
clientConnection:
  acceptContentTypes: ""
  burst: 0
  contentType: ""
  kubeconfig: ""
  qps: 0
kind: KubeSchedulerConfiguration
leaderElection:
  leaderElect: true
  leaseDuration: 15s
  renewDeadline: 10s
  resourceLock: leases
  resourceName: kube-scheduler
  resourceNamespace: test
  retryPeriod: 2s
profiles:
- schedulerName: test
`))
	})

	It("returns correct hash", func() {
		Expect(sha256).To(Equal("e3a9d797e9b5a7044504026878aac1105a641c97efd6a49aad65db1b070f78a6"))
	})
})