# Frontend Usage Example

## How Frontend Can Access Fee Waived Information

The `FeeWaivedSummary` structure provides both requirements in a single, clean structure:

### 1. **Total Waived/Saved Fee** → `feeWaivedSummary.totalSaved`
### 2. **Details of Each Fee Waive** → `feeWaivedSummary.platformDetails[]`

## Frontend Implementation Examples

### React/TypeScript Example

```typescript
interface FeeWaivedSummary {
  userId: number;
  mainWalletAddr: string;
  totalSaved: number;        // ← REQUIREMENT 1: Total waived/saved fee
  totalVolume: number;
  overallReduction: number;
  maxFeeReduction: number;
  benefitSources?: NFTBenefitSources;
  platformDetails: PlatformFeeDetail[];  // ← REQUIREMENT 2: Details of each fee waive
  calculatedAt: string;
  nextUpdateAt?: string;
}

interface PlatformFeeDetail {
  platform: string;
  exchangeNameId?: number;
  walletAddress: string;
  tradingVolume: number;
  standardFeeRate: number;
  discountedFeeRate: number;
  feeReductionRate: number;
  feeSaved: number;          // ← Individual platform fee saved
  benefitSource: string;
  lastUpdated?: string;
}

// Component usage
const UserNFTDashboard: React.FC = () => {
  const [nftInfo, setNftInfo] = useState<GetUserNftInfoData | null>(null);

  useEffect(() => {
    fetchUserNFTInfo().then(setNftInfo);
  }, []);

  if (!nftInfo) return <div>Loading...</div>;

  const { feeWaivedSummary } = nftInfo;

  return (
    <div className="nft-dashboard">
      {/* REQUIREMENT 1: Display Total Waived/Saved Fee */}
      <div className="total-savings-card">
        <h2>Total Fee Savings</h2>
        <div className="amount">${feeWaivedSummary.totalSaved.toFixed(2)}</div>
        <div className="reduction">
          {(feeWaivedSummary.overallReduction * 100).toFixed(1)}% reduction
        </div>
      </div>

      {/* REQUIREMENT 2: Display Details of Each Fee Waive */}
      <div className="platform-details">
        <h3>Fee Savings by Platform</h3>
        {feeWaivedSummary.platformDetails.map((platform, index) => (
          <div key={index} className="platform-card">
            <div className="platform-header">
              <h4>{platform.platform.toUpperCase()}</h4>
              <span className="savings">${platform.feeSaved.toFixed(2)} saved</span>
            </div>
            <div className="platform-stats">
              <div>Volume: ${platform.tradingVolume.toLocaleString()}</div>
              <div>Standard Rate: {(platform.standardFeeRate * 100).toFixed(3)}%</div>
              <div>Discounted Rate: {(platform.discountedFeeRate * 100).toFixed(3)}%</div>
              <div>Reduction: {(platform.feeReductionRate * 100).toFixed(1)}%</div>
              <div>Source: {platform.benefitSource}</div>
            </div>
          </div>
        ))}
      </div>

      {/* Additional NFT benefit information */}
      {feeWaivedSummary.benefitSources && (
        <div className="benefit-sources">
          <h3>NFT Benefit Sources</h3>
          {feeWaivedSummary.benefitSources.tieredNft && (
            <div className="tiered-nft">
              <h4>Tiered NFT: {feeWaivedSummary.benefitSources.tieredNft.name}</h4>
              <p>Tier {feeWaivedSummary.benefitSources.tieredNft.tier}</p>
              <p>Discount: {(feeWaivedSummary.benefitSources.tieredNft.tradingFeeDiscount * 100).toFixed(1)}%</p>
            </div>
          )}
          {feeWaivedSummary.benefitSources.bestCompetitionNft && (
            <div className="competition-nft">
              <h4>Competition NFT: {feeWaivedSummary.benefitSources.bestCompetitionNft.name}</h4>
              <p>Source: {feeWaivedSummary.benefitSources.bestCompetitionNft.competitionSource}</p>
              <p>Discount: {(feeWaivedSummary.benefitSources.bestCompetitionNft.tradingFeeDiscount * 100).toFixed(1)}%</p>
            </div>
          )}
        </div>
      )}
    </div>
  );
};
```

### Vue.js Example

