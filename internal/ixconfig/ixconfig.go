// Copyright 2021 Google Inc.
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

// Package ixconfig contains generated IxNetwork structs (along with some
// hand-written helper code) and implements an Ixia JSON config client using
// those structs.
//
// Since the autogenerated golang documentation for IxNetwork config structs in
// this package is long and hard to parse, see the OpenAPI spec for
// information about those structs instead:
// https://openixia.github.io/ixnetwork_openapi/
//
// Basic usage examples for the client can be found in README.md.
package ixconfig

import (
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"

	"github.com/openconfig/ondatra/binding/ixweb"
)

type ixSession interface {
	Config() config
}

type config interface {
	Export(context.Context) (string, error)
	Import(context.Context, string, bool) error
	QueryIDs(context.Context, ...string) (map[string]string, error)
}

type sessionWrapper struct {
	*ixweb.Session
}

func (sw *sessionWrapper) Config() config {
	return sw.Session.Config()
}

// Client implements an API for interacting with an Ixia session using a JSON-based config representation.
type Client struct {
	sess         ixSession
	lastImported *Ixnetwork
	xPathToID    map[string]string
}

// New returns a new Ixia config Client for a specific session for the given Ixia controller connection.
func New(sess *ixweb.Session) *Client {
	return &Client{sess: &sessionWrapper{sess}}
}

// Session returns the IxNetwork session used by the config client.
func (c *Client) Session() *ixweb.Session {
	return c.sess.(*sessionWrapper).Session
}

// NodeID returns the updated ID for the specified node. Returns an error if the
// node is not part of an imported config or the node ID has not been updated.
func (c *Client) NodeID(node IxiaCfgNode) (string, error) {
	xp := node.XPath()
	if xp == nil {
		return "", fmt.Errorf("node of type %T not yet imported", node)
	}
	id, ok := c.xPathToID[xp.String()]
	if !ok {
		return "", fmt.Errorf("node at %q has no updated ID", xp)
	}
	return id, nil
}

// ExportConfig exports the current full configuration of the IxNetwork session.
func (c *Client) ExportConfig(ctx context.Context) (*Ixnetwork, error) {
	cfgStr, err := c.sess.Config().Export(ctx)
	if err != nil {
		return nil, err
	}
	cfg := &Ixnetwork{}
	if err := json.Unmarshal([]byte(cfgStr), cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Ixia config object from config string %q: %w", cfgStr, err)
	}
	return cfg, nil
}

// ImportConfig imports the specified config into the IxNetwork session.
// The second argument is the root config node, and the third is the specific config to apply.
// If overwrite is 'true', the existing config is completely replaced with the contents of cfgNode.
// If overwrite is 'false, all values configured at and below the given config node are updated.
// For values that are a list of config nodes, only the nodes that are specified are updated. (Eg.
// you cannot remove a config node from a list using this function with overwrite set to 'false'.)
// All XPaths in the config are updated before this function returns.
func (c *Client) ImportConfig(ctx context.Context, cfg *Ixnetwork, node IxiaCfgNode, overwrite bool) error {
	c.xPathToID = map[string]string{}
	cfg.updateAllXPaths()

	jsonCfg, err := json.Marshal(node)
	if err != nil {
		return fmt.Errorf("could not marshal Ixnetwork config to JSON: %w", err)
	}
	if err := c.sess.Config().Import(ctx, string(jsonCfg), overwrite); err != nil {
		return err
	}

	// Record the config that was pushed.
	c.lastImported = cfg.Copy()
	return nil
}

// LastImportedConfig returns a copy of the last config push attempt using this
// client. Returns 'nil' if there has not been a config push. A new copy of the
// last config is returned on every invocation and does not have its XPaths set.
func (c *Client) LastImportedConfig() *Ixnetwork {
	return c.lastImported.Copy()
}

// UpdateIDs updates recorded REST IDs for the target nodes in the config.
// If the ID for the node is already updated, this is a noop for that node.
// This query can be expensive if used with many different types of objects.
func (c *Client) UpdateIDs(ctx context.Context, cfg *Ixnetwork, nodes ...IxiaCfgNode) error {
	// Update XPaths because they may be lost as *Ixnetwork config objects
	// are copied around (such as a config returned from
	// 'LastImportedConfig' or a user may have constructed a read-only
	// config subobject that was never imported.
	cfg.updateAllXPaths()

	var xPathsMissing []string
	for _, n := range nodes {
		xp := n.XPath().String()
		if _, ok := c.xPathToID[xp]; !ok {
			xPathsMissing = append(xPathsMissing, xp)
		}
	}
	newIDs, err := c.sess.Config().QueryIDs(ctx, xPathsMissing...)
	if err != nil {
		return err
	}
	for xp, id := range newIDs {
		c.xPathToID[xp] = id
	}
	return nil
}
