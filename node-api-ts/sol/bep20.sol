// SPDX-License-Identifier: MIT
pragma solidity >=0.6.0;

// ### 什么是ERC-20?
// ERC-20被认为是最重要的以太坊代币之一，是由EIP-20提案投票产生。 ERC-20已成为同质化代币的技术标准。它被用于在以太坊区块链上的智能合约，并提供所有同质化代币所需的基础功能。
// 同质化代币指的是，相同的代币，就像比特币，以太坊等其他数字货币，与之相对的是非同质化代币，也就是唯一的代币。
// ERC-20是一种数字资产，并且可以通过发送和接收的方式进行交易。它最主要的区别就是它是以区块链作为基础载体，在以太坊上生成并且进行交易。
// 在ERC-20当中，有六个主要的函数来承担其所需的基本功能。

contract TestToken { 

    string public name = "Test Token";  //代币名称
    string  public symbol = "ts";
    uint256 public totalSupply_ = 1000000000000000000000000; // 1 million tokens 总体货币供应量
    uint8   public decimals = 18;
    
    event Transfer(
        address indexed _from,
        address indexed _to,
        uint256 _value
    );

    event Approval(
        address indexed _owner,
        address indexed _spender,
        uint256 _value
    );
    
    mapping(address => uint256) public balances;
    mapping(address => mapping(address => uint256)) public allowed;
    
    constructor() public {
        balances[msg.sender] = totalSupply_;
    }

    function totalSupply() public view returns (uint256) {
        return totalSupply_;
    }


    //每个地址的余额，是一张hash表
    function balanceOf(address _owner) public view returns (uint256) {
        return balances[_owner];
    }
    // 从自己的账户发送代币到另外一个地址
    function transfer(address _to, uint256 _value) public returns (bool success) {
        require(balances[msg.sender] >= _value);
        balances[msg.sender] = balances[msg.sender] - _value;
        balances[_to] =  balances[_to] + _value;
        emit Transfer(msg.sender, _to, _value);
        return true;
    }
    //基于某个地址操作自己账户上余额的权限
    function approve(address _spender, uint256 _value) public returns (bool success) {
        allowed[msg.sender][_spender] = _value; 
        emit Approval(msg.sender, _spender, _value);
        return true;
    }
    //为了增加一个代币的默认应用并且在同一个合约内编写功能，我增加了一个airDrop的函数，当被调用到airDrop函数时，我们就会往这个交易发起者的地址上打一笔代币，并且在代币总数目上加上这笔代币
    function airDropToken() public {
        balances[msg.sender] += 100000000000000000000;
        totalSupply_ += 100000000000000000000;
    }
    //通过第三方账户发送代币
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success) {
        require(_value <= balances[_from]);
        require(_value <= allowed[_from][msg.sender]);
        balances[_from] -= _value;
        balances[_to] += _value;
        allowed[_from][msg.sender] -= _value; 
        emit Transfer(_from, _to, _value);
        return true;
    }
    //每个地址允许别人操作余额的记录，也是一张hash表
    function allowance(address _owner, address _spender) public view returns (uint256 remaining) {
        return allowed[_owner][_spender];
    }
}
