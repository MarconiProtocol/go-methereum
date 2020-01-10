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
	"enode://f5ffe0bfb44a58db3adb583299464cae392d9adc312e601479fda04b3dbfcc64e382ec603d351d3f60e13ff6715417d1a1c97970e30571e1215e56609ea303a8@34.234.104.38:42424",		// n8.p25.c0.bn.prd.vm.tc
	"enode://9e44ead52e601d32ffe1f3d0a9029b5dc671cba823d44d34e07bbd86e6ce222b0dfc427dddd929591e9e7dfa32417508d114c2f990432d761f27e58431121c19@13.56.71.52:42424",	// n9.p25.c0.bn.prd.vm.tc
	"enode://534bce33347964a8154b68a43073a460125ac40fd04f17ca84948ecc99f70a86110ab25256bb8f59c9e0a040adf4687f4d7983e3e403e8d1012eaf223870dd6d@13.59.179.155:42424",	// n10.p25.c0.bn.prd.vm.tc
	"enode://48c7d3297c47bbc0238213278ac5ee52f3950dedb82bf3e25680c6cab631d30e2380f75c2466619ec06661bacf804197e5f32b60ad5946f041d41c80309d46b8@54.245.50.96:42424", // n11.p25.c0.bn.prd.vm.tc
	"enode://349c0ac8bc6488a55cda0d5bd528b7a7e043e1c560d601f68249f88e0c623d86d14d3cf4ec94ef0ec3bdb17f6a1e4b492c3312c34ec6715a2d01961009ba11d9@13.52.143.247:42424", // n0.p125.c0.bn.prd.vm.tc:
	"enode://ba976fb2da7acbcb68dae4aa38e4bae18df1f72c0011728b8d5d57c3ec4d3df6472b0bf68203d73cf51898465dac1250216ec8221618b4a433fca8c0578b916b@13.52.17.180:42424", // n1.p125.c0.bn.prd.vm.tc
	"enode://0c1aa197a8281955d2685f54629720ea2e07256813ebb3d7fb3e65b859a0f9655317c5c608aa69a4a24e58c8cace66409453ac1d3ac502b9650aa7eec2e6f108@13.56.106.66:42424", // n2.p125.c0.bn.prd.vm.tc
	"enode://c74308e27a0d8bf3b7096681132442fce424c3a354a17e850bdc0e344914da7e1335ff1f0fc9dd52ae23c1d5e0cabb06cb392c268a7ea3b901f399bcee88b2ee@54.67.21.47:42424", // n3.p125.c0.bn.prd.vm.tc
	"enode://a10c54297be1fffaf1e53a2fc8557cfa18d244375034a380cb09ed78ffc69f2d576013b883666b762d43c502ed83382ed97ca8f478b508af2f35c65147febee2@51.81.29.11:42424", // n0.p37.c0.bn.prd.vm.tc
	"enode://64f7d37a328530290bcbd2caf4a7a700041b75c6ae7806a3a0cb9c930d728bfe768c6b882148a930314ca2d58ad9c0fdca675ad3dfb8617fef45182d1b7e9111@51.81.29.49:42424", // n1.p37.c0.bn.prd.vm.tc
	"enode://1486727e492989e147ecf9fd31a3124a3e317e48e548da5a05d9f51610f1be02f3c6afc1312dfc76baf408350fd3adf65a36593b6eda2f9caa4b40e7d6fcfadc@51.81.1.100:42424", // n2.p37.c0.bn.prd.vm.tc
	"enode://d20017090df13e84d7ba5fa0251cf412678d3d594ba74c9a80b517c23a2b3cd92411ad53156d95517f8198e689fd29a10383d4a7151056237150d04ee47f1d8b@51.81.1.101:42424", // n3.p37.c0.bn.prd.vm.tc
	"enode://cfc3bdafeca518c0e6ee0fdb619ef3e80eabd2565393353b20075fb11bb913b9a5e7fdd8c106ddf75d1658978f0dddd65b4f33b03bb02aa4277ddc3748cf5712@147.135.81.250:42424", //
	"enode://2fea06c4d9981b6916f8dcc6569b94db8c60dd34182ffaccfe3b72ebaf6a3b0564954eee5d610d866937ee8130b224e2357304d00d1aa72290d41e3312187609@51.81.224.250:42424", //
}

// Faraday test network.
var TestnetBootnodes = []string{
	"enode://3e112c527fc6e672808de5129d6d493b77191706d59dda984496c34b49109ae09ff1f35398aa6f31e44cb3ab65ec91443e94ab7151d1bed618933acae34069a1@54.39.118.148:23200",
	"enode://4e4a8f23ade0c45bd79fdebb19ecbb22c6eb68855682cf150b3601ace1a3db201cfbe40e202bb15139f82498485b1fa647c01962905d9bfe5fea6ba7e258576c@3.16.178.211:23200",
	"enode://2ae6112bdcc30543111188c40eecc7c0ca7d056304c3cb8db314565571917db4331aa1a438f6b92ce1c6294038fed630ae2becc4c7f773bb7c1dec980a712c17@54.177.152.179:23200",
	"enode://c9ac3b611129aa4b717b2ea784d0af2e6896f76e284fa8d8cb671daad2854ee759f692275f4e15ceb8833758e4eadaea2132ddfde4dd4a3900c21e7b2854a122@54.39.8.60:23200",
}

var RinkebyBootnodes = []string{
	"enode://a0e40c0e4228e7a78288ecdc7f100912ae075305a39f0b382541b1689867632facb7d0671df53e35d04fa24af147a9b4b9687acc75d58fef9c2895c880cb73fc@44.224.78.43:23200",
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstrem bootnodes
	"enode://011f758e6552d105183b1761c5e2dea0111bc20fd5f6422bc7f91e0fabbec9a6595caf6239b37feb773dddd3f87240d99d859431891e4a642cf2a0a9e6cbb98a@51.141.78.53:30303",
	"enode://176b9417f511d05b6b2cf3e34b756cf0a7096b3094572a8f6ef4cdcb9d1f9d00683bf0f83347eebdf3b81c3521c2332086d9592802230bf528eaf606a1d9677b@13.93.54.137:30303",
	"enode://46add44b9f13965f7b9875ac6b85f016f341012d84f975377573800a863526f4da19ae2c620ec73d11591fa9510e992ecc03ad0751f53cc02f7c7ed6d55c7291@94.237.54.114:30313",
	"enode://c1f8b7c2ac4453271fa07d8e9ecf9a2e8285aa0bd0c07df0131f47153306b0736fd3db8924e7a9bf0bed6b1d8d4f87362a71b033dc7c64547728d953e43e59b2@52.64.155.147:30303",
	"enode://f4a9c6ee28586009fb5a96c8af13a58ed6d8315a9eee4772212c1d4d9cebe5a8b8a78ea4434f318726317d04a3f531a1ef0420cf9752605a562cfe858c46e263@213.186.16.82:30303",

	// Ethereum Foundation bootnode
	"enode://573b6607cd59f241e30e4c4943fd50e99e2b6f42f9bd5ca111659d309c06741247f4f1e93843ad3e8c8c18b6e2d94c161b7ef67479b3938780a97134b618b5ce@52.56.136.200:30303",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://4abad20ca8bc8571fbb64c31e961cbec1d61bcf6c6485f874a54a748cee0198cc9784fb93ab28a859afdcb5ea65b6a1b0fd81a05fe373937adaa33984a037990@13.52.50.48:42424",
}
