<h1 align=center><code>c4udit</code></h1>

## Add
Add `blob/main` folder to get a legal link

![image](https://user-images.githubusercontent.com/84496536/187116396-2aed1fcd-3f45-45f7-881e-177615f90d71.png)


## Introduction

`c4udit` is a static analyzer for solidity contracts based on regular
expressions specifically crafted for [Code4Rena](https://code4rena.com) contests.

It is capable of finding low risk issues and gas optimization documented in the
[c4-common-issues](https://github.com/byterocket/c4-common-issues) repository.

Note that `c4udit` uses [c4-common-issues](https://github.com/byterocket/c4-common-issues)'s issue identifiers.

## Installation

First you need to have the Go toolchain installed. You can find instruction [here](https://go.dev/doc/install).

Then install `c4udit` with:

```
$ go install github.com/byterocket/c4udit@latest
```

To just build the binary:

```
$ git clone https://github.com/byterocket/c4udit
$ cd c4udit/
$ go build .
```

Now you should be able to run `c4udit` with:

```
$ ./c4udit
```

## Usage

```
Usage:
	c4udit [flags] [files...]

Flags:
	-h    Print help text.
	-s    Save report as file.
```

## Example

Running `c4udit` against the `examples` directory:

```
$ ./c4udit examples/
Files analyzed:
- examples/Test.sol

Issues found:
 G001:
  examples/Test.sol::4 => uint256 a = 0;
  examples/Test.sol::12 => for (uint256 i = 0; i < array.length; i++) {

 G002:
  examples/Test.sol::12 => for (uint256 i = 0; i < array.length; i++) {

 G007:
  examples/Test.sol::6 => string b = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";

 G008:
  examples/Test.sol::13 => i = i / 2;

 L001:
  examples/Test.sol::16 => token.transferFrom(msg.sender, address(this), 100);

 L003:
  examples/Test.sol::1 => pragma solidity ^0.8.0;
```

Using the `-s` flag, `c4udit` will create a report in markdown format.
For an example check out the report in the `examples` directory [here](./examples/c4udit-report.md).

## License

Note that this tool is licensed as [free software](./LICENSE)!
