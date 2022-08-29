# Gas | QA

## G-X: Don't Initialize Variables with Default Value

### Description

Issue Information: [G001](https://github.com/byterocket/c4-common-issues/blob/main/0-Gas-Optimizations.md#g001---dont-initialize-variables-with-default-value)

### Findings:

[for (uint256 i = 0; i < reqLength; ) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/Kernel.sol#L397)

[for (uint i = 0; i < 90; i++) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/scripts/Deploy.sol#L239)

[for (uint256 i = 0; i < userNum; i++) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/UserFactory.sol#L25)

[for (uint256 i = 0; i < 5; ) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/utils/KernelUtils.sol#L43)

[for (uint256 i = 0; i < 32; ) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/utils/KernelUtils.sol#L58)

## G-X: Cache Array Length Outside of Loop

### Description

Issue Information: [G002](https://github.com/byterocket/c4-common-issues/blob/main/0-Gas-Optimizations.md#g002---cache-array-length-outside-of-loop)

### Findings:

[getPolicyIndex[policy\_] = activePolicies.length - 1;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/Kernel.sol#L300)

[uint256 depLength = dependencies.length;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/Kernel.sol#L304)

[getDependentIndex[keycode][policy\_] = moduleDependents[keycode].length - 1;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/Kernel.sol#L310)

[Policy lastPolicy = activePolicies[activePolicies.length - 1];](https://github.com/code-423n4/2022-08-olympus/blob/main/src/Kernel.sol#L334)

[uint256 keycodeLen = allKeycodes.length;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/Kernel.sol#L352)

[uint256 policiesLen = activePolicies.length;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/Kernel.sol#L361)

[uint256 depLength = dependents.length;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/Kernel.sol#L380)

[uint256 reqLength = requests\_.length;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/Kernel.sol#L396)

[uint256 depcLength = dependencies.length;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/Kernel.sol#L411)

[Policy lastPolicy = dependents[dependents.length - 1];](https://github.com/code-423n4/2022-08-olympus/blob/main/src/Kernel.sol#L418)

[revert("ECDSA: invalid signature length");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L105)

[// Check the signature length](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L138)

[if (signature.length == 65) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L141)

[} else if (signature.length == 64) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L153)

[// 32 is the length in bytes of hash,](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L285)

[/// @dev Should be used if length exceeds max to query entire array](https://github.com/code-423n4/2022-08-olympus/blob/main/src/interfaces/IBondAggregator.sol#L65)

[/// @dev 8. Is fixed term ? Vesting length (seconds) : Vesting expiration (timestamp).](https://github.com/code-423n4/2022-08-olympus/blob/main/src/interfaces/IBondAuctioneer.sol#L36)

[require(success && (data.length == 0 || abi.decode(data, (bool))), "TRANSFER_FROM_FAILED");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/TransferHelper.sol#L20)

[require(success && (data.length == 0 || abi.decode(data, (bool))), "TRANSFER_FAILED");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/TransferHelper.sol#L32)

[require(success && (data.length == 0 || abi.decode(data, (bool))), "APPROVE_FAILED");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/TransferHelper.sol#L44)

[uint256 length = instructions\_.length;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/modules/INSTR.sol#L43)

[if (length == 0) revert INSTR_InstructionsCannotBeEmpty();](https://github.com/code-423n4/2022-08-olympus/blob/main/src/modules/INSTR.sol#L48)

[for (uint256 i; i < length; ) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/modules/INSTR.sol#L50)

[} else if (instruction.action == Actions.ChangeExecutor && i != length - 1) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/modules/INSTR.sol#L61)

[/// @param startObservations\_ - Array of observations to initialize the moving average with. Must be of length numObservations.](https://github.com/code-423n4/2022-08-olympus/blob/main/src/modules/PRICE.sol#L201)

[uint256 numObs = observations.length;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/modules/PRICE.sol#L212)

[if (startObservations*.length != numObs || lastObservationTime* > uint48(block.timestamp))](https://github.com/code-423n4/2022-08-olympus/blob/main/src/modules/PRICE.sol#L215)

[uint256 len = tokens\_.length;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/BondCallback.sol#L155)

[if (instructions.length == 0) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Governance.sol#L188)

[for (uint256 step; step < instructions.length; ) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Governance.sol#L278)

[/// @param startObservations\_ Array of observations to initialize the moving average with. Must be of length numObservations.](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/PriceConfig.sol#L41)

[uint256 len = tokens\_.length;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/TreasuryCustodian.sol#L58)

[uint256 len = requests\_.length;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/ModuleTestFixtureGenerator.sol#L18)

[uint256 len = \_requests.length;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/ModuleTestFixtureGenerator.sol#L36)

[uint256 num = selectors.length;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/ModuleTestFixtureGenerator.sol#L63)

[uint256 len = mkts.length;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/BondAggregator.sol#L156)

[uint256 len = forPayout.length;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/BondAggregator.sol#L182)

[uint256 len = ids.length;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/BondAggregator.sol#L213)

[uint256 len = tokenIds\_.length;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/BondFixedTermTeller.sol#L160)

[length: secondsToConclusion,](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseCDA.sol#L199)

[// Initial target debt is equal to capacity scaled by the ratio of the debt decay interval and the length of the market.](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseCDA.sol#L216)

[uint256(meta.length) - timeRemaining,](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseCDA.sol#L543)

[uint256(meta.length)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseCDA.sol#L544)

[// Calculate target debt from the timeNeutralCapacity and the ratio of debt decay interval and the length of the market](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseCDA.sol#L564)

[uint256(meta.length)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseCDA.sol#L567)

[uint256 len = tokens\_.length;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L101)

[/// @dev Should be used if length exceeds max to query entire array](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/interfaces/IBondAggregator.sol#L65)

[/// @dev 8. Is fixed term ? Vesting length (seconds) : Vesting expiration (timestamp).](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/interfaces/IBondAuctioneer.sol#L36)

[uint48 vesting; // length of time from deposit to maturity if fixed-term, vesting timestamp if fixed-expiry](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/interfaces/IBondCDA.sol#L28)

[uint32 length; // time from creation to conclusion.](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/interfaces/IBondCDA.sol#L37)

[// l = length of program](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/interfaces/IBondCDA.sol#L75)

[to.code.length == 0](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L65)

[uint256 idsLength = ids.length; // Saves MLOADs.](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L80)

[require(idsLength == amounts.length, "LENGTH_MISMATCH");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L82)

[// An array can't have a total length](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L97)

[to.code.length == 0](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L107)

[uint256 ownersLength = owners.length; // Saves MLOADs.](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L126)

[require(ownersLength == ids.length, "LENGTH_MISMATCH");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L128)

[to.code.length == 0](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L167)

[uint256 idsLength = ids.length; // Saves MLOADs.](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L186)

[require(idsLength == amounts.length, "LENGTH_MISMATCH");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L188)

[// An array can't have a total length](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L193)

[to.code.length == 0](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L203)

[uint256 idsLength = ids.length; // Saves MLOADs.](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L221)

[require(idsLength == amounts.length, "LENGTH_MISMATCH");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L223)

[// An array can't have a total length](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L228)

[uint256 len = response.length;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/quabi/Quabi.sol#L35)

[uint256 len = requests\_.length;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/mocks/MockModuleWriter.sol#L15)

[uint256 len = \_requests.length;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/mocks/MockModuleWriter.sol#L30)

[if (output.length == 0) revert();](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/mocks/MockModuleWriter.sol#L42)

## G-X: Use != 0 instead of > 0 for Unsigned Integer Comparison

### Description

Issue Information: [G003](https://github.com/byterocket/c4-common-issues/blob/main/0-Gas-Optimizations.md#g003---use--0-instead-of--0-for-unsigned-integer-comparison)

### Findings:

[if (uint256(s) > 0x7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF5D576E7357A4501DDFE92F46681B20A0) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L245)

[require(b > 0, errorMessage);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L611)

[require(denominator > 0);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L35)

[if (mulmod(a, b, denominator) > 0) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L122)

[if (userVotesForProposal[activeProposal.proposalId][msg.sender] > 0) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Governance.sol#L247)

[vm.assume(amount\_ > 0);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L98)

[vm.assume(amount\_ > 0);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L108)

[vm.assume(amount\_ > 0);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L126)

[vm.assume(amount\_ > 0);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L143)

[if (char < 0x41 || char > 0x5A) revert InvalidKeycode(keycode\_); // A-Z only](https://github.com/code-423n4/2022-08-olympus/blob/main/src/utils/KernelUtils.sol#L46)

[if ((char < 0x61 || char > 0x7A) && char != 0x5f && char != 0x00) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/utils/KernelUtils.sol#L60)

## G-X: Use immutable for OpenZeppelin AccessControl's Roles Declarations

### Description

Issue Information: [G006](https://github.com/byterocket/c4-common-issues/blob/main/0-Gas-Optimizations.md#g006---use-immutable-for-openzeppelin-accesscontrols-roles-declarations)

### Findings:

[return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L287)

[return keccak256(abi.encodePacked("\x19\x01", domainSeparator, structHash));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L304)

[\* they need in their contracts using a combination of `abi.encode` and `keccak256`.](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L315)

[bytes32 hashedName = keccak256(bytes(name));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L360)

[bytes32 hashedVersion = keccak256(bytes(version));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L361)

[bytes32 typeHash = keccak256(](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L362)

[return keccak256(abi.encode(typeHash, nameHash, versionHash, chainID, address(this)));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L398)

[\* bytes32 digest = \_hashTypedDataV4(keccak256(abi.encode(](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L408)

[\* keccak256("Mail(address to,string contents)"),](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L409)

[\* keccak256(bytes(mailContents))](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L411)

[bytes32 private constant ERC20TOKEN_ERC1820_INTERFACE_ID = keccak256("ERC20Token");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L665)

[keccak256(](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L835)

[bytes32 structHash = keccak256(](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L860)

[address(bytes20(uint160(uint256(keccak256("hevm cheat code")))));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/UserFactory.sol#L9)

[bytes32 internal nextUser = keccak256(abi.encodePacked("users"));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/UserFactory.sol#L13)

[nextUser = keccak256(abi.encodePacked(nextUser));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/UserFactory.sol#L18)

[keccak256(abi.encodePacked(underlying*, expiry* / uint48(1 days)))](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/BondFixedTermTeller.sol#L231)

[address(bytes20(uint160(uint256(keccak256("hevm cheat code")))));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/larping.sol#L9)

[Vm internal constant vm = Vm(address(bytes20(uint160(uint256(keccak256("hevm cheat code"))))));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/quabi/Quabi.sol#L8)

[change = int256(uint256(keccak256(abi.encodePacked(nonce, i)))) %!i(MISSING)nt256(1000);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/PRICE.t.sol#L101)

[change = int256(uint256(keccak256(abi.encodePacked(nonce, i)))) %!i(MISSING)nt256(1000);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/PRICE.t.sol#L128)

[change = int256(uint256(keccak256(abi.encodePacked(nonce, i)))) %!i(MISSING)nt256(1000);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/PriceConfig.t.sol#L124)

## G-X: Long Revert Strings

### Description

Issue Information: [G007](https://github.com/byterocket/c4-common-issues/blob/main/0-Gas-Optimizations.md#g007---long-revert-strings)

### Findings:

[revert("ECDSA: invalid signature 's' value");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L107)

[revert("ECDSA: invalid signature 'v' value");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L109)

["EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L363)

[require(c / a == b, "SafeMath: multiplication overflow");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L597)

[\_allowances[sender][msg.sender].sub(amount, "ERC20: transfer amount exceeds allowance")](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L738)

["ERC20: decreased allowance below zero"](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L758)

[require(sender != address(0), "ERC20: transfer from the zero address");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L769)

[require(recipient != address(0), "ERC20: transfer to the zero address");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L770)

[\_balances[sender] = \_balances[sender].sub(amount, "ERC20: transfer amount exceeds balance");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L774)

[require(account != address(0), "ERC20: burn from the zero address");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L788)

[\_balances[account] = \_balances[account].sub(amount, "ERC20: burn amount exceeds balance");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L792)

[require(owner != address(0), "ERC20: approve from the zero address");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L802)

[require(spender != address(0), "ERC20: approve to the zero address");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L803)

["Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)"](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L836)

["ERC20: burn amount exceeds allowance"](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L925)

[import {AggregatorV2V3Interface} from "interfaces/AggregatorV2V3Interface.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/modules/PRICE.sol#L4)

[import {ReentrancyGuard} from "solmate/utils/ReentrancyGuard.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/modules/TRSRY.sol#L5)

[ERC20("OlympusDAO Dummy Voting Tokens", "VOTES", 0)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/modules/VOTES.sol#L18)

[import {ReentrancyGuard} from "solmate/utils/ReentrancyGuard.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/BondCallback.sol#L5)

[import {ReentrancyGuard} from "solmate/utils/ReentrancyGuard.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Heart.sol#L4)

[import {IOperator} from "policies/interfaces/IOperator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Heart.sol#L7)

[import {ReentrancyGuard} from "solmate/utils/ReentrancyGuard.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Operator.sol#L4)

[import {IOperator} from "policies/interfaces/IOperator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Operator.sol#L7)

[import {AggregatorV2V3Interface} from "interfaces/AggregatorV2V3Interface.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/scripts/Deploy.sol#L4)

[console2.log("RESERVE-ETH Price Feed deployed to:", address(reserveEthPriceFeed));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/scripts/Deploy.sol#L271)

[err = abi.encodeWithSignature("InvalidKeycode(bytes5)", Keycode.wrap("inval"));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L77)

[err = abi.encodeWithSignature("InvalidKeycode(bytes5)", Keycode.wrap(""));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L81)

[err = abi.encodeWithSignature("InvalidRole(bytes32)", Role.wrap("INVALID_ID"));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L89)

[err = abi.encodeWithSignature("InvalidKeycode(bytes5)", Keycode.wrap("badkc"));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L150)

["Kernel_ModuleAlreadyInstalled(bytes5)",](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L156)

[err = abi.encodeWithSignature("Policy_ModuleDoesNotExist(bytes5)", testKeycode);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L169)

[err = abi.encodeWithSignature("Kernel_PolicyAlreadyActivated(address)", address(policy));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L187)

[err = abi.encodeWithSignature("Kernel_PolicyAlreadyActivated(address)", address(policy));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L246)

[err = abi.encodeWithSignature("Module_PolicyNotPermitted(address)", address(policy));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L254)

[err = abi.encodeWithSignature("Kernel_InvalidModuleUpgrade(bytes5)", Keycode.wrap("MOCKY"));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L275)

[err = abi.encodeWithSignature("Kernel_InvalidModuleUpgrade(bytes5)", Keycode.wrap("MOCKY"));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L281)

[err = abi.encodeWithSignature("Policy_OnlyRole(bytes32)", Role.wrap("tester"));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L359)

[import {IBondFixedTermTeller} from "./interfaces/IBondFixedTermTeller.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/BondFixedTermTeller.sol#L7)

[import {ReentrancyGuard} from "solmate/utils/ReentrancyGuard.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseCDA.sol#L5)

[import {IBondAggregator} from "../interfaces/IBondAggregator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseCDA.sol#L11)

[import {ReentrancyGuard} from "solmate/utils/ReentrancyGuard.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L5)

[import {IBondAggregator} from "../interfaces/IBondAggregator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L10)

[import {IBondAuctioneer} from "../interfaces/IBondAuctioneer.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L11)

[import {IBondAuctioneer} from "../interfaces/IBondAuctioneer.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/interfaces/IBondCDA.sol#L5)

[inputs[2] = string(bytes.concat("./src/test/lib/quabi/jq.sh ", bytes(query), " ", bytes(path), ""));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/quabi/Quabi.sol#L17)

[inputs[2] = string(bytes.concat("./src/test/lib/quabi/path.sh ", bytes(contractName), ".json", ""));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/quabi/Quabi.sol#L27)

[string memory query = "'[.ast.nodes[-1].nodes[] | if .nodeType == \"FunctionDefinition\" and .kind == \"function\" then .functionSelector else empty end ]'";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/quabi/Quabi.sol#L48)

[string memory query = string(bytes.concat("'[.ast.nodes[-1].nodes[] | if .nodeType == \"FunctionDefinition\" and .kind == \"function\" and ([.modifiers[] | .modifierName.name == \"", bytes(modifierName), "\" ] | any ) then .functionSelector else empty end ]'"));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/quabi/Quabi.sol#L55)

[import {ReentrancyGuard} from "solmate/utils/ReentrancyGuard.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/mocks/Faucet.sol#L6)

[import {MockERC20, ERC20} from "solmate/test/utils/mocks/MockERC20.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/mocks/MockPrice.sol#L4)

[import {AggregatorV2V3Interface} from "interfaces/AggregatorV2V3Interface.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/mocks/MockPriceFeed.sol#L4)

[import {ModuleTestFixtureGenerator} from "test/lib/ModuleTestFixtureGenerator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/INSTR.t.sol#L7)

[import {MockValidUpgradedModule} from "test/mocks/MockValidUpgradedModule.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/INSTR.t.sol#L15)

[import {ModuleTestFixtureGenerator} from "test/lib/ModuleTestFixtureGenerator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/MINTR.t.sol#L10)

[import {ModuleTestFixtureGenerator} from "test/lib/ModuleTestFixtureGenerator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/PRICE.t.sol#L7)

[import {MockERC20, ERC20} from "solmate/test/utils/mocks/MockERC20.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/PRICE.t.sol#L9)

[import {ModuleTestFixtureGenerator} from "test/lib/ModuleTestFixtureGenerator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/RANGE.t.sol#L7)

[import {MockERC20, ERC20} from "solmate/test/utils/mocks/MockERC20.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/RANGE.t.sol#L9)

[import {ModuleTestFixtureGenerator} from "test/lib/ModuleTestFixtureGenerator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L7)

[import {MockERC20} from "solmate/test/utils/mocks/MockERC20.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L9)

[import {ModuleTestFixtureGenerator} from "test/lib/ModuleTestFixtureGenerator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/VOTES.t.sol#L8)

[import {BondFixedTermCDA} from "test/lib/bonds/BondFixedTermCDA.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L8)

[import {BondAggregator} from "test/lib/bonds/BondAggregator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L9)

[import {BondFixedTermTeller} from "test/lib/bonds/BondFixedTermTeller.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L10)

[import {IBondAuctioneer as LibAuctioneer} from "test/lib/bonds/interfaces/IBondAuctioneer.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L11)

[import {RolesAuthority, Authority as SolmateAuthority} from "solmate/auth/authorities/RolesAuthority.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L12)

[import {MockERC20, ERC20} from "solmate/test/utils/mocks/MockERC20.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L14)

["Callback_MarketNotSupported(uint256)",](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L375)

["Callback_MarketNotSupported(uint256)",](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L431)

[import {ModuleTestFixtureGenerator} from "../lib/ModuleTestFixtureGenerator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Governance.t.sol#L7)

[governance.submitProposal(instructions\_, "proposalName", "This is the proposal URI");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Governance.t.sol#L118)

[governance.submitProposal(instructions\_, "proposalName", "This is the proposal URI");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Governance.t.sol#L128)

[governance.submitProposal(instructions\_, "proposalName", "This is the proposal URI");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Governance.t.sol#L139)

[governance.submitProposal(instructions\_, "proposalName", "This is the proposal URI");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Governance.t.sol#L150)

[import {MockERC20, ERC20} from "solmate/test/utils/mocks/MockERC20.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Heart.t.sol#L8)

[import {IOperator, ERC20, IBondAuctioneer, IBondCallback} from "policies/interfaces/IOperator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Heart.t.sol#L17)

[import {BondFixedTermCDA} from "test/lib/bonds/BondFixedTermCDA.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L8)

[import {BondAggregator} from "test/lib/bonds/BondAggregator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L9)

[import {BondFixedTermTeller} from "test/lib/bonds/BondFixedTermTeller.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L10)

[import {RolesAuthority, Authority as SolmateAuthority} from "solmate/auth/authorities/RolesAuthority.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L11)

[import {MockERC20, ERC20} from "solmate/test/utils/mocks/MockERC20.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L13)

["Operator_AmountLessThanMinimum(uint256,uint256)",](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L409)

["Operator_AmountLessThanMinimum(uint256,uint256)",](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L424)

[import {MockERC20, ERC20} from "solmate/test/utils/mocks/MockERC20.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/PriceConfig.t.sol#L8)

[import {MockERC20} from "solmate/test/utils/mocks/MockERC20.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/TreasuryCustodian.t.sol#L7)

[import {TreasuryCustodian} from "src/policies/TreasuryCustodian.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/TreasuryCustodian.t.sol#L13)

## G-X: Use Shift Right/Left instead of Division/Multiplication if possible

### Description

Issue Information: [G008](https://github.com/byterocket/c4-common-issues/blob/main/0-Gas-Optimizations.md/#g008---use-shift-rightleft-instead-of-divisionmultiplication-if-possible)

### Findings:

[// vice versa. If your library also generates signatures with 0/1 for v instead 27/28, add 27 to v to accept](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L243)

[\* https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L520)

[// this feature: see https://github.com/ethereum/solidity/issues/4637](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L641)

[/// @dev Should be calculated as: (payoutDecimals - quoteDecimals) - ((payoutPriceDecimals - quotePriceDecimals) / 2)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/interfaces/IBondAuctioneer.sol#L41)

[/// @dev Credit to Remco Bloemen under MIT license https://xn--2-umb.com/21/muldiv](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L13)

[// Compute the product mod 2**256 and mod 2**256 - 1](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L21)

[// variables such that product = prod1 \* 2\*\*256 + prod0](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L24)

[// Make sure the result is less than 2\*\*256.](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L42)

[// to flip `twos` such that it is 2\*\*256 / twos.](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L76)

[// Invert denominator mod 2\*\*256](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L83)

[// modulo 2**256 such that denominator \* inv = 1 mod 2**256.](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L85)

[// correct for four bits. That is, denominator \* inv = 1 mod 2\*\*4](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L87)

[inv _= 2 - denominator _ inv; // inverse mod 2\*\*8](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L92)

[inv _= 2 - denominator _ inv; // inverse mod 2\*\*256](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L97)

[// correct result modulo 2\*\*256. Since the precoditions guarantee](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L101)

[// that the outcome is less than 2\*\*256, this is the final result.](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L102)

[int8 scaleAdjustment = int8(ohmDecimals) - int8(reserveDecimals) + (priceDecimals / 2);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Operator.sol#L372)

[uint256 invCushionPrice = 10\*_(oracleDecimals _ 2) / range.cushion.low.price;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Operator.sol#L419)

[uint256 invWallPrice = 10\*_(oracleDecimals _ 2) / range.wall.low.price;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Operator.sol#L420)

[int8 scaleAdjustment = int8(reserveDecimals) - int8(ohmDecimals) + (priceDecimals / 2);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Operator.sol#L427)

[) _ (FACTOR_SCALE + RANGE.spread(true) _ 2)) /](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Operator.sol#L786)

[uint32(7) // regenObserve // 21](https://github.com/code-423n4/2022-08-olympus/blob/main/src/scripts/Deploy.sol#L144)

[// scaleAdjustment should be equal to (payoutDecimals - quoteDecimals) - ((payoutPriceDecimals - quotePriceDecimals) / 2)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseCDA.sol#L157)

[// 2. If a tune interval has passed since last tune adjustment and the market is undersold](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseCDA.sol#L520)

[// 2. Calculate protocol fee as the total expected fee amount minus the referrer fee](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L134)

[int256 num1 = \_\_days + 68569 + 2440588; // 2440588 = OFFSET19700101](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L251)

[num1 = num1 - (146097 \* num2 + 3) / 4;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L253)

[num1 = num1 - (1461 \* \_year) / 4 + 31;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L255)

[int256 \_month = (80 \* num1) / 2447;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L256)

[int256 \_day = num1 - (2447 \* \_month) / 80;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L257)

[/// @dev Should be calculated as: (payoutDecimals - quoteDecimals) - ((payoutPriceDecimals - quotePriceDecimals) / 2)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/interfaces/IBondAuctioneer.sol#L41)

[// ,address](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/larping.sol#L13)

[// ,bool](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/larping.sol#L38)

[// ,bytes32](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/larping.sol#L63)

[// ,string](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/larping.sol#L88)

[// ,uint256](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/larping.sol#L113)

[// ,uint8](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/larping.sol#L138)

[vm.warp(51 _ 365 _ 24 _ 60 _ 60); // Set timestamp at roughly Jan 1, 2021 (51 years since Unix epoch)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/PRICE.t.sol#L34)

[vm.warp(51 _ 365 _ 24 _ 60 _ 60); // Set timestamp at roughly Jan 1, 2021 (51 years since Unix epoch)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/RANGE.t.sol#L36)

[TRSRY.setDebt(ngmi, debtor, INITIAL_TOKEN_AMOUNT / 2);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L174)

[assertEq(TRSRY.reserveDebt(ngmi, debtor), INITIAL_TOKEN_AMOUNT / 2);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L176)

[assertEq(TRSRY.totalDebt(ngmi), INITIAL_TOKEN_AMOUNT / 2);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L177)

[TRSRY.setDebt(ngmi, debtor, INITIAL_TOKEN_AMOUNT / 2);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L191)

[vm.warp(51 _ 365 _ 24 _ 60 _ 60); // Set timestamp at roughly Jan 1, 2021 (51 years since Unix epoch)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L81)

[ohm.mint(alice, testOhm \* 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L200)

[reserve.mint(alice, testReserve \* 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L201)

[ohm.approve(address(operator), testOhm \* 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L207)

[reserve.approve(address(operator), testReserve \* 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L209)

[ohm.approve(address(teller), testOhm \* 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L212)

[reserve.approve(address(teller), testReserve \* 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L214)

[// 2. Internal bond (OHM -> OHM)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L221)

[// 4. Regular OHM bond that will not be whitelisted](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L225)

[uint256 minimumPrice = (priceSignificand / 2) \*](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L271)

[vm.warp(51 _ 365 _ 24 _ 60 _ 60); // Set timestamp at roughly Jan 1, 2021 (51 years since Unix epoch)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Heart.t.sol#L64)

[vm.warp(51 _ 365 _ 24 _ 60 _ 60); // Set timestamp at roughly Jan 1, 2021 (51 years since Unix epoch)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L73)

[ohm.mint(alice, testOhm \* 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L185)

[reserve.mint(alice, testReserve \* 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L186)

[ohm.approve(address(operator), testOhm \* 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L192)

[reserve.approve(address(operator), testReserve \* 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L194)

[ohm.approve(address(teller), testOhm \* 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L197)

[reserve.approve(address(teller), testReserve \* 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L199)

[uint256 amountIn = auctioneer.maxAmountAccepted(id, guardian) / 2;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L869)

[uint256 amountIn = auctioneer.maxAmountAccepted(id, guardian) / 2;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L900)

[) _ (1e4 + range.spread(true) _ 2)) / 1e4;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L2044)

[vm.warp(51 _ 365 _ 24 _ 60 _ 60); // Set timestamp at roughly Jan 1, 2021 (51 years since Unix epoch)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/PriceConfig.t.sol#L38)

## Unsafe ERC20 Operation(s)

### Description

Issue Information: [L001](https://github.com/byterocket/c4-common-issues/blob/main/2-Low-Risk.md#l001---unsafe-erc20-operations)

### Findings:

[VOTES.transferFrom(msg.sender, address(this), userVotes);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Governance.sol#L259)

[VOTES.transferFrom(address(this), msg.sender, userVotes);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Governance.sol#L312)

[underlying*.transferFrom(msg.sender, address(this), amount*);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/BondFixedTermTeller.sol#L115)

[ohm.approve(address(MINTR), amount\_);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/MINTR.t.sol#L87)

[ohm.approve(users[0], amount\_);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/MINTR.t.sol#L107)

[ngmi.approve(address(TRSRY), amount\_);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L156)

[VOTES.transfer(address(0), 10);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/VOTES.t.sol#L50)

[VOTES.transferFrom(address(1), address(2), 3);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/VOTES.t.sol#L74)

[ohm.approve(address(operator), testOhm \* 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L207)

[reserve.approve(address(operator), testReserve \* 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L209)

[ohm.approve(address(teller), testOhm \* 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L212)

[reserve.approve(address(teller), testReserve \* 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L214)

[VOTES.approve(address(governance), type(uint256).max);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Governance.t.sol#L96)

[VOTES.approve(address(governance), type(uint256).max);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Governance.t.sol#L98)

[VOTES.approve(address(governance), type(uint256).max);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Governance.t.sol#L100)

[VOTES.approve(address(governance), type(uint256).max);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Governance.t.sol#L102)

[VOTES.approve(address(governance), type(uint256).max);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Governance.t.sol#L104)

[ohm.approve(address(operator), testOhm \* 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L192)

[reserve.approve(address(operator), testReserve \* 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L194)

[ohm.approve(address(teller), testOhm \* 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L197)

[reserve.approve(address(teller), testReserve \* 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L199)

## Unspecific Compiler Version Pragma

### Description

Issue Information: [L003](https://github.com/byterocket/c4-common-issues/blob/main/2-Low-Risk.md#l003---unspecific-compiler-version-pragma)

### Findings:

[pragma solidity >=0.7.5;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L2)

[pragma solidity ^0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/interfaces/AggregatorV2V3Interface.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/interfaces/IBondAggregator.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/interfaces/IBondAuctioneer.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/interfaces/IBondCallback.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/interfaces/IBondTeller.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/interfaces/IWETH9.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/TransferHelper.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/interfaces/IHeart.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/interfaces/IOperator.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/ModuleTestFixtureGenerator.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/UserFactory.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/interfaces/IBondAggregator.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/interfaces/IBondAuctioneer.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/interfaces/IBondCDA.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/interfaces/IBondCallback.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/interfaces/IBondFixedTermTeller.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/interfaces/IBondTeller.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/larping.sol#L1)

[pragma solidity ^0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/mocks/MockPriceFeed.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/INSTR.t.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/PRICE.t.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/RANGE.t.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/VOTES.t.sol#L3)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Governance.t.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Heart.t.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/PriceConfig.t.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/TreasuryCustodian.t.sol#L2)

[pragma solidity >=0.8.0;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/VoterRegistration.t.sol#L2)

## Do not use Deprecated Library Functions

### Description

Issue Information: [L005](https://github.com/byterocket/c4-common-issues/blob/main/2-Low-Risk.md#l005---do-not-use-deprecated-library-functions)

### Findings:

[function safeApprove(](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/TransferHelper.sol#L35)

[ohm.safeApprove(address(MINTR), type(uint256).max);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/BondCallback.sol#L57)

[ohm.safeApprove(address(MINTR), type(uint256).max);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Operator.sol#L167)
