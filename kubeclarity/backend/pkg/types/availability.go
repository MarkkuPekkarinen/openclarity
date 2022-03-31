// Copyright © 2022 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
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

package types

type Availability string

const (

	// AvailabilityNONE captures enum value "NONE".
	AvailabilityNONE Availability = "NONE"

	// AvailabilityLOW captures enum value "LOW".
	AvailabilityLOW Availability = "LOW"

	// AvailabilityHIGH captures enum value "HIGH".
	AvailabilityHIGH Availability = "HIGH"
)
