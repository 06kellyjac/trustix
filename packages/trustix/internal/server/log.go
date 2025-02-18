// Copyright (C) 2021 Tweag IO
//
// This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.

package server

import (
	"context"
	"fmt"

	"github.com/nix-community/trustix/packages/trustix-proto/api"
	"github.com/nix-community/trustix/packages/trustix-proto/schema"
	"github.com/nix-community/trustix/packages/trustix/interfaces"
	"github.com/nix-community/trustix/packages/trustix/internal/pool"
)

// LogAPIServer wraps kvStoreLogApi and turns it into a gRPC implementation
type LogAPIServer struct {
	api.UnimplementedLogAPIServer

	clients *pool.ClientPool

	// Keep a set of published logs around to filter on, we don't want to leak what we subscribe to
	published map[string]struct{}
}

func NewLogAPIServer(logs []*api.Log, clients *pool.ClientPool) *LogAPIServer {
	published := make(map[string]struct{})
	for _, log := range logs {
		published[*log.LogID] = struct{}{}
	}

	return &LogAPIServer{
		published: published,
		clients:   clients,
	}
}

func (s *LogAPIServer) getClient(logID string) (interfaces.LogAPI, error) {
	_, ok := s.published[logID]
	if !ok {
		return nil, fmt.Errorf("Log with ID '%s' is not published", logID)
	}

	client, err := s.clients.Get(logID)
	if err != nil {
		return nil, err
	}

	return client.LogAPI, nil
}

func (s *LogAPIServer) GetHead(ctx context.Context, req *api.LogHeadRequest) (*schema.LogHead, error) {
	log, err := s.getClient(*req.LogID)
	if err != nil {
		return nil, err
	}

	return log.GetHead(ctx, req)
}

func (s *LogAPIServer) GetLogConsistencyProof(ctx context.Context, req *api.GetLogConsistencyProofRequest) (*api.ProofResponse, error) {
	log, err := s.getClient(*req.LogID)
	if err != nil {
		return nil, err
	}
	return log.GetLogConsistencyProof(ctx, req)
}

func (s *LogAPIServer) GetLogAuditProof(ctx context.Context, req *api.GetLogAuditProofRequest) (*api.ProofResponse, error) {
	log, err := s.getClient(*req.LogID)
	if err != nil {
		return nil, err
	}
	return log.GetLogAuditProof(ctx, req)
}

func (s *LogAPIServer) GetLogEntries(ctx context.Context, req *api.GetLogEntriesRequest) (*api.LogEntriesResponse, error) {
	log, err := s.getClient(*req.LogID)
	if err != nil {
		return nil, err
	}
	return log.GetLogEntries(ctx, req)
}

func (s *LogAPIServer) GetMapValue(ctx context.Context, req *api.GetMapValueRequest) (*api.MapValueResponse, error) {
	log, err := s.getClient(*req.LogID)
	if err != nil {
		return nil, err
	}
	return log.GetMapValue(ctx, req)
}

func (s *LogAPIServer) GetMHLogConsistencyProof(ctx context.Context, req *api.GetLogConsistencyProofRequest) (*api.ProofResponse, error) {
	log, err := s.getClient(*req.LogID)
	if err != nil {
		return nil, err
	}
	return log.GetMHLogConsistencyProof(ctx, req)
}

func (s *LogAPIServer) GetMHLogAuditProof(ctx context.Context, req *api.GetLogAuditProofRequest) (*api.ProofResponse, error) {
	log, err := s.getClient(*req.LogID)
	if err != nil {
		return nil, err
	}
	return log.GetMHLogAuditProof(ctx, req)
}

func (s *LogAPIServer) GetMHLogEntries(ctx context.Context, req *api.GetLogEntriesRequest) (*api.LogEntriesResponse, error) {
	log, err := s.getClient(*req.LogID)
	if err != nil {
		return nil, err
	}
	return log.GetMHLogEntries(ctx, req)
}
