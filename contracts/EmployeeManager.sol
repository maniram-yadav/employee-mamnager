// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

contract EmployeeManager {
  //address owner;
  constructor() public {
  }
  function getContractName() public pure returns (string memory){
    return "EmployeeManager";
  }
}
