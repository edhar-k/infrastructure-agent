// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
package agent

import (
	"time"

	"github.com/newrelic/infrastructure-agent/pkg/backend/identityapi"
	"github.com/newrelic/infrastructure-agent/pkg/entity"
)

type MockIdentityRegisterClient struct{}

func (icc *MockIdentityRegisterClient) RegisterEntitiesRemoveMe(agentEntityID entity.ID, entities []identityapi.RegisterEntityRemoveMe) (r []identityapi.RegisterEntityResponse, retryAfter time.Duration, err error) {
	return
}
