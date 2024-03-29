// Copyright 2018 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package remote

import (
	"github.com/google/go-containerregistry/name"
	"github.com/google/go-containerregistry/v1"
)

// MountableLayer wraps a v1.Layer in a shim that enables the layer to be
// "mounted" when published to another registry.
type MountableLayer struct {
	v1.Layer

	Repository name.Repository
}

// mountableImage wraps the v1.Layer references returned by the embedded v1.Image
// in MountableLayer's so that remote.Write might attempt to mount them from their
// source repository.
type mountableImage struct {
	v1.Image

	Repository name.Repository
}

// Layers implements v1.Image
func (mi *mountableImage) Layers() ([]v1.Layer, error) {
	ls, err := mi.Image.Layers()
	if err != nil {
		return nil, err
	}
	mls := make([]v1.Layer, 0, len(ls))
	for _, l := range ls {
		mls = append(mls, &MountableLayer{
			Layer:      l,
			Repository: mi.Repository,
		})
	}
	return mls, nil
}

// LayerByDigest implements v1.Image
func (mi *mountableImage) LayerByDigest(d v1.Hash) (v1.Layer, error) {
	l, err := mi.Image.LayerByDigest(d)
	if err != nil {
		return nil, err
	}
	return &MountableLayer{
		Layer:      l,
		Repository: mi.Repository,
	}, nil
}

// LayerByDiffID implements v1.Image
func (mi *mountableImage) LayerByDiffID(d v1.Hash) (v1.Layer, error) {
	l, err := mi.Image.LayerByDiffID(d)
	if err != nil {
		return nil, err
	}
	return &MountableLayer{
		Layer:      l,
		Repository: mi.Repository,
	}, nil
}
