import {
  defineConfig,
  presetWind3
} from 'unocss'

export default defineConfig({
  presets: [
    presetWind3({
      dark: 'class',
    }),
  ],
  theme: {
    colors: {
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
    {
      'header-shortcut': 'sticky top-0 z-100 w-full bg-background border-b border-border m-0 p-0 pb-4',
      'header-container': 'max-w-4xl mx-auto px-6 h-16 flex items-center justify-between',
      'body-shortcut': 'bg-background text-foreground flex flex-col min-h-screen transition-colors duration-300 font-sans antialiased m-0 p-0',
      'container-main': 'max-w-4xl mx-auto px-6',
      'section-hero': 'py-16 md:py-24 flex flex-col-reverse md:flex-row items-center justify-between gap-12',
      'hero-image-container': 'w-full md:w-1/2 flex justify-center',
      'hero-content': 'w-full md:w-1/2 flex flex-col items-center md:items-start text-center md:text-start',
      'app-name': 'text-4xl md:text-6xl font-extrabold mb-6 tracking-tight text-foreground text-center md:text-start',
      'app-description': 'text-lg md:text-xl text-muted-foreground max-w-lg mb-10 leading-relaxed text-center md:text-start',
      // Cleaned up shortcut
      'app-stores': 'flex flex-col md:flex-row items-center justify-center md:justify-start gap-4 w-full mx-auto md:mx-0',
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
      'footer-links-space': 'flex flex-wrap justify-center gap-x-4 gap-y-2 underline underline-offset-4 text-sm font-medium text-foreground underline underline-offset-4 decoration-muted-foreground/40 hover:decoration-foreground',
      'control-btn': `
        inline-flex items-center justify-center rounded-md text-sm font-medium 
        ring-offset-background transition-colors focus-visible:outline-none 
        focus-visible:ring-2 focus-visible:ring-offset-2 
        disabled:pointer-events-none disabled:opacity-50 
        border border-border bg-background hover:bg-accent hover:text-accent-foreground 
        h-10 w-10 active:scale-95 border-none bg-transparent shadow-none
      `,
      'options-btn': 'fixed top-6 [right:1.5rem] [left:auto] z-50 flex items-center gap-1 p-1 rounded-lg border border-border bg-background/80 backdrop-blur shadow-sm select-none',
      'separator': 'w-[1px] h-6 bg-border mx-1',
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
        color-scheme: dark;
        /* Lofi Soft Charcoal/Blue - replace #09090b */
        --background: #1a1b1e; 
        /* Softer foreground (off-white) - replace #fafafa */
        --foreground: #d1d5db; 
        
        /* Cards slightly lighter than background for depth */
        --card: #222327;
        --card-foreground: #e5e7eb;
        
        /* Softer borders to reduce the "grid" feel */
        --border: #2d2e33;
        --input: #2d2e33;
        
        /* Primary buttons: less "piercing" white */
        --primary: #e5e7eb;
        --primary-foreground: #1a1b1e;
        
        /* Muted elements: softer grays */
        --muted: #2a2b30;
        --muted-foreground: #9ca3af;
        
        /* Accent/Hover: subtle shift */
        --accent: #2d2e33;
        --accent-foreground: #ffffff;
      }

        /* SYNCED THEME TRANSITION */
        body, header {
          background-color: var(--background) !important;
          transition: background-color 0.3s ease, border-color 0.3s ease !important;
          margin: 0;
          padding: 0;
        }

        header {
          position: sticky;
          top: 0;
          z-index: 100;
        }

        html {
           color-scheme: dark light;
           scroll-behavior: smooth;
        }
        
        body {
          font-family: "Inter", system-ui, sans-serif;
          -webkit-font-smoothing: antialiased;
        }

        a { color: inherit; text-decoration: none; }

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

        .font-arabic {
          font-family: 'NotoArabic', system-ui, sans-serif !important;
        }

        [dir="rtl"] body {
          font-family: 'NotoArabic', system-ui, sans-serif !important;
        }

        .preload * {
          transition: none !important;
        }

        *, ::before, ::after {
          box-sizing: border-box;
        }
        
        .dropdown-menu {
          transform-origin: top right;
        }

        .dropdown-item {
          -webkit-appearance: none;
          appearance: none;
          user-select: none;
        }

        .fixed.top-6.right-6 {
          direction: rtl !important;
          overflow: visible !important; 
        }

        .options-btn {
          right: 1.5rem !important; 
          left: auto !important;
          direction: rtl !important; /* This keeps icons inside the button from flipping */
        }

        /* Remove the mobile tap highlight globally */
      * {
        -webkit-tap-highlight-color: transparent;
        -webkit-touch-callout: none; /* Prevents callout menu on long-press if desired */
      }

        [x-cloak], [un-cloak] { 
          display: none !important; 
        }     
      `,
    },
  ],
})