```vue
<template>
  <div class="nft-dashboard">
    <!-- REQUIREMENT 1: Total Waived/Saved Fee -->
    <div class="total-savings-card">
      <h2>Total Fee Savings</h2>
      <div class="amount">${{ totalSaved }}</div>
      <div class="reduction">{{ overallReductionPercent }}% reduction</div>
    </div>

    <!-- REQUIREMENT 2: Details of Each Fee Waive -->
    <div class="platform-details">
      <h3>Fee Savings by Platform</h3>
      <div 
        v-for="(platform, index) in platformDetails" 
        :key="index" 
        class="platform-card"
      >
        <div class="platform-header">
          <h4>{{ platform.platform.toUpperCase() }}</h4>
          <span class="savings">${{ platform.feeSaved.toFixed(2) }} saved</span>
        </div>
        <div class="platform-stats">
          <div>Volume: ${{ platform.tradingVolume.toLocaleString() }}</div>
          <div>Standard Rate: {{ (platform.standardFeeRate * 100).toFixed(3) }}%</div>
          <div>Discounted Rate: {{ (platform.discountedFeeRate * 100).toFixed(3) }}%</div>
          <div>Reduction: {{ (platform.feeReductionRate * 100).toFixed(1) }}%</div>
          <div>Source: {{ platform.benefitSource }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';

const nftInfo = ref<GetUserNftInfoData | null>(null);

// REQUIREMENT 1: Total waived/saved fee
const totalSaved = computed(() => 
  nftInfo.value?.feeWaivedSummary.totalSaved.toFixed(2) || '0.00'
);

const overallReductionPercent = computed(() => 
  ((nftInfo.value?.feeWaivedSummary.overallReduction || 0) * 100).toFixed(1)
);

// REQUIREMENT 2: Details of each fee waive
const platformDetails = computed(() => 
  nftInfo.value?.feeWaivedSummary.platformDetails || []
);

onMounted(async () => {
  nftInfo.value = await fetchUserNFTInfo();
});
</script>
```

### Simple JavaScript Example

```javascript
// Fetch user NFT info
async function displayFeeWaivedInfo() {
  const response = await fetch('/api/user/nft-info', {
    headers: { 'Authorization': `Bearer ${token}` }
  });
  const data = await response.json();
  
  const { feeWaivedSummary } = data.data;
  
  // REQUIREMENT 1: Display total waived/saved fee
  document.getElementById('total-saved').textContent = 
    `$${feeWaivedSummary.totalSaved.toFixed(2)}`;
  
  document.getElementById('overall-reduction').textContent = 
    `${(feeWaivedSummary.overallReduction * 100).toFixed(1)}%`;
  
  // REQUIREMENT 2: Display details of each fee waive
  const platformContainer = document.getElementById('platform-details');
  platformContainer.innerHTML = '';
  
  feeWaivedSummary.platformDetails.forEach(platform => {
    const platformDiv = document.createElement('div');
    platformDiv.className = 'platform-card';
    platformDiv.innerHTML = `
      <h4>${platform.platform.toUpperCase()}</h4>
      <div>Fee Saved: $${platform.feeSaved.toFixed(2)}</div>
      <div>Volume: $${platform.tradingVolume.toLocaleString()}</div>
      <div>Standard Rate: ${(platform.standardFeeRate * 100).toFixed(3)}%</div>
      <div>Discounted Rate: ${(platform.discountedFeeRate * 100).toFixed(3)}%</div>
      <div>Reduction: ${(platform.feeReductionRate * 100).toFixed(1)}%</div>
      <div>Source: ${platform.benefitSource}</div>
    `;
    platformContainer.appendChild(platformDiv);
  });
}
```

## Key Benefits for Frontend

### ✅ **Single Source of Truth**
- One API call gets all fee waived information
- No need to make multiple requests or combine data

### ✅ **Complete Information**
- **Total saved**: `feeWaivedSummary.totalSaved`
- **Platform breakdown**: `feeWaivedSummary.platformDetails[]`
- **NFT benefit sources**: `feeWaivedSummary.benefitSources`

### ✅ **Rich Analytics**
- Platform-specific performance
- Fee reduction rates and sources
- Trading volume breakdown
- Real-time calculations

### ✅ **Easy to Use**
- Clean, predictable structure
- TypeScript-friendly interfaces
- Straightforward data access patterns

This structure perfectly meets both requirements while providing a clean, maintainable frontend implementation!