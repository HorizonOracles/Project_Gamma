// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Script.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../src/OutcomeToken.sol";
import "../src/HorizonPerks.sol";
import "../src/FeeSplitter.sol";
import "../src/ResolutionModule.sol";
import "../src/MarketFactory.sol";
import "../test/mocks/MockERC20.sol";
// Import market types to ensure they compile
import "../src/markets/LimitOrderMarket.sol";
import "../src/markets/MultiChoiceMarket.sol";
import "../src/markets/PooledLiquidityMarket.sol";

/**
 * @title DeployLocal
 * @notice Deployment script for local Anvil testing environment
 * @dev Deploys all core contracts + mock tokens for development and testing
 *      Use this script for testing new market types on a local Anvil node
 * 
 * Usage:
 *   1. Start Anvil in terminal 1:
 *      anvil
 * 
 *   2. Deploy contracts in terminal 2:
 *      forge script script/DeployLocal.s.sol --fork-url http://localhost:8545 --broadcast
 * 
 *   3. Save the deployed addresses from output for testing
 */
contract DeployLocal is Script {
    // ============ Deployed Contracts ============

    IERC20 public horizonToken;
    OutcomeToken public outcomeToken;
    HorizonPerks public horizonPerks;
    FeeSplitter public feeSplitter;
    ResolutionModule public resolutionModule;
    MarketFactory public marketFactory;
    
    // Mock tokens for testing
    MockERC20 public usdc;
    MockERC20 public usdt;
    
    // Use external Token3 address (set via environment variable or parameter)
    address public tokenAddress;

    // ============ Configuration ============

    // Anvil default accounts
    address public deployer = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266; // Anvil account #0
    address public treasury = 0x70997970C51812dc3A010C7d01b50e0d17dc79C8; // Anvil account #1
    address public arbitrator = 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC; // Anvil account #2
    
    // Protocol parameters (relaxed for local testing)
    uint256 public constant HORIZON_INITIAL_SUPPLY = 1_000_000_000 ether;
    string public constant OUTCOME_TOKEN_URI = "https://localhost/{id}";
    uint256 public constant MIN_CREATOR_STAKE = 100 ether;
    uint256 public constant MIN_RESOLUTION_BOND = 1000 ether;
    uint256 public constant DISPUTE_WINDOW = 1 hours; // Short window for testing

    // ============ Main Deployment Function ============

    function run() external {
        // Get Token3 address from environment variable
        tokenAddress = vm.envOr("TOKEN_ADDRESS", address(0));
        require(tokenAddress != address(0), "TOKEN_ADDRESS environment variable not set");
        
        horizonToken = IERC20(tokenAddress);
        
        console.log("\n=== DEPLOYING TO LOCAL ANVIL ===");
        console.log("Deployer:", deployer);
        console.log("Chain ID:", block.chainid);
        console.log("Using Token3 at:", tokenAddress);
        
        vm.startBroadcast();

        // Phase 1: Deploy mock tokens
        console.log("\n=== PHASE 1: DEPLOYING MOCK TOKENS ===");
        deployMockTokens();

        // Phase 2: Deploy core contracts
        console.log("\n=== PHASE 2: DEPLOYING CORE CONTRACTS ===");
        deployCoreContracts();

        // Phase 3: Configure authorizations
        console.log("\n=== PHASE 3: CONFIGURING AUTHORIZATIONS ===");
        configureAuthorizations();

        // Phase 4: Setup test environment
        console.log("\n=== PHASE 4: SETTING UP TEST ENVIRONMENT ===");
        setupTestEnvironment();

        vm.stopBroadcast();

        // Output deployment info
        console.log("\n=== DEPLOYMENT COMPLETE ===");
        outputDeploymentSummary();
    }

    // ============ Deployment Functions ============

    /**
     * @notice Deploy mock ERC20 tokens for testing
     */
    function deployMockTokens() internal {
        console.log("Deploying USDC mock...");
        usdc = new MockERC20("USD Coin", "USDC");
        console.log("  USDC deployed at:", address(usdc));

        console.log("Deploying USDT mock...");
        usdt = new MockERC20("Tether USD", "USDT");
        console.log("  USDT deployed at:", address(usdt));
    }

    /**
     * @notice Deploy core protocol contracts
     */
    function deployCoreContracts() internal {
        // 1. HorizonToken already deployed (Token3), using address from environment
        console.log("Using HorizonToken (Token3) at:", address(horizonToken));

        // 2. Deploy OutcomeToken
        console.log("Deploying OutcomeToken...");
        outcomeToken = new OutcomeToken(OUTCOME_TOKEN_URI);
        console.log("  OutcomeToken deployed at:", address(outcomeToken));

        // 3. Deploy HorizonPerks
        console.log("Deploying HorizonPerks...");
        horizonPerks = new HorizonPerks(address(horizonToken));
        console.log("  HorizonPerks deployed at:", address(horizonPerks));

        // 4. Deploy FeeSplitter
        console.log("Deploying FeeSplitter...");
        feeSplitter = new FeeSplitter(treasury);
        console.log("  FeeSplitter deployed at:", address(feeSplitter));

        // 5. Deploy ResolutionModule
        console.log("Deploying ResolutionModule...");
        resolutionModule = new ResolutionModule(
            address(outcomeToken),
            address(horizonToken),
            arbitrator
        );
        console.log("  ResolutionModule deployed at:", address(resolutionModule));
        
        // Configure resolution parameters
        resolutionModule.setMinBond(MIN_RESOLUTION_BOND);
        resolutionModule.setDisputeWindow(DISPUTE_WINDOW);
        console.log("  Resolution parameters configured");

        // 6. Deploy MarketFactory
        console.log("Deploying MarketFactory...");
        marketFactory = new MarketFactory(
            address(outcomeToken),
            address(feeSplitter),
            address(horizonPerks),
            address(horizonToken)
        );
        console.log("  MarketFactory deployed at:", address(marketFactory));
        
        // Configure factory parameters
        marketFactory.setMinCreatorStake(MIN_CREATOR_STAKE);
        console.log("  Factory parameters configured");
    }

    /**
     * @notice Configure contract authorizations
     */
    function configureAuthorizations() internal {
        console.log("Setting ResolutionModule authorization...");
        outcomeToken.setResolutionAuthorization(address(resolutionModule), true);
        console.log("  ResolutionModule authorized");

        console.log("Transferring OutcomeToken ownership to MarketFactory...");
        outcomeToken.transferOwnership(address(marketFactory));
        console.log("  Ownership transferred");
    }

    /**
     * @notice Setup test environment with funded accounts
     */
    function setupTestEnvironment() internal {
        // Anvil test accounts (accounts 3-9 for testing)
        address[] memory testAccounts = new address[](7);
        testAccounts[0] = 0x90F79bf6EB2c4f870365E785982E1f101E93b906; // Account #3
        testAccounts[1] = 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65; // Account #4
        testAccounts[2] = 0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc; // Account #5
        testAccounts[3] = 0x976EA74026E726554dB657fA54763abd0C3a0aa9; // Account #6
        testAccounts[4] = 0x14dC79964da2C08b23698B3D3cc7Ca32193d9955; // Account #7
        testAccounts[5] = 0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f; // Account #8
        testAccounts[6] = 0xa0Ee7A142d267C1f36714E4a8F75612F20a79720; // Account #9

        console.log("Funding test accounts with tokens...");
        
        // Fund each test account with tokens
        for (uint256 i = 0; i < testAccounts.length; i++) {
            address account = testAccounts[i];
            
            // Give each account 1M USDC, 1M USDT
            usdc.mint(account, 1_000_000 ether);
            usdt.mint(account, 1_000_000 ether);
            
            // Transfer HORIZON tokens (using Token3 transfer which should work in NORMAL mode)
            horizonToken.transfer(account, 100_000 ether);
            
            console.log("  Funded account:", account);
        }
        
        console.log("Test environment ready!");
    }

    // ============ Output Functions ============

    /**
     * @notice Output deployment summary
     */
    function outputDeploymentSummary() internal view {
        console.log("\n========================================");
        console.log("DEPLOYMENT SUMMARY");
        console.log("========================================");
        console.log("");
        console.log("Core Contracts:");
        console.log("  HorizonToken:       ", address(horizonToken));
        console.log("  OutcomeToken:       ", address(outcomeToken));
        console.log("  HorizonPerks:       ", address(horizonPerks));
        console.log("  FeeSplitter:        ", address(feeSplitter));
        console.log("  ResolutionModule:   ", address(resolutionModule));
        console.log("  MarketFactory:      ", address(marketFactory));
        console.log("");
        console.log("Mock Tokens:");
        console.log("  USDC:               ", address(usdc));
        console.log("  USDT:               ", address(usdt));
        console.log("");
        console.log("Admin Addresses:");
        console.log("  Deployer:           ", deployer);
        console.log("  Treasury:           ", treasury);
        console.log("  Arbitrator:         ", arbitrator);
        console.log("");
        console.log("Configuration:");
        console.log("  Min Creator Stake:  ", MIN_CREATOR_STAKE);
        console.log("  Min Resolution Bond:", MIN_RESOLUTION_BOND);
        console.log("  Dispute Window:     ", DISPUTE_WINDOW, "seconds");
        console.log("");
        console.log("========================================");
        console.log("Next Steps:");
        console.log("1. Save these addresses for your tests");
        console.log("2. Use Anvil accounts #3-9 for testing");
        console.log("3. Create markets via MarketFactory");
        console.log("========================================");
    }

    // ============ Helper Functions ============

    /**
     * @notice Helper to create a test market (useful for quick testing)
     * @param collateralToken Collateral token address
     * @param closeTime Market close time
     * @param category Market category
     * @param metadataURI Market metadata URI
     * @return marketId Created market ID
     */
    function createTestMarket(
        address collateralToken,
        uint256 closeTime,
        string memory category,
        string memory metadataURI
    ) external returns (uint256 marketId) {
        require(address(marketFactory) != address(0), "Deploy contracts first");
        
        // Approve HORIZON tokens
        horizonToken.approve(address(marketFactory), MIN_CREATOR_STAKE);
        
        // Create market
        MarketFactory.MarketParams memory params = MarketFactory.MarketParams({
            collateralToken: collateralToken,
            closeTime: closeTime,
            category: category,
            metadataURI: metadataURI,
            creatorStake: MIN_CREATOR_STAKE
        });
        
        marketId = marketFactory.createMarket(params);
        console.log("Created test market:", marketId);
    }
}
