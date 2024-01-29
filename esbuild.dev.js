#!/usr/bin/env node

import esbuild from 'esbuild';
import esbuildConfig from './esbuild.config.js';

let ctx = await esbuild.context({
  ...esbuildConfig,
  minify: false,
  sourcemap: true,
});


await ctx.watch();
await ctx.serve({
  servedir: esbuildConfig.outdir,
  port: 1420,
});
