package users

import "math/rand"

var tokens=[] string{
	"1072d95046f18e67463ce40d645a9b8d",
	"a6484e08c7c9e7fd9bec784163b9ef18",
	"b2b8f791570d96f6cf46898747e7ce1a",
	"9aebc7723eee1ef07a73f5ddfc0c1df4",
	"eac0051c446319b99c65da060f3e83e4",
	"86c9e10dfe14ced9f43901b6f4c9e983",
	"85b88ce10c15f390ee75bf571688b3b7",
	"03a74352f8945d4e011b1914e0527514",
	"9024bb5d2e154746bb513878231cc0cf",
	"c1d6b49adb2ba817109873dbc13becb4",
	"28b495e4df789d971d2ae77b01a55a55",
	"997487c2aa6dc93d84169f293ae2073d",
	"1dfcf1a2b70604242eb5c6abe8ee9703",
	"19a3cd1b08ddbfeb3701f58b4843bac0",
	"9ce21c72792f415f5ce0b96e29e00d52",
	"b0ad70dc306d789204ddb4ec0b7c2b4d",
	"854a6c21a3f15f27dab5b2d676ad2321",
	"037732e0f605c2b98884616a4584d38f",
	"b976a2cd81f5fab373ced07d17b9aa81",
	"a4cccd543a7a9fbe8b85f04746eb2753",
}

func Tokens() []string {
	return tokens
}

func RandomToken() string{
	return  tokens[rand.Intn(len(tokens))]
}
