// @ts-check
import { defineConfig } from 'astro/config';
import lottie from "astro-integration-lottie";
import tailwindcss from "@tailwindcss/vite";
import path from 'path';

// https://astro.build/config
export default defineConfig({
    integrations: [
    lottie(), 
  ],
  vite: {
     resolve: {
      alias: {
        '@': path.resolve('./src'), 
      },
    },
     plugins: [tailwindcss()],
  },
});
