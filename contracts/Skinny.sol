// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.11;

contract Skinny {
 
  address public manager;
  uint public minimumContribution;
  uint public contractBalance = address(this).balance;

  function Campaign (uint minimum, address creator) public {
      manager = creator;
      minimumContribution = minimum;
  }

  function contribute() public payable {
      require(msg.value > minimumContribution);
  }
}