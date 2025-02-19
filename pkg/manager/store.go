/*
 * Copyright (c) 2020. Ant Group. All rights reserved.
 * Copyright (c) 2022. Nydus Developers. All rights reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package manager

import (
	"context"

	"github.com/containerd/nydus-snapshotter/pkg/daemon"
	"github.com/containerd/nydus-snapshotter/pkg/rafs"
	"github.com/containerd/nydus-snapshotter/pkg/store"
)

// Nydus daemons and fs instances persistence storage.
type Store interface {
	// If the daemon is inserted to DB before, return error ErrAlreadyExisted.
	AddDaemon(d *daemon.Daemon) error
	UpdateDaemon(d *daemon.Daemon) error
	DeleteDaemon(id string) error
	WalkDaemons(ctx context.Context, cb func(*daemon.ConfigState) error) error
	CleanupDaemons(ctx context.Context) error

	AddRafsInstance(r *rafs.Rafs) error
	DeleteRafsInstance(snapshotID string) error
	WalkRafsInstances(ctx context.Context, cb func(*rafs.Rafs) error) error

	NextInstanceSeq() (uint64, error)

	AddInfo(supplementInfo *daemon.NydusdSupplementInfo) error
	GetInfo(daemonID string) (*daemon.NydusdSupplementInfo, error)
}

var _ Store = &store.DaemonRafsStore{}
