package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _faucet_html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x3a\x7f\x73\xdb\x36\x96\x7f\x3b\x33\xf9\x0e\x2f\xdc\xba\xa4\xce\x26\x29\x39\x69\xda\x4a\xa2\x3b\x69\x36\xdd\xf3\xcd\xee\xb6\xd3\x26\xb3\x77\xd3\xed\xdd\x40\xc4\x93\x08\x1b\x04\x58\x00\x94\xac\xed\xf8\xbb\xdf\x00\x04\x29\x52\xa6\x9d\x64\x93\x7f\x2c\xf2\xe1\xe1\xfd\xfe\x05\xd0\xcb\x67\x7f\xfe\xf1\xf5\xdb\xff\xf9\xe9\x0d\x14\xa6\xe4\x97\x4f\x9f\x2c\xed\x2f\x70\x22\x36\x59\x80\x22\xb8\x7c\xfa\xc4\xc2\x90\xd0\xcb\xa7\x4f\x4e\x96\x25\x1a\x02\x79\x41\x94\x46\x93\x05\xb5\x59\xc7\xdf\x04\x87\x85\xc2\x98\x2a\xc6\xdf\x6b\xb6\xcd\x82\xff\x8e\xdf\xbd\x8a\x5f\xcb\xb2\x22\x86\xad\x38\x06\x90\x4b\x61\x50\x98\x2c\xb8\x7a\x93\x21\xdd\x60\x6f\x9f\x20\x25\x66\xc1\x96\xe1\xae\x92\xca\xf4\x50\x77\x8c\x9a\x22\xa3\xb8\x65\x39\xc6\xee\xe5\x1c\x98\x60\x86\x11\x1e\xeb\x9c\x70\xcc\x66\x8d\x84\x27\x4b\xc3\x0c\xc7\xcb\x77\xdf\x5f\xbd\x86\x1f\x48\x9d\xa3\x59\xa6\x0d\xa8\x59\xe6\x4c\xdc\x40\xa1\x70\x9d\x05\x56\x4a\x3d\x4f\xd3\x9c\x8a\x6b\x9d\xe4\x5c\xd6\x74\xcd\x89\xc2\x24\x97\x65\x4a\xae\xc9\x6d\xca\xd9\x4a\xa7\x66\xc7\x8c\x41\x15\xaf\xa4\x34\xda\x28\x52\xa5\xcf\x93\xe7\xc9\xd7\x69\xae\x75\xda\xc1\x92\x92\x89\x24\xd7\x3a\x00\x85\x3c\x0b\xb4\xd9\x73\xd4\x05\xa2\x09\x20\xbd\xfc\xf7\xf8\xae\xa5\x30\x31\xd9\xa1\x96\x25\xa6\x2f\x92\xaf\x93\xa9\x63\xd9\x07\xbf\x87\xab\xe5\xab\x73\xc5\x2a\x03\x5a\xe5\x1f\xcc\xf8\xfa\xf7\x1a\xd5\x3e\x7d\x9e\xcc\x92\x99\x7f\x71\x8c\xae\x75\x70\xb9\x4c\x1b\x82\x97\x9f\x44\x3b\x16\xd2\xec\xd3\x8b\xe4\x45\x32\x4b\x2b\x92\xdf\x90\x0d\xd2\x96\x93\x5d\x4a\x5a\xe0\x67\xe3\xfb\x90\x13\xaf\x8f\x7d\xf8\x39\x98\x95\xb2\x44\x61\x92\x6b\x9d\x5e\x24\xb3\x6f\x92\x69\x0b\x18\xa1\xef\x38\x58\xb7\x59\x5e\x27\xc9\x16\x95\x61\x39\xe1\x71\x8e\xc2\xa0\x82\x3f\x2c\xf4\xa4\x64\x22\x2e\x90\x6d\x0a\x33\x87\xd9\x74\x7a\xba\x18\x83\x6e\x8b\x06\x4c\x99\xae\x38\xd9\xcf\x61\xcd\xf1\xb6\x01\x11\xce\x36\x22\x66\x06\x4b\x3d\x87\x86\xb2\x5b\xb8\x73\xfc\x4f\x92\x4a\xc9\x8d\x42\xad\x3d\xbb\x4a\x6a\x66\x98\x14\x73\x1b\x55\xc4\xb0\x2d\x8e\x63\xeb\x8a\x88\x7b\x5b\xc8\x4a\x4b\x5e\x1b\x3c\x12\x66\xc5\x65\x7e\xd3\xc0\x5c\xfa\xf6\x15\xc9\x25\x97\x6a\x0e\xbb\x82\x99\x01\xa7\x4a\x61\x4b\x9e\x50\xca\xc4\x66\x0e\x2f\x2b\xaf\x52\x49\xd4\x86\x89\x39\x4c\xdb\x0d\x27\xcb\xb4\xb5\xe3\x32\xf5\x65\xea\xe9\x93\xe5\x4a\xd2\x7d\xf3\x44\xd9\x16\x18\xcd\x82\xa6\xa6\xd0\xef\x25\xdd\x07\x90\x73\xa2\x75\x16\x1c\x99\xbd\xa9\x48\x76\x83\x5f\xb7\x7b\x08\x13\x7e\x65\xb0\xa4\xe4\x2e\x00\xc7\x39\x0b\x1a\xa1\xe2\x95\x34\x46\x96\x73\x98\x59\x69\x9b\x1d\x47\xd4\x78\xcc\x37\xf1\xec\xc2\xaf\x9d\x2c\x8b\x59\x4b\xc2\xe0\xad\x89\x9d\xbf\x3a\x4f\x05\x97\x4b\xd6\x6e\x5d\x13\x58\x93\x78\x45\x4c\x11\x00\x51\x8c\xc4\x05\xa3\x14\x45\x16\x18\x55\xa3\x0d\x2c\x76\x09\x83\xd2\x57\xcc\x5a\x1e\x56\x80\xc7\x98\xfc\x40\xb4\x01\x22\xa8\x75\x3a\x23\x2b\x8e\x09\x4c\x93\xaf\xc0\xa0\x36\x8e\x64\x4a\xc9\x3e\x59\xa6\x94\x6d\xbd\x46\xdd\x63\xef\xa9\xc7\xa4\xef\xe6\x45\xe7\xc0\x17\xd3\xea\x16\x66\xdf\x9c\x2e\x56\x52\x51\x54\x73\x98\x55\xb7\xa0\x25\x67\x14\xfe\x94\xe7\xb9\x07\xc7\x8a\x50\x56\xeb\x39\x5c\x4c\x87\x26\x6c\x74\x39\x19\xd8\x7b\x0e\x53\x20\xb5\x91\x8b\x61\xe4\x5f\xd7\xda\xb0\xf5\x3e\xf6\x3d\xa4\xd3\x74\x2c\x19\xbc\x78\x9d\xe7\x1c\xdb\x31\x6f\x7a\xf1\x3a\xd8\x41\x78\x4a\xe9\x22\x18\xb1\xb5\xa7\xa2\x7c\x9e\x5e\x38\x7d\xfe\x41\x38\x47\x03\xaf\x28\xb5\x89\xd4\x33\xea\x60\xab\x55\x63\x0e\xb3\x8e\xec\xc9\x92\x89\xaa\x36\x2e\x8a\x6b\xc5\x03\xdf\x2e\xdd\xa3\xd9\x57\xde\xb1\x5d\x50\xaf\xa5\x2a\x9d\xf6\x4a\xf2\xc0\x53\x38\x01\x80\x8a\x93\x1c\x0b\xc9\x29\xaa\x2c\xf8\x89\x23\xd1\x08\x0d\xe1\xbd\xac\x15\xbc\x79\xfb\x9f\xb0\x6b\xe4\x23\x8d\x7c\x49\xa7\xd7\x98\xf7\x9d\xc8\x1e\xe1\x47\x91\xa3\xa5\x02\x66\x87\x68\x60\xc7\x4c\xd1\x10\xf5\x94\x1a\x88\x29\x10\xa8\x62\x15\xac\x5d\x90\xc2\xbb\x9f\xff\xea\x02\xcf\x20\xe7\x50\x6b\xd8\x15\xa8\x10\x8c\x84\x35\xb3\xd0\x02\x1b\x72\xe7\x96\xf0\x8e\x71\x0e\x1b\x34\x16\xdc\xf0\xec\x02\x34\x19\x13\x0c\x8e\x42\xc5\x7a\x11\xa6\x07\x4f\x11\x90\x22\xe7\x2c\xbf\xc9\x02\x59\xa1\x78\xdb\xb4\x89\x77\x8a\x47\x93\xe0\xf2\x2f\x28\x50\x11\x83\x07\x19\x60\x2d\x15\x94\xe8\xc4\xbd\x0a\x39\x07\x5d\x10\x85\xcf\x96\x29\x19\x35\xcb\x07\x14\x86\x51\x29\x3f\x39\xa0\x0f\x21\xf3\xbe\x50\x7c\xeb\xd4\x7a\xf7\xf3\x5f\xfb\x51\xf8\x78\x18\xf6\xe3\xd0\x59\xe5\xdd\x21\x18\x0f\xef\x1f\x1a\x91\x1f\x10\x92\xa6\x95\x31\x39\xe8\xd5\xcf\x99\x07\xc3\xf2\x31\xfb\x7f\x82\x71\x1b\xb6\x4e\x7f\x85\x39\xa9\x4c\x5e\x90\xff\x43\x8e\xb6\xc9\xdb\x22\x3c\x2a\xc3\x68\xd1\xf5\x0a\xac\x6a\x63\xa4\xf0\x26\x6b\x5e\x3a\xa3\xad\x8c\x80\x95\x11\x71\xa5\x58\x49\xd4\x3e\x38\x04\xac\xae\x57\x25\x33\x36\x50\x7f\x41\x41\xe1\x6f\x08\x6f\xdb\x5c\x58\xa6\x0d\x95\x07\x4b\x75\xfb\xd0\xfe\x3e\x7d\xb2\x3c\xcc\x3c\x69\x0a\x7f\xe1\x72\x45\x38\x6c\x6d\x93\x59\x71\xd4\x36\x1d\xad\x83\x5c\x2a\xe4\xb5\x52\x28\x0c\x68\x43\x4c\xad\x41\xae\x1d\xb4\xc9\xe6\xa7\x4f\x4e\xb6\x44\x01\x31\x06\xcb\xca\x40\xd6\x74\x68\x0b\xd2\xa8\xb6\xcd\xe0\x61\xdf\x0c\x43\x35\x58\xdd\x97\x2b\xc9\x21\x83\xc0\xca\x1f\xb4\x60\x85\xbf\xd7\xa8\x8d\x86\x0c\x7e\xfd\xad\x05\x7a\x93\x2f\xdc\x98\x90\xa6\xf0\x67\x5c\x33\x81\x40\x60\x5d\x8b\xdc\x8e\x21\x60\x0a\x62\x20\x57\x48\x0c\x6a\xc8\xb9\xd4\xb5\x6a\x74\xa0\x4a\x56\x60\xf5\x68\x09\x7b\x92\x16\x5e\x39\x89\x3a\x1a\x51\x41\x74\x31\x69\x26\x10\x85\xa6\x56\xa2\xb7\xe6\xe1\x27\xb6\x1e\x44\x96\x00\x73\xca\x00\x83\x65\x47\x3a\xe1\x28\x36\xa6\x58\x00\x3b\x3b\x6b\xf1\x4f\xd8\x1a\xa2\x16\xe1\x57\xf6\x5b\x62\x6e\x13\xcb\x07\xb2\x0c\x7a\xfc\x4e\x2c\x4b\x4f\x45\x57\x9c\xe5\x18\xb1\x73\x98\x4d\x16\x7e\x71\xa5\x90\xf8\x71\xca\x0d\x3f\xfe\xaf\xfd\x73\xd7\x98\xa5\x31\x9e\xb7\xd4\x6b\xc2\xf9\x8a\xe4\x37\x03\xf5\x3c\xab\xc2\xe3\x24\x0a\x05\x45\x15\x85\xf7\x42\x3a\x3c\xf7\x42\x85\x9a\x19\xbc\xc1\x7d\x38\x87\x10\x57\xb3\xaf\xd7\x39\xa5\xf1\xec\x65\xfe\x4d\xfc\x82\x50\x12\x93\x97\xe4\xdb\x78\x95\xaf\xbf\xa5\x5f\x13\x32\xfd\xf6\x25\x86\xe7\xcd\xb6\xdc\xb3\x0f\xe7\x40\xb4\x66\x1b\xd1\xca\xe3\xd6\xef\x9c\x56\x3d\xb1\x87\x38\x03\x99\x37\x5e\x34\x2f\xbb\x7f\x83\x0c\x36\x87\x98\xe8\x93\x1a\xd6\xf4\x31\xf5\x1d\x43\xdf\x9d\x32\xf8\x22\x0a\xfe\x64\xfb\xe9\xe4\xd7\xe9\x6f\xc9\x96\xf0\x1a\x9b\x51\xd4\x7a\xad\xed\x86\x25\x31\x79\x11\xa5\xff\x1b\x4d\x6f\x27\xdf\xfd\x3a\x8d\xbf\x25\xf1\xfa\x55\xfc\xc3\x6f\x7f\xbc\x98\xde\x7d\x91\x4e\x5a\x0f\xee\x98\xa0\x72\x97\x58\x09\xa2\xb0\x3d\x38\xf8\x63\x88\x3b\x2f\x30\x57\x69\x52\x57\xdc\xbe\xb3\xd5\x21\xfb\xb9\x71\x39\x13\x9b\xd3\x8b\x69\x93\x4e\xf6\xa1\x16\x54\x9f\x5e\x4c\x99\x30\xf2\xf4\x62\x1a\xc2\x59\x27\xf0\x19\x84\xa7\x17\x53\x29\x4e\x2f\xa6\xa6\xc0\xd3\x8b\xe9\xe9\xc5\xf3\x77\x2b\x96\x9f\x5e\x4c\x5f\xdb\xc3\x89\x83\x4c\x6d\x87\x3c\xbd\x98\x0a\x34\x3b\xa9\x6e\x92\x70\xe2\xac\x0e\xc8\x75\x3b\x5f\xdb\x33\x57\xf4\x07\x70\xb2\x97\xb5\x99\x43\x68\x64\xf5\xda\x95\xa8\xf0\x1c\xac\x64\x73\x08\xaf\xc4\x96\xd8\x41\xc7\xb3\x7e\x66\x57\xf6\x15\xda\x48\x50\x4a\x3a\x44\x56\xa2\xdb\xfe\xd5\x74\x7a\x0e\xed\x29\xe1\x7b\xa2\xe6\x60\xa7\x53\x68\xfc\xdc\xc4\xe8\x51\xe6\xba\xf2\x31\xc8\xdd\xa6\xb6\x69\x20\xb0\x61\xda\x40\xad\xb8\xcd\x5e\x8b\xd7\x94\x91\xb6\x6a\x38\xb4\x87\xfc\x6a\xfa\x9e\xb7\xae\xed\xba\xd3\x98\x7f\x9f\x3d\x3b\xe0\xb7\x4e\x6c\x98\x25\x1a\x05\x8d\xfe\xeb\x97\x1f\xff\x9e\x68\xa3\x98\xd8\xb0\xf5\x3e\xf2\x79\x5a\x2b\x3e\xef\x31\x3a\xf7\x85\xec\xdc\x15\xb8\xf3\xb6\x54\x35\x29\x3a\xf1\xf9\xdb\x4b\x39\x8d\x26\xf2\x76\xf9\x48\x87\xbc\x29\x2b\xb3\x6f\x39\xdb\xe6\xf8\x89\x2e\xb9\x5f\x4d\x4b\x34\x85\xa4\xd6\xee\x0a\x73\x29\x04\xe6\x06\xea\x4a\x0a\xef\x02\xe0\x52\xeb\x43\xa5\xf1\x08\x23\xae\xf0\xe8\x19\x08\xdc\xc1\x3f\x70\xf5\x8b\xcc\x6f\xd0\x44\x51\xe4\x53\x84\xcb\x9c\xd8\x0d\xf6\x64\x69\x64\x6e\xbb\x40\x96\x81\x3f\x6c\x07\x13\xf8\x0e\x82\x9d\xb6\xd9\x13\xc0\xdc\x3e\xda\xa7\x09\x9c\xc1\xf1\xf6\x42\x6a\x03\x67\x10\xa4\x4d\xea\xc4\xba\x24\xca\xa4\xa4\x62\xc1\xa4\x51\xae\x75\xa7\x14\x25\x6a\x4d\x36\x38\x90\x16\xb7\x28\x4c\xeb\x78\xab\x54\xa9\x37\x90\x81\x73\x7b\x45\x94\xc6\x06\x23\xa1\xc4\x10\xef\x48\x1b\x36\x0e\x2b\xcb\x40\xd4\xbc\x0b\x1b\xdf\x2f\x16\x6d\x65\xee\x23\x27\x2e\x9f\xe1\x59\x96\x41\x2d\xa8\xb3\x35\xed\xf6\xd9\x20\x75\xeb\xc1\x24\xb1\x7e\x3e\x6c\x98\x2c\x0e\x65\x7e\x40\x0a\xe9\x7b\x68\x21\x3d\x26\x86\x74\x94\x5a\x85\xa8\x1e\x13\xcc\xad\xf7\x69\x39\xc0\x28\x29\x51\x97\x2b\x54\x8f\xd0\x72\xf7\x01\x2d\x2d\x67\xdd\x2b\x61\x7a\x5b\xcf\x61\xf6\x72\x32\x4a\xda\xc5\xf6\x03\x94\xdf\x9f\x36\x1d\x85\x47\x73\xe5\xb1\x64\x39\x12\x47\xd7\x79\x6e\x4b\xf1\xa7\x08\xe4\x69\x74\x22\xf9\xf7\x7f\x5f\xa8\x6e\x6a\x1a\x48\x05\x5f\x7e\x09\xf7\x56\x07\x61\x9b\xa6\xf0\x37\xa2\x6e\x80\x70\x0e\x95\xc2\x2d\x93\xb5\x3e\x8c\x60\x25\xd3\x9a\x89\x0d\x10\x0d\x54\x0a\x7f\xfc\xfa\xe8\x01\xe8\x9e\x90\x1e\x0b\x2e\x61\x7a\x2c\xa1\xad\xd2\xbd\x01\x69\x64\x6e\x3a\x90\x1d\xcc\x44\xde\x1e\x23\xe3\x16\x2b\x11\x9e\x65\x10\x04\xbd\x9d\xf7\x10\xec\x7a\x4b\xe9\x44\xa3\x79\xdb\xb8\x21\xf2\x53\xe2\xd8\x00\x37\x39\x87\xe7\xd3\xe9\x74\x72\x24\xc0\x5d\x67\xd8\x57\x55\x65\xa7\x74\x22\xf6\xae\x0e\x76\x56\xb5\x6d\x1d\xec\x29\xc7\xd6\x31\x0e\xb9\xe4\x1c\x5d\x45\x6a\x76\xba\x79\x57\x96\xa5\x14\x90\x41\x3c\x5b\xdc\x9f\x22\x7b\xf6\x3b\xe8\x34\xe6\x96\x11\xa3\x1f\xb9\x66\x68\xae\x63\x06\x31\xcc\x06\xee\x18\x78\x6a\xd4\x25\x27\x9d\xe4\xac\x33\xe7\xd0\x4f\x9d\xa3\x86\xf6\x1a\x88\xef\x89\x9c\xc1\xec\x43\xf5\xe8\x56\xab\x5a\x17\xd1\x91\x9c\x93\xc5\xb1\x6b\xae\x4c\x73\xcc\x97\xb6\x4d\x59\x57\xa0\x30\x4c\xe1\x3d\x8f\xf8\xeb\xb1\xb8\x99\x94\xdb\xa9\x85\xda\x9c\x30\xf6\x94\xd4\xf7\x98\x1b\xee\xfa\x81\xf4\xf1\x89\xe2\xce\x25\x52\xd8\x70\xbc\x1f\xa1\xfd\x10\xb5\x88\xc8\x49\xa5\x91\x42\x06\xcd\x65\x73\x34\x49\x6a\xc1\x6e\xa3\x09\xc4\x2d\xe4\x98\x48\x8b\xe1\x9b\xa3\x73\x57\x23\xf6\x59\x06\xc1\xd2\x28\x7b\xc2\x0d\x03\x38\x1b\x4b\x3c\xdb\x68\xc3\xcb\x4e\x84\xfe\x4e\x80\xa5\xa1\x97\xbd\xb3\xef\x3f\x03\x3b\xc8\x6f\x94\xac\x05\x9d\xdb\x59\x2e\xba\x47\x95\x6c\x89\x21\xca\x11\x9d\x2c\xe0\x80\x1e\x6b\xf6\x2f\x9c\x43\x6e\x5d\xb3\x80\xe6\x4e\xf1\xf9\x45\x75\xbb\x80\xf6\xea\xbb\x79\x3b\xba\x37\x7c\x51\xdd\x2e\xfe\xd9\x1e\xc5\x97\xa9\xa1\x8f\x4a\x5a\x29\xbc\xbc\x27\x50\x9e\xcb\x5a\xb8\x79\x62\x99\x5a\x84\xf7\x50\xe9\x54\xed\xdf\x7b\xc2\xc8\xa1\x1f\xba\xcb\x66\x0f\x2f\x19\xa5\x1c\xad\xb8\x1d\x75\x9b\x86\xd6\xf3\xbd\x54\x1a\x32\x04\xf0\x9d\xbd\x2b\x37\xfd\xe9\x71\x04\x7d\xe9\x2e\xeb\xad\x3f\xad\xe7\x63\xab\x2d\x73\xd6\xf6\xd7\x0c\x0e\xac\x42\x67\x06\xff\xb5\x82\xd6\xca\xcd\x55\x51\xec\x63\xeb\x1c\x42\x6d\xe7\x3c\xaa\xc3\x49\x52\xd4\x25\x11\xec\x5f\x18\xd9\x36\x34\x69\xcc\x64\x79\x1c\x94\xb8\x1b\x35\x15\x40\xef\x2a\x3c\x6c\x1b\x5a\xe8\xcd\x17\xb6\x5e\xb5\x0e\x84\xc3\x25\x7f\xf8\x51\xb6\x19\xe7\x11\xaf\x88\x82\xfe\x4b\xdc\xf6\x59\x50\xd2\xf2\x6e\xd7\x56\x44\x85\xcd\xed\xba\x3b\x21\x08\xb9\xcb\xc2\xe7\xd3\x4e\xc4\xc6\xc1\xce\xbf\xa1\x8f\xb0\x63\x37\x58\x19\xdb\x84\xbc\x84\xe7\xd3\xcf\x20\x2b\x25\x62\x83\xc7\xf2\x1b\xc5\x2a\xa4\x40\x72\xc3\xb6\xf8\xf9\xd5\xf8\x74\x03\x7f\xb4\x80\x36\xfe\x5a\xcb\xb9\xf0\x1c\x48\x6b\x57\x3b\xc3\xfe\x87\xcd\x31\x48\x9d\x79\xcf\x20\x18\x53\xe3\xa1\x08\x1c\xa2\x1d\xe5\xf2\x83\x79\xbe\x4c\x8d\xea\x56\xee\x0e\xb3\x6c\x5b\x36\x82\x49\x52\x98\x92\x47\xc1\xd2\xb8\xcf\x4e\x56\xda\x6e\xbf\xdb\xde\x80\xfb\x23\xdb\xdd\xe0\x58\x92\x73\xa9\x8f\x0e\x25\x4a\xb5\xc1\x93\x4b\xa1\x25\xc7\x84\xcb\x4d\x14\x7a\xdc\xf0\xdc\x22\xf8\xa3\x6a\x37\xa6\x74\xa7\xb1\x76\x26\x81\xbb\x45\x73\xea\xb6\xdd\xee\x17\x43\x94\x01\x02\xef\xae\xa0\xae\x28\xb1\xa7\x47\x23\xc1\x76\x4a\xd7\xd1\xba\xcf\x7b\x2b\xa2\x34\xac\xa5\xda\x11\x45\xa1\x16\x86\x71\xbb\xbe\x07\xa2\xb0\x9d\xff\x34\x9a\x2b\x5b\xd4\xb6\x84\x47\xf7\x8e\x7d\x5f\x44\x61\xd2\x8f\x86\x70\x92\x20\xc9\x8b\xfb\x88\xae\x7f\x75\x6c\x33\xf8\xbb\x3b\x01\x44\x5f\x44\xa6\x60\x7a\x92\x10\x63\x54\x14\x0e\xe2\x24\x9c\x4c\xdc\x38\xd0\x0d\xbd\xdd\xee\x65\x3f\xdb\x1e\xa3\x70\x18\xa6\xdb\x79\xa0\xc5\xce\xb5\x8e\x9a\x78\x0b\xcf\x7b\x94\x87\xe1\x16\x9e\x86\xad\x1b\x0f\x29\x7f\xd0\x21\x1b\x13\x63\x40\x38\xb4\x99\x17\x1e\xf3\x26\x94\xbe\xb6\x49\x15\x05\x23\xc9\x3f\x0c\x9c\x49\x6b\xe3\xa6\x72\x3f\x6a\x5c\x26\x28\xde\x3e\x64\x59\x46\xc3\x49\xa2\xeb\x55\x73\xb5\x11\x7d\xd5\x1e\xbb\x5a\x2c\x17\xd3\xc7\x3d\xe1\xde\x48\x61\x39\x0c\xc7\x8a\xc3\xe0\xd1\x42\x1e\x69\x20\xfe\x22\xc2\x2a\x75\x77\x6e\x2d\x3d\x9d\x74\x17\x12\x6f\xb4\x9d\xae\x98\x2e\x80\xc0\x0e\x57\xda\x5d\x1f\x80\x0f\x71\x77\x63\xd4\xdc\x0c\xbd\xfa\xe9\xea\x70\x3b\xd4\xe5\x80\x9b\x6f\x7a\x5f\xdc\x47\x3f\xe8\x5f\xeb\xa4\xbb\x94\xc9\x65\x99\xce\x52\x52\xb1\xe4\x5a\x7f\x27\x05\x97\x84\x66\xf7\x6e\x52\xbf\x6c\x66\xc0\x0c\x6f\x2b\xce\x72\x66\x02\x20\x7a\x2f\x72\xa0\xb8\x46\xd5\xff\xc0\xbf\x4c\x0f\x9f\xa0\x53\xf7\xaf\x35\xff\x1f\x00\x00\xff\xff\x1a\xb3\xd4\x76\x6a\x23\x00\x00")

func faucet_html() ([]byte, error) {
	return bindata_read(
		_faucet_html,
		"faucet.html",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"faucet.html": faucet_html,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"faucet.html": {faucet_html, map[string]*_bintree_t{}},
}}
