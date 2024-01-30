export default {
  entryPoints: ['./assets/css/index.css'],
  outdir: './tracker/static',
  bundle: true,
  logLevel: 'info',
  loader: {
    '.css': 'css',
    '.otf': 'file',
    '.ttf': 'file',
  }
};
