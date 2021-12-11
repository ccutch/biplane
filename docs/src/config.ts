export const SITE = {
  title: 'Your Documentation Website',
  description: 'Your website description.',
  defaultLanguage: 'en_US',
};

export const OPEN_GRAPH = {
  image: {
    src: 'https://github.com/withastro/astro/blob/main/assets/social/banner.jpg?raw=true',
    alt: 'astro logo on a starry expanse of space,' + ' with a purple saturn-like planet floating in the right foreground',
  },
  twitter: 'astrodotbuild',
};

export const KNOWN_LANGUAGES = {
  English: 'en',
};

export const GITHUB_EDIT_URL = `https://github.com/ccutch/biplane/blob/main/docs/`;

// Uncomment this to add an "Join our Community" button to every page of documentation.
// export const COMMUNITY_INVITE_URL = `https://astro.build/chat`;

export const SIDEBAR = {
  en: [
    { text: '', header: true },
    { text: 'Setup', header: true },
    { text: 'Getting Started', link: 'en/getting-started' },

    { text: 'MVC Framework', header: true },
    { text: 'Models', link: 'en/mvc/models' },
    { text: 'Views', link: 'en/mvc/views' },
    { text: 'Controllers', link: 'en/mvc/controllers' },
    { text: 'Objects', link: 'en/mvc/objects' },
    { text: 'Auth', link: 'en/mvc/auth' },
  ],
};
