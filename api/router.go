package main

import (
	"github.com/aiw3/nft-solana-api/nfts"
	"github.com/swaggest/rest/web"
)

// ==========================================
// API ROUTES SETUP
// ==========================================

func setupAPIRoutes(s *web.Service) {
	// ==========================================
	// 🎯 FRONTEND USER ENDPOINTS (NFT Related)
	// ==========================================

	// NFT Data & Management
	s.Get("/api/user/nft-info", nfts.GetUserNftInfo()) // Complete NFT portfolio + badge summary
	//s.Get("/api/user/nft-avatars", nfts.GetNftAvatars())       // Available NFT avatars for profile
	//s.Post("/api/user/nft/claim", nfts.ClaimNFT())             // Claim Level 1 NFT
	//s.Get("/api/user/nft/can-upgrade", nfts.CanUpgradeNFT())   // Check upgrade eligibility
	//s.Post("/api/user/nft/upgrade", nfts.UpgradeNFT())         // Upgrade to higher level
	//s.Post("/api/user/nft/activate", nfts.ActivateTieredNFT()) // Activate NFT benefits
	//
	//// Badge Data & Management
	//s.Get("/api/user/badges", badges.GetUserBadges())          // Complete badge portfolio
	//s.Get("/api/badges/{level}", badges.GetBadgesByLevel())    // Level-specific badges
	//s.Post("/api/user/badge/activate", badges.ActivateBadge()) // Activate earned badge
	//
	//// Badge Task System
	//s.Post("/api/badge/task-complete", badges.CompleteTask())       // Complete badge task with anti-gaming
	//s.Get("/api/badge/status", badges.GetBadgeStatus())             // Get badge status and progress
	//s.Post("/api/badge/activate", badges.ActivateBadgeForUpgrade()) // Activate badge for NFT upgrades
	//s.Get("/api/badge/list", badges.GetBadgeList())                 // Get all available badges
	//
	//// ==========================================
	//// 👑 ADMIN ENDPOINTS
	//// ==========================================
	//
	//// NFT Management
	//s.Post("/api/admin/nft/upload-image", admin.UploadTierImage())     // Upload NFT images to IPFS
	//s.Get("/api/admin/users/nft-status", admin.GetAllUsersNftStatus()) // User NFT status overview
	//
	//// Competition Management
	//s.Post("/api/admin/competition-nfts/award", admin.AwardCompetitionNFTs()) // Award competition NFTs
	//
	//// Avatar Management
	//s.Post("/api/admin/profile-avatars/upload", admin.UploadAvatar())        // Upload profile avatars
	//s.Get("/api/admin/profile-avatars/list", admin.ListAvatars())            // List profile avatars
	//s.Put("/api/admin/profile-avatars/{id}/update", admin.UpdateAvatar())    // Update profile avatar
	//s.Delete("/api/admin/profile-avatars/{id}/delete", admin.DeleteAvatar()) // Delete profile avatar
}
