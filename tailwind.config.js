/** @type {import('tailwindcss').Config} */
const colors = require('tailwindcss/colors');

module.exports = {
  content: ["*.html", "./views/**/*.html",],
  theme: {
    colors: {
      ...colors,
      primary: '#83a598',
      secondary: '#d3869b',
      accent: '#bdae93',
      highlight: '#fe8019',

      background: '#282828',
      'background-light': '#3c3836',
      'background-dark': '#1d2021',

      text: '#ebdbb2',
      'text-light': '#ebdbb2',
      'text-dark': '#282828',

      link: '#83a598',
      'link-hover': '#d3869b',
      'link-visited': '#bdae93',

      button: '#83a598',
      'button-hover': '#d3869b',
      'button-active': '#fe8019',

      border: '#a89984',
      'border-light': '#3c3836',
      'border-dark': '#282828',

      error: '#fb4934',
      warning: '#fe8019',
      success: '#b8bb26',
      info: '#83a598',

      disabled: '#504945',

      header: '#282828',
      footer: '#282828',

      nav: '#3c3836',
      'nav-hover': '#504945',

      'dark-shadow': 'rgba(0, 0, 0, 0.5)',

      overlay: 'rgba(40, 40, 40, 0.8)',
      modal: '#3c3836',

      placeholder: '#928374',

      subtle: '#a89984',

      cta: '#fe8019',
      'cta-hover': '#d3869b',
      'cta-active': '#83a598',

      gradient: {
        start: '#3c3836',
        end: '#1d2021',
      },

      muted: '#928374',

      'light-primary': '#8f3f71',
      'light-secondary': '#076678',
      'light-accent': '#af3a03',
      'light-highlight': '#d65d0e',

      'light-background': '#fbf1c7',
      'light-background-light': '#f2e5bc',
      'light-background-dark': '#ebdbb2',

      'light-text': '#282828',
      'light-text-light': '#3c3836',
      'light-text-dark': '#504945',

      'light-link': '#076678',
      'light-link-hover': '#8f3f71',
      'light-link-visited': '#af3a03',

      'light-button': '#076678',
      'light-button-hover': '#8f3f71',
      'light-button-active': '#d65d0e',

      'light-border': '#bdae93',
      'light-border-light': '#f2e5bc',
      'light-border-dark': '#ebdbb2',

      'light-error': '#cc241d',
      'light-warning': '#d65d0e',
      'light-success': '#98971a',
      'light-info': '#458588',

      'light-disabled': '#928374',

      'light-header': '#fbf1c7',
      'light-footer': '#fbf1c7',

      'light-nav': '#f2e5bc',
      'light-nav-hover': '#ebdbb2',

      'light-shadow': 'rgba(0, 0, 0, 0.5)',

      'light-overlay': 'rgba(251, 241, 199, 0.8)',
      'light-modal': '#f2e5bc',

      'light-placeholder': '#504945',

      'light-subtle': '#3c3836',

      'light-cta': '#d65d0e',
      'light-cta-hover': '#8f3f71',
      'light-cta-active': '#076678',

      'light-gradient': {
        'light-start': '#f2e5bc',
        'light-end': '#ebdbb2',
      },

      'light-muted': '#504945',

    },
    extend: {},
  },
  plugins: [],
  darkMode: "class",
}

