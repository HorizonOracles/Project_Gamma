// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Script.sol";
import "../src/HorizonToken.sol";
import "../src/OutcomeToken.sol";
import "../src/HorizonPerks.sol";
import "../src/FeeSplitter.sol";
import "../src/ResolutionModule.sol";
import "../src/AIOracleAdapter.sol";
import "../src/MarketFactory.sol";

/**
 * @title Deploy
 * @notice Comprehensive deployment script for Horizon prediction market protocol
 * @dev Handles sequential deployment with proper dependency management and authorization setup
 */
contract Deploy is Script {
    // ============ Deployment Configuration ============

    struct DeploymentConfig {
        // Token configuration
        uint256 horizonInitialSupply;
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

    HorizonToken public horizonToken;
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

        // Start broadcasting transactions
        vm.startBroadcast();

        // 1. Deploy core token contracts
        console.log("\n=== PHASE 1: DEPLOYING CORE TOKENS ===");
        deployTokens(config);

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
     * @notice Deploy core token contracts (HorizonToken, OutcomeToken, HorizonPerks)
     */
    function deployTokens(DeploymentConfig memory config) internal {
        console.log("Deploying HorizonToken...");
        horizonToken = new HorizonToken(config.horizonInitialSupply);
        console.log("  HorizonToken deployed at:", address(horizonToken));

        console.log("Deploying OutcomeToken...");
        outcomeToken = new OutcomeToken(config.outcomeTokenURI);
        console.log("  OutcomeToken deployed at:", address(outcomeToken));

        console.log("Deploying HorizonPerks...");
        horizonPerks = new HorizonPerks(address(horizonToken));
        console.log("  HorizonPerks deployed at:", address(horizonPerks));
    }

    /**
     * @notice Deploy protocol infrastructure (FeeSplitter, ResolutionModule, AIOracleAdapter)
     */
    function deployInfrastructure(DeploymentConfig memory config) internal {
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
    }

    /**
     * @notice Configure all authorizations and role assignments
     */
    function configureAuthorizations(DeploymentConfig memory config) internal {
        console.log("Setting up MarketFactory authorizations...");
        // MarketFactory needs to be authorized as a minter for HorizonToken (for staking)
        horizonToken.addMinter(address(marketFactory));

        console.log("Setting up ResolutionModule authorizations...");
        // ResolutionModule needs to be authorized to set winning outcomes
        outcomeToken.setResolutionAuthorization(address(resolutionModule), true);

        console.log("Transferring OutcomeToken and FeeSplitter ownership to MarketFactory...");
        // MarketFactory needs ownership to register markets and authorize AMMs
        outcomeToken.transferOwnership(address(marketFactory));
        // MarketFactory also needs to register markets in FeeSplitter
        feeSplitter.transferOwnership(address(marketFactory));

        console.log("Setting up ownership transfers...");
        // Transfer ownership to protocol owner if different from deployer
        if (config.protocolOwner != msg.sender && config.protocolOwner != address(0)) {
            console.log("  Transferring HorizonToken ownership to:", config.protocolOwner);
            horizonToken.transferOwnership(config.protocolOwner);

            console.log("  Transferring HorizonPerks ownership to:", config.protocolOwner);
            horizonPerks.transferOwnership(config.protocolOwner);

            console.log("  Transferring ResolutionModule ownership to:", config.protocolOwner);
            resolutionModule.transferOwnership(config.protocolOwner);

            console.log("  Transferring AIOracleAdapter ownership to:", config.protocolOwner);
            aiOracleAdapter.transferOwnership(config.protocolOwner);

            console.log("  Transferring MarketFactory ownership to:", config.protocolOwner);
            marketFactory.transferOwnership(config.protocolOwner);
        }
    }

    /**
     * @notice Verify that all contracts are deployed correctly with proper configuration
     */
    function verifyDeployment(DeploymentConfig memory config) internal view {
        console.log("Verifying HorizonToken...");
        require(address(horizonToken) != address(0), "HorizonToken not deployed");
        require(horizonToken.totalSupply() == config.horizonInitialSupply, "HorizonToken supply mismatch");
        require(horizonToken.MAX_SUPPLY() == 10_000_000_000 * 10 ** 18, "HorizonToken max supply mismatch");

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

        console.log("CORE TOKENS:");
        console.log("  HorizonToken:           ", address(horizonToken));
        console.log("  OutcomeToken:           ", address(outcomeToken));
        console.log("  HorizonPerks:           ", address(horizonPerks));

        console.log("\nPROTOCOL INFRASTRUCTURE:");
        console.log("  FeeSplitter:            ", address(feeSplitter));
        console.log("  ResolutionModule:       ", address(resolutionModule));
        console.log("  AIOracleAdapter:        ", address(aiOracleAdapter));

        console.log("\nMARKET SYSTEM:");
        console.log("  MarketFactory:          ", address(marketFactory));

        console.log("\nADMIN ADDRESSES:");
        console.log("  Protocol Owner:         ", config.protocolOwner);
        console.log("  Protocol Treasury:      ", config.protocolTreasury);
        console.log("  Arbitrator:             ", config.arbitrator);
        console.log("  AI Signer:              ", config.aiSigner);

        console.log("\nPROTOCOL PARAMETERS:");
        console.log("  HORIZON Initial Supply: ", config.horizonInitialSupply / 10 ** 18, "HORIZON");
        console.log("  Min Creator Stake:      ", config.minCreatorStake / 10 ** 18, "HORIZON");
        console.log("  Min Resolution Bond:    ", config.minResolutionBond / 10 ** 18, "HORIZON");
        console.log("  Dispute Window:         ", config.disputeWindow, "seconds");

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
        // Token configuration
        config.horizonInitialSupply = vm.envOr("HORIZON_INITIAL_SUPPLY", uint256(100_000_000 * 10 ** 18)); // 100M default
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
        require(config.horizonInitialSupply > 0, "Invalid initial supply");
        require(config.horizonInitialSupply <= 10_000_000_000 * 10 ** 18, "Initial supply exceeds max");
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
     * @notice Helper to deploy to a specific network
     * @dev Can be called with: forge script script/Deploy.s.sol:Deploy --sig "deployToNetwork(string)" "bsc"
     */
    function deployToNetwork(string memory network) external {
        console.log("Deploying to network:", network);
        this.run();
    }

    /**
     * @notice Helper to simulate deployment without broadcasting
     * @dev Useful for testing: forge script script/Deploy.s.sol:Deploy --sig "simulate()"
     */
    function simulate() external {
        DeploymentConfig memory config = loadConfig();
        validateConfig(config);

        console.log("\n=== SIMULATING DEPLOYMENT ===");
        console.log("This is a dry run - no transactions will be broadcast\n");

        console.log("Configuration:");
        console.log("  HORIZON Initial Supply:", config.horizonInitialSupply / 10 ** 18, "HORIZON");
        console.log("  Outcome Token URI:", config.outcomeTokenURI);
        console.log("  Protocol Owner:", config.protocolOwner);
        console.log("  Protocol Treasury:", config.protocolTreasury);
        console.log("  Arbitrator:", config.arbitrator);
        console.log("  AI Signer:", config.aiSigner);
        console.log("  Min Creator Stake:", config.minCreatorStake / 10 ** 18, "HORIZON");
        console.log("  Min Resolution Bond:", config.minResolutionBond / 10 ** 18, "HORIZON");
        console.log("  Dispute Window:", config.disputeWindow, "seconds");

        console.log("\nConfiguration validated successfully!");
        console.log("Ready to deploy with: forge script script/Deploy.s.sol:Deploy --rpc-url <RPC_URL> --broadcast");
    }
}
