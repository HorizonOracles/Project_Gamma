// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title HorizonToken
 * @notice Horizon utility token for Project Gamma
 * @dev Used for:
 *      - Creator stakes (locked when creating markets)
 *      - Fee tier discounts (higher balance = lower fees)
 *      - Dispute bonds (locked during dispute period)
 *      - AI proposal bonds (locked when AI proposes outcomes)
 */
contract HorizonToken is ERC20, ERC20Burnable, Ownable {
    // ============ Events ============

    event MinterAdded(address indexed minter);
    event MinterRemoved(address indexed minter);

    // ============ State Variables ============

    /// @notice Mapping of addresses authorized to mint tokens
    mapping(address => bool) public minters;

    /// @notice Maximum supply cap (10 billion tokens)
    uint256 public constant MAX_SUPPLY = 10_000_000_000 * 10 ** 18;

    // ============ Constructor ============

    /**
     * @notice Initializes the HORIZON token
     * @param initialSupply Initial supply to mint to deployer
     */
    constructor(uint256 initialSupply) ERC20("Horizon", "HORIZON") Ownable(msg.sender) {
        require(initialSupply <= MAX_SUPPLY, "Initial supply exceeds max supply");
        _mint(msg.sender, initialSupply);
    }

    // ============ Admin Functions ============

    /**
     * @notice Adds an address as an authorized minter
     * @param minter Address to authorize
     */
    function addMinter(address minter) external onlyOwner {
        minters[minter] = true;
        emit MinterAdded(minter);
    }

    /**
     * @notice Removes an address from authorized minters
     * @param minter Address to deauthorize
     */
    function removeMinter(address minter) external onlyOwner {
        minters[minter] = false;
        emit MinterRemoved(minter);
    }

    // ============ Minting Functions ============

    /**
     * @notice Mints new tokens (only callable by authorized minters or owner)
     * @param to Address to mint tokens to
     * @param amount Amount of tokens to mint
     */
    function mint(address to, uint256 amount) external {
        require(minters[msg.sender] || msg.sender == owner(), "Not authorized to mint");
        require(totalSupply() + amount <= MAX_SUPPLY, "Exceeds max supply");
        _mint(to, amount);
    }

    /**
     * @notice Burns tokens from the caller's balance
     * @dev Overrides ERC20Burnable to add supply tracking
     * @param amount Amount of tokens to burn
     */
    function burn(uint256 amount) public override {
        super.burn(amount);
    }

    /**
     * @notice Burns tokens from a specific address (requires allowance)
     * @dev Overrides ERC20Burnable to add supply tracking
     * @param account Address to burn tokens from
     * @param amount Amount of tokens to burn
     */
    function burnFrom(address account, uint256 amount) public override {
        super.burnFrom(account, amount);
    }
}
