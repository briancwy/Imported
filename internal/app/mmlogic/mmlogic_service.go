// Copyright 2019 Google LLC
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

package mmlogic

import (
	"open-match.dev/open-match/internal/config"
	"open-match.dev/open-match/internal/pb"
	"open-match.dev/open-match/internal/statestore"
)

// The MMLogic API provides utility functions for common MMF functionality such
// as retreiving Tickets from state storage.
type mmlogicService struct {
	cfg   config.View
	store statestore.Service
}

// newMmlogic creates and initializes the mmlogic service.
func newMmlogic(cfg config.View) (*mmlogicService, error) {
	ms := &mmlogicService{
		cfg: cfg,
	}

	// Initialize the state storage interface.
	var err error
	ms.store, err = statestore.New(cfg)
	if err != nil {
		return nil, err
	}

	return ms, nil
}

// GetPoolTickets gets the list of Tickets that match every Filter in the
// specified Pool.
func (s *mmlogicService) QueryTickets(req *pb.QueryTicketsRequest, stream pb.MmLogic_QueryTicketsServer) error {
	ctx := stream.Context()

	for {
		select {
		case <-ctx.Done():
			return nil

		default:
			err := stream.Send(&pb.QueryTicketsResponse{})
			if err != nil {
				return err
			}
		}
	}
}