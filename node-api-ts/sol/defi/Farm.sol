pragma solidity >=0.6.0;
// Import contracts for both Dapp and DAI token.
import "./RevenueToken.sol";
import "./ChainToken.sol";
// 智能合约实现质押分红
// https://learnblockchain.cn/article/3152
contract TokenFarm{
    string public name = "TokenFarm";
    address public owner;
    RevenueToken public revenueToken;
    ChainToken public chainToken;

    mapping(address=>uint) public stakingBalance;
    mapping(address=>bool) public hasStaked;
    mapping(address=>bool) public isStaking;
    address[] public staker;

    constructor (
        RevenueToken _revenueToken,
        ChainToken _chainToken
    ) public{
        revenueToken = _revenueToken;
        chainToken = _chainToken;
        owner = msg.sender; // address of the owner of the contract
    }
    
    /// @param _amount The amount of the tokens you want to stake.
    /// stakeToken方法是将CT抵押给这个合约
    function stakeToken(uint _amount) public {
        
        // check, amount should be greater than zero. There should be some tokens to be staked.
        require(_amount>0,"amount need to be more than 0");         
        
        // this refers to the instance of the contract where the call is made (you can have multiple instances of the same contract).
        // address(this) refers to the address of the instance of the contract where the call is being made.
        // msg. sender refers to the address where the contract is being called from.
        // @param _amount, the amount of tokens you want to stake .
        chainToken.transferFrom(msg.sender, address(this), _amount); 
        
        // The balance of the owner of the contract, after staking the coins.
        stakingBalance[msg.sender] = stakingBalance[msg.sender] + _amount;
        
        if(!hasStaked[msg.sender]){
            staker.push(msg.sender);
        }
        isStaking[msg.sender] = true;
        hasStaked[msg.sender] = true;
    }
       
       
    //require checks if the condition is true, thows the exceptionotherwise 'trader is not owner'.
    //if the require condition is true, then all the tokens that are staked, are unstaked .
    // untakeTokens是取出抵押的数字货币
    function unstakeToken() public {
        require(isStaking[msg.sender] == true,"You have nothing to unstake.");
        uint balance = stakingBalance[msg.sender];
        stakingBalance[msg.sender] = 0;
        chainToken.transfer(msg.sender,balance);
        isStaking[msg.sender] = false;
    }
    //@param _owner is the address of the owner which is msg.sender
    // returns the staking balance
    function stakeAmount(address _owner) public view returns(uint) {
        return stakingBalance[_owner];
    }
    /// issueToken是根据抵押的CT产生RT利息
    function issusToken() public {
        require(msg.sender==owner,"trader is not owner");
        for(uint i=0; i<staker.length;i++){
            address recipient = staker[i];
            if(isStaking[recipient] == true){
                uint balance = stakingBalance[recipient];
                revenueToken.transfer(recipient, balance);
            }
        }
    }

}