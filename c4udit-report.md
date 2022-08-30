# Gas | QA

## ✅ G-X: Don't Initialize Variables with Default Value

### 📝 Description

Uninitialized variables are assigned with the types default value. Explicitly initializing a variable with it's default value costs unnecesary gas. Not overwriting the default for [stack variables](https://gist.github.com/IllIllI000/e075d189c1b23dce256cd166e28f3397) saves 8 gas. Storage and memory variables have larger savings

### 💡 Recommendation

Delete useless variable declarations to save gas.

### 🔍 Findings:

`examples/Test.sol#L4` [uint256 a = 0;](https://github.com/code-423n4/examples/Test.sol#L4)

`examples/Test.sol#L13` [for (uint256 i = 0; i < array.length; i++) {](https://github.com/code-423n4/examples/Test.sol#L13)

## ✅ G-X: Cache Array Length Outside of Loop

### 📝 Description

Caching the array length outside a loop saves reading it on each iteration, as long as the array's length is not changed during the loop.
You save 3 gas by not reading `array.length` - 3 gas per instance - 27 gas saved

### 💡 Recommendation

Reco

### 🔍 Findings:

`examples/Test.sol#L12` [require(array.length > 0,"Edfdafjldka;fjadsljfdaslkjflkas;aaasjkjadlkfjdslalldkjadl");](https://github.com/code-423n4/examples/Test.sol#L12)

`examples/Test.sol#L13` [for (uint256 i = 0; i < array.length; i++) {](https://github.com/code-423n4/examples/Test.sol#L13)

## ✅ G-X: Use != 0 instead of > 0 for Unsigned Integer Comparison

### 📝 Description

Use != 0 when comparing uint variables to zero, which cannot hold values below zero

### 💡 Recommendation

You should change from `> 0` to `!=0`.

### 🔍 Findings:

`examples/Test.sol#L12` [require(array.length > 0,"Edfdafjldka;fjadsljfdaslkjflkas;aaasjkjadlkfjdslalldkjadl");](https://github.com/code-423n4/examples/Test.sol#L12)

## ✅ G-X: Long Revert Strings

### 📝 Description

Shortening revert strings to fit in 32 bytes will decrease gas costs for deployment and gas costs when the revert condition has been met.

### 💡 Recommendation

Use short statements instead of long statement.

### 🔍 Findings:

`examples/Test.sol#L6` [string b = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";](https://github.com/code-423n4/examples/Test.sol#L6)

`examples/Test.sol#L12` [require(array.length > 0,"Edfdafjldka;fjadsljfdaslkjflkas;aaasjkjadlkfjdslalldkjadl");](https://github.com/code-423n4/examples/Test.sol#L12)

## ✅ G-X: Use Shift Right/Left instead of Division/Multiplication if possible

### 📝 Description

A division/multiplication by any number `x` being a power of 2 can be calculated by shifting `log2(x)` to the right/left.
While the `DIV` opcode uses 5 gas, the `SHR` opcode only uses 3 gas.
urthermore, Solidity's division operation also includes a division-by-0 prevention which is bypassed using shifting.

### 💡 Recommendation

Reco

### 🔍 Findings:

`examples/Test.sol#L14` [i = i / 2;](https://github.com/code-423n4/examples/Test.sol#L14)

## ✅ G-X: Use assembly for calculation (add, sub, mul, div)

### 📝 Description

See [this issue](https://github.com/code-423n4/2022-06-putty-findings/issues/59)
You can save about 4000 gas per instance in deploying contracrt.
You can save about 40~60 gas per instance if using assembly to execute the function.

### 💡 Recommendation

Please follow [this link](https://gist.github.com/Tomosuke0930/a18e03c4ef20ff284dfca62e7ccf6f91) to make corrections

### 🔍 Findings:

`examples/Test.sol#L12` [require(array.length > 0,"Edfdafjldka;fjadsljfdaslkjflkas;aaasjkjadlkfjdslalldkjadl");](https://github.com/code-423n4/examples/Test.sol#L12)
