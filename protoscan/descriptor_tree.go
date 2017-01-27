// Copyright © 2016 Zenly <hello@zen.ly>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package protoscan

import (
	"crypto/sha1"
	"fmt"
	"reflect"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/pkg/errors"
)

// -----------------------------------------------------------------------------

type Descriptor struct {
	// descr is either a DescriptorProto (Message) or an
	// EnumDescriptorProto (Enum).
	descr proto.Message

	// hashSingle is the the SHA-1 hash of `this.descr`.
	// hashSingle = SHA1(d.descr)
	//
	// It is used to uniquely and deterministically identify a Message or Enum
	// type that has been instanciated by a protobuf library.
	//
	// hashSingle is necessary for the internal DescriptorTree machinery to work,
	// but it is never exposed to the end-user of the protoscan package.
	hashSingle string
	// hashRecursive is the the SHA-1 hash of `this.descr` + the SHA-1 hashes
	// of every type it depends of.
	// To be deterministic, this list of hashes is sorted alphabetically first.
	// hashRecursive = SHA1(SORT(d.hashSingle, d.AllDependencyHashes))
	//
	// It is the unique, deterministic and versioned identifier for a Message or
	// Enum type.
	// This is the actual value that is exposed to the end-user of the protoscan
	// package, and is traditionally used as key inside a schema registry.
	hashRecursive string
}

// Descr returns the Message/Enum descriptor.
func (d Descriptor) Descr() proto.Message { return d.descr }

// UID returns the unique, deterministic, versioned identifier of a Descriptor.
func (d Descriptor) UID() string { return d.hashRecursive }

// -----------------------------------------------------------------------------

// A DescriptorTree is a dependency tree of Message/Enum descriptors.
type DescriptorTree struct {
	*Descriptor
	children []*DescriptorTree
}

func NewDescriptorTrees(
	fdps map[string]*descriptor.FileDescriptorProto,
) (map[string]*DescriptorTree, error) {
	/* DEBUG */
	for file, fdp := range fdps {
		fmt.Println("file:", file)
		for _, dep := range fdp.GetDependency() {
			fmt.Println("\tdependency:", dep)
		}
		for _, et := range fdp.GetEnumType() {
			fmt.Println("\tenum type:", et.GetName())
		}
		for _, mt := range fdp.GetMessageType() {
			fmt.Println("\tmessage type:", mt.GetName())
		}
	}

	dtsByName, err := collectDescriptorTypes(fdps)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for name := range dtsByName {
		fmt.Println(name)
	}

	for _, dt := range dtsByName {
		if err := dt.ComputeDependencies(dtsByName); err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return nil, nil
}

// newDescriptorTree returns a new DescriptorTree with its `hashSingle`
// field already computed.
// At this stage, dependencies have not been calculated, hence the `children`
// slice of the Descriptor is nil.
func newDescriptorTree(descr proto.Message) (*DescriptorTree, error) {
	dt := &DescriptorTree{Descriptor: &Descriptor{descr: descr}}

	if err := checkDescriptorType(descr); err != nil {
		return nil, errors.WithStack(err)
	}

	descrBytes, err := proto.Marshal(descr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	h := sha1.New()
	_, err = h.Write(descrBytes)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	dt.Descriptor.hashSingle = fmt.Sprintf("%x", h.Sum(nil))

	return dt, nil
}

func (dt *DescriptorTree) ComputeDependencies(
	dtsByName map[string]*DescriptorTree,
) error {
	switch descr := dt.descr.(type) {
	case *descriptor.DescriptorProto:
	case *descriptor.EnumDescriptorProto:
	default:
		return errors.Errorf("`%v`: illegal type", reflect.TypeOf(descr))
	}
	return nil
}

// -----------------------------------------------------------------------------

// collectDescriptorTypes recursively goes through all the FileDescriptorProtos
// that are passed to it in order to collect all of the Message, Enum & Nested
// types that can be found in these descriptors.
//
// The result is a flattened map of single-level (i.e. `children` is nil)
// DescriptorTrees arranged by the types' absolute names
// (e.g. `.google.protobuf.Timestamp`).
//
// This map is a necessary intermediate representation for computing the final
// dependency trees.
func collectDescriptorTypes(
	fdps map[string]*descriptor.FileDescriptorProto,
) (map[string]*DescriptorTree, error) {
	// Pre-allocating len(fdps)*3, assuming an average of 3 Message/Enum
	// types per file.
	// This mapping is arranged by the descriptors' respective name and is
	// a necessary intermediate representation for the final computation of
	// dependency trees.
	dtsByName := make(map[string]*DescriptorTree, len(fdps)*3)

	// recurse through a protoFile and pushes every Message/Enum/Nested
	// type to the `dtsByName` map
	var recurseMessages func(mt *descriptor.DescriptorProto, parent string) error
	recurseMessages = func(mt *descriptor.DescriptorProto, parent string) error {
		dt, err := newDescriptorTree(mt)
		if err != nil {
			return errors.WithStack(err)
		}
		parent += ("." + mt.GetName())
		dtsByName[parent] = dt
		for _, et := range mt.GetEnumType() {
			dt, err := newDescriptorTree(et)
			if err != nil {
				return errors.WithStack(err)
			}
			dtsByName[parent+"."+et.GetName()] = dt
		}
		for _, nmt := range mt.GetNestedType() {
			if err := recurseMessages(nmt, parent); err != nil {
				return errors.WithStack(err)
			}
		}
		return nil
	}

	// walk through the protoFiles
	for _, fdp := range fdps {
		parent := "." + fdp.GetPackage()
		for _, et := range fdp.GetEnumType() {
			dt, err := newDescriptorTree(et)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			dtsByName[parent+"."+et.GetName()] = dt
		}
		for _, mt := range fdp.GetMessageType() {
			if err := recurseMessages(mt, parent); err != nil {
				return nil, errors.WithStack(err)
			}
		}
	}

	return dtsByName, nil
}

// -----------------------------------------------------------------------------

func checkDescriptorType(descr proto.Message) error {
	switch descr.(type) {
	case *descriptor.DescriptorProto:
	case *descriptor.EnumDescriptorProto:
	default:
		return errors.Errorf("`%v`: illegal type", reflect.TypeOf(descr))
	}
	return nil
}
