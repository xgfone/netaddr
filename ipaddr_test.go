// Copyright 2018 xgfone
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

package netaddr

import (
	"fmt"
)

func ExampleIPAddress() {
	// IPv6
	ip, err := NewIPAddress("fe80::d656:43a8:fc42:948c")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ip, ip.IsValid())

	s := ip.Value()
	fmt.Println(ip.Equal(MustNewIPAddress(s, ip.Version())))

	fmt.Println(ip.IsGlobalUnicast())
	fmt.Println(ip.IsInterfaceLocalMulticast())
	fmt.Println(ip.IsLinkLocalMulticast())
	fmt.Println(ip.IsLinkLocalUnicast())
	fmt.Println(ip.IsLoopback())
	fmt.Println(ip.IsMulticast())

	// Ipv4
	ip, err = NewIPAddress("192.168.10.10")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ip, ip.IsValid())
	fmt.Println(ip.Bytes())

	s = ip.Value()
	fmt.Println(ip.Equal(MustNewIPAddress(s, ip.Version())))

	fmt.Println(ip.IsGlobalUnicast())
	fmt.Println(ip.IsInterfaceLocalMulticast())
	fmt.Println(ip.IsLinkLocalMulticast())
	fmt.Println(ip.IsLinkLocalUnicast())
	fmt.Println(ip.IsLoopback())
	fmt.Println(ip.IsMulticast())

	// Loopback
	ip, _ = NewIPAddress("127.0.0.1")
	fmt.Println(ip.IsLoopback())
	ip, _ = NewIPAddress("::1")
	fmt.Println(ip.IsLoopback())

	// Output:
	// fe80::d656:43a8:fc42:948c true
	// true
	// false
	// false
	// false
	// true
	// false
	// false
	// 192.168.10.10 true
	// [192 168 10 10]
	// true
	// true
	// false
	// false
	// false
	// false
	// false
	// true
	// true
}

func ExampleIPAddress_Value() {
	ip1, _ := NewIPAddress("fe80::")
	ip2, _ := NewIPAddress(ip1.Value(), ip1.Version())

	fmt.Println(ip1)
	fmt.Println(ip2)
	fmt.Println(ip1.Equal(ip2))

	// Output:
	// fe80::
	// fe80::
	// true
}

func ExampleIPAddress_Less() {
	ip1 := MustNewIPAddress("192.168.10.10")
	ip2 := MustNewIPAddress("192.168.10.11")
	fmt.Println(ip1.Less(ip2))

	// Output:
	// true
}

func ExampleIPAddress_Network() {
	ip := MustNewIPAddress("192.168.10.10")
	net := ip.Network()
	fmt.Println(net)

	ip = MustNewIPAddress("fe80::d656:43a8:fc42:948c")
	net = ip.Network()
	fmt.Println(net)

	// Output:
	// 192.168.10.10/32
	// fe80::d656:43a8:fc42:948c/128
}
