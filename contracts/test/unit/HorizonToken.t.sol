// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Test.sol";
import "../../src/HorizonToken.sol";

contract HorizonTokenTest is Test {
    HorizonToken public token;

    address public owner = address(this);
    address public minter = address(0x1);
    address public user1 = address(0x2);
    address public user2 = address(0x3);
    address public unauthorized = address(0x4);

    uint256 public constant INITIAL_SUPPLY = 1_000_000_000 * 10 ** 18; // 1 billion

    event MinterAdded(address indexed minter);
    event MinterRemoved(address indexed minter);

    function setUp() public {
        token = new HorizonToken(INITIAL_SUPPLY);
        token.addMinter(minter);
    }

    // ============ Constructor Tests ============

    function test_Constructor() public view {
        assertEq(token.name(), "Horizon");
        assertEq(token.symbol(), "HORIZON");
        assertEq(token.decimals(), 18);
        assertEq(token.totalSupply(), INITIAL_SUPPLY);
        assertEq(token.balanceOf(owner), INITIAL_SUPPLY);
        assertEq(token.owner(), owner);
    }

    function test_Constructor_MaxSupply() public view {
        assertEq(token.MAX_SUPPLY(), 10_000_000_000 * 10 ** 18);
    }

    function test_RevertWhen_Constructor_ExceedsMaxSupply() public {
        uint256 maxSupply = 10_000_000_000 * 10 ** 18;
        vm.expectRevert("Initial supply exceeds max supply");
        new HorizonToken(maxSupply + 1);
    }

    // ============ Minter Management Tests ============

    function test_AddMinter() public {
        address newMinter = address(0x5);

        vm.expectEmit(true, false, false, false);
        emit MinterAdded(newMinter);

        token.addMinter(newMinter);
        assertTrue(token.minters(newMinter));
    }

    function test_RemoveMinter() public {
        vm.expectEmit(true, false, false, false);
        emit MinterRemoved(minter);

        token.removeMinter(minter);
        assertFalse(token.minters(minter));
    }

    function test_RevertWhen_AddMinter_Unauthorized() public {
        vm.prank(unauthorized);
        vm.expectRevert();
        token.addMinter(address(0x5));
    }

    function test_RevertWhen_RemoveMinter_Unauthorized() public {
        vm.prank(unauthorized);
        vm.expectRevert();
        token.removeMinter(minter);
    }

    // ============ Minting Tests ============

    function test_Mint_ByOwner() public {
        uint256 amount = 1000 * 10 ** 18;
        token.mint(user1, amount);
        assertEq(token.balanceOf(user1), amount);
        assertEq(token.totalSupply(), INITIAL_SUPPLY + amount);
    }

    function test_Mint_ByAuthorizedMinter() public {
        uint256 amount = 1000 * 10 ** 18;

        vm.prank(minter);
        token.mint(user1, amount);

        assertEq(token.balanceOf(user1), amount);
        assertEq(token.totalSupply(), INITIAL_SUPPLY + amount);
    }

    function test_RevertWhen_Mint_Unauthorized() public {
        uint256 amount = 1000 * 10 ** 18;

        vm.prank(unauthorized);
        vm.expectRevert("Not authorized to mint");
        token.mint(user1, amount);
    }

    function test_RevertWhen_Mint_ExceedsMaxSupply() public {
        uint256 remaining = token.MAX_SUPPLY() - token.totalSupply();

        vm.expectRevert("Exceeds max supply");
        token.mint(user1, remaining + 1);
    }

    function testFuzz_Mint(uint256 amount) public {
        uint256 maxMintable = token.MAX_SUPPLY() - token.totalSupply();
        vm.assume(amount > 0 && amount <= maxMintable);

        token.mint(user1, amount);
        assertEq(token.balanceOf(user1), amount);
    }

    // ============ Burning Tests ============

    function test_Burn() public {
        uint256 burnAmount = 1000 * 10 ** 18;

        token.burn(burnAmount);

        assertEq(token.balanceOf(owner), INITIAL_SUPPLY - burnAmount);
        assertEq(token.totalSupply(), INITIAL_SUPPLY - burnAmount);
    }

    function test_Burn_ReducesTotalSupply() public {
        uint256 burnAmount = 1000 * 10 ** 18;
        uint256 initialTotal = token.totalSupply();

        token.burn(burnAmount);

        assertEq(token.totalSupply(), initialTotal - burnAmount);
    }

    function test_BurnFrom() public {
        uint256 burnAmount = 1000 * 10 ** 18;

        // Transfer tokens to user1
        token.transfer(user1, burnAmount);

        // User1 approves owner to burn
        vm.prank(user1);
        token.approve(owner, burnAmount);

        // Owner burns from user1
        token.burnFrom(user1, burnAmount);

        assertEq(token.balanceOf(user1), 0);
        assertEq(token.totalSupply(), INITIAL_SUPPLY - burnAmount);
    }

    function test_RevertWhen_Burn_InsufficientBalance() public {
        vm.prank(user1);
        vm.expectRevert();
        token.burn(100 * 10 ** 18);
    }

    function testFuzz_Burn(uint256 amount) public {
        vm.assume(amount > 0 && amount <= INITIAL_SUPPLY);

        token.burn(amount);
        assertEq(token.balanceOf(owner), INITIAL_SUPPLY - amount);
        assertEq(token.totalSupply(), INITIAL_SUPPLY - amount);
    }

    // ============ Transfer Tests ============

    function test_Transfer() public {
        uint256 amount = 1000 * 10 ** 18;

        token.transfer(user1, amount);

        assertEq(token.balanceOf(user1), amount);
        assertEq(token.balanceOf(owner), INITIAL_SUPPLY - amount);
    }

    function test_TransferFrom() public {
        uint256 amount = 1000 * 10 ** 18;

        // Approve user2 to spend owner's tokens
        token.approve(user2, amount);

        // User2 transfers from owner to user1
        vm.prank(user2);
        token.transferFrom(owner, user1, amount);

        assertEq(token.balanceOf(user1), amount);
        assertEq(token.balanceOf(owner), INITIAL_SUPPLY - amount);
    }

    function testFuzz_Transfer(uint256 amount) public {
        vm.assume(amount > 0 && amount <= INITIAL_SUPPLY);

        token.transfer(user1, amount);
        assertEq(token.balanceOf(user1), amount);
    }

    // ============ Approval Tests ============

    function test_Approve() public {
        uint256 amount = 1000 * 10 ** 18;

        token.approve(user1, amount);

        assertEq(token.allowance(owner, user1), amount);
    }

    // ============ Integration Tests ============

    function test_MintAndBurn_MaintainsSupplyCap() public {
        uint256 mintAmount = 1_000_000 * 10 ** 18;

        // Mint tokens
        token.mint(user1, mintAmount);
        uint256 supplyAfterMint = token.totalSupply();

        // Burn tokens
        vm.prank(user1);
        token.burn(mintAmount / 2);

        assertEq(token.totalSupply(), supplyAfterMint - (mintAmount / 2));
    }

    function test_MultipleMinters() public {
        address minter2 = address(0x6);
        token.addMinter(minter2);

        uint256 amount = 1000 * 10 ** 18;

        vm.prank(minter);
        token.mint(user1, amount);

        vm.prank(minter2);
        token.mint(user2, amount);

        assertEq(token.balanceOf(user1), amount);
        assertEq(token.balanceOf(user2), amount);
    }

    function test_MinterRemovalPreventsMinting() public {
        token.removeMinter(minter);

        vm.prank(minter);
        vm.expectRevert("Not authorized to mint");
        token.mint(user1, 1000 * 10 ** 18);
    }

    function test_FullLifecycle() public {
        // 1. Transfer to user
        token.transfer(user1, 10000 * 10 ** 18);

        // 2. User approves user2
        vm.prank(user1);
        token.approve(user2, 5000 * 10 ** 18);

        // 3. User2 transfers from user1
        vm.prank(user2);
        token.transferFrom(user1, user2, 3000 * 10 ** 18);

        // 4. User2 burns some tokens
        vm.prank(user2);
        token.burn(1000 * 10 ** 18);

        // Verify final balances
        assertEq(token.balanceOf(user1), 7000 * 10 ** 18);
        assertEq(token.balanceOf(user2), 2000 * 10 ** 18);
        assertEq(token.totalSupply(), INITIAL_SUPPLY - 1000 * 10 ** 18);
    }
}
