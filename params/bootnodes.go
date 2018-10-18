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

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://a979fb575495b8d6db44f750317d0f4622bf4c2aa3365d6af7c284339968eef29b69ad0dce72a4d8db5ebb4968de0e3bec910127f134779fbcb0cb6d3331163c@52.16.188.185:30303", // IE
	"enode://3f1d12044546b76342d59d4a05532c14b85aa669704bfe1f864fe079415aa2c02d743e03218e57a33fb94523adb54032871a6c51b2cc5514cb7c7e35b3ed0a99@13.93.211.84:30303",  // US-WEST
	"enode://78de8a0916848093c73790ead81d1928bec737d565119932b98c6b100d944b7a95e94f847f689fc723399d2e31129d182f7ef3863f2b4c820abbf3ab2722344d@191.235.84.50:30303", // BR
	"enode://158f8aab45f6d19c6cbf4a089c2670541a8da11978a2f90dbf6a502a4a3bab80d288afdbeb7ec0ef6d92de563767f3b1ea9e8e334ca711e9f8e2df5a0385e8e6@13.75.154.138:30303", // AU
	"enode://1118980bf48b0a3640bdba04e0fe78b1add18e1cd99bf22d53daac1fd9972ad650df52176e7c7d89d1114cfef2bc23a2959aa54998a46afcf7d91809f0855082@52.74.57.123:30303",  // SG

	// Ethereum Foundation C++ Bootnodes
	"enode://979b7fa28feeb35a4741660a16076f1943202cb72b6af70d327f053e248bab9ba81760f39d0701ef1d8f89cc1fbd2cacba0710a12cd5314d5e0c9021aa3637f9@5.1.83.226:30303", // DE
}

// EllaismBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the Ellaism network.
var EllaismBootnodes = []string{
	"enode://0d88e242aa0b01ee306ca43e956174677c96ec8eba4197f4d8be6fd7d4f2e57731e95d533b88229b66eb1a44399d870e99b7a4fe6547c8c80cdf00407a986e14@94.130.237.158:30303",
	"enode://4be9e419d3efb0214faf3ef1794a0c33ebbd7633ece734a0a956faa166fefc496b2692a2a485adc66af805e461ba3e12f8d3941ec207e56bb9f3d3626787a705@94.130.237.158:60606",
	"enode://834246cc2a7584df29ccdcf3b5366f118a0e291264980376769e809665a02c4caf0d68c43eecf8390dbeaf861823b05583807af0a62542a1f3f717046b958a76@45.77.106.33:30303",
	"enode://d8059dcb137cb52b8960ca82613eeba1d121105572decd8f1d3ea22b09070645eeab548d2a3cd2914f206e1331c7870bd2bd5a231ebac6b3d4886ec3b8e627e5@173.212.216.105:30303",
	"enode://9215ad77bd081e35013cb42a8ceadff9d8e94a78fcc680dff1752a54e7484badff0904e331c4b40a68be593782e55acfd800f076d22f9d2832e8483733ade149@213.14.82.125:30303",
	"enode://5dd35866da95aea15211fb1f98684f6e8c4e355e6aa3cc17585680ed53fa164477b8c52cb6ca4b24ec4d80f3d48ff9212b53feb131d825c7945a3abaaf02d24d@178.79.189.58:60606",
	"enode://6c585c18024eb902ca093278af73b04863ac904caabc39ac2920c23532307c572ad92afd828a990c980d272b1f26307f2409cc97aec3ff9fe866732cae49a8c2@144.217.163.224:31337",
	"enode://edd90c4cc64528802ad52fd127d80b641ff80fd43fa5292fb111c8bd2914482dffee288fd1b0d26440c6b2c669b10a53cbcd37c895ba0d6194110e100a965b2d@188.166.179.159:30303",
	"enode://9d960373335c1cc38ca696dea8f2893e2a071c8f21524f21e8aae22be032acc3b67797b1d21e866f9d832943ae7d9555b8466c6ab34f473d21e547114952df37@213.32.53.183:30303",
}

// AtheiosBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the Atheios network.
var AtheiosBootnodes = []string{
	"enode://4c0b55e67005b76136ac81e89e126c05cb2d6b5ab8aaf73c12d94c28026ab484fe6d852f56263a674fa996a9b6d465fc8cf17d15d5c3a42f188459cf275617ce@35.230.80.37:30696",
	"enode://002197c1271ab789548ce78d5f9be3c08100c9fa189c533a2ed14caf01d8ce9d6d691bddf13105b3958f339852b1b2da28d76b18dc2123be7b70bb837bb78852@35.234.112.176:30696",
	"enode://f7c4467562a52a90336e72385cc91efa09866cd09e6ea13393dbf5a5d42836263bb2678d4160890d39aca9edf244b2852e2636ef92ae1a15ff9f958fa86a0fc0@35.229.212.25:30696",
	"enode://ddd52f85230ebd07259a3e22187091373f0c74b2dc7c777f1687cbafa78fb49c484c0d80b479455d63987819c5aa91e252ce18fe51f2270e75cef77acd29fd01@35.199.65.5:30696",
	"enode://95bee7fc08916698a691f53911e557121c09d45742f6bd9965bf38de3c85bd27cc3ed7d568ded858a1c3b7861e351f826e54be0dc820b8f2d8350771ac181703@35.201.8.2:30696",
	"enode://4c5021c5ff91b022470f56936aa291a93f18afe22819309d14e6887c9e1483cc285ced7a115c5ce0407db91ce462078987883a753a4b18768b03bbc732c91c2c@35.203.118.142:30696",
	"enode://505f5ee260b6d2ca815c68118e8865036aa16ba5833f93a795dd9a4a4f27fb0fe669f1495a8877eb79a9dd35d5fb7fc843a9f990a22311a3c01e6bd3d9488263@35.200.57.196:30696",
}

// ClassicBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the Ethereum Classic network.
var ClassicBootnodes = []string{
	"enode://e809c4a2fec7daed400e5e28564e23693b23b2cc5a019b612505631bbe7b9ccf709c1796d2a3d29ef2b045f210caf51e3c4f5b6d3587d43ad5d6397526fa6179@174.112.32.157:30303",
	"enode://6e538e7c1280f0a31ff08b382db5302480f775480b8e68f8febca0ceff81e4b19153c6f8bf60313b93bef2cc34d34e1df41317de0ce613a201d1660a788a03e2@52.206.67.235:30303",
	"enode://5fbfb426fbb46f8b8c1bd3dd140f5b511da558cd37d60844b525909ab82e13a25ee722293c829e52cb65c2305b1637fa9a2ea4d6634a224d5f400bfe244ac0de@162.243.55.45:30303",
	"enode://42d8f29d1db5f4b2947cd5c3d76c6d0d3697e6b9b3430c3d41e46b4bb77655433aeedc25d4b4ea9d8214b6a43008ba67199374a9b53633301bca0cd20c6928ab@104.155.176.151:30303",
	"enode://814920f1ec9510aa9ea1c8f79d8b6e6a462045f09caa2ae4055b0f34f7416fca6facd3dd45f1cf1673c0209e0503f02776b8ff94020e98b6679a0dc561b4eba0@104.154.136.117:30303",
	"enode://72e445f4e89c0f476d404bc40478b0df83a5b500d2d2e850e08eb1af0cd464ab86db6160d0fde64bd77d5f0d33507ae19035671b3c74fec126d6e28787669740@104.198.71.200:30303",
	"enode://39abab9d2a41f53298c0c9dc6bbca57b0840c3ba9dccf42aa27316addc1b7e56ade32a0a9f7f52d6c5db4fe74d8824bcedfeaecf1a4e533cacb71cf8100a9442@144.76.238.49:30303",
	"enode://f50e675a34f471af2438b921914b5f06499c7438f3146f6b8936f1faeb50b8a91d0d0c24fb05a66f05865cd58c24da3e664d0def806172ddd0d4c5bdbf37747e@144.76.238.49:30306",
}

// SocialBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the Ethereum Social network.
var SocialBootnodes = []string{
	"enode://54d0824a268747046b6cabc7ee3afda48edba319f0d175e9e505aa9d425a1872b8b6f9ebf8f3b0a10dc7611a4c44ddec0fc691e5a5cde23e06fc4e4b3ff9dbef@13.125.185.147:30303",
	"enode://7e150d47637177f675e20d663fc2500987f2149332caf23da522d92363be8a7880ef9150a6183e9031288a441e0457239474967a111eafce17e19a4288076ea9@18.219.40.235:30303",
	"enode://6244c9d9cd288015d7ff165e90f3bb5649e34467e095a47c6d3c56e8fb8c849b3b4db683ff3c7ae8a654bbdc07ef12ee2fd7d72831ac213723281c1b0cc90599@13.250.220.98:30303",
	"enode://e39f162b9f4b6ed6f098550f7867c2fb068fc66f362b3db0f45124c43ea18508f5ceef4e0e4de53d301e14a6f1683226aeb931d7401b4e83b5a583153ffdd7fd@52.57.98.157:30303",
	"enode://54b4a117d66dc3aa93358dec1b31d4f38e72e4381b3e28a65ac6f1aaac3b304ebbe41d32cc864fa69a9a6815c34cf9b8965690dc174a5f72af14547b601b7924@222.239.255.71:30303",
	"enode://851f14c5cc86cbc0a81acfcbe5dd99ad5c823435357219df736932c5f89ad4318f6973a553857a32d97a71793f5a35c062d46320be282aa0a80b06b9c6b624e4@13.125.232.71:30303",
}

// MixBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the MIX network.
var MixBootnodes = []string{
	"enode://aeb6070deb50efeb41c5e4283a6a3b08ff701fef90e3522161c145f30df2852af3dfc51ba74591f7c9d96b11ca4c3c2b354bf58dd243f2d877f6eecc2373fd1d@139.162.15.124:30313",
	"enode://e0c926dcdc5c1cf58b2ecba371c577c28c28c91f9b210093178a812389b65e5b53f0e478753b94fceb0b36645b779a915ca57c0c48507fe4d7f786508653656c@74.207.240.177:30313",
	"enode://a2a2adb8c12b9b189306050013a44f28db30f92fb3670db9675a049b98b96eb18901d6ff7b961b6e96cfa3923ac29e8f647ef452f0a23ddfef3903ac1cf826af@173.255.195.214:30313",
	"enode://5460fd1ad217941befd0f8d060e6729a0535a0738770aba56827d1313c09aeb68e3098d458aace59faba2c6780b8c9c30cb140b80cd8e30ca3a074ce6d3344d3@50.116.38.52:30313",
	"enode://31dcc30d6864687b404bbfb77b93008073187386c452abd17d9d509bf1dcec560d5c808c680d2fca9e72e78d164945683ed671514311b4bfd1bddb580e067131@45.79.128.151:30313",
	"enode://fd80e04c75559cfdd9ed8c08ef2c39c5bc95021f7cbaf31acb601914bc7dac7c34b470b90a05e519bc8a8435a46e1ce51053ae07fac31a83567285c34a79c6bf@139.162.224.203:30313",
	"enode://4742134a153c108855eb16563424887ed3aa5b6b74e4b713c8e93a10c376d954ff3041442716bdf9ee28fab2ea09f04d07e3366f834ea472c19820b7337eb27a@172.104.130.233:30313",
	"enode://799d0a8836e17ef7fcc58b3d5ced5bb1fe474b31a09851f938d381f4556fa8954ca308f6a178d22ed56769a8b878ac8f9cc62c889f9cafab45a3bd4f6024bb29@172.104.68.7:30313",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303",    // US-Azure geth
	"enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:30303",  // US-Azure parity
	"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", // Parity
	"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303", // @gpip
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://06051a5573c81934c9554ef2898eb13b33a34b94cf36b202b69fde139ca17a85051979867720d4bdae4323d4943ddf9aeeb6643633aa656e0be843659795007a@35.177.226.168:30303",
	"enode://0cc5f5ffb5d9098c8b8c62325f3797f56509bff942704687b6530992ac706e2cb946b90a34f1f19548cd3c7baccbcaea354531e5983c7d1bc0dee16ce4b6440b@40.118.3.223:30304",
	"enode://1c7a64d76c0334b0418c004af2f67c50e36a3be60b5e4790bdac0439d21603469a85fad36f2473c9a80eb043ae60936df905fa28f1ff614c3e5dc34f15dcd2dc@40.118.3.223:30306",
	"enode://85c85d7143ae8bb96924f2b54f1b3e70d8c4d367af305325d30a61385a432f247d2c75c45c6b4a60335060d072d7f5b35dd1d4c45f76941f62a4f83b6e75daaf@40.118.3.223:30307",
}
