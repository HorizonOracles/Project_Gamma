# DegenArena Design Guidelines

## Design Approach

**Reference-Based:** Inspired by BetDEX's modern betting interface, adapted with DegenArena's distinctive black/yellow/gold identity. This creates a premium, high-contrast sports betting experience that emphasizes live market dynamics and social engagement.

**Core Principles:**
- Bold, high-contrast visual hierarchy
- Real-time data emphasis with dynamic odds display
- Social proof through public betting feed
- Gamified engagement elements
- Crypto-native transparency

## Color Palette

**Dark Mode Primary (Default):**
- **Background Base:** 0 0% 0% (Pure Black)
- **Background Elevated:** 0 0% 8% (Cards, modals, elevated surfaces)
- **Background Subtle:** 0 0% 12% (Hover states, input backgrounds)

**Accent Colors:**
- **Yellow Primary:** 48 100% 50% (Primary CTAs, active states, highlights)
- **Gold Secondary:** 45 100% 38% (Secondary actions, borders, icons)
- **Yellow Muted:** 48 90% 45% (Disabled states, subtle accents)

**Text Hierarchy:**
- **Primary Text:** 0 0% 100% (White - headlines, key data)
- **Secondary Text:** 48 20% 85% (Warm off-white for body text)
- **Accent Text:** Yellow Primary (Links, odds, important metrics)

**Functional Colors:**
- **Success/Win:** 48 100% 50% (Yellow - winning bets, positive changes)
- **Loss/Negative:** 0 0% 40% (Dark grey - losing bets, negative changes)
- **Live/Active:** 48 100% 60% (Bright yellow pulse - live markets)
- **Locked/Pending:** 45 60% 30% (Muted gold - locked markets)

## Typography

**Font Family:**
- **Primary:** 'Inter', 'SF Pro Display', system-ui, sans-serif (Modern, crypto-native aesthetic)
- **Numeric/Data:** 'JetBrains Mono', 'Roboto Mono', monospace (Odds, balances, transaction data)

**Type Scale:**
- **Hero/Display:** text-5xl to text-7xl, font-bold (Landing hero, major announcements)
- **Page Headers:** text-3xl to text-4xl, font-bold (Dashboard sections)
- **Section Headers:** text-xl to text-2xl, font-semibold (Market categories)
- **Body Large:** text-base, font-medium (Game cards, primary content)
- **Body Regular:** text-sm, font-normal (Descriptions, secondary info)
- **Captions/Meta:** text-xs, font-medium (Timestamps, bet IDs, small labels)
- **Odds Display:** text-2xl to text-3xl, font-bold, monospace (Large, readable odds)

## Layout System

**Spacing Primitives:** Use Tailwind units of 2, 4, 6, 8, 12, 16, 20 for consistent rhythm.

**Grid Structure:**
- **Desktop (1280px+):** Three-column layout
  - Left Sidebar: 240px fixed (Sports navigation)
  - Center Content: flex-1 (Main betting interface)
  - Right Sidebar: 320px fixed (Featured bets, social feed)
- **Tablet (768px-1279px):** Two-column, collapsible sidebar
- **Mobile (<768px):** Single column, bottom navigation

**Container Widths:**
- **Max Content Width:** max-w-screen-2xl (Game cards, listings)
- **Betting Slips:** max-w-md (Focused bet placement)
- **Modals:** max-w-2xl (Wallet, settings, detailed views)

**Vertical Rhythm:**
- Section spacing: py-8 to py-12
- Card spacing: p-4 to p-6
- Component gaps: gap-4 to gap-6

## Component Library

### Navigation & Layout

**Left Sidebar (Sports Menu):**
- Fixed position, black background with gold border-r
- Sport categories with yellow icons and event counts in gold badges
- Hover: Background subtle with yellow left border (border-l-2)
- Active sport: Yellow background at 10% opacity with gold text

**Top Navigation Bar:**
- Black background, gold bottom border
- Logo: Yellow/gold wordmark on left
- Wallet balance: Monospace font, yellow text, black pill background with gold border
- User menu: Profile icon in gold, dropdown with black elevated background
- Mobile: Hamburger menu in yellow

**Right Sidebar:**
- Featured Matches section with yellow gradient header
- Public Bet Feed: Scrollable cards showing recent bets
- Card style: Black elevated background, gold borders, yellow accent on bet amounts

### Market Display Cards

**Live Game Cards:**
- **Container:** Black elevated background (bg 8%), rounded-xl, gold border on hover
- **Header:** 
  - Sport icon (yellow), league name (secondary text)
  - LIVE badge: Animated yellow pulse, gold border, uppercase text
  - Time/Status in muted gold
- **Matchup Display:**
  - Team names in white, bold, text-lg
  - "VS" separator in gold, uppercase
  - Score (if live) in large yellow monospace
- **Odds Buttons (Dual Pool):**
  - **Pool A Button:** Yellow background, black text, hover brightens 10%
  - **Pool B Button:** Gold background, black text, hover brightens 10%
  - Button layout: Full-width split (grid-cols-2 gap-2)
  - Odds value: Large monospace font (text-2xl)
  - Pool size: Small text below odds in muted color
  - Dynamic update animation: Brief yellow glow on change

**Market Types Grid:**
- Each market type as a card with gold border-t accent
- Icon in yellow, title in white, description in secondary text
- Hover: Lift with shadow and yellow border glow

### Betting Interface

