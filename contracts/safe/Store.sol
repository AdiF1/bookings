// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.11;

contract MyStore {
    event ItemSet(bytes32 key, bytes32 value);

    string public version;
    string public adf = "007";
    
    mapping (bytes32 => bytes32) public items;

    function MyVersion(string memory _version ) public {
        version = _version;
    }

    function setItem(bytes32 key, bytes32 value) external {
        items[key] = value;
        emit ItemSet(key, value);
    }

    function Hello() public pure returns (string memory) {
        return "Hello Store Front!";
    }
   
}

// solc --optimize --abi ./contracts/Store.sol -o build --overwrite
// solc --optimize --bin ./contracts/Store.sol -o build 
// abigen --abi=./build/MyStore.abi --bin=./build/MyStore.bin --pkg=Store --out=./contracts/Store.go