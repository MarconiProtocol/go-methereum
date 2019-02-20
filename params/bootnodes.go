// Copyright 2015 The go-ethereum Authors
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

package params

// Marconi main Net
var MainnetBootnodes = []string{
	//WIP: missing n41
	"enode://a10c54297be1fffaf1e53a2fc8557cfa18d244375034a380cb09ed78ffc69f2d576013b883666b762d43c502ed83382ed97ca8f478b508af2f35c65147febee2@51.75.153.99:13200",		// n0
	"enode://64f7d37a328530290bcbd2caf4a7a700041b75c6ae7806a3a0cb9c930d728bfe768c6b882148a930314ca2d58ad9c0fdca675ad3dfb8617fef45182d1b7e9111@139.99.179.17:13200",		// n33
	"enode://349c0ac8bc6488a55cda0d5bd528b7a7e043e1c560d601f68249f88e0c623d86d14d3cf4ec94ef0ec3bdb17f6a1e4b492c3312c34ec6715a2d01961009ba11d9@79.137.123.238:13200",	// n35
	"enode://0c1aa197a8281955d2685f54629720ea2e07256813ebb3d7fb3e65b859a0f9655317c5c608aa69a4a24e58c8cace66409453ac1d3ac502b9650aa7eec2e6f108@51.38.86.240:13200",		// n37
	"enode://450f0985be5e71fe40326d5e00be096ef9fe1edb3812b4091e09c241bc6255f9067d9a51e29c9109631ddf5619acb463dfbddd14b5e73d849a13b594c7a5d1b5@147.135.178.162:13200",	// n39
	"enode://7588b939c8170c3611def2fe83799d8ec8f5cf9aaa38eafa5fd992311aaa5c3fc6720e560563a5099e6fdadd8c9b4151894a64e50fe8347caad4180449d45a39@149.56.117.96:13200",	// n41
}

// Faraday test network.
var TestnetBootnodes = []string{
	"enode://3e112c527fc6e672808de5129d6d493b77191706d59dda984496c34b49109ae09ff1f35398aa6f31e44cb3ab65ec91443e94ab7151d1bed618933acae34069a1@54.39.118.148:23200",
	"enode://4e4a8f23ade0c45bd79fdebb19ecbb22c6eb68855682cf150b3601ace1a3db201cfbe40e202bb15139f82498485b1fa647c01962905d9bfe5fea6ba7e258576c@3.16.178.211:23200",
	"enode://2ae6112bdcc30543111188c40eecc7c0ca7d056304c3cb8db314565571917db4331aa1a438f6b92ce1c6294038fed630ae2becc4c7f773bb7c1dec980a712c17@54.177.152.179:23200",
	"enode://c9ac3b611129aa4b717b2ea784d0af2e6896f76e284fa8d8cb671daad2854ee759f692275f4e15ceb8833758e4eadaea2132ddfde4dd4a3900c21e7b2854a122@54.39.8.60:23200",
}

var RinkebyBootnodes = []string{
	"enode://4abad20ca8bc8571fbb64c31e961cbec1d61bcf6c6485f874a54a748cee0198cc9784fb93ab28a859afdcb5ea65b6a1b0fd81a05fe373937adaa33984a037990@13.52.50.48:42424",
}

var DiscoveryV5Bootnodes = []string{
	"enode://4abad20ca8bc8571fbb64c31e961cbec1d61bcf6c6485f874a54a748cee0198cc9784fb93ab28a859afdcb5ea65b6a1b0fd81a05fe373937adaa33984a037990@13.52.50.48:42424",
}