**Bet Slip Modal:**
- **Background:** Black elevated with gold border
- **Header:** Yellow gradient background, white text "Your Bet"
- **Market Summary:** 
  - Team/outcome name in white bold
  - Current odds in large yellow monospace
  - "Odds may change" warning in gold italic
- **Amount Input:**
  - Large input field, black subtle background, yellow border on focus
  - BNB symbol prefix in gold
  - Max button: Small gold border pill
- **Potential Payout Display:**
  - Large yellow text with monospace font
  - "Potential Return" label in secondary text
  - Live recalculation with yellow pulse animation
- **Confirm Button:** 
  - Full-width, yellow background, black bold text
  - Loading state: Gold spinner
  - Success: Brief checkmark animation in yellow

### Wallet Management

**Wallet Dashboard:**
- **Balance Card:**
  - Large balance in yellow monospace (text-5xl)
  - "BNB" label in gold
  - Secondary balances (USD equivalent) in secondary text
- **Deposit Address Display:**
  - Monospace font, black subtle background
  - Copy button with yellow icon
  - QR code in gold on white background (modal popup)
- **Transaction History:**
  - Table with alternating row backgrounds (subtle black)
  - Type icons in yellow (deposit/withdrawal/bet)
  - Amounts in monospace, yellow for incoming, muted for outgoing
  - BscScan links in gold with external link icon
- **Action Buttons:**
  - Deposit: Yellow solid button
  - Withdraw: Gold border button with black background

### Social Features

**Public Bet Card (Feed):**
- Black elevated background, rounded-lg
- User avatar with gold ring border
- Username in white, timestamp in muted gold
- Bet details: "UserX bet 5 BNB on Lakers Win" with yellow bet amount
- Odds at time of bet in small monospace gold
- Outcome status badge (if settled): Yellow for win, muted for loss

**User Profile:**
- **Header:** Black background with yellow accent gradient top border
- **Stats Row:** 
  - Win Rate: Large percentage in yellow
  - Total Wagered: Monospace gold
  - Badges/Achievements: Yellow icons with gold labels
- **Loyalty Tier Display:**
  - Bronze/Silver/Gold icon in corresponding color (gold for Gold tier)
  - Progress bar: Black background, yellow fill
  - Next tier benefits in secondary text with yellow bullet points

**Achievement Badges:**
- Circular icons with yellow/gold gradient backgrounds
- White icon symbols
- Tooltip on hover with black elevated background, gold border

### Data Visualization

**Pool Distribution Chart:**
- Horizontal stacked bar showing pool allocation
- Yellow fill for selected outcome
- Gold fill for opposite outcome
- Percentage labels in white on bars
- Total pool size in large monospace below

**Odds History Graph:**
- Line chart with yellow line, gold shadow
- Black grid lines at 10% opacity
- Y-axis: Odds values in monospace
- X-axis: Time labels in small secondary text
- Current odds highlighted with yellow dot

### Forms & Inputs

**Text Inputs:**
- Black subtle background, gold border
- Yellow border on focus with subtle glow
- Placeholder text in muted gold
- Error state: Red border (only functional color exception)

**Buttons:**
- **Primary CTA:** Yellow background, black bold text, rounded-lg
- **Secondary:** Gold border, yellow text, black background
- **Tertiary:** Text only, yellow color
- **Disabled:** Muted yellow background at 30% opacity, grey text

**Toggles/Switches:**
- Track: Black subtle background
- Active: Yellow fill
- Thumb: White circle with gold shadow

## Animations & Interactions

**Micro-interactions (Subtle):**
- Odds update: 200ms yellow pulse on value change
- LIVE badge: Gentle 2s infinite yellow opacity pulse (80%-100%)
- Button hover: 150ms brightness increase
- Card hover: 200ms lift with gold border glow
- Loading states: Yellow spinner or skeleton with gold shimmer

**Transitions:**
- Page navigation: 300ms fade
- Modal appearance: 200ms scale from 0.95 with fade
- Sidebar collapse: 250ms slide with easing

## Images & Visual Assets

**Landing Page Hero:**
- Large, dramatic hero image showing sports action (basketball, football, soccer montage)
- Dark overlay at 70% opacity to maintain readability
- Yellow gradient overlay from bottom (20% opacity)
- Hero CTAs as yellow solid buttons with blurred backgrounds
- Image: 1920x800px minimum, focal point center

**Sport Category Icons:**
- Line-style icons in yellow/gold
- Size: 24px standard, 32px for headers
- Source: Heroicons or custom SVG set

**Achievement Badges:**
- Circular icons with gold/yellow gradients
- Trophy, star, flame, target symbols
- Size: 48px for profile display, 32px for compact views

**Promotional Banners:**
- Gradient backgrounds (yellow to gold diagonal)
- Bold typography overlay in black
- Subtle pattern texture at 5% opacity

## Accessibility & Usability

**Contrast Ratios:**
- Yellow on Black: Excellent contrast (>7:1)
- White text on Black: Maximum contrast
- Gold borders clearly visible against black

**Focus States:**
- All interactive elements have yellow 2px outline on keyboard focus
- Focus outline offset by 2px for clarity

**Responsive Breakpoints:**
- Mobile: 640px (single column, bottom nav)
- Tablet: 768px (collapsible sidebar)
- Desktop: 1024px (three-column layout)
- Large: 1280px+ (optimized spacing)

This design creates a bold, premium sports betting platform that leverages high contrast for clarity while maintaining a cohesive, crypto-native aesthetic with the distinctive black/yellow/gold palette.