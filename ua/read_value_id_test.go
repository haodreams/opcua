// Copyright 2018-2019 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ua

import (
	"testing"
)

func TestReadValueID(t *testing.T) {
	cases := []CodecTestCase{
		{
			Name: "Normal",
			Struct: &ReadValueID{
				NodeID:       NewFourByteNodeID(0, 2256),
				AttributeID:  AttributeIDValue,
				DataEncoding: &QualifiedName{},
			},
			Bytes: []byte{
				// NodeID
				0x01,
				0x00,
				0xd0, 0x08,
				// AttributeID
				0x0d, 0x00, 0x00, 0x00,
				// Index Range
				0xff, 0xff, 0xff, 0xff,
				// qualified name
				0x00, 0x00,
				0xff, 0xff, 0xff, 0xff,
			},
		},
	}
	RunCodecTest(t, cases)
}

func TestReadValueIDArray(t *testing.T) {
	cases := []CodecTestCase{
		{
			Name:   "empty",
			Struct: []*ReadValueID{},
			Bytes: []byte{
				// length
				0x00, 0x00, 0x00, 0x00,
			},
		},
		{
			Name: "Normal",
			Struct: []*ReadValueID{
				&ReadValueID{
					NodeID:       NewStringNodeID(1, "Temperature"),
					AttributeID:  AttributeIDNodeClass,
					DataEncoding: &QualifiedName{},
				},
				&ReadValueID{
					NodeID:       NewStringNodeID(1, "Temperature"),
					AttributeID:  AttributeIDBrowseName,
					DataEncoding: &QualifiedName{},
				},
				&ReadValueID{
					NodeID:       NewStringNodeID(1, "Temperature"),
					AttributeID:  AttributeIDDisplayName,
					DataEncoding: &QualifiedName{},
				},
			},
			Bytes: []byte{
				// Length
				0x03, 0x00, 0x00, 0x00,

				// NodeID
				0x03,
				0x01, 0x00,
				0x0b, 0x00, 0x00, 0x00,
				0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65,
				// AttributeID
				0x02, 0x00, 0x00, 0x00,
				// IndexRange
				0xff, 0xff, 0xff, 0xff,
				// QualifiedName
				0x00, 0x00,
				0xff, 0xff, 0xff, 0xff,

				// NodeID
				0x03,
				0x01, 0x00,
				0x0b, 0x00, 0x00, 0x00,
				0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65,
				// AttributeID
				0x03, 0x00, 0x00, 0x00,
				// IndexRange
				0xff, 0xff, 0xff, 0xff,
				// QualifiedName
				0x00, 0x00,
				0xff, 0xff, 0xff, 0xff,

				// NodeID
				0x03,
				0x01, 0x00,
				0x0b, 0x00, 0x00, 0x00,
				0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65,
				// AttributeID
				0x04, 0x00, 0x00, 0x00,
				// IndexRange
				0xff, 0xff, 0xff, 0xff,
				// QualifiedName
				0x00, 0x00,
				0xff, 0xff, 0xff, 0xff,
			},
		},
	}
	RunCodecTest(t, cases)
}
