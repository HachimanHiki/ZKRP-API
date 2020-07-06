pragma solidity ^0.4.24;

import "./Ownable.sol";

contract Health is Ownable{
    string public commitment;
    string public merkleTreeRoot;

    // "0x19be17213c2ee781defeef4abc2d6964f1418177f8d54418c55412e198eebf2107c6e8cc1a43e9be4e7f331f7b729e53a2e14e37aa874383628cf89be3cd0ef6"
    // "d248f0355b955dad7d88be03cf279654bd8ebbbbc8d6302ae19ff34c26143eae"
    constructor(string c, string m) public{
        commitment = c;
        merkleTreeRoot = m;
    }

    function updateCommitment(string c) public onlyOwner{
        commitment = c;
    }

    function updateMerkleTreeRoot(string m) public onlyOwner{
        merkleTreeRoot = m;
    }
}