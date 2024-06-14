import {defineConfig} from "vite";
import react from "@vitejs/plugin-react";
import basicSsl from "@vitejs/plugin-basic-ssl";
import { fileURLToPath, URL } from 'url'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [react(), basicSsl()],
    build: {
        outDir: "./docs",
    },
    base: "./",
    resolve: {
        alias: {
            '@': fileURLToPath(new URL('./src', import.meta.url)),
            '@components': fileURLToPath(new URL('./src/components', import.meta.url)),
            '@utils': fileURLToPath(new URL('./src/utils', import.meta.url)),
            '@styles': fileURLToPath(new URL('./src/styles', import.meta.url)),
            '@assets': fileURLToPath(new URL('./src/assets', import.meta.url)),
        },
    },
});
