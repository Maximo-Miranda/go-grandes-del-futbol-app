import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import vuetify from "vite-plugin-vuetify";
import { resolve } from "path";

export default defineConfig(({ command }) => ({
  plugins: [vue(), vuetify({ autoImport: true })],
  resolve: {
    alias: {
      "@": resolve(__dirname, "resources/js"),
    },
  },
  base: command === "build" ? "/build/" : "/",
  publicDir: false,
  build: {
    manifest: true,
    outDir: "public/build",
    rollupOptions: {
      input: "resources/js/app.ts",
    },
  },
  server: {
    origin: "http://localhost:5173",
    cors: true,
  },
  optimizeDeps: {
    include: ["@inertiajs/vue3", "@inertiajs/core", "vue"],
  },
}));
