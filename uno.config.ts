import {
  defineConfig,
  presetWind3
} from 'unocss'

export default defineConfig({
  presets: [
    presetWind3({
      dark: 'class', // This tells UnoCSS to look for the .dark class on <html>
    }),
  ],
  theme: {
    // Define your brand colors and semantic tokens here
  colors: {
      // These map to the CSS variables defined in preflights below
      background: 'var(--background)',
      foreground: 'var(--foreground)',
      card: 'var(--card)',
      'card-foreground': 'var(--card-foreground)',
      border: 'var(--border)',
      input: 'var(--input)',
      primary: 'var(--primary)',
      'primary-foreground': 'var(--primary-foreground)',
      muted: 'var(--muted)',
      'muted-foreground': 'var(--muted-foreground)',
      accent: 'var(--accent)',
      'accent-foreground': 'var(--accent-foreground)',
    },
  },
  shortcuts: [
    // Use the Array-of-Objects syntax for better reliability in some environments
    {
            'body-shortcut': 'bg-background text-foreground flex flex-col min-h-screen transition-colors duration-300 font-sans antialiased',
      
      'container-main': 'max-w-4xl mx-auto',
      'section-hero': 'text-center pt-24 pb-20 px-4',
      
      // 2. HEADINGS & TEXT
      'app-name': 'text-5xl font-extrabold mb-6 tracking-tight text-foreground',
      'app-description': 'text-xl text-muted-foreground max-w-2xl mx-auto mb-10 leading-relaxed',
      
      // 3. BUTTONS (Shadcn Style)
      //    Uses 'primary' color (White in Dark mode, Black in Light mode)
      'app-stores': 'flex flex-col sm:flex-row justify-center gap-4 items-center',
      'app-store-button': 'inline-flex items-center bg-primary text-primary-foreground h-12 px-6 rounded-lg hover:opacity-90 transition-all active:scale-95 min-w-[180px] no-underline shadow-sm',
      'app-store-svg': 'w-6 h-6 me-3 fill-current',
      'app-store-button-text': 'flex flex-col items-start leading-none',
      'app-store-button-text-1': 'text-[0.65rem] uppercase tracking-wide opacity-80',
      'app-store-button-text-2': 'font-bold text-base',
      
      // 4. CARDS (The Hover Fix)
      //    border-border: subtle border in both modes
      //    hover:bg-accent: subtle gray shift in both modes
      'section-sub-hero': 'py-20',
      'container-sub-main': 'max-w-6xl mx-auto px-6',
      'app-features-cards': 'grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6',
      'feature-card': 'p-6 rounded-xl border border-border bg-card text-card-foreground shadow-sm hover:bg-accent hover:text-accent-foreground transition-all duration-200',
      'feature-card-emoji': 'text-4xl block mb-4',
      'feature-card-text': 'text-lg font-bold mb-2',
      'feature-card-describtion': 'text-sm text-muted-foreground leading-relaxed',

       'footer-shortcut': 'mt-auto py-10 text-center border-t border-border text-muted-foreground text-sm',
      'footer-container': 'max-w-4xl mx-auto px-4',
      'footer-text': 'mb-2',
      'footer-links-space': 'flex flex-wrap justify-center gap-x-4 gap-y-2 underline underline-offset-4 text-sm font-medium text-foreground underline underline-offset-4 decoration-muted-foreground/40 hover:decoration-foreground', // gap-x-8 provides wide space
      // 'footer-link': 'text-blue-600 hover:underline',
      'control-btn': `
      inline-flex items-center justify-center rounded-md text-sm font-medium 
      ring-offset-background transition-colors focus-visible:outline-none 
      focus-visible:ring-2 focus-visible:ring-offset-2 
      disabled:pointer-events-none disabled:opacity-50 
      border border-border bg-background hover:bg-accent hover:text-accent-foreground 
      h-10 w-10 active:scale-95 border-none bg-transparent shadow-none
    `,
    'options-btn': 'fixed top-6 right-6 z-50 flex items-center gap-1 p-1 rounded-lg border border-border bg-background/80 backdrop-blur shadow-sm select-none',
    'separator': 'w-[1px] h-6 bg-border mx-1',
    // The "Pill" look for dropdown items
    'dropdown-menu': `
      absolute mt-3 min-w-[10rem] 
      p-1 rounded-md border border-border bg-card shadow-lg 
      z-50 flex flex-col overflow-hidden
      backdrop-blur-md bg-card/95
    `,
    'dropdown-item': `
      flex items-center justify-between w-full px-2 py-2
      text-sm rounded-sm transition-colors duration-150
      hover:bg-accent/80 hover:text-accent-foreground
      cursor-pointer no-underline whitespace-nowrap
    `,
    },
  ],
  preflights: [
    {
      getCSS: () => `
        :root {
          /* LIGHT MODE (Zinc/Slate) */
          color-scheme: light;
          --background: #ffffff;
          --foreground: #09090b;
          --card: #ffffff;
          --card-foreground: #09090b;
          --border: #e4e4e7;
          --input: #e4e4e7;
          --primary: #18181b;
          --primary-foreground: #fafafa;
          --muted: #f4f4f5;
          --muted-foreground: #71717a;
          --accent: #e2e8f0;
          --accent-foreground: #18181b;
        }

        .dark {
          /* DARK MODE (Zinc/Slate - Shadcn Style) */
          color-scheme: dark;
          --background: #09090b; /* Very dark zinc */
          --foreground: #fafafa;
          --card: #09090b;
          --card-foreground: #fafafa;
          --border: #27272a;      /* Subtle dark border */
          --input: #27272a;
          --primary: #fafafa;     /* White buttons in dark mode */
          --primary-foreground: #18181b; /* Black text on buttons */
          --muted: #27272a;
          --muted-foreground: #a1a1aa;
          --accent: #27272a;      /* Hover state color */
          --accent-foreground: #fafafa;
        }

        /* Base Setup */
        html {
           color-scheme: dark light;
           scroll-behavior: smooth;
        }
        
        body {
          font-family: "Inter", system-ui, sans-serif;
          -webkit-font-smoothing: antialiased;
        }

        /* Prevent unstyled link colors */
        a { color: inherit; text-decoration: none; }

        /* Arabic Override */
        @font-face {
          font-family: 'NotoArabic';
          src: url('/assets/fonts/noto-kufi-arabic-regular.woff2') format('woff2');
          font-weight: 400;
          font-display: swap;
        }
        @font-face {
          font-family: 'NotoArabic';
          src: url('/assets/fonts/noto-kufi-arabic-bold.woff2') format('woff2');
          font-weight: 700;
          font-display: swap;
        }

        [dir="rtl"] body {
          font-family: 'NotoArabic', system-ui, sans-serif !important;
        }

      /* Prevent transitions on initial load to stop the "fading" flicker */
      .preload * {
        -webkit-transition: none !important;
        -moz-transition: none !important;
        -ms-transition: none !important;
        -o-transition: none !important;
        transition: none !important;
      }

      *, ::before, ::after {
        box-sizing: border-box;
        }
        
        /* Fix for right-side clipping on small screens */
        .dropdown-menu {
          /* If right-0 is too aggressive, this acts as a safety */
          transform-origin: top right;
        }

      /* This specifically ensures that the hover effect stays 
         within the rounded corners of the dropdown-item */
      .dropdown-item {
        -webkit-appearance: none;
        appearance: none;
        user-select: none;
      }

     /* Force RTL for the control group */
      .fixed.top-6.right-6 {
        direction: rtl !important;
        /* Ensure children (the menu) aren't clipped */
        overflow: visible !important; 
      }

      [x-cloak] { 
    display: none !important; 
      }     

      `,
    },
  ],
})