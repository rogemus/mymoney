#!/usr/bin/env node

import esbuild from 'esbuild';
import esbuildConfig from './esbuild.config.js';

await esbuild
  .build({
    ...esbuildConfig,
    minify: true,
    sourcemap: false,
  })
  .catch(() => (
    process.exit(1)
  ));
