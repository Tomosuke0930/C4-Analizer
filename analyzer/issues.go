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
			`.length; `,// 改善は必要 
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
			"\".{33,}\"", // Anything between "'s with at least 33 characters　requireとかにしよう！
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
			"fdafdafdasfadsfdasf", //数字とかの計算についてはわからない +, -, / , *
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
			"G-X: Don't use SafeMath lib 0.8.0~",
			"Version 0.8.0~, it can calculate as a ```SafeMath``` without using it.",
			"Delete the import to save gas.",
			"SafeMath",
		},
		{
			"G014",
			GASOP,
			"G-X: Use abi.encodePacked() not abi.encode()",
			"`abi.encode` pads extra null bytes at the end of the call data which is normally unnecessary. In general, `abi.encodePacked` is more gas-efficient.",
			"Changing abi.encode to abi.encodePacked can save gas.",
			"abi.encode(",
		},
		{
			"G015",
			GASOP,
			"G-X: Solidity Version should be the latest",
			"Use a solidity version of at least 0.8.0 to get overflow protection without SafeMath\nUse a solidity version of at least 0.8.2 to get simple compiler automatic inlining\nUse a solidity version of at least 0.8.3 to get better struct packing and cheaper multiple storage reads\nUse a solidity version of at least 0.8.4 to get custom errors, which are cheaper at deployment than revert()/require() strings\nUse a solidity version of at least 0.8.10 to have external calls skip contract existence checks if the external call has a return value \n",
			"You should consider with your team members whether you should choose the latest version.\nThe latest version has its advantages, but also it has disadvantages, such as the possibility of unknown bugs.",
			"0.6, 0.7, 0.8.0~9",
		},
		{
			"G016",
			GASOP,
			"G-X: Empty blocks should be removed or emit something",
			"The code should be refactored such that they no longer exist, or the block should do something useful, such as emitting an event or reverting.",
			"Empty blocks should be removed or emit something (The code should be refactored such that they no longer exist, or the block should do something useful, such as emitting an event or reverting.",
			"{}",
		},
		{
			"G017",
			GASOP,
			"G-X: Functions guaranteed to revert when called by normal users can be marked payable",
			"See [this issue](https://github.com/code-423n4/2022-06-putty-findings/issues/59) \nIf a function modifier such as `onlyOwner` is used, the function will revert if a normal user tries to pay the function. Marking the function as payable will lower the gas cost for legitimate callers because the compiler will not include checks for whether a payment was provided. The extra opcodes avoided are CALLVALUE(2),DUP1(3),ISZERO(3),PUSH2(3),JUMPI(10),PUSH1(3),DUP1(3),REVERT(0),JUMPDEST(1),POP(2), which costs an average of about 21 gas per call to the function, in addition to the extra deployment cost \nSaves 2400 gas per instance in deploying contracrt.\nSaves about 20 gas per instance if using assembly to executing the function.",
			"You should add the keyword `payable` to those functions.",
			"only",
		},
		{
			"G018",
			GASOP,
			"G-X: Use custom errors rather than revert()/require() strings to save gas",
			"See [this issue](https://github.com/code-423n4/2022-06-putty-findings/issues/59).\nYou can save 7000 gas per instance in deploying contracrt.\nYou can save about 60 gas per instance if using assembly to execute the function.",
			"You should change require to if&custom error in checking input value etc… \nPlease follow [this link](https://gist.github.com/Tomosuke0930/55d137c504d88f17e89e91111e22dcd6) to make corrections",
			"require(",
		},
		{
			"G019",
			GASOP,
			"G-X: Splitting require() statements that use && saves gas",
			"See [this issue](https://github.com/code-423n4/2022-01-xdefi-findings/issues/128) which describes the fact that there is a larger deployment gas cost, but with enough runtime calls, the change ends up being cheaper",
			"You should change one require which has `&&` to two simple require() statements to save gas",
			"require( &&)",
		},
		{
			"G020",
			GASOP,
			"G-X: Use ++i instead of i++",
			"You can get cheaper for loops (at least 25 gas, however can be up to 80 gas under certain conditions), by rewriting\nThe post-increment operation (i++) boils down to saving the original value of i, incrementing it and saving that to a temporary place in memory, and then returning the original value; only after that value is returned is the value of i actually updated (4 operations).On the other hand, in pre-increment doesn't need to store temporary(2 operations) so, the pre-increment operation uses less opcodes and is thus more gas efficient.",
			"You should change from i++ to ++i.",
			"i++, i ++,index ++, index++,",// stringとか--とか
		},
		{
			"G021",
			GASOP,
			"G-X: Use x=x+y instad of x+=y",
			"You can save about 35 gas per instance if you change from `x+=y**`** to `x=x+y`\nThis works equally well for subtraction, multiplication and division.",
			"You should change from `x+=y` to `x=x+y`.",
			"+=, -= *=, /=",
		},
		{
			"G022",
			GASOP,
			"G-X: Use `uint256` instead of `bool`",
			"Booleans are more expensive than uint256 or any type that takes up a full word because each write operation emits an extra SLOAD to first read the slot's contents, replace the bits taken up by the boolean, and then write back. This is the compiler's defense against contract upgrades and pointer aliasing, and it cannot be disabled.",
			"You should change from bool to uint256 to save gas",
			"bool",
		},
		{
			"G023",
			GASOP,
			"G-X: Use `indexed` for uint, bool, and address",
			"Using the indexed keyword for value types such as uint, bool, and address saves gas costs, as seen in the example below. However, this is only the case for value types, whereas indexing bytes and strings are more expensive than their unindexed version.\nAlso indexed keyword has more merits.\nIt can be useful to have a way to monitor the contract's activity after it was deployed. One way to accomplish this is to look at all transactions of the contract, however that may be insufficient, as message calls between contracts are not recorded in the blockchain. Moreover, it shows only the input parameters, not the actual changes being made to the state. Also events could be used to trigger functions in the user interface.",
			"Up to three `indexed` can be used per event and should be added.",
			"event( indexed )",//non indexed
		},
		// {
		// 	"G024",
		// 	GASOP,
		// 	"G-X: TILE",
		// 	"DESCRIPTION",
		// 	"RECO",
		// 	"THIS_IS_TARGET",
		// },
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
