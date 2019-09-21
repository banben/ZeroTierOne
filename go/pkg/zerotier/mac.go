/*
 * Copyright (c)2019 ZeroTier, Inc.
 *
 * Use of this software is governed by the Business Source License included
 * in the LICENSE.TXT file in the project's root directory.
 *
 * Change Date: 2023-01-01
 *
 * On the date above, in accordance with the Business Source License, use
 * of this software will be governed by version 2.0 of the Apache License.
 */
/****/

package zerotier

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// MAC represents an Ethernet hardware address
type MAC uint64

// NewMACFromString decodes a MAC address in canonical colon-separated hex format
func NewMACFromString(s string) (MAC, error) {
	ss := strings.Split(s, ":")
	if len(ss) != 6 {
		return MAC(0), ErrInvalidMACAddress
	}
	var m uint64
	for i := 0; i < 6; i++ {
		m <<= 8
		c, _ := strconv.ParseUint(ss[i], 16, 64)
		if c > 0xff {
			return MAC(0), ErrInvalidMACAddress
		}
		m |= (c & 0xff)
	}
	return MAC(m), nil
}

// String returns this MAC address in canonical human-readable form
func (m MAC) String() string {
	return fmt.Sprintf("%.2x:%.2x:%.2x:%.2x:%.2x:%.2x", (uint64(m)>>40)&0xff, (uint64(m)>>32)&0xff, (uint64(m)>>24)&0xff, (uint64(m)>>16)&0xff, (uint64(m)>>8)&0xff, uint64(m)&0xff)
}

// MarshalJSON marshals this MAC as a string
func (m MAC) MarshalJSON() ([]byte, error) {
	return []byte("\"" + m.String() + "\""), nil
}

// UnmarshalJSON unmarshals this MAC from a string
func (m *MAC) UnmarshalJSON(j []byte) error {
	var s string
	err := json.Unmarshal(j, &s)
	if err != nil {
		return err
	}
	*m, err = NewMACFromString(s)
	return err
}
