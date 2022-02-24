// SPDX-License-Identifier: MIT
pragma solidity >=0.6.2;

contract hello {
 
    constructor() public{

    }

    function say() public view returns (string memory) {
        return "Hello World!";
    }
}