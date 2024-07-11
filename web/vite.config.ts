import { vitePlugin as remix } from "@remix-run/dev";
import { installGlobals } from "@remix-run/node";
import { defineConfig } from "vite";
import tsconfigPaths from "vite-tsconfig-paths";
import svgr from "vite-plugin-svgr";

installGlobals();

export default defineConfig({
  server: {
    proxy: {
      "/api": {
        target: "http://localhost:14090",
        changeOrigin: true,
      },
    },
  },
  plugins: [remix(), tsconfigPaths(), svgr()],
});
