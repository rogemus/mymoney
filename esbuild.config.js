export default {
  entryPoints: ['./assets/main.css'],
  outdir: './tracker/static',
  bundle: true,
  logLevel: 'info',
  loader: {
    '.css': 'css',
    '.otf': 'file',
    '.ttf': 'file',
  }
};
