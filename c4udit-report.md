# Gas | QA 


## ✅ G-X: Don't Initialize Variables with Default Value

### 📝 Description
Uninitialized variables are assigned with the types default value. Explicitly initializing a variable with it's default value costs unnecesary gas. Not overwriting the default for [stack variables](https://gist.github.com/IllIllI000/e075d189c1b23dce256cd166e28f3397) saves 8 gas. Storage and memory variables have larger savings

### 💡 Recommendation
Delete useless variable declarations to save gas.

### 🔍 Findings:
```2022-08-olympus/blob/main/src/Kernel.sol#L397``` [for (uint256 i = 0; i < reqLength; ) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/Kernel.sol#L397 )

```2022-08-olympus/blob/main/src/scripts/Deploy.sol#L239``` [for (uint i = 0; i < 90; i++) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/scripts/Deploy.sol#L239 )

```2022-08-olympus/blob/main/src/test/lib/UserFactory.sol#L25``` [for (uint256 i = 0; i < userNum; i++) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/UserFactory.sol#L25 )

```2022-08-olympus/blob/main/src/utils/KernelUtils.sol#L43``` [for (uint256 i = 0; i < 5; ) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/utils/KernelUtils.sol#L43 )

```2022-08-olympus/blob/main/src/utils/KernelUtils.sol#L58``` [for (uint256 i = 0; i < 32; ) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/utils/KernelUtils.sol#L58 )


## ✅ G-X: Cache Array Length Outside of Loop

### 📝 Description
Caching the array length outside a loop saves reading it on each iteration, as long as the array's length is not changed during the loop.
You save 3 gas by not reading `array.length` - 3 gas per instance - 27 gas saved

### 💡 Recommendation
Reco

### 🔍 Findings:
```2022-08-olympus/blob/main/src/modules/INSTR.sol#L50``` [for (uint256 i; i < length; ) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/modules/INSTR.sol#L50 )

```2022-08-olympus/blob/main/src/policies/Governance.sol#L278``` [for (uint256 step; step < instructions.length; ) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Governance.sol#L278 )

```2022-08-olympus/blob/main/src/test/lib/bonds/interfaces/IBondCDA.sol#L37``` [uint32 length; // time from creation to conclusion.](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/interfaces/IBondCDA.sol#L37 )

```2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L80``` [uint256 idsLength = ids.length; // Saves MLOADs.](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L80 )

```2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L126``` [uint256 ownersLength = owners.length; // Saves MLOADs.](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L126 )

```2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L186``` [uint256 idsLength = ids.length; // Saves MLOADs.](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L186 )

```2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L221``` [uint256 idsLength = ids.length; // Saves MLOADs.](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/lib/ERC1155.sol#L221 )


## ✅ G-X: Use != 0 instead of > 0 for Unsigned Integer Comparison

### 📝 Description
Use != 0 when comparing uint variables to zero, which cannot hold values below zero

### 💡 Recommendation
You should change from `> 0` to  `!=0`.

### 🔍 Findings:
```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L245``` [if (uint256(s) > 0x7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF5D576E7357A4501DDFE92F46681B20A0) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L245 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L611``` [require(b > 0, errorMessage);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L611 )

```2022-08-olympus/blob/main/src/libraries/FullMath.sol#L35``` [require(denominator > 0);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L35 )

```2022-08-olympus/blob/main/src/libraries/FullMath.sol#L122``` [if (mulmod(a, b, denominator) > 0) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L122 )

```2022-08-olympus/blob/main/src/policies/Governance.sol#L247``` [if (userVotesForProposal[activeProposal.proposalId][msg.sender] > 0) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Governance.sol#L247 )

```2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L98``` [vm.assume(amount_ > 0);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L98 )

```2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L108``` [vm.assume(amount_ > 0);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L108 )

```2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L126``` [vm.assume(amount_ > 0);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L126 )

```2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L143``` [vm.assume(amount_ > 0);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L143 )

```2022-08-olympus/blob/main/src/utils/KernelUtils.sol#L46``` [if (char < 0x41 || char > 0x5A) revert InvalidKeycode(keycode_); // A-Z only](https://github.com/code-423n4/2022-08-olympus/blob/main/src/utils/KernelUtils.sol#L46 )

```2022-08-olympus/blob/main/src/utils/KernelUtils.sol#L60``` [if ((char < 0x61 || char > 0x7A) && char != 0x5f && char != 0x00) {](https://github.com/code-423n4/2022-08-olympus/blob/main/src/utils/KernelUtils.sol#L60 )


## ✅ G-X: Long Revert Strings

### 📝 Description
Shortening revert strings to fit in 32 bytes will decrease gas costs for deployment and gas costs when the revert condition has been met.

### 💡 Recommendation
Use short statements instead of long statement.

### 🔍 Findings:
```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L107``` [revert("ECDSA: invalid signature 's' value");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L107 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L109``` [revert("ECDSA: invalid signature 'v' value");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L109 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L363``` ["EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L363 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L597``` [require(c / a == b, "SafeMath: multiplication overflow");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L597 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L738``` [_allowances[sender][msg.sender].sub(amount, "ERC20: transfer amount exceeds allowance")](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L738 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L758``` ["ERC20: decreased allowance below zero"](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L758 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L769``` [require(sender != address(0), "ERC20: transfer from the zero address");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L769 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L770``` [require(recipient != address(0), "ERC20: transfer to the zero address");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L770 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L774``` [_balances[sender] = _balances[sender].sub(amount, "ERC20: transfer amount exceeds balance");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L774 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L788``` [require(account != address(0), "ERC20: burn from the zero address");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L788 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L792``` [_balances[account] = _balances[account].sub(amount, "ERC20: burn amount exceeds balance");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L792 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L802``` [require(owner != address(0), "ERC20: approve from the zero address");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L802 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L803``` [require(spender != address(0), "ERC20: approve to the zero address");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L803 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L836``` ["Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)"](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L836 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L925``` ["ERC20: burn amount exceeds allowance"](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L925 )

```2022-08-olympus/blob/main/src/modules/PRICE.sol#L4``` [import {AggregatorV2V3Interface} from "interfaces/AggregatorV2V3Interface.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/modules/PRICE.sol#L4 )

```2022-08-olympus/blob/main/src/modules/TRSRY.sol#L5``` [import {ReentrancyGuard} from "solmate/utils/ReentrancyGuard.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/modules/TRSRY.sol#L5 )

```2022-08-olympus/blob/main/src/modules/VOTES.sol#L18``` [ERC20("OlympusDAO Dummy Voting Tokens", "VOTES", 0)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/modules/VOTES.sol#L18 )

```2022-08-olympus/blob/main/src/policies/BondCallback.sol#L5``` [import {ReentrancyGuard} from "solmate/utils/ReentrancyGuard.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/BondCallback.sol#L5 )

```2022-08-olympus/blob/main/src/policies/Heart.sol#L4``` [import {ReentrancyGuard} from "solmate/utils/ReentrancyGuard.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Heart.sol#L4 )

```2022-08-olympus/blob/main/src/policies/Heart.sol#L7``` [import {IOperator} from "policies/interfaces/IOperator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Heart.sol#L7 )

```2022-08-olympus/blob/main/src/policies/Operator.sol#L4``` [import {ReentrancyGuard} from "solmate/utils/ReentrancyGuard.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Operator.sol#L4 )

```2022-08-olympus/blob/main/src/policies/Operator.sol#L7``` [import {IOperator} from "policies/interfaces/IOperator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Operator.sol#L7 )

```2022-08-olympus/blob/main/src/scripts/Deploy.sol#L4``` [import {AggregatorV2V3Interface} from "interfaces/AggregatorV2V3Interface.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/scripts/Deploy.sol#L4 )

```2022-08-olympus/blob/main/src/scripts/Deploy.sol#L271``` [console2.log("RESERVE-ETH Price Feed deployed to:", address(reserveEthPriceFeed));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/scripts/Deploy.sol#L271 )

```2022-08-olympus/blob/main/src/test/Kernel.t.sol#L77``` [err = abi.encodeWithSignature("InvalidKeycode(bytes5)", Keycode.wrap("inval"));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L77 )

```2022-08-olympus/blob/main/src/test/Kernel.t.sol#L81``` [err = abi.encodeWithSignature("InvalidKeycode(bytes5)", Keycode.wrap(""));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L81 )

```2022-08-olympus/blob/main/src/test/Kernel.t.sol#L89``` [err = abi.encodeWithSignature("InvalidRole(bytes32)", Role.wrap("INVALID_ID"));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L89 )

```2022-08-olympus/blob/main/src/test/Kernel.t.sol#L150``` [err = abi.encodeWithSignature("InvalidKeycode(bytes5)", Keycode.wrap("badkc"));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L150 )

```2022-08-olympus/blob/main/src/test/Kernel.t.sol#L156``` ["Kernel_ModuleAlreadyInstalled(bytes5)",](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L156 )

```2022-08-olympus/blob/main/src/test/Kernel.t.sol#L169``` [err = abi.encodeWithSignature("Policy_ModuleDoesNotExist(bytes5)", testKeycode);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L169 )

```2022-08-olympus/blob/main/src/test/Kernel.t.sol#L187``` [err = abi.encodeWithSignature("Kernel_PolicyAlreadyActivated(address)", address(policy));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L187 )

```2022-08-olympus/blob/main/src/test/Kernel.t.sol#L246``` [err = abi.encodeWithSignature("Kernel_PolicyAlreadyActivated(address)", address(policy));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L246 )

```2022-08-olympus/blob/main/src/test/Kernel.t.sol#L254``` [err = abi.encodeWithSignature("Module_PolicyNotPermitted(address)", address(policy));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L254 )

```2022-08-olympus/blob/main/src/test/Kernel.t.sol#L275``` [err = abi.encodeWithSignature("Kernel_InvalidModuleUpgrade(bytes5)", Keycode.wrap("MOCKY"));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L275 )

```2022-08-olympus/blob/main/src/test/Kernel.t.sol#L281``` [err = abi.encodeWithSignature("Kernel_InvalidModuleUpgrade(bytes5)", Keycode.wrap("MOCKY"));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L281 )

```2022-08-olympus/blob/main/src/test/Kernel.t.sol#L359``` [err = abi.encodeWithSignature("Policy_OnlyRole(bytes32)", Role.wrap("tester"));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/Kernel.t.sol#L359 )

```2022-08-olympus/blob/main/src/test/lib/bonds/BondFixedTermTeller.sol#L7``` [import {IBondFixedTermTeller} from "./interfaces/IBondFixedTermTeller.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/BondFixedTermTeller.sol#L7 )

```2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseCDA.sol#L5``` [import {ReentrancyGuard} from "solmate/utils/ReentrancyGuard.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseCDA.sol#L5 )

```2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseCDA.sol#L11``` [import {IBondAggregator} from "../interfaces/IBondAggregator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseCDA.sol#L11 )

```2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L5``` [import {ReentrancyGuard} from "solmate/utils/ReentrancyGuard.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L5 )

```2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L10``` [import {IBondAggregator} from "../interfaces/IBondAggregator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L10 )

```2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L11``` [import {IBondAuctioneer} from "../interfaces/IBondAuctioneer.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L11 )

```2022-08-olympus/blob/main/src/test/lib/bonds/interfaces/IBondCDA.sol#L5``` [import {IBondAuctioneer} from "../interfaces/IBondAuctioneer.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/interfaces/IBondCDA.sol#L5 )

```2022-08-olympus/blob/main/src/test/lib/quabi/Quabi.sol#L17``` [inputs[2] = string(bytes.concat("./src/test/lib/quabi/jq.sh ", bytes(query), " ", bytes(path), ""));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/quabi/Quabi.sol#L17 )

```2022-08-olympus/blob/main/src/test/lib/quabi/Quabi.sol#L27``` [inputs[2] = string(bytes.concat("./src/test/lib/quabi/path.sh ", bytes(contractName), ".json", ""));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/quabi/Quabi.sol#L27 )

```2022-08-olympus/blob/main/src/test/lib/quabi/Quabi.sol#L48``` [string memory query = "'[.ast.nodes[-1].nodes[] | if .nodeType == \"FunctionDefinition\" and .kind == \"function\" then .functionSelector else empty end ]'";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/quabi/Quabi.sol#L48 )

```2022-08-olympus/blob/main/src/test/lib/quabi/Quabi.sol#L55``` [string memory query = string(bytes.concat("'[.ast.nodes[-1].nodes[] | if .nodeType == \"FunctionDefinition\" and .kind == \"function\" and ([.modifiers[] | .modifierName.name == \"", bytes(modifierName), "\" ] | any ) then .functionSelector else empty end ]'"));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/quabi/Quabi.sol#L55 )

```2022-08-olympus/blob/main/src/test/mocks/Faucet.sol#L6``` [import {ReentrancyGuard} from "solmate/utils/ReentrancyGuard.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/mocks/Faucet.sol#L6 )

```2022-08-olympus/blob/main/src/test/mocks/MockPrice.sol#L4``` [import {MockERC20, ERC20} from "solmate/test/utils/mocks/MockERC20.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/mocks/MockPrice.sol#L4 )

```2022-08-olympus/blob/main/src/test/mocks/MockPriceFeed.sol#L4``` [import {AggregatorV2V3Interface} from "interfaces/AggregatorV2V3Interface.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/mocks/MockPriceFeed.sol#L4 )

```2022-08-olympus/blob/main/src/test/modules/INSTR.t.sol#L7``` [import {ModuleTestFixtureGenerator} from "test/lib/ModuleTestFixtureGenerator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/INSTR.t.sol#L7 )

```2022-08-olympus/blob/main/src/test/modules/INSTR.t.sol#L15``` [import {MockValidUpgradedModule} from "test/mocks/MockValidUpgradedModule.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/INSTR.t.sol#L15 )

```2022-08-olympus/blob/main/src/test/modules/MINTR.t.sol#L10``` [import {ModuleTestFixtureGenerator} from "test/lib/ModuleTestFixtureGenerator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/MINTR.t.sol#L10 )

```2022-08-olympus/blob/main/src/test/modules/PRICE.t.sol#L7``` [import {ModuleTestFixtureGenerator} from "test/lib/ModuleTestFixtureGenerator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/PRICE.t.sol#L7 )

```2022-08-olympus/blob/main/src/test/modules/PRICE.t.sol#L9``` [import {MockERC20, ERC20} from "solmate/test/utils/mocks/MockERC20.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/PRICE.t.sol#L9 )

```2022-08-olympus/blob/main/src/test/modules/RANGE.t.sol#L7``` [import {ModuleTestFixtureGenerator} from "test/lib/ModuleTestFixtureGenerator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/RANGE.t.sol#L7 )

```2022-08-olympus/blob/main/src/test/modules/RANGE.t.sol#L9``` [import {MockERC20, ERC20} from "solmate/test/utils/mocks/MockERC20.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/RANGE.t.sol#L9 )

```2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L7``` [import {ModuleTestFixtureGenerator} from "test/lib/ModuleTestFixtureGenerator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L7 )

```2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L9``` [import {MockERC20} from "solmate/test/utils/mocks/MockERC20.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L9 )

```2022-08-olympus/blob/main/src/test/modules/VOTES.t.sol#L8``` [import {ModuleTestFixtureGenerator} from "test/lib/ModuleTestFixtureGenerator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/VOTES.t.sol#L8 )

```2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L8``` [import {BondFixedTermCDA} from "test/lib/bonds/BondFixedTermCDA.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L8 )

```2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L9``` [import {BondAggregator} from "test/lib/bonds/BondAggregator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L9 )

```2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L10``` [import {BondFixedTermTeller} from "test/lib/bonds/BondFixedTermTeller.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L10 )

```2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L11``` [import {IBondAuctioneer as LibAuctioneer} from "test/lib/bonds/interfaces/IBondAuctioneer.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L11 )

```2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L12``` [import {RolesAuthority, Authority as SolmateAuthority} from "solmate/auth/authorities/RolesAuthority.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L12 )

```2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L14``` [import {MockERC20, ERC20} from "solmate/test/utils/mocks/MockERC20.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L14 )

```2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L375``` ["Callback_MarketNotSupported(uint256)",](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L375 )

```2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L431``` ["Callback_MarketNotSupported(uint256)",](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L431 )

```2022-08-olympus/blob/main/src/test/policies/Governance.t.sol#L7``` [import {ModuleTestFixtureGenerator} from "../lib/ModuleTestFixtureGenerator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Governance.t.sol#L7 )

```2022-08-olympus/blob/main/src/test/policies/Governance.t.sol#L118``` [governance.submitProposal(instructions_, "proposalName", "This is the proposal URI");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Governance.t.sol#L118 )

```2022-08-olympus/blob/main/src/test/policies/Governance.t.sol#L128``` [governance.submitProposal(instructions_, "proposalName", "This is the proposal URI");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Governance.t.sol#L128 )

```2022-08-olympus/blob/main/src/test/policies/Governance.t.sol#L139``` [governance.submitProposal(instructions_, "proposalName", "This is the proposal URI");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Governance.t.sol#L139 )

```2022-08-olympus/blob/main/src/test/policies/Governance.t.sol#L150``` [governance.submitProposal(instructions_, "proposalName", "This is the proposal URI");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Governance.t.sol#L150 )

```2022-08-olympus/blob/main/src/test/policies/Heart.t.sol#L8``` [import {MockERC20, ERC20} from "solmate/test/utils/mocks/MockERC20.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Heart.t.sol#L8 )

```2022-08-olympus/blob/main/src/test/policies/Heart.t.sol#L17``` [import {IOperator, ERC20, IBondAuctioneer, IBondCallback} from "policies/interfaces/IOperator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Heart.t.sol#L17 )

```2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L8``` [import {BondFixedTermCDA} from "test/lib/bonds/BondFixedTermCDA.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L8 )

```2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L9``` [import {BondAggregator} from "test/lib/bonds/BondAggregator.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L9 )

```2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L10``` [import {BondFixedTermTeller} from "test/lib/bonds/BondFixedTermTeller.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L10 )

```2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L11``` [import {RolesAuthority, Authority as SolmateAuthority} from "solmate/auth/authorities/RolesAuthority.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L11 )

```2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L13``` [import {MockERC20, ERC20} from "solmate/test/utils/mocks/MockERC20.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L13 )

```2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L409``` ["Operator_AmountLessThanMinimum(uint256,uint256)",](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L409 )

```2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L424``` ["Operator_AmountLessThanMinimum(uint256,uint256)",](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L424 )

```2022-08-olympus/blob/main/src/test/policies/PriceConfig.t.sol#L8``` [import {MockERC20, ERC20} from "solmate/test/utils/mocks/MockERC20.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/PriceConfig.t.sol#L8 )

```2022-08-olympus/blob/main/src/test/policies/TreasuryCustodian.t.sol#L7``` [import {MockERC20} from "solmate/test/utils/mocks/MockERC20.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/TreasuryCustodian.t.sol#L7 )

```2022-08-olympus/blob/main/src/test/policies/TreasuryCustodian.t.sol#L13``` [import {TreasuryCustodian} from "src/policies/TreasuryCustodian.sol";](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/TreasuryCustodian.t.sol#L13 )


## ✅ G-X: Use Shift Right/Left instead of Division/Multiplication if possible

### 📝 Description
A division/multiplication by any number `x` being a power of 2 can be calculated by shifting `log2(x)` to the right/left.
While the `DIV` opcode uses 5 gas, the `SHR` opcode only uses 3 gas.
urthermore, Solidity's division operation also includes a division-by-0 prevention which is bypassed using shifting.

### 💡 Recommendation
Reco

### 🔍 Findings:
```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L243``` [// vice versa. If your library also generates signatures with 0/1 for v instead 27/28, add 27 to v to accept](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L243 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L520``` [* https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L520 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L641``` [// this feature: see https://github.com/ethereum/solidity/issues/4637](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L641 )

```2022-08-olympus/blob/main/src/interfaces/IBondAuctioneer.sol#L41``` [/// @dev                        Should be calculated as: (payoutDecimals - quoteDecimals) - ((payoutPriceDecimals - quotePriceDecimals) / 2)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/interfaces/IBondAuctioneer.sol#L41 )

```2022-08-olympus/blob/main/src/libraries/FullMath.sol#L13``` [/// @dev Credit to Remco Bloemen under MIT license https://xn--2-umb.com/21/muldiv](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L13 )

```2022-08-olympus/blob/main/src/libraries/FullMath.sol#L21``` [// Compute the product mod 2**256 and mod 2**256 - 1](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L21 )

```2022-08-olympus/blob/main/src/libraries/FullMath.sol#L24``` [// variables such that product = prod1 * 2**256 + prod0](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L24 )

```2022-08-olympus/blob/main/src/libraries/FullMath.sol#L42``` [// Make sure the result is less than 2**256.](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L42 )

```2022-08-olympus/blob/main/src/libraries/FullMath.sol#L76``` [// to flip `twos` such that it is 2**256 / twos.](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L76 )

```2022-08-olympus/blob/main/src/libraries/FullMath.sol#L83``` [// Invert denominator mod 2**256](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L83 )

```2022-08-olympus/blob/main/src/libraries/FullMath.sol#L85``` [// modulo 2**256 such that denominator * inv = 1 mod 2**256.](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L85 )

```2022-08-olympus/blob/main/src/libraries/FullMath.sol#L87``` [// correct for four bits. That is, denominator * inv = 1 mod 2**4](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L87 )

```2022-08-olympus/blob/main/src/libraries/FullMath.sol#L92``` [inv *= 2 - denominator * inv; // inverse mod 2**8](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L92 )

```2022-08-olympus/blob/main/src/libraries/FullMath.sol#L97``` [inv *= 2 - denominator * inv; // inverse mod 2**256](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L97 )

```2022-08-olympus/blob/main/src/libraries/FullMath.sol#L101``` [// correct result modulo 2**256. Since the precoditions guarantee](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L101 )

```2022-08-olympus/blob/main/src/libraries/FullMath.sol#L102``` [// that the outcome is less than 2**256, this is the final result.](https://github.com/code-423n4/2022-08-olympus/blob/main/src/libraries/FullMath.sol#L102 )

```2022-08-olympus/blob/main/src/policies/Operator.sol#L372``` [int8 scaleAdjustment = int8(ohmDecimals) - int8(reserveDecimals) + (priceDecimals / 2);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Operator.sol#L372 )

```2022-08-olympus/blob/main/src/policies/Operator.sol#L419``` [uint256 invCushionPrice = 10**(oracleDecimals * 2) / range.cushion.low.price;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Operator.sol#L419 )

```2022-08-olympus/blob/main/src/policies/Operator.sol#L420``` [uint256 invWallPrice = 10**(oracleDecimals * 2) / range.wall.low.price;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Operator.sol#L420 )

```2022-08-olympus/blob/main/src/policies/Operator.sol#L427``` [int8 scaleAdjustment = int8(reserveDecimals) - int8(ohmDecimals) + (priceDecimals / 2);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Operator.sol#L427 )

```2022-08-olympus/blob/main/src/policies/Operator.sol#L786``` [) * (FACTOR_SCALE + RANGE.spread(true) * 2)) /](https://github.com/code-423n4/2022-08-olympus/blob/main/src/policies/Operator.sol#L786 )

```2022-08-olympus/blob/main/src/scripts/Deploy.sol#L144``` [uint32(7) // regenObserve    // 21](https://github.com/code-423n4/2022-08-olympus/blob/main/src/scripts/Deploy.sol#L144 )

```2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseCDA.sol#L157``` [// scaleAdjustment should be equal to (payoutDecimals - quoteDecimals) - ((payoutPriceDecimals - quotePriceDecimals) / 2)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseCDA.sol#L157 )

```2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseCDA.sol#L520``` [// 2. If a tune interval has passed since last tune adjustment and the market is undersold](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseCDA.sol#L520 )

```2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L134``` [// 2. Calculate protocol fee as the total expected fee amount minus the referrer fee](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L134 )

```2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L251``` [int256 num1 = __days + 68569 + 2440588; // 2440588 = OFFSET19700101](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L251 )

```2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L253``` [num1 = num1 - (146097 * num2 + 3) / 4;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L253 )

```2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L255``` [num1 = num1 - (1461 * _year) / 4 + 31;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L255 )

```2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L256``` [int256 _month = (80 * num1) / 2447;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L256 )

```2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L257``` [int256 _day = num1 - (2447 * _month) / 80;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/bases/BondBaseTeller.sol#L257 )

```2022-08-olympus/blob/main/src/test/lib/bonds/interfaces/IBondAuctioneer.sol#L41``` [/// @dev                        Should be calculated as: (payoutDecimals - quoteDecimals) - ((payoutPriceDecimals - quotePriceDecimals) / 2)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/interfaces/IBondAuctioneer.sol#L41 )

```2022-08-olympus/blob/main/src/test/lib/larping.sol#L13``` [// ,address](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/larping.sol#L13 )

```2022-08-olympus/blob/main/src/test/lib/larping.sol#L38``` [// ,bool](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/larping.sol#L38 )

```2022-08-olympus/blob/main/src/test/lib/larping.sol#L63``` [// ,bytes32](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/larping.sol#L63 )

```2022-08-olympus/blob/main/src/test/lib/larping.sol#L88``` [// ,string](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/larping.sol#L88 )

```2022-08-olympus/blob/main/src/test/lib/larping.sol#L113``` [// ,uint256](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/larping.sol#L113 )

```2022-08-olympus/blob/main/src/test/lib/larping.sol#L138``` [// ,uint8](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/larping.sol#L138 )

```2022-08-olympus/blob/main/src/test/modules/PRICE.t.sol#L34``` [vm.warp(51 * 365 * 24 * 60 * 60); // Set timestamp at roughly Jan 1, 2021 (51 years since Unix epoch)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/PRICE.t.sol#L34 )

```2022-08-olympus/blob/main/src/test/modules/RANGE.t.sol#L36``` [vm.warp(51 * 365 * 24 * 60 * 60); // Set timestamp at roughly Jan 1, 2021 (51 years since Unix epoch)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/RANGE.t.sol#L36 )

```2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L174``` [TRSRY.setDebt(ngmi, debtor, INITIAL_TOKEN_AMOUNT / 2);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L174 )

```2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L176``` [assertEq(TRSRY.reserveDebt(ngmi, debtor), INITIAL_TOKEN_AMOUNT / 2);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L176 )

```2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L177``` [assertEq(TRSRY.totalDebt(ngmi), INITIAL_TOKEN_AMOUNT / 2);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L177 )

```2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L191``` [TRSRY.setDebt(ngmi, debtor, INITIAL_TOKEN_AMOUNT / 2);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/TRSRY.t.sol#L191 )

```2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L81``` [vm.warp(51 * 365 * 24 * 60 * 60); // Set timestamp at roughly Jan 1, 2021 (51 years since Unix epoch)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L81 )

```2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L200``` [ohm.mint(alice, testOhm * 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L200 )

```2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L201``` [reserve.mint(alice, testReserve * 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L201 )

```2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L207``` [ohm.approve(address(operator), testOhm * 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L207 )

```2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L209``` [reserve.approve(address(operator), testReserve * 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L209 )

```2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L212``` [ohm.approve(address(teller), testOhm * 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L212 )

```2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L214``` [reserve.approve(address(teller), testReserve * 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L214 )

```2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L221``` [// 2. Internal bond (OHM -> OHM)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L221 )

```2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L225``` [// 4. Regular OHM bond that will not be whitelisted](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L225 )

```2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L271``` [uint256 minimumPrice = (priceSignificand / 2) *](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/BondCallback.t.sol#L271 )

```2022-08-olympus/blob/main/src/test/policies/Heart.t.sol#L64``` [vm.warp(51 * 365 * 24 * 60 * 60); // Set timestamp at roughly Jan 1, 2021 (51 years since Unix epoch)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Heart.t.sol#L64 )

```2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L73``` [vm.warp(51 * 365 * 24 * 60 * 60); // Set timestamp at roughly Jan 1, 2021 (51 years since Unix epoch)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L73 )

```2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L185``` [ohm.mint(alice, testOhm * 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L185 )

```2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L186``` [reserve.mint(alice, testReserve * 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L186 )

```2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L192``` [ohm.approve(address(operator), testOhm * 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L192 )

```2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L194``` [reserve.approve(address(operator), testReserve * 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L194 )

```2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L197``` [ohm.approve(address(teller), testOhm * 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L197 )

```2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L199``` [reserve.approve(address(teller), testReserve * 20);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L199 )

```2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L869``` [uint256 amountIn = auctioneer.maxAmountAccepted(id, guardian) / 2;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L869 )

```2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L900``` [uint256 amountIn = auctioneer.maxAmountAccepted(id, guardian) / 2;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L900 )

```2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L2044``` [) * (1e4 + range.spread(true) * 2)) / 1e4;](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/Operator.t.sol#L2044 )

```2022-08-olympus/blob/main/src/test/policies/PriceConfig.t.sol#L38``` [vm.warp(51 * 365 * 24 * 60 * 60); // Set timestamp at roughly Jan 1, 2021 (51 years since Unix epoch)](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/PriceConfig.t.sol#L38 )


## ✅ G-X: Use assembly to hash

### 📝 Description
See [this issue](https://github.com/code-423n4/2022-06-putty-findings/issues/59)
You can save about 5000 gas per instance in deploying contracrt.
You can save about 80 gas per instance if using assembly to execute the function.

### 💡 Recommendation
Please follow [this link](https://gist.github.com/Tomosuke0930/72beb469f08c954b232e58b8da03ef28) to make corrections

### 🔍 Findings:
```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L287``` [return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L287 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L304``` [return keccak256(abi.encodePacked("\x19\x01", domainSeparator, structHash));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L304 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L315``` [* they need in their contracts using a combination of `abi.encode` and `keccak256`.](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L315 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L360``` [bytes32 hashedName = keccak256(bytes(name));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L360 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L361``` [bytes32 hashedVersion = keccak256(bytes(version));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L361 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L362``` [bytes32 typeHash = keccak256(](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L362 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L398``` [return keccak256(abi.encode(typeHash, nameHash, versionHash, chainID, address(this)));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L398 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L408``` [* bytes32 digest = _hashTypedDataV4(keccak256(abi.encode(](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L408 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L409``` [*     keccak256("Mail(address to,string contents)"),](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L409 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L411``` [*     keccak256(bytes(mailContents))](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L411 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L665``` [bytes32 private constant ERC20TOKEN_ERC1820_INTERFACE_ID = keccak256("ERC20Token");](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L665 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L835``` [keccak256(](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L835 )

```2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L860``` [bytes32 structHash = keccak256(](https://github.com/code-423n4/2022-08-olympus/blob/main/src/external/OlympusERC20.sol#L860 )

```2022-08-olympus/blob/main/src/test/lib/UserFactory.sol#L9``` [address(bytes20(uint160(uint256(keccak256("hevm cheat code")))));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/UserFactory.sol#L9 )

```2022-08-olympus/blob/main/src/test/lib/UserFactory.sol#L13``` [bytes32 internal nextUser = keccak256(abi.encodePacked("users"));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/UserFactory.sol#L13 )

```2022-08-olympus/blob/main/src/test/lib/UserFactory.sol#L18``` [nextUser = keccak256(abi.encodePacked(nextUser));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/UserFactory.sol#L18 )

```2022-08-olympus/blob/main/src/test/lib/bonds/BondFixedTermTeller.sol#L231``` [keccak256(abi.encodePacked(underlying_, expiry_ / uint48(1 days)))](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/bonds/BondFixedTermTeller.sol#L231 )

```2022-08-olympus/blob/main/src/test/lib/larping.sol#L9``` [address(bytes20(uint160(uint256(keccak256("hevm cheat code")))));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/larping.sol#L9 )

```2022-08-olympus/blob/main/src/test/lib/quabi/Quabi.sol#L8``` [Vm internal constant vm = Vm(address(bytes20(uint160(uint256(keccak256("hevm cheat code"))))));](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/lib/quabi/Quabi.sol#L8 )

```2022-08-olympus/blob/main/src/test/modules/PRICE.t.sol#L101``` [change = int256(uint256(keccak256(abi.encodePacked(nonce, i)))) %!i(MISSING)nt256(1000);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/PRICE.t.sol#L101 )

```2022-08-olympus/blob/main/src/test/modules/PRICE.t.sol#L128``` [change = int256(uint256(keccak256(abi.encodePacked(nonce, i)))) %!i(MISSING)nt256(1000);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/modules/PRICE.t.sol#L128 )

```2022-08-olympus/blob/main/src/test/policies/PriceConfig.t.sol#L124``` [change = int256(uint256(keccak256(abi.encodePacked(nonce, i)))) %!i(MISSING)nt256(1000);](https://github.com/code-423n4/2022-08-olympus/blob/main/src/test/policies/PriceConfig.t.sol#L124 )


