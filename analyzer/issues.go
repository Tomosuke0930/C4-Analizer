package analyzer

// AllIssues returns the list of all issues.
func AllIssues() []Issue {
	return append(GasOpIssues(), LowRiskIssues()...)
}

// GasOpIssues returns the list of all gas optimization issues.
func GasOpIssues() []Issue {
	return []Issue{
		// G001 - Don't Initialize Variables with Default Value
		{
			"G001",
			GASOP,
			"G-X: Don't Initialize Variables with Default Value",
			"Uninitialized variables are assigned with the types default value. Explicitly initializing a variable with it's default value costs unnecesary gas. Not overwriting the default for [stack variables](https://gist.github.com/IllIllI000/e075d189c1b23dce256cd166e28f3397) saves 8 gas. Storage and memory variables have larger savings",
			"Delete useless variable declarations to save gas.",
			`(uint[0-9]*[[:blank:]][a-z,A-Z,0-9]*.?=.?0;)|(bool.[a-z,A-Z,0-9]*.?=.?false;)`,
		},
		// G002 - Cache Array Length Outside of Loop
		{
			"G002",
			GASOP,
			"G-X: Cache Array Length Outside of Loop",
			"Caching the array length outside a loop saves reading it on each iteration, as long as the array's length is not changed during the loop.\nYou save 3 gas by not reading `array.length` - 3 gas per instance - 27 gas saved",
			"Reco",
			".length",
		},
		// G003 - Use != 0 instead of > 0 for Unsigned Integer Comparison
		{
			"G003",
			GASOP,
			"G-X: Use != 0 instead of > 0 for Unsigned Integer Comparison",
			"Use != 0 when comparing uint variables to zero, which cannot hold values below zero",
			"You should change from `> 0` to  `!=0`.",
			"(>0|> 0)",
		},
		// G007 - Long Revert Strings
		{
			"G007",
			GASOP,
			"G-X: Long Revert Strings",
			"Shortening revert strings to fit in 32 bytes will decrease gas costs for deployment and gas costs when the revert condition has been met.",
			"Use short statements instead of long statement.",
			"\".{33,}\"", // Anything between "'s with at least 33 characters
		},
		// G008 - Use Shift Right/Left instead of Division/Multiplication if possible
		{
			"G008",
			GASOP,
			"G-X: Use Shift Right/Left instead of Division/Multiplication if possible",
			"A division/multiplication by any number `x` being a power of 2 can be calculated by shifting `log2(x)` to the right/left.\nWhile the `DIV` opcode uses 5 gas, the `SHR` opcode only uses 3 gas.\nurthermore, Solidity's division operation also includes a division-by-0 prevention which is bypassed using shifting.",
			"Reco",
			`(/[2,4,8]|/ [2,4,8]|\*[2,4,8]|\* [2,4,8])`,
		},
		{
			"G009",
			GASOP,
			"G-X: Use assembly balance ETH",
			"You can save 159 gas per instance if using assembly to check for balance of address",
			"Please follow [this link](https://gist.github.com/Tomosuke0930/4547e012ece4575e88e2a6d2baf498fa) to make corrections",
			").balance", 
		},
		{
			"G010",
			GASOP,
			"G-X: Use assembly to check for address(0)",
			"See [this issue](https://github.com/code-423n4/2022-06-putty-findings/issues/59)\nYou can save 16000 gas per instance in deploying contracrt.\nYou can save about 6 gas per instance if using assembly to check for address (0)",
			"Please follow [this link](https://github.com/code-423n4/2022-06-putty-findings/issues/59) to make corrections",
			"address(0)",
		},
		{
			"G011",
			GASOP,
			"G-X: Use assembly for calculation (add, sub, mul, div)",
			"See [this issue](https://github.com/code-423n4/2022-06-putty-findings/issues/59)\nYou can save about 4000 gas per instance in deploying contracrt.\nYou can save about 40~60  gas per instance if using assembly to execute the function.",
			"Please follow [this link](https://gist.github.com/Tomosuke0930/a18e03c4ef20ff284dfca62e7ccf6f91) to make corrections",
			"(>0|> 0)", //数字とかの計算についてはわからない +, -, / , *
		},
		{
			"G012",
			GASOP,
			"G-X: Use assembly to hash",
			"See [this issue](https://github.com/code-423n4/2022-06-putty-findings/issues/59)\nYou can save about 5000 gas per instance in deploying contracrt.\nYou can save about 80 gas per instance if using assembly to execute the function.",
			"Please follow [this link](https://gist.github.com/Tomosuke0930/72beb469f08c954b232e58b8da03ef28) to make corrections",
			"keccak256",
		},
		{
			"G013",
			GASOP,
			"G-X: TITLE",
			"Description",
			"Reco",
			"afdafadfdafadfdsa",
		},
	}
}

// LowRiskIssues returns the list of all low risk issues.
func LowRiskIssues() []Issue {
	return []Issue{
		// // L001 - Unsafe ERC20 Operation(s)
		// {
		// 	"L001",
		// 	LOW,
		// 	"Unsafe ERC20 Operation(s)",
		// 	"https://github.com/byterocket/c4-common-issues/blob/main/2-Low-Risk.md#l001---unsafe-erc20-operations",
		// 	"Reco",
		// 	`\.transfer\(|\.transferFrom\(|\.approve\(`, // ".tranfer(", ".transferFrom(" or ".approve("
		// },
		// // L003 - Unspecific Compiler Version Pragma
		// {
		// 	"L003",
		// 	LOW,
		// 	"Unspecific Compiler Version Pragma",
		// 	"https://github.com/byterocket/c4-common-issues/blob/main/2-Low-Risk.md#l003---unspecific-compiler-version-pragma",
		// 	"Reco",
		// 	"pragma solidity (\\^|>)", // "pragma solidity ^" or "pragma solidity >"
		// },
		// // L005 - Do not use Deprecated Library Functions
		// {
		// 	"L005",
		// 	LOW,
		// 	"Do not use Deprecated Library Functions",
		// 	"https://github.com/byterocket/c4-common-issues/blob/main/2-Low-Risk.md#l005---do-not-use-deprecated-library-functions",
		// 	"Reco",
		// 	`_setupRole\(|safeApprove\(`, // _setupRole and safeApprove are common deprecated lib functions
		// },
	}
}
