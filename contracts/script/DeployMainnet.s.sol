// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Script.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../src/OutcomeToken.sol";
import "../src/HorizonPerks.sol";
import "../src/FeeSplitter.sol";
import "../src/ResolutionModule.sol";
import "../src/AIOracleAdapter.sol";
import "../src/MarketFactory.sol";
// Import all market types to ensure they compile
import "../src/markets/LimitOrderMarket.sol";
import "../src/markets/MultiChoiceMarket.sol";
import "../src/markets/PooledLiquidityMarket.sol";

/**
 * @title DeployMainnet
 * @notice Mainnet deployment script that uses existing Token3 and redeploys all other contracts
 * @dev Usage:
 *      1. Set HORIZON_TOKEN_ADDRESS in .env to existing Token3 address
 *      2. Run: forge script script/DeployMainnet.s.sol --rpc-url $RPC_URL --broadcast --verify
 * 
 *      This script will:
 *      - Use existing HorizonToken (Token3) from mainnet
 *      - Redeploy all protocol contracts (OutcomeToken, HorizonPerks, FeeSplitter, etc.)
 *      - Redeploy MarketFactory with updated configuration
 *      - Ensure all market types (LimitOrder, MultiChoice, PooledLiquidity) compile correctly
 */
contract DeployMainnet is Script {
    // ============ Deployment Configuration ============

    struct DeploymentConfig {
        // Token configuration
        address horizonTokenAddress;  // Existing token on mainnet
        string outcomeTokenURI;
        // Admin addresses
        address protocolOwner;
        address protocolTreasury;
        address arbitrator;
        address aiSigner;
        // Protocol parameters
        uint256 minCreatorStake;
        uint256 minResolutionBond;
        uint256 disputeWindow;
    }

    // ============ Deployed Contracts ============

    IERC20 public horizonToken;
    OutcomeToken public outcomeToken;
    HorizonPerks public horizonPerks;
    FeeSplitter public feeSplitter;
    ResolutionModule public resolutionModule;
    AIOracleAdapter public aiOracleAdapter;
    MarketFactory public marketFactory;

    // ============ Main Deployment Function ============

    /**
     * @notice Main deployment function that coordinates the entire deployment process
     * @dev Deploys all contracts in the correct order and sets up authorizations
     */
    function run() external {
        // Load configuration from environment
        DeploymentConfig memory config = loadConfig();

        // Validate configuration
        validateConfig(config);

        console.log("\n========================================");
        console.log("  MAINNET DEPLOYMENT - PROJECT GAMMA");
        console.log("========================================\n");
        console.log("Using existing HorizonToken (Token3):", config.horizonTokenAddress);
        console.log("Deployer:", msg.sender);
        console.log("Chain ID:", block.chainid);

        // Start broadcasting transactions
        vm.startBroadcast();

        // 1. Set existing token reference
        console.log("\n=== PHASE 1: TOKEN SETUP ===");
        setupToken(config);

        // 2. Deploy protocol infrastructure
        console.log("\n=== PHASE 2: DEPLOYING PROTOCOL INFRASTRUCTURE ===");
        deployInfrastructure(config);

        // 3. Deploy market creation system
        console.log("\n=== PHASE 3: DEPLOYING MARKET SYSTEM ===");
        deployMarketSystem(config);

        // 4. Configure authorizations and roles
        console.log("\n=== PHASE 4: CONFIGURING AUTHORIZATIONS ===");
        configureAuthorizations(config);

        // 5. Verify deployment
        console.log("\n=== PHASE 5: VERIFYING DEPLOYMENT ===");
        verifyDeployment(config);

        vm.stopBroadcast();

        // 6. Output deployment summary
        console.log("\n=== DEPLOYMENT COMPLETE ===");
        outputDeploymentSummary(config);
    }

    // ============ Deployment Phases ============

    /**
     * @notice Setup reference to existing HorizonToken (Token3)
     * @dev Does not deploy a new token - uses existing mainnet deployment
     */
    function setupToken(DeploymentConfig memory config) internal {
        console.log("Setting up existing HorizonToken reference...");
        horizonToken = IERC20(config.horizonTokenAddress);
        
        // Verify the token exists by checking balance
        uint256 deployerBalance = horizonToken.balanceOf(msg.sender);
        console.log("  Token address:", address(horizonToken));
        console.log("  Deployer balance:", deployerBalance / 1e18, "HORIZON");
        
        require(address(horizonToken).code.length > 0, "Token address has no code");
    }

    /**
     * @notice Deploy protocol infrastructure (OutcomeToken, HorizonPerks, FeeSplitter, ResolutionModule, AIOracleAdapter)
     */
    function deployInfrastructure(DeploymentConfig memory config) internal {
        console.log("Deploying OutcomeToken...");
        outcomeToken = new OutcomeToken(config.outcomeTokenURI);
        console.log("  OutcomeToken deployed at:", address(outcomeToken));

        console.log("Deploying HorizonPerks...");
        horizonPerks = new HorizonPerks(address(horizonToken));
        console.log("  HorizonPerks deployed at:", address(horizonPerks));

        console.log("Deploying FeeSplitter...");
        feeSplitter = new FeeSplitter(config.protocolTreasury);
        console.log("  FeeSplitter deployed at:", address(feeSplitter));

        console.log("Deploying ResolutionModule...");
        resolutionModule = new ResolutionModule(
            address(outcomeToken),
            address(horizonToken),
            config.arbitrator
        );
        console.log("  ResolutionModule deployed at:", address(resolutionModule));

        // Set resolution parameters
        console.log("  Setting resolution parameters...");
        resolutionModule.setMinBond(config.minResolutionBond);
        resolutionModule.setDisputeWindow(config.disputeWindow);

        console.log("Deploying AIOracleAdapter...");
        aiOracleAdapter = new AIOracleAdapter(
            address(resolutionModule),
            address(horizonToken),
            config.aiSigner
        );
        console.log("  AIOracleAdapter deployed at:", address(aiOracleAdapter));
    }

    /**
     * @notice Deploy market creation system (MarketFactory)
     * @dev Factory will support all market types: Standard AMM, LimitOrder, MultiChoice, PooledLiquidity
     */
    function deployMarketSystem(DeploymentConfig memory config) internal {
        console.log("Deploying MarketFactory...");
        marketFactory = new MarketFactory(
            address(outcomeToken),
            address(feeSplitter),
            address(horizonPerks),
            address(horizonToken)
        );
        console.log("  MarketFactory deployed at:", address(marketFactory));

        // Set market creation parameters
        console.log("  Setting market creation parameters...");
        marketFactory.setMinCreatorStake(config.minCreatorStake);
        
        console.log("\n  Supported Market Types:");
        console.log("  - Standard AMM (via MarketAMM)");
        console.log("  - Limit Order Markets (LimitOrderMarket)");
        console.log("  - Multi-Choice Markets (MultiChoiceMarket)");
        console.log("  - Pooled Liquidity Markets (PooledLiquidityMarket)");
    }

    /**
     * @notice Configure all authorizations and role assignments
     */
    function configureAuthorizations(DeploymentConfig memory config) internal {
        console.log("Setting up ResolutionModule authorizations...");
        // ResolutionModule needs to be authorized to set winning outcomes
        outcomeToken.setResolutionAuthorization(address(resolutionModule), true);

        console.log("Transferring OutcomeToken ownership to MarketFactory...");
        // MarketFactory needs ownership to register markets and authorize AMMs
        outcomeToken.transferOwnership(address(marketFactory));

        console.log("Transferring FeeSplitter ownership to MarketFactory...");
        // MarketFactory also needs to register markets in FeeSplitter
        feeSplitter.transferOwnership(address(marketFactory));

        // Transfer ownership to protocol owner if different from deployer
        if (config.protocolOwner != msg.sender && config.protocolOwner != address(0)) {
            console.log("Transferring ownership to protocol owner:", config.protocolOwner);
            
            console.log("  Transferring HorizonPerks ownership...");
            horizonPerks.transferOwnership(config.protocolOwner);

            console.log("  Transferring ResolutionModule ownership...");
            resolutionModule.transferOwnership(config.protocolOwner);

            console.log("  Transferring AIOracleAdapter ownership...");
            aiOracleAdapter.transferOwnership(config.protocolOwner);

            console.log("  Transferring MarketFactory ownership...");
            marketFactory.transferOwnership(config.protocolOwner);
        }
    }

    /**
     * @notice Verify that all contracts are deployed correctly with proper configuration
     */
    function verifyDeployment(DeploymentConfig memory config) internal view {
        console.log("Verifying HorizonToken...");
        require(address(horizonToken) == config.horizonTokenAddress, "HorizonToken address mismatch");

        console.log("Verifying OutcomeToken...");
        require(address(outcomeToken) != address(0), "OutcomeToken not deployed");

        console.log("Verifying HorizonPerks...");
        require(address(horizonPerks) != address(0), "HorizonPerks not deployed");
        require(address(horizonPerks.horizonToken()) == address(horizonToken), "HorizonPerks token mismatch");

        console.log("Verifying FeeSplitter...");
        require(address(feeSplitter) != address(0), "FeeSplitter not deployed");
        require(feeSplitter.protocolTreasury() == config.protocolTreasury, "FeeSplitter treasury mismatch");

        console.log("Verifying ResolutionModule...");
        require(address(resolutionModule) != address(0), "ResolutionModule not deployed");
        require(address(resolutionModule.outcomeToken()) == address(outcomeToken), "ResolutionModule token mismatch");
        require(address(resolutionModule.bondToken()) == address(horizonToken), "ResolutionModule bond token mismatch");
        require(resolutionModule.arbitrator() == config.arbitrator, "ResolutionModule arbitrator mismatch");
        require(resolutionModule.minBond() == config.minResolutionBond, "ResolutionModule min bond mismatch");
        require(resolutionModule.disputeWindow() == config.disputeWindow, "ResolutionModule dispute window mismatch");

        console.log("Verifying AIOracleAdapter...");
        require(address(aiOracleAdapter) != address(0), "AIOracleAdapter not deployed");
        require(address(aiOracleAdapter.resolutionModule()) == address(resolutionModule), "AIOracleAdapter resolution module mismatch");
        require(address(aiOracleAdapter.bondToken()) == address(horizonToken), "AIOracleAdapter bond token mismatch");

        console.log("Verifying MarketFactory...");
        require(address(marketFactory) != address(0), "MarketFactory not deployed");
        require(address(marketFactory.outcomeToken()) == address(outcomeToken), "MarketFactory outcome token mismatch");
        require(address(marketFactory.feeSplitter()) == address(feeSplitter), "MarketFactory fee splitter mismatch");
        require(address(marketFactory.horizonPerks()) == address(horizonPerks), "MarketFactory horizon perks mismatch");
        require(address(marketFactory.horizonToken()) == address(horizonToken), "MarketFactory horizon token mismatch");
        require(marketFactory.minCreatorStake() == config.minCreatorStake, "MarketFactory min stake mismatch");

        console.log("All contracts verified successfully!");
    }

    /**
     * @notice Output deployment summary with all contract addresses and configuration
     */
    function outputDeploymentSummary(DeploymentConfig memory config) internal view {
        console.log("\n==========================================");
        console.log("        DEPLOYMENT SUMMARY");
        console.log("==========================================\n");

        console.log("TOKEN (EXISTING - NOT REDEPLOYED):");
        console.log("  HorizonToken (Token3):  ", address(horizonToken));

        console.log("\nNEWLY DEPLOYED CONTRACTS:");
        console.log("  OutcomeToken:           ", address(outcomeToken));
        console.log("  HorizonPerks:           ", address(horizonPerks));
        console.log("  FeeSplitter:            ", address(feeSplitter));
        console.log("  ResolutionModule:       ", address(resolutionModule));
        console.log("  AIOracleAdapter:        ", address(aiOracleAdapter));
        console.log("  MarketFactory:          ", address(marketFactory));

        console.log("\nADMIN ADDRESSES:");
        console.log("  Protocol Owner:         ", config.protocolOwner);
        console.log("  Protocol Treasury:      ", config.protocolTreasury);
        console.log("  Arbitrator:             ", config.arbitrator);
        console.log("  AI Signer:              ", config.aiSigner);

        console.log("\nPROTOCOL PARAMETERS:");
        console.log("  Min Creator Stake:      ", config.minCreatorStake / 10 ** 18, "HORIZON");
        console.log("  Min Resolution Bond:    ", config.minResolutionBond / 10 ** 18, "HORIZON");
        console.log("  Dispute Window:         ", config.disputeWindow, "seconds");

        console.log("\nSUPPORTED MARKET TYPES:");
        console.log("  1. Standard AMM         (MarketAMM - constant product)");
        console.log("  2. Limit Order Market   (LimitOrderMarket - CLOB)");
        console.log("  3. Multi-Choice Market  (MultiChoiceMarket - LMSR)");
        console.log("  4. Pooled Liquidity     (PooledLiquidityMarket - Uniswap V3 style)");

        console.log("\n==========================================");
        console.log("Copy these addresses to your .env file:");
        console.log("==========================================\n");

        console.log("HORIZON_TOKEN_ADDRESS=", address(horizonToken));
        console.log("OUTCOME_TOKEN_ADDRESS=", address(outcomeToken));
        console.log("HORIZON_PERKS_ADDRESS=", address(horizonPerks));
        console.log("FEE_SPLITTER_ADDRESS=", address(feeSplitter));
        console.log("RESOLUTION_MODULE_ADDRESS=", address(resolutionModule));
        console.log("AI_ORACLE_ADAPTER_ADDRESS=", address(aiOracleAdapter));
        console.log("MARKET_FACTORY_ADDRESS=", address(marketFactory));

        console.log("\n==========================================\n");
    }

    // ============ Configuration Helpers ============

    /**
     * @notice Load deployment configuration from environment variables
     * @return config Populated deployment configuration struct
     */
    function loadConfig() internal view returns (DeploymentConfig memory config) {
        // Token address (REQUIRED - must be set in .env)
        config.horizonTokenAddress = vm.envAddress("HORIZON_TOKEN_ADDRESS");
        require(config.horizonTokenAddress != address(0), "HORIZON_TOKEN_ADDRESS not set in .env");

        // Token configuration
        config.outcomeTokenURI = vm.envOr("OUTCOME_TOKEN_URI", string("https://horizon.markets/api/metadata/{id}.json"));

        // Admin addresses
        config.protocolOwner = vm.envOr("PROTOCOL_OWNER", msg.sender);
        config.protocolTreasury = vm.envOr("PROTOCOL_TREASURY", msg.sender);
        config.arbitrator = vm.envOr("ARBITRATOR_ADDRESS", msg.sender);
        config.aiSigner = vm.envOr("AI_SIGNER_ADDRESS", msg.sender);

        // Protocol parameters
        config.minCreatorStake = vm.envOr("MIN_CREATOR_STAKE", uint256(10_000 * 10 ** 18)); // 10k HORIZON
        config.minResolutionBond = vm.envOr("MIN_RESOLUTION_BOND", uint256(1_000 * 10 ** 18)); // 1k HORIZON
        config.disputeWindow = vm.envOr("DISPUTE_WINDOW", uint256(172800)); // 48 hours
    }

    /**
     * @notice Validate deployment configuration
     * @param config Configuration to validate
     */
    function validateConfig(DeploymentConfig memory config) internal pure {
        require(config.horizonTokenAddress != address(0), "Invalid token address");
        require(bytes(config.outcomeTokenURI).length > 0, "Invalid URI");
        require(config.protocolTreasury != address(0), "Invalid treasury");
        require(config.arbitrator != address(0), "Invalid arbitrator");
        require(config.aiSigner != address(0), "Invalid AI signer");
        require(config.minCreatorStake > 0, "Invalid creator stake");
        require(config.minResolutionBond > 0, "Invalid resolution bond");
        require(config.disputeWindow >= 3600, "Dispute window too short"); // At least 1 hour
        require(config.disputeWindow <= 604800, "Dispute window too long"); // At most 7 days
    }

    // ============ Utility Functions ============

    /**
     * @notice Helper to simulate deployment without broadcasting
     * @dev Useful for testing: forge script script/DeployMainnet.s.sol:DeployMainnet --sig "simulate()"
     */
    function simulate() external view {
        DeploymentConfig memory config = loadConfig();
        validateConfig(config);

        console.log("\n=== SIMULATING MAINNET DEPLOYMENT ===");
        console.log("This is a dry run - no transactions will be broadcast\n");

        console.log("Configuration:");
        console.log("  Existing Token Address:", config.horizonTokenAddress);
        console.log("  Outcome Token URI:", config.outcomeTokenURI);
        console.log("  Protocol Owner:", config.protocolOwner);
        console.log("  Protocol Treasury:", config.protocolTreasury);
        console.log("  Arbitrator:", config.arbitrator);
        console.log("  AI Signer:", config.aiSigner);
        console.log("  Min Creator Stake:", config.minCreatorStake / 10 ** 18, "HORIZON");
        console.log("  Min Resolution Bond:", config.minResolutionBond / 10 ** 18, "HORIZON");
        console.log("  Dispute Window:", config.disputeWindow, "seconds");

        console.log("\nConfiguration validated successfully!");
        console.log("Ready to deploy with: forge script script/DeployMainnet.s.sol:DeployMainnet --rpc-url <RPC_URL> --broadcast --verify");
    }
}
