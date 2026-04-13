/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{ts,tsx,js,jsx}'],
  theme: {
    extend: {
      colors: {
        primary: 'var(--tzone-primary)',
        'primary-light': 'var(--tzone-primary-light)',
        accent: 'var(--tzone-accent)',
        'accent-light': 'var(--tzone-accent-light)',
        success: 'var(--tzone-success)',
        warning: 'var(--tzone-warning)',
        danger: 'var(--tzone-danger)',
        'danger-dark': 'var(--tzone-danger-dark)',
        surface: 'var(--tzone-surface)',
        'surface-light': 'var(--tzone-surface-light)',
        'surface-lighter': 'var(--tzone-surface-lighter)',
        border: 'var(--tzone-border)',
        'text-primary': 'var(--tzone-text-primary)',
        'text-secondary': 'var(--tzone-text-secondary)',
        'text-muted': 'var(--tzone-text-muted)',
      },
      fontFamily: {
        sans: [
          'Inter',
          'ui-sans-serif',
          'system-ui',
          '-apple-system',
          'BlinkMacSystemFont',
          'Segoe UI',
          'sans-serif',
        ],
      },
      boxShadow: {
        soft: '0 12px 40px rgba(2, 6, 23, 0.22)',
        glow: '0 24px 80px rgba(2, 6, 23, 0.45)',
      },
    },
  },
};

