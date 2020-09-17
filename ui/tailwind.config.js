module.exports = {
  future: {
    purgeLayersByDefault: true,
    removeDeprecatedGapUtilities: true
  },
  purge: ['./public/**/*.html', './src/**/*.vue', '../templates/**/*.html'],
  theme: {
    extend: {
      fontFamily: {
        sans: [
          'Inter var',
          ...require('tailwindcss/defaultTheme').fontFamily.sans
        ]
      }
    }
  },
  variants: {},
  plugins: [require('@tailwindcss/ui')]
}
