// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

// Constants containing the genesis allocation of built-in genesis blocks.
// Their content is an RLP-encoded list of (address, balance) tuples.
// Use mkalloc.go to create/update them.

const mainnetAllocData = "\xf8\xa8\xe1\x94\xbb\xfd\xcb\xbd\x22\x96\x0b\x6f\xcf\x4a\x0a\x10\x1b\x81\x66\x14\xaa\x55\x1c\x4b\x8b\x02\x06\x79\x2b\x1a\x29\x93\x21\x74\x00\x00\xe1\x94\xbc\x45\x17\xbc\x2d\xde\x77\x47\x81\xe3\xd7\xb4\x96\x77\xde\x34\x49\xd4\xd5\x81\x8b\x01\xa7\x84\x37\x9d\x99\xdb\x42\x00\x00\x00\xe1\x94\x62\xa8\x7d\x97\x16\xb5\x82\x60\x63\xd9\x82\x94\x68\x8e\xc7\x6f\x77\x40\x34\xd6\x8b\x04\xf6\x8c\xa6\xd8\xcd\x91\xc6\x00\x00\x00\xe0\x94\x81\x75\x70\xe7\xe0\x83\x8c\xa0\xc6\xc1\x36\xbf\x97\x01\x96\x2f\xf7\xa6\xe5\x62\x8a\xd3\xc2\x1b\xce\xcc\xed\xa1\x00\x00\x00\xe0\x94\xbd\x27\x46\xc1\x32\x39\x3f\xd8\x22\xd9\x71\xee\xca\xf7\xf4\xcd\x77\x0a\x54\x72\x8a\xd3\xc2\x1b\xce\xcc\xed\xa1\x00\x00\x00"
const testnetAllocData = "\xf8\xab\xe1\x94\xbb\xfd\xcb\xbd\x22\x96\x0b\x6f\xcf\x4a\x0a\x10\x1b\x81\x66\x14\xaa\x55\x1c\x4b\x8b\xca\x87\x54\xd6\x38\x3d\x79\x11\x50\x00\x00\xe1\x94\xbc\x45\x17\xbc\x2d\xde\x77\x47\x81\xe3\xd7\xb4\x96\x77\xde\x34\x49\xd4\xd5\x81\x8b\xa5\x6f\xa5\xb9\x90\x19\xa5\xc8\x00\x00\x00\xe2\x94\x62\xa8\x7d\x97\x16\xb5\x82\x60\x63\xd9\x82\x94\x68\x8e\xc7\x6f\x77\x40\x34\xd6\x8c\x01\xf0\x4e\xf1\x2c\xb0\x4c\xf1\x58\x00\x00\x00\xe1\x94\x81\x75\x70\xe7\xe0\x83\x8c\xa0\xc6\xc1\x36\xbf\x97\x01\x96\x2f\xf7\xa6\xe5\x62\x8b\x52\xb7\xd2\xdc\xc8\x0c\xd2\xe4\x00\x00\x00\xe1\x94\xbd\x27\x46\xc1\x32\x39\x3f\xd8\x22\xd9\x71\xee\xca\xf7\xf4\xcd\x77\x0a\x54\x72\x8b\x52\xb7\xd2\xdc\xc8\x0c\xd2\xe4\x00\x00\x00"
const rinkebyAllocData = "\xf8\xab\xe1\x94\xbb\xfd\xcb\xbd\x22\x96\x0b\x6f\xcf\x4a\x0a\x10\x1b\x81\x66\x14\xaa\x55\x1c\x4b\x8b\xca\x87\x54\xd6\x38\x3d\x79\x11\x50\x00\x00\xe1\x94\xbc\x45\x17\xbc\x2d\xde\x77\x47\x81\xe3\xd7\xb4\x96\x77\xde\x34\x49\xd4\xd5\x81\x8b\xa5\x6f\xa5\xb9\x90\x19\xa5\xc8\x00\x00\x00\xe2\x94\x62\xa8\x7d\x97\x16\xb5\x82\x60\x63\xd9\x82\x94\x68\x8e\xc7\x6f\x77\x40\x34\xd6\x8c\x01\xf0\x4e\xf1\x2c\xb0\x4c\xf1\x58\x00\x00\x00\xe1\x94\x81\x75\x70\xe7\xe0\x83\x8c\xa0\xc6\xc1\x36\xbf\x97\x01\x96\x2f\xf7\xa6\xe5\x62\x8b\x52\xb7\xd2\xdc\xc8\x0c\xd2\xe4\x00\x00\x00\xe1\x94\xbd\x27\x46\xc1\x32\x39\x3f\xd8\x22\xd9\x71\xee\xca\xf7\xf4\xcd\x77\x0a\x54\x72\x8b\x52\xb7\xd2\xdc\xc8\x0c\xd2\xe4\x00\x00\x00"
const devAllocData     = "\xf8\xab\xe1\x94\xbb\xfd\xcb\xbd\x22\x96\x0b\x6f\xcf\x4a\x0a\x10\x1b\x81\x66\x14\xaa\x55\x1c\x4b\x8b\xca\x87\x54\xd6\x38\x3d\x79\x11\x50\x00\x00\xe1\x94\xbc\x45\x17\xbc\x2d\xde\x77\x47\x81\xe3\xd7\xb4\x96\x77\xde\x34\x49\xd4\xd5\x81\x8b\xa5\x6f\xa5\xb9\x90\x19\xa5\xc8\x00\x00\x00\xe2\x94\x62\xa8\x7d\x97\x16\xb5\x82\x60\x63\xd9\x82\x94\x68\x8e\xc7\x6f\x77\x40\x34\xd6\x8c\x01\xf0\x4e\xf1\x2c\xb0\x4c\xf1\x58\x00\x00\x00\xe1\x94\x81\x75\x70\xe7\xe0\x83\x8c\xa0\xc6\xc1\x36\xbf\x97\x01\x96\x2f\xf7\xa6\xe5\x62\x8b\x52\xb7\xd2\xdc\xc8\x0c\xd2\xe4\x00\x00\x00\xe1\x94\xbd\x27\x46\xc1\x32\x39\x3f\xd8\x22\xd9\x71\xee\xca\xf7\xf4\xcd\x77\x0a\x54\x72\x8b\x52\xb7\xd2\xdc\xc8\x0c\xd2\xe4\x00\x00\x00"