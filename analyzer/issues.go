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
		// L001 - Unsafe ERC20 Operation(s)
		{
			"L001",
			LOW,
			"L-X: Unsafe use of transfer()/transferFrom() with IERC20",
			"Some tokens do not implement the ERC20 standard properly but are still accepted by most code that accepts ERC20 tokens.\nexample Tether (USDT)'s transfer() and transferFrom() functions do not return booleans as the specification requires, and instead have no return value. When these sorts of tokens are cast to IERC20, their function signatures do not match and therefore the calls made, revert.",
			"Use OpenZeppelin's `SafeERC20` safeTransfer()/safeTransferFrom() instead",
			`\.transfer\(|\.transferFrom\(|\.approve\(`, // ".tranfer(", ".transferFrom(" or ".approve("
		},
		{
			"L002",
			GASOP,
			"L-X: `ecrecover()` not checked for signer address of zero",
			"The `ecrecover()` function returns an address of zero when the signature does not match. This can cause problems if address zero is ever the owner of assets, and someone uses the permit function on address zero. ",
			"You should check whether the return value is address(0) or not in terms of the security.",
			"THIS_IS_TARGET", //ガンバ！
		},
		{
			"L003",
			GASOP,
			"L-X: `_safeMint()` should be used rather than `_mint()` wherever possible",
			"`_mint()` is discouraged in favor of `_safeMint()` which ensures that the recipient is either an EOA or implements `IERC721Receiver`. ",
			"You should change from` _mint` to `_safeMint` in terms of security.",
			"_mint",
		},
		{
			"L004",
			GASOP,
			"L-X: Vulnerable to cross-chain replay attacks due to static DOMAIN_SEPARATOR/domainSeparator",
			"Should Ethereum fork in the feature, the chainid will change however the one used by the permits will not enabling a user to use any new permits on both chains thus breaking the token on the forked chain permanently.",
			"Please consult EIP1344 for more details: https://eips.ethereum.org/EIPS/eip-1344#rationale\nAnd the mitigation action that should be applied is the calculation of the chainid dynamically on each permit invocation. As a gas optimization, the deployment pre-calculated hash for the permits can be stored to an immutable variable and a validation can occur on the permit function that ensure the current chainid is equal to the one of the cached hash and if not, to re-calculate it on the spot.",
			"EIP712_DOMAIN_TYPEHASH",
		},
		{
			"L005",
			GASOP,
			"L-X: Unused/empty `receive()`/`fallback()` function",
			"If the intention is for the Ether to be used, the function should call another function, otherwise, it should revert ",
			"This is one example to fix this problem\ne.g. `require(msg.sender == address(weth))`",
			"(fallback() external payable {} | fallback() external payable {}| fallback() external {})",
		},
		{
			"L006",
			GASOP,
			"L-X: `abi.encodePacked()`should not be used in `keccak256()`",
			"You have a risk of hash collisions.",
			"You should use `abi.encode()` instead which will pad items to 32 bytes. Unless there is a compelling reason, `abi.encode` should be preferred. \nIf there is only one argument to `abi.encodePacked()` it can often be cast to bytes() or bytes32() instead.",
			"keccak256(abi.encodePacked(",
		},
		{
			"L007",
			GASOP,
			"L-X: rescueETH() cannot rescue Ether",
			"rescueETH() sends `msg.value` to the destination address, which means it requires the caller of `rescueETH()` to provide the Ether to send. Essentially the owner is directly paying the destination address, and the Ether in the contract remains untouched",
			"You should discuss how to solve this problem with your team.",
			"rescueETH",
		},
		{
			"L008",
			GASOP,
			"L-X: Check zero denominator",
			"When a division is computed, it must be ensured that the denominator is non-zero to prevent failure of the function call.",
			"Before doing these computations, add a non-zero check to these variables. Or alternatively, add a non-zero check inupdatePenalties().",
			"/STRING",
		},
		{
			"L009",
			GASOP,
			"L-X: `require()` should be used instead of `assert()`",
			"Prior to solidity version 0.8.0, hitting an assert consumes the remainder of the transaction's available gas rather than returning it, as `require()`/`revert()` do. `assert()` should be avoided even past solidity version 0.8.0 as its [documentation](https://docs.soliditylang.org/en/v0.8.14/control-structures.html#panic-via-assert-and-error-via-require) states that The assert function creates an error of type Panic(uint256). ... Properly functioning code should never create a Panic, not even on invalid external input. If this happens, then there is a bug in your contract which you should fix",
			"You should change from `assert()` to `require()`",
			"THIS_IS_TARGET",
		},
		{
			"L0010",
			GASOP,
			"L-X: pragma experimental ABIEncoderV2 is deprecated",
			"The `pragma experimental ABIEncoderV2` is deprecated version.",
			"You should use `pragma abicoder v2` instead instead of `pragma experimental ABIEncoderV2`",
			"pragma experimental ABIEncoderV2",
		},
		{
			"N001",
			LOW,
			"N-X: Do not use Deprecated Library Functions",
			"[Deprecated](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/bfff03c0d2a59bcd8e2ead1da9aed9edf0080d05/contracts/token/ERC20/utils/SafeERC20.sol#L38-L44) in favor of safeIncreaseAllowance() and safeDecreaseAllowance(). If only setting the initial allowance to the value that means infinite, safeIncreaseAllowance() can be used instead",
			"Use `safeIncreaseAllowance()` and `safeDecreaseAllowance()` instead of `safeApprove`",
			`_setupRole\(|safeApprove\(`, // _setupRole and safeApprove are common deprecated lib functions
		},
		{
			"N002",
			LOW,
			"N-X: Consider addings checks for signature malleability",
			"You should consider addings checks for signature malleability",
			"Use OpenZeppelin's `ECDSA` contract rather than calling `ecrecover()` directly",
			" ecrecover",
		},
		{
			"N003",
			LOW,
			"N-X: `require()/revert()` statements should have descriptive reason strings",
			"If there is no statement, the user and frontend have no way to get the reason transaction failed.",
			"You should add the statement.",
			"require(_success);",//No string
		},
		{
			"N004",
			LOW,
			"N-X: `type(uint<n>).max` should be used instead of`uint<n>(-1)`",
			"Earlier versions of solidity can use `uint<n>(-1)` instead. Expressions not including the `- 1` can often be re-written to accomodate the change (e.g. by using a `>` rather than a `>=`, which will also save some gas)",
			"You should change from `uint<n>(-1)` to `type(uint<n>).max`",
			"uint<n>(-1)", //nに数字
		},
		{
			"N005",
			LOW,
			"N-X: Non-assembly method available",
			"You don't need to use assembly. `assembly` should be avoided as much as possible due to various risks",
			"You can change as follows.\n `assembly{ id := chainid() } => uint256 id = block.chainid`, `assembly { size := extcodesize() } => uint256 size = address().code.length`",
			"(id := chainid | size := extcodesize)",
		},
		{
			"N006",
			LOW,
			"N-X: Constants should be defined rather than using magic numbers",
			"Even [`assembly`](https://github.com/code-423n4/2022-05-opensea-seaport/blob/9d7ce4d08bf3c3010304a0476a785c70c0e90ae7/contracts/lib/TokenTransferrer.sol#L35-L39) can benefit from using readable constants instead of hex/numeric literals.\nAlso, it is bad practice to use numbers directly in code without explanation.",
			"You should declare the variable instead of magic number.",
			"NUMBERを検出",// むずいね
		},
		{
			"N007",
			LOW,
			"N-X: Missing event and or timelock for critical parameter change",
			"Events help non-contract tools to track changes, and events prevent users from being surprised by changes",
			"You should add the event in the function to change the critical parameter.",
			"set { non-event }",
		},
		{
			"N008",
			LOW,
			"N-X: Expressions for constant values such as a call to `keccak256()`, should use `immutable` rather than `constant`",
			"Expressions for constant values such as a call to `keccak256()`, should use `immutable` rather than `constant`",
			"You should use immutable instead of constant",
			"constant ** = keccak256",
		},
		{
			"N009",
			LOW,
			"N-X: Non-library/interface files should use fixed compiler versions, not floating ones",
			"Non-library/interface files should use fixed compiler versions, not floating ones",
			"Delete the floating keyword `^`.",
			"pragma solidity (\\^|>)", // "pragma solidity ^" or "pragma solidity >"
		},
		{
			"N010",
			LOW,
			"N-X: Use scientific notation (e.g.`1e18`) rather than exponentiation (e.g.`10**18`)",
			"Use scientific notation (e.g.`1e18`) rather than exponentiation (e.g.`10**18`)",
			"Use scientific notation instead of exponentiation.",
			"10**",
		},
		{
			"N011",
			LOW,
			"N-X: Variable names that consist of all capital letters should be reserved for `constant/immutable` variables",
			"Variable names that consist of all capital letters should be reserved for `constant/immutable` variables",
			"Variables that are not constant/immutable should be declared in lower case",
			"大文字",
		},
		{
			"N012",
			LOW,
			"N-X: Open Todos",
			"Code architecture, incentives, and error handling/reporting questions/issues should be resolved before deployment",
			"Delete TODO keyword",
			"TODO",
		},
		{
			"N013",
			LOW,
			"N-X: No use of two-phase ownership transfers",
			"Consider adding a two-phase transfer, where the current owner nominates the next owner, and the next owner has to call `accept*()` to become the new owner. This prevents passing the ownership to an account that is unable to use it.",
			"Consider implementing a two step process where the admin nominates an account and the nominated account needs to call an acceptOwnership() function for the transfer of admin to fully succeed. This ensures the nominated EOA account is a valid and active account.",
			"(admin = | owner = |Ownable)",
		},
		{
			"N014",
			LOW,
			"N-X: Return values of `approve()` not checked",
			"Not all `IERC20` implementations `revert()` when there's a failure in `transfer()/transferFrom()`. The function signature has a `boolean` return value and they indicate errors that way instead. By not checking the return value, operations that should have marked as failed, may potentially go through without actually making a payment",
			"You should check whether the return value is true or false in terms of the security.",
			".approve",// returnを確認していない
		},
		{
			"N015",
			LOW,
			"N-X: Use a more recent version of solidity",
			"Use a solidity version of at least 0.8.4 to get bytes.concat() instead of abi.encodePacked (<bytes>, <bytes>)\nUse a solidity version of at least 0.8.12 to get string.concat() instead of abi.encodePacked (<str>, <str>)\nUse a solidity version of at least 0.8.13 to get the ability to use using for with a list of free functions",
			"Use more recent version of solidity.",
			"0.6, 0.7, 0.8.0~9",
		},
		{
			"N016",
			LOW,
			"N-X: Use `string.concat()` or`bytes.concat()`",
			"Solidity version 0.8.4 introduces `bytes.concat()` (vs `abi.encodePacked(<bytes>,<bytes>)`)Solidity version 0.8.12 introduces `string.concat()`(vs `abi.encodePacked(<str>,<str>)`)",
			"Use concat instead of abi.encodePacked",
			"abi.encodePacked(",
		},
		{
			"N017",
			LOW,
			"N-X: approve should be replaced with safeIncreaseAllowance() / safeDecreaseAllowance()",
			"approve is subject to a known front-running attack. ",
			"Consider using safeIncreaseAllowance () or safeDecreaseAllowance() instead.",
			".approve",
		},
	}
}